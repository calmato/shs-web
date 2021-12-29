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
)

func (h *apiV1Handler) ListShiftSummaries(ctx *gin.Context) {
	c := util.SetMetadata(ctx)
	const (
		defaultLimit  = "30"
		defaultOffset = "0"
		defaultStatus = "0"
	)

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
	status, err := strconv.ParseInt(ctx.DefaultQuery("status", defaultStatus), 10, 32)
	if err != nil {
		badRequest(ctx, err)
		return
	}
	shiftStatus, _ := entity.ShiftStatus(status).LessonShiftStatus()

	in := &lesson.ListShiftSummariesRequest{
		Limit:  limit,
		Offset: offset,
		Status: shiftStatus,
	}
	out, err := h.lesson.ListShiftSummaries(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}
	summaries := gentity.NewShiftSummaries(out.Summaries)

	res := &response.ShiftSummariesResponse{
		Summaries: entity.NewShiftSummaries(summaries),
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV1Handler) CreateShifts(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	req := &request.CreateShiftsRequest{}
	if err := ctx.BindJSON(req); err != nil {
		badRequest(ctx, err)
		return
	}
	yearMonth, err := strconv.ParseInt(req.YearMonth, 10, 64)
	if err != nil {
		badRequest(ctx, err)
		return
	}
	openAt, err := jst.ParseFromYYYYMMDD(req.OpenDate)
	if err != nil {
		badRequest(ctx, err)
		return
	}
	endAt, err := jst.ParseFromYYYYMMDD(req.EndDate)
	if err != nil {
		badRequest(ctx, err)
		return
	}

	in := &lesson.CreateShiftsRequest{
		YearMonth:   int32(yearMonth),
		OpenAt:      openAt.Unix(),
		EndAt:       endAt.AddDate(0, 0, 1).Unix(),
		ClosedDates: req.ClosedDates,
	}
	out, err := h.lesson.CreateShifts(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}
	gsummary := gentity.NewShiftSummary(out.Summary)
	gshifts := gentity.NewShifts(out.Shifts)

	summary := entity.NewShiftSummary(gsummary)
	shifts := entity.NewShifts(gshifts)

	shiftsMap, err := shifts.GroupByDate()
	if err != nil {
		httpError(ctx, err)
		return
	}
	res := &response.ShiftsResponse{
		Summary: summary,
		Shifts:  entity.NewShiftDetailsForMonth(summary, shiftsMap),
	}
	ctx.JSON(http.StatusOK, res)
}
