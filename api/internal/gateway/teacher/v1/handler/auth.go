package handler

import (
	"net/http"

	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/response"
	"github.com/calmato/shs-web/api/internal/gateway/util"
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
