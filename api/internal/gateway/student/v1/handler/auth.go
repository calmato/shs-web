package handler

import (
	"context"
	"net/http"

	gentity "github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/internal/gateway/student/v1/entity"
	"github.com/calmato/shs-web/api/internal/gateway/student/v1/response"
	"github.com/calmato/shs-web/api/internal/gateway/util"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/gin-gonic/gin"
)

func (h *apiV1Handler) GetAuth(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	studentID := getStudentID(ctx)
	student, err := h.getStudent(c, studentID)
	if err != nil {
		httpError(ctx, err)
		return
	}

	res := &response.AuthResponse{
		Auth: entity.NewAuth(student),
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV1Handler) getStudent(ctx context.Context, studentID string) (*gentity.Student, error) {
	in := &user.GetStudentRequest{Id: studentID}
	out, err := h.user.GetStudent(ctx, in)
	if err != nil {
		return nil, err
	}
	student := gentity.NewStudent(out.Student)
	return student, nil
}
