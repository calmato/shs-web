package handler

import (
	"fmt"
	"sync"
	"time"

	"github.com/calmato/shs-web/api/internal/gateway/util"
	"github.com/calmato/shs-web/api/pkg/firebase/authentication"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/**
 * ###############################################
 * handler
 * ###############################################
 */
type APIV1Handler interface {
	AuthRoutes(rg *gin.RouterGroup)   // 認証済みでアクセス可能なエンドポイント一覧
	AdminRoutes(rg *gin.RouterGroup)  // 管理者のみアクセス可能なエンドポイント一覧
	NoAuthRoutes(rg *gin.RouterGroup) // 未認証でもアクセス可能なエンドポイント一覧
	Authentication() gin.HandlerFunc  // 認証情報の検証
	Authorization() gin.HandlerFunc   // 認可情報の検証
}

type apiV1Handler struct {
	now         func() time.Time
	logger      *zap.Logger
	sharedGroup *singleflight.Group
	waitGroup   *sync.WaitGroup
	auth        authentication.Client
	classroom   classroom.ClassroomServiceClient
	lesson      lesson.LessonServiceClient
	user        user.UserServiceClient
}

type Params struct {
	Auth             authentication.Client
	ClassroomService classroom.ClassroomServiceClient
	LessonService    lesson.LessonServiceClient
	UserService      user.UserServiceClient
	Logger           *zap.Logger
	WaitGroup        *sync.WaitGroup
}

func NewAPIV1Handler(params *Params) APIV1Handler {
	return &apiV1Handler{
		now:         jst.Now,
		logger:      params.Logger,
		waitGroup:   params.WaitGroup,
		auth:        params.Auth,
		classroom:   params.ClassroomService,
		lesson:      params.LessonService,
		user:        params.UserService,
		sharedGroup: &singleflight.Group{},
	}
}

/**
 * ###############################################
 * routes
 * ###############################################
 */
func (h *apiV1Handler) AuthRoutes(rg *gin.RouterGroup) {
	rg.GET("/v1/me", h.GetAuth)
	rg.PATCH("/v1/me/subjects", h.UpdateMySubjects)
	rg.PATCH("/v1/me/mail", h.UpdateMyMail)
	rg.PATCH("/v1/me/password", h.UpdateMyPassword)
	rg.GET("/v1/teachers", h.ListTeachers)
	rg.GET("/v1/teachers/:teacherId", h.GetTeacher)
	rg.GET("/v1/teachers/:teacherId/submissions", h.ListTeacherSubmissions)
	rg.GET("/v1/teachers/:teacherId/submissions/:summaryId", h.ListTeacherShifts)
	rg.POST("/v1/teachers/:teacherId/submissions/:summaryId", h.UpsertTeacherShifts)
	rg.GET("v1/students", h.ListStudents)
	rg.GET("/v1/students/:studentId", h.GetStudent)
	rg.GET("/v1/subjects", h.ListSubjects)
	rg.GET("/v1/schedules", h.ListSchedules)
	rg.GET("/v1/rooms", h.GetRoomsTotal)
	rg.GET("/v1/shifts", h.ListShiftSummaries)
	rg.GET("/v1/shifts/:shiftId", h.ListShifts)
}

func (h *apiV1Handler) AdminRoutes(rg *gin.RouterGroup) {
	rg.POST("/v1/teachers", h.CreateTeacher)
	rg.DELETE("/v1/teachers/:teacherId", h.DeleteTeacher)
	rg.PATCH("/v1/teachers/:teacherId/mail", h.UpdateTeacherMail)
	rg.PATCH("/v1/teachers/:teacherId/password", h.UpdateTeacherPassword)
	rg.PATCH("/v1/teachers/:teacherId/role", h.UpdateTeacherRole)
	rg.PATCH("/v1/teachers/:teacherId/subjects", h.UpdateTeacherSubject)
	rg.POST("/v1/students", h.CreateStudent)
	rg.DELETE("v1/students/:studentId", h.DeleteStudent)
	rg.POST("/v1/subjects", h.CreateSubject)
	rg.PATCH("/v1/subjects/:subjectId", h.UpdateSubject)
	rg.DELETE("/v1/subjects/:subjectId", h.DeleteSubject)
	rg.PATCH("/v1/schedules", h.UpdateSchedules)
	rg.PATCH("/v1/rooms", h.UpdateRoomsTotal)
	rg.POST("/v1/shifts", h.CreateShifts)
	rg.PATCH("/v1/shifts/:shiftId/schedule", h.UpdateShiftSummarySchedule)
	rg.PATCH("/v1/shifts/:shiftId/decided", h.UpdateShiftSummaryDecided)
	rg.DELETE("/v1/shifts/:shiftId", h.DeleteShiftSummary)
	rg.GET("/v1/shifts/:shiftId/submissions/:submissionId", h.ListShiftSubmissions)
	rg.GET("/v1/shifts/:shiftId/teachers/:teacherId", h.ListEnabledTeacherShifts)
	rg.GET("/v1/shifts/:shiftId/students/:studentId", h.ListEnabledStudentShifts)
	rg.GET("/v1/shifts/:shiftId/lessons", h.ListLessons)
	rg.POST("/v1/shifts/:shiftId/lessons", h.CreateLesson)
	rg.PATCH("/v1/shifts/:shiftId/lessons/:lessonId", h.UpdateLesson)
	rg.DELETE("/v1/shifts/:shiftId/lessons/:lessonId", h.DeleteLesson)
}

func (h *apiV1Handler) NoAuthRoutes(rg *gin.RouterGroup) {}

/**
 * ###############################################
 * error handling
 * ###############################################
 */
func httpError(ctx *gin.Context, err error) {
	res, code := util.NewErrorResponse(err)
	ctx.JSON(code, res)
	ctx.Abort()
}

func badRequest(ctx *gin.Context, err error) {
	httpError(ctx, status.Error(codes.InvalidArgument, err.Error()))
}

func unauthorized(ctx *gin.Context, err error) {
	httpError(ctx, status.Error(codes.Unauthenticated, err.Error()))
}

func forbidden(ctx *gin.Context, err error) {
	httpError(ctx, status.Error(codes.PermissionDenied, err.Error()))
}

/**
 * ###############################################
 * other
 * ###############################################
 */
func (h *apiV1Handler) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := util.GetAuthToken(ctx)
		if err != nil {
			unauthorized(ctx, err)
			return
		}

		teacherID, err := h.auth.VerifyIDToken(ctx, token)
		if err != nil || teacherID == "" {
			unauthorized(ctx, err)
			return
		}

		setAuth(ctx, teacherID)

		ctx.Next()
	}
}

func (h *apiV1Handler) Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := util.GetAuthToken(ctx)
		if err != nil {
			unauthorized(ctx, err)
			return
		}

		teacherID, err := h.auth.VerifyIDToken(ctx, token)
		if err != nil || teacherID == "" {
			unauthorized(ctx, err)
			return
		}

		c := util.SetMetadata(ctx)
		teacher, err := h.getTeacher(c, teacherID)
		if err != nil {
			unauthorized(ctx, err)
			return
		}

		if !teacher.AdminRole() {
			err := fmt.Errorf("handler: permission denied")
			forbidden(ctx, err)
			return
		}

		setAuth(ctx, teacherID)

		ctx.Next()
	}
}

func setAuth(ctx *gin.Context, userID string) {
	if userID != "" {
		ctx.Request.Header.Set("userId", userID)
	}
}

func getTeacherID(ctx *gin.Context) string {
	return ctx.GetHeader("userId")
}
