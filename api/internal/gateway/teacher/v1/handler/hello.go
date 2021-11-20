package handler

import (
	"net/http"

	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/request"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/response"
	"github.com/calmato/shs-web/api/internal/gateway/util"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/gin-gonic/gin"
)

func (h *apiV1Handler) Hello(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	req := &request.HelloRequest{}
	if err := ctx.BindJSON(req); err != nil {
		badRequest(ctx, err)
		return
	}

	in := &user.HelloRequest{Name: req.Name}
	out, err := h.user.Hello(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}

	res := &response.HelloResponse{Message: out.Message}
	ctx.JSON(http.StatusOK, res)
}
