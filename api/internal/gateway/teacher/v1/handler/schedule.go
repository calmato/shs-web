package handler

import (
	"net/http"

	gentity "github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/request"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/response"
	"github.com/calmato/shs-web/api/internal/gateway/util"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/gin-gonic/gin"
)

func (h *apiV1Handler) ListSchedules(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	in := &classroom.ListSchedulesRequest{}
	out, err := h.classroom.ListSchedules(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}
	schedules := gentity.NewSchedules(out.Schedules)

	res := &response.SchedulesResponse{
		Schedules: entity.NewSchedules(schedules),
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV1Handler) UpdateSchedules(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	req := &request.UpdateSchedulesRequest{}
	if err := ctx.BindJSON(req); err != nil {
		badRequest(ctx, err)
		return
	}

	in := &classroom.UpdateSchedulesRequest{
		Schedules: entity.NewSchedulesToUpdate(req.Schedules),
	}
	_, err := h.classroom.UpdateSchedules(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}
