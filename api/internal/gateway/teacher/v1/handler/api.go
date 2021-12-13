package handler

import (
	"sync"
	"time"

	"github.com/calmato/shs-web/api/internal/gateway/util"
	"github.com/calmato/shs-web/api/pkg/firebase/authentication"
	"github.com/calmato/shs-web/api/pkg/jst"
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
}

type apiV1Handler struct {
	now         func() time.Time
	logger      *zap.Logger
	sharedGroup *singleflight.Group
	waitGroup   *sync.WaitGroup
	auth        authentication.Client
	user        user.UserServiceClient
}

type Params struct {
	Auth        authentication.Client
	UserService user.UserServiceClient
	Logger      *zap.Logger
	WaitGroup   *sync.WaitGroup
}

func NewAPIV1Handler(params *Params) APIV1Handler {
	return &apiV1Handler{
		now:         jst.Now,
		logger:      params.Logger,
		waitGroup:   params.WaitGroup,
		auth:        params.Auth,
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
}

func (h *apiV1Handler) AdminRoutes(rg *gin.RouterGroup) {
	rg.POST("/v1/teachers", h.CreateTeacher)
}

func (h *apiV1Handler) NoAuthRoutes(rg *gin.RouterGroup) {
	rg.POST("/v1/hello", h.Hello)
}

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

func setAuth(ctx *gin.Context, userID string) {
	if userID != "" {
		ctx.Set("userId", userID)
	}
}

func getTeacherID(ctx *gin.Context) string {
	return ctx.GetHeader("userId")
}
