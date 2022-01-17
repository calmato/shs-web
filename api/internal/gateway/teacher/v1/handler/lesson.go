package handler

import (
	"net/http"
	"strconv"

	gentity "github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/request"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/response"
	"github.com/calmato/shs-web/api/internal/gateway/util"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/gin-gonic/gin"
)

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
