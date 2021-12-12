package handler

import (
	"net/http"

	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/response"
	"github.com/gin-gonic/gin"
)

// mock
func (h *apiV1Handler) GetAuth(ctx *gin.Context) {
	teacherID := getTeacherID(ctx)

	// TODO: get teacher

	res := &response.AuthResponse{
		Auth: &entity.Auth{
			ID:            teacherID,
			LastName:      "中村",
			FirstName:     "広大",
			LastNameKana:  "なかむら",
			FirstNameKana: "こうだい",
			Mail:          "teacher-test001@calmato.jp",
			Role:          entity.Role(entity.RoleTeacher),
		},
	}
	ctx.JSON(http.StatusOK, res)
}
