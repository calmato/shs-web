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
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
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
	subjects, err := h.multiGetTeacherSubjects(c, teachers.IDs())
	if err != nil {
		httpError(ctx, err)
		return
	}

	res := &response.TeachersResponse{
		Teachers: entity.NewTeachers(teachers, subjects),
		Total:    out.Total,
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV1Handler) GetTeacher(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	teacherID := ctx.Param("teacherId")

	eg, ectx := errgroup.WithContext(c)
	var teacher *gentity.Teacher
	eg.Go(func() (err error) {
		teacher, err = h.getTeacher(c, teacherID)
		return
	})
	var subjects gentity.Subjects
	eg.Go(func() (err error) {
		subjects, err = h.getTeacherSubject(ectx, teacherID)
		return
	})
	if err := eg.Wait(); err != nil {
		httpError(ctx, err)
		return
	}

	res := &response.TeacherResponse{
		Teacher: entity.NewTeacher(teacher, subjects),
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
	role, err := req.Role.UserRole()
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
		Teacher: entity.NewTeacher(teacher, nil),
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV1Handler) UpdateTeacherMail(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	teacherID := ctx.Param("teacherId")
	req := &request.UpdateTeacherMailRequest{}
	if err := ctx.BindJSON(req); err != nil {
		badRequest(ctx, err)
		return
	}

	in := &user.UpdateTeacherMailRequest{
		Id:   teacherID,
		Mail: req.Mail,
	}
	_, err := h.user.UpdateTeacherMail(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (h *apiV1Handler) UpdateTeacherPassword(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	teacherID := ctx.Param("teacherId")
	req := &request.UpdateTeacherPasswordRequest{}
	if err := ctx.BindJSON(req); err != nil {
		badRequest(ctx, err)
		return
	}

	in := &user.UpdateTeacherPasswordRequest{
		Id:                   teacherID,
		Password:             req.Password,
		PasswordConfirmation: req.PasswordConfirmation,
	}
	_, err := h.user.UpdateTeacherPassword(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (h *apiV1Handler) UpdateTeacherRole(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	teacherID := ctx.Param("teacherId")
	req := &request.UpdateTeacherRoleRequest{}
	if err := ctx.BindJSON(req); err != nil {
		badRequest(ctx, err)
		return
	}

	role, err := req.Role.UserRole()
	if err != nil {
		badRequest(ctx, err)
		return
	}

	in := &user.UpdateTeacherRoleRequest{
		Id:   teacherID,
		Role: role,
	}
	_, err = h.user.UpdateTeacherRole(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (h *apiV1Handler) UpdateTeacherSubject(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	teacherID := ctx.Param("teacherId")

	req := &request.UpdateTeacherSubjectRequest{}
	if err := ctx.BindJSON(req); err != nil {
		badRequest(ctx, err)
		return
	}
	schoolType, err := req.SchoolType.ClassroomSchoolType()
	if err != nil {
		badRequest(ctx, err)
		return
	}

	in := &classroom.UpsertTeacherSubjectRequest{
		TeacherId:  teacherID,
		SubjectIds: req.SubjectIDs,
		SchoolType: schoolType,
	}
	_, err = h.classroom.UpsertTeacherSubject(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (h *apiV1Handler) DeleteTeacher(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	teacherID := ctx.Param("teacherId")
	in := &user.DeleteTeacherRequest{
		Id: teacherID,
	}
	_, err := h.user.DeleteTeacher(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
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
