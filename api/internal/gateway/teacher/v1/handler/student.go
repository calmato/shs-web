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

func (h *apiV1Handler) ListStudents(ctx *gin.Context) {
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

	in := &user.ListStudentsRequest{
		Limit:  limit,
		Offset: offset,
	}

	out, err := h.user.ListStudents(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}
	students := gentity.NewStudents(out.Students)
	subjects, err := h.multiGetStudentSubjects(c, students.IDs())
	if err != nil {
		httpError(ctx, err)
		return
	}

	res := &response.StudentsResponse{
		Students: entity.NewStudents(students, subjects),
		Total:    out.Total,
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV1Handler) GetStudent(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	studentID := ctx.Param("studentId")

	eg, ectx := errgroup.WithContext(c)
	var student *gentity.Student
	eg.Go(func() (err error) {
		student, err = h.getStudent(c, studentID)
		return
	})
	var subjects gentity.Subjects
	eg.Go(func() (err error) {
		subjects, err = h.getStudentSubject(ectx, studentID)
		return
	})
	if err := eg.Wait(); err != nil {
		httpError(ctx, err)
		return
	}

	res := &response.StudentResponse{
		Student: entity.NewStudent(student, subjects),
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV1Handler) CreateStudent(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	req := &request.CreateStudentRequest{}
	if err := ctx.BindJSON(req); err != nil {
		badRequest(ctx, err)
		return
	}
	schoolType, err := req.SchoolType.UserSchoolType()
	if err != nil {
		badRequest(ctx, err)
		return
	}

	in := &user.CreateStudentRequest{
		LastName:             req.LastName,
		FirstName:            req.FirstName,
		LastNameKana:         req.LastNameKana,
		FirstNameKana:        req.FirstNameKana,
		Mail:                 req.Mail,
		Password:             req.Password,
		PasswordConfirmation: req.PasswordConfirmation,
		SchoolType:           schoolType,
		Grade:                req.Grade,
	}
	out, err := h.user.CreateStudent(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}
	student := gentity.NewStudent(out.Student)

	res := &response.StudentResponse{
		Student: entity.NewStudent(student, nil),
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV1Handler) UpdateStudentMail(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	studentID := ctx.Param("studentId")
	req := &request.UpdateStudentMailRequest{}
	if err := ctx.BindJSON(req); err != nil {
		badRequest(ctx, err)
		return
	}

	in := &user.UpdateStudentMailRequest{
		Id:   studentID,
		Mail: req.Mail,
	}
	_, err := h.user.UpdateStudentMail(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (h *apiV1Handler) UpdateStudentPassword(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	studentID := ctx.Param("studentId")
	req := &request.UpdateStudentPasswordRequest{}
	if err := ctx.BindJSON(req); err != nil {
		badRequest(ctx, err)
		return
	}

	in := &user.UpdateStudentPasswordRequest{
		Id:                   studentID,
		Password:             req.Password,
		PasswordConfirmation: req.PasswordConfirmation,
	}
	_, err := h.user.UpdateStudentPassword(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (h *apiV1Handler) UpdateStudentSubject(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	studentID := ctx.Param("studentId")

	req := &request.UpdateStudentSubjectRequest{}
	if err := ctx.BindJSON(req); err != nil {
		badRequest(ctx, err)
		return
	}
	schoolType, err := req.SchoolType.ClassroomSchoolType()
	if err != nil {
		badRequest(ctx, err)
		return
	}

	in := &classroom.UpsertStudentSubjectRequest{
		StudentId:  studentID,
		SubjectIds: req.SubjectIDs,
		SchoolType: schoolType,
	}
	_, err = h.classroom.UpsertStudentSubject(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (h *apiV1Handler) multiGetStudents(ctx context.Context, studentIDs []string) (gentity.Students, error) {
	in := &user.MultiGetStudentsRequest{
		Ids: studentIDs,
	}
	out, err := h.user.MultiGetStudents(ctx, in)
	if err != nil {
		return nil, err
	}
	return gentity.NewStudents(out.Students), nil
}

func (h *apiV1Handler) getStudent(ctx context.Context, studentID string) (*gentity.Student, error) {
	in := &user.GetStudentRequest{
		Id: studentID,
	}
	out, err := h.user.GetStudent(ctx, in)
	if err != nil {
		return nil, err
	}
	return gentity.NewStudent(out.Student), nil
}

func (h *apiV1Handler) DeleteStudent(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	studentID := ctx.Param("studentId")
	in := &user.DeleteStudentRequest{
		Id: studentID,
	}

	_, err := h.user.DeleteStudent(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}
