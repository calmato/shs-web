package handler

import (
	"net/http"
	"strconv"

	gentity "github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/internal/gateway/student/v1/entity"
	"github.com/calmato/shs-web/api/internal/gateway/student/v1/response"
	"github.com/calmato/shs-web/api/internal/gateway/util"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/gin-gonic/gin"
)

func (h *apiV1Handler) ListSubjects(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	typ, err := strconv.ParseInt(ctx.DefaultQuery("type", "0"), 10, 32)
	if err != nil {
		badRequest(ctx, err)
		return
	}
	schoolType, _ := entity.SchoolType(typ).ClassroomSchoolType()

	in := &classroom.ListSubjectsRequest{
		SchoolType: schoolType,
	}
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
