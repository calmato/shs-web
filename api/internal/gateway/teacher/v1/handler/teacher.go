package handler

import (
	"context"
	"net/http"
	"strconv"

	gentity "github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/request"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/response"
	"github.com/calmato/shs-web/api/internal/gateway/util"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/gin-gonic/gin"
)

func (h *apiV1Handler) ListTeachers(ctx *gin.Context) {
	c := util.SetMetadata(ctx)
	const (
		defaultLimit  = "30"
		defaultOffset = "0"
	)

	limit, err := strconv.ParseInt(ctx.DefaultQuery("limit", defaultLimit), 10, 64)
	if err != nil {
		badRequest(ctx, err)
		return
	}
	offset, err := strconv.ParseInt(ctx.DefaultQuery("offset", defaultOffset), 10, 64)
	if err != nil {
		badRequest(ctx, err)
		return
	}

	in := &user.ListTeachersRequest{
		Limit:  limit,
		Offset: offset,
	}
	out, err := h.user.ListTeachers(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}
	teachers := gentity.NewTeachers(out.Teachers)

	res := &response.TeachersResponse{
		Teachers: entity.NewTeachers(teachers),
		Total:    out.Total,
	}
	ctx.JSON(http.StatusOK, res)
}

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

func (h *apiV1Handler) CreateTeacher(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	req := &request.CreateTeacherRequest{}
	if err := ctx.BindJSON(req); err != nil {
		badRequest(ctx, err)
		return
	}
	role, err := entity.Role(req.Role).UserRole()
	if err != nil {
		badRequest(ctx, err)
		return
	}

	in := &user.CreateTeacherRequest{
		LastName:             req.LastName,
		FirstName:            req.FirstName,
		LastNameKana:         req.LastNameKana,
		FirstNameKana:        req.FirstNameKana,
		Mail:                 req.Mail,
		Role:                 role,
		Password:             req.Password,
		PasswordConfirmation: req.PasswordConfirmation,
	}
	out, err := h.user.CreateTeacher(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}
	teacher := gentity.NewTeacher(out.Teacher)

	res := &response.TeacherResponse{
		Teacher: entity.NewTeacher(teacher),
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
