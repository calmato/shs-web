package handler

import (
	"net/http"

	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/request"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/response"
	"github.com/calmato/shs-web/api/internal/gateway/util"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/gin-gonic/gin"
)

func (h *apiV1Handler) GetRoomsTotal(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	in := &classroom.GetRoomsTotalRequest{}
	out, err := h.classroom.GetRoomsTotal(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}

	res := &response.RoomsTotalResponse{
		Total: out.Total,
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV1Handler) UpdateRoomsTotal(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	req := &request.UpdateRoomsTotalRequest{}
	if err := ctx.BindJSON(req); err != nil {
		badRequest(ctx, err)
		return
	}

	in := &classroom.UpdateRoomsTotalRequest{
		Total: req.Total,
	}
	_, err := h.classroom.UpdateRoomsTotal(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}
