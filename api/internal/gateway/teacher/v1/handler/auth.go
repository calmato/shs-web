package handler

import (
	"context"
	"net/http"

	gentity "github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/response"
	"github.com/calmato/shs-web/api/internal/gateway/util"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/gin-gonic/gin"
)

func (h *apiV1Handler) GetAuth(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	teacherID := getTeacherID(ctx)
	teacher, err := h.getTeacher(c, teacherID)
	if err != nil {
		httpError(ctx, err)
		return
	}

	res := &response.AuthResponse{
		Auth: entity.NewAuth(teacher),
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV1Handler) getTeacher(ctx context.Context, teacherID string) (*gentity.Teacher, error) {
	in := &user.GetTeacherRequest{
		Id: teacherID,
	}
	out, err := h.user.GetTeacher(ctx, in)
	if err != nil {
		return nil, err
	}
	return gentity.NewTeacher(out.Teacher), nil
}
