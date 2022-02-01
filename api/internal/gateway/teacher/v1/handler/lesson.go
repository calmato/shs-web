package handler

import (
	"net/http"
	"strconv"

	gentity "github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/request"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/response"
	"github.com/calmato/shs-web/api/internal/gateway/util"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

func (h *apiV1Handler) ListLessons(ctx *gin.Context) {
	c := util.SetMetadata(ctx)
	now := h.now()
	var (
		defaultSince = jst.FormatYYYYMMDD(jst.BeginningOfMonth(now.Year(), int(now.Month())))
		defaultUntil = jst.FormatYYYYMMDD(jst.EndOfMonth(now.Year(), int(now.Month())))
	)

	in := &lesson.ListLessonsByDurationRequest{
		TeacherId: getTeacherID(ctx),
		Since:     ctx.DefaultQuery("since", defaultSince),
		Until:     ctx.DefaultQuery("until", defaultUntil),
	}
	out, err := h.lesson.ListLessonsByDuration(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}
	glessons := gentity.NewLessons(out.Lessons)
	gShifts := gentity.NewShifts(out.Shifts)

	eg, ectx := errgroup.WithContext(c)
	var gteachers gentity.Teachers
	eg.Go(func() (err error) {
		gteachers, err = h.multiGetTeachers(ectx, glessons.TeacherIDs())
		return
	})
	var gstudents gentity.Students
	eg.Go(func() (err error) {
		gstudents, err = h.multiGetStudents(ectx, glessons.StudentIDs())
		return
	})
	if err := eg.Wait(); err != nil {
		httpError(ctx, err)
		return
	}

	lessons, err := entity.NewLessons(glessons, gShifts.Map())
	if err != nil {
		httpError(ctx, err)
		return
	}
	res := &response.LessonsResponse{
		Lessons:  lessons,
		Teachers: entity.NewTeachers(gteachers, nil), // パフォーマンスを考慮して担当科目は取得しない
		Students: entity.NewStudents(gstudents, nil), // パフォーマンスを考慮して受講科目は取得しない
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV1Handler) CreateLesson(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	req := &request.CreateLessonRequest{}
	if err := ctx.BindJSON(req); err != nil {
		badRequest(ctx, err)
		return
	}
	summaryID, err := strconv.ParseInt(ctx.Param("shiftId"), 10, 64)
	if err != nil {
		badRequest(ctx, err)
		return
	}

	in := &lesson.CreateLessonRequest{
		ShiftSummaryId: summaryID,
		ShiftId:        req.ShiftID,
		SubjectId:      req.SubjectID,
		RoomId:         req.Room,
		TeacherId:      req.TeacherID,
		StudentId:      req.StudentID,
	}
	out, err := h.lesson.CreateLesson(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}
	glesson := gentity.NewLesson(out.Lesson)
	gshift := gentity.NewShift(out.Shift)

	lesson, err := entity.NewLesson(glesson, gshift)
	if err != nil {
		httpError(ctx, err)
		return
	}
	res := &response.LessonResponse{
		Lesson: lesson,
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV1Handler) UpdateLesson(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	req := &request.UpdateLessonRequest{}
	if err := ctx.BindJSON(req); err != nil {
		badRequest(ctx, err)
		return
	}
	summaryID, err := strconv.ParseInt(ctx.Param("shiftId"), 10, 64)
	if err != nil {
		badRequest(ctx, err)
		return
	}
	lessonID, err := strconv.ParseInt(ctx.Param("lessonId"), 10, 64)
	if err != nil {
		badRequest(ctx, err)
		return
	}

	in := &lesson.UpdateLessonRequest{
		LessonId:       lessonID,
		ShiftSummaryId: summaryID,
		ShiftId:        req.ShiftID,
		SubjectId:      req.SubjectID,
		RoomId:         req.Room,
		TeacherId:      req.TeacherID,
		StudentId:      req.StudentID,
	}
	_, err = h.lesson.UpdateLesson(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (h *apiV1Handler) DeleteLesson(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	lessonID, err := strconv.ParseInt(ctx.Param("lessonId"), 10, 64)
	if err != nil {
		badRequest(ctx, err)
		return
	}

	in := &lesson.DeleteLessonRequest{
		LessonId: lessonID,
	}
	_, err = h.lesson.DeleteLesson(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}
