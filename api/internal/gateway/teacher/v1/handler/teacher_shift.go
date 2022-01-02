package handler

import (
	"net/http"
	"strconv"

	gentity "github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/response"
	"github.com/calmato/shs-web/api/internal/gateway/util"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/gin-gonic/gin"
)

func (h *apiV1Handler) ListTeacherSubmissions(ctx *gin.Context) {
	c := util.SetMetadata(ctx)
	const (
		defaultLimit  = "30"
		defaultOffset = "0"
	)

	teacherID := ctx.Param("teacherId")
	limit, err := strconv.ParseInt(ctx.DefaultQuery("limit", defaultLimit), 10, 64)
	if err != nil {
		badRequest(ctx, err)
		return
	}
	offset, err := strconv.ParseInt(ctx.DefaultQuery("offset", defaultOffset), 10, 64)
	if err != nil {
		badRequest(ctx, err)
		return
	}

	shiftsIn := &lesson.ListShiftSummariesRequest{
		Status:  lesson.ShiftStatus_SHIFT_STATUS_ACCEPTING,
		OrderBy: lesson.ListShiftSummariesRequest_ORDER_BY_YEAR_MONTH_DESC,
		Limit:   limit,
		Offset:  offset,
	}
	shiftsOut, err := h.lesson.ListShiftSummaries(c, shiftsIn)
	if err != nil {
		httpError(ctx, err)
		return
	}
	summaries := gentity.NewShiftSummaries(shiftsOut.Summaries)

	submissionsIn := &lesson.ListTeacherSubmissionsByShiftSummaryIDsRequest{
		TeacherId:       teacherID,
		ShiftSummaryIds: summaries.IDs(),
	}
	submissionsOut, err := h.lesson.ListTeacherSubmissionsByShiftSummaryIDs(c, submissionsIn)
	if err != nil {
		httpError(ctx, err)
		return
	}
	submissions := gentity.NewTeacherSubmissions(submissionsOut.Submissions)

	res := &response.TeacherSubmissionsResponse{
		Summaries: entity.NewTeacherSubmissions(summaries, submissions.MapByShiftSummaryID()),
	}
	ctx.JSON(http.StatusOK, res)
}
