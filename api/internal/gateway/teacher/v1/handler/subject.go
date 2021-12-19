package handler

import (
	"net/http"

	gentity "github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/response"
	"github.com/calmato/shs-web/api/internal/gateway/util"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/gin-gonic/gin"
)

func (h *apiV1Handler) ListSubjects(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	in := &classroom.ListSubjectsRequest{}
	out, err := h.classroom.ListSubjects(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}
	subjects := gentity.NewSubjects(out.Subjects)

	res := &response.SubjectsResponse{
		Subjects: entity.NewSubjects(subjects),
	}
	ctx.JSON(http.StatusOK, res)
}