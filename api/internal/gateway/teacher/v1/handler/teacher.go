package handler

import (
	"net/http"

	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/request"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/response"
	"github.com/gin-gonic/gin"
)

// mock
func (h *apiV1Handler) CreateTeacher(ctx *gin.Context) {
	req := &request.CreateTeacherRequest{}
	if err := ctx.BindJSON(req); err != nil {
		badRequest(ctx, err)
		return
	}

	res := &response.TeacherResponse{
		Teacher: &entity.Teacher{
			ID:            "123456789012345678901",
			LastName:      "中村",
			FirstName:     "広大",
			LastNameKana:  "なかむら",
			FirstNameKana: "こうだい",
			Mail:          "teacher-test001@calmato.jp",
			Role:          int32(entity.RoleTeacher),
			CreatedAt:     h.now().String(),
			UpdatedAt:     h.now().String(),
		},
	}
	ctx.JSON(http.StatusOK, res)
}
