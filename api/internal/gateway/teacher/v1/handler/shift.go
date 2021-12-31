package handler

import (
	"context"
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
		Limit:   limit,
		Offset:  offset,
		Status:  shiftStatus,
		OrderBy: lesson.ListShiftSummariesRequest_ORDER_BY_YEAR_MONTH_DESC,
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

func (h *apiV1Handler) UpdateShiftSummarySchedule(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	req := &request.UpdateShiftSummaryScheduleRequest{}
	if err := ctx.BindJSON(req); err != nil {
		badRequest(ctx, err)
		return
	}
	shiftSummaryID, err := strconv.ParseInt(ctx.Param("shiftId"), 10, 64)
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

	in := &lesson.UpdateShiftSummaryScheduleRequest{
		Id:     shiftSummaryID,
		OpenAt: openAt.Unix(),
		EndAt:  endAt.AddDate(0, 0, 1).Unix() - 1,
	}
	_, err = h.lesson.UpdateShiftSummarySchedule(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (h *apiV1Handler) DeleteShiftSummary(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	shiftSummaryID, err := strconv.ParseInt(ctx.Param("shiftId"), 10, 64)
	if err != nil {
		badRequest(ctx, err)
		return
	}

	in := &lesson.DeleteShiftSummaryRequest{
		Id: shiftSummaryID,
	}
	_, err = h.lesson.DeleteShiftSummary(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (h *apiV1Handler) ListShifts(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	shiftSummaryID, err := strconv.ParseInt(ctx.Param("shiftId"), 10, 64)
	if err != nil {
		badRequest(ctx, err)
		return
	}

	eg, ectx := errgroup.WithContext(c)
	var summary *entity.ShiftSummary
	eg.Go(func() error {
		gsummary, err := h.getShiftSummary(ectx, shiftSummaryID)
		if err != nil {
			return err
		}
		summary = entity.NewShiftSummary(gsummary)
		return nil
	})
	var gshifts gentity.Shifts
	eg.Go(func() (err error) {
		gshifts, err = h.listShiftsBySummaryID(ectx, shiftSummaryID)
		return
	})
	if err := eg.Wait(); err != nil {
		httpError(ctx, err)
		return
	}

	shiftsMap, err := gshifts.GroupByDate()
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
		EndAt:       endAt.AddDate(0, 0, 1).Unix() - 1,
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
	shiftsMap, err := gshifts.GroupByDate()
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

func (h *apiV1Handler) getShiftSummary(ctx context.Context, summaryID int64) (*gentity.ShiftSummary, error) {
	in := &lesson.GetShiftSummaryRequest{
		Id: summaryID,
	}
	out, err := h.lesson.GetShiftSummary(ctx, in)
	if err != nil {
		return nil, err
	}
	summary := gentity.NewShiftSummary(out.Summary)
	return summary, nil
}

func (h *apiV1Handler) listShiftsBySummaryID(ctx context.Context, summaryID int64) (gentity.Shifts, error) {
	in := &lesson.ListShiftsRequest{
		ShiftSummaryId: summaryID,
	}
	out, err := h.lesson.ListShifts(ctx, in)
	if err != nil {
		return nil, err
	}
	shifts := gentity.NewShifts(out.Shifts)
	return shifts, nil
}
