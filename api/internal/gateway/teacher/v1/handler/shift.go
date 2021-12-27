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
	summary := gentity.NewShiftSummary(out.Summary)
	shifts := gentity.NewShifts(out.Shifts)

	res := &response.ShiftsResponse{
		Summary: entity.NewShiftSummary(summary),
		Shifts:  entity.NewShifts(shifts).GroupByDate(),
	}
	ctx.JSON(http.StatusOK, res)
}
