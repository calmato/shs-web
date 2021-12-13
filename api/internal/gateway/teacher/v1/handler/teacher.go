package handler

import (
	"context"
	"net/http"

	gentity "github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/request"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/response"
	"github.com/calmato/shs-web/api/internal/gateway/util"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/gin-gonic/gin"
)

func (h *apiV1Handler) GetTeacher(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	teacherID := ctx.Param("teacherId")
	teacher, err := h.getTeacher(c, teacherID)
	if err != nil {
		httpError(ctx, err)
		return
	}

	res := &response.TeacherResponse{
		Teacher: entity.NewTeacher(teacher),
	}
	ctx.JSON(http.StatusOK, res)
}

// mock
func (h *apiV1Handler) CreateTeacher(ctx *gin.Context) {
	req := &request.CreateTeacherRequest{}
	if err := ctx.BindJSON(req); err != nil {
		badRequest(ctx, err)
		return
	}

	// TODO: get teacher

	res := &response.TeacherResponse{
		Teacher: &entity.Teacher{
			ID:            "123456789012345678901",
			LastName:      "中村",
			FirstName:     "広大",
			LastNameKana:  "なかむら",
			FirstNameKana: "こうだい",
			Mail:          "teacher-test001@calmato.jp",
			Role:          entity.RoleTeacher,
			CreatedAt:     h.now(),
			UpdatedAt:     h.now(),
		},
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
