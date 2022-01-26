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
	"golang.org/x/sync/errgroup"
)

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
