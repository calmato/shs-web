package handler

import (
	"context"
	"net/http"

	gentity "github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/internal/gateway/student/v1/entity"
	"github.com/calmato/shs-web/api/internal/gateway/student/v1/request"
	"github.com/calmato/shs-web/api/internal/gateway/student/v1/response"
	"github.com/calmato/shs-web/api/internal/gateway/util"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

func (h *apiV1Handler) GetAuth(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	studentID := getStudentID(ctx)

	eg, ectx := errgroup.WithContext(c)
	var student *gentity.Student
	eg.Go(func() (err error) {
		student, err = h.getStudent(ectx, studentID)
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

	schoolType := entity.NewSchoolTypeFromUser(student.SchoolType)
	res := &response.AuthResponse{
		Auth:     entity.NewAuth(student),
		Subjects: entity.NewSubjects(subjects).FiterBySchoolType(schoolType),
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV1Handler) UpdateMyMail(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	studentID := getStudentID(ctx)
	req := &request.UpdateMyMailRequest{}
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

func (h *apiV1Handler) UpdateMyPassword(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	studentID := getStudentID(ctx)
	req := &request.UpdateMyPasswordRequest{}
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

func (h *apiV1Handler) getStudent(ctx context.Context, studentID string) (*gentity.Student, error) {
	in := &user.GetStudentRequest{Id: studentID}
	out, err := h.user.GetStudent(ctx, in)
	if err != nil {
		return nil, err
	}
	student := gentity.NewStudent(out.Student)
	return student, nil
}

func (h *apiV1Handler) getStudentSubject(
	ctx context.Context, studentID string,
) (gentity.Subjects, error) {
	in := &classroom.GetStudentSubjectRequest{
		StudentId: studentID,
	}
	out, err := h.classroom.GetStudentSubject(ctx, in)
	if err != nil {
		return nil, err
	}
	subjects := gentity.NewSubjects(out.Subjects)
	return subjects, nil
}
