package handler

import (
	"context"
	"net/http"

	gentity "github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"
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

func (h *apiV1Handler) multiGetStudents(ctx context.Context, studentIDs []string) (gentity.Students, error) {
	// TODO: 実装
	return gentity.Students{}, nil
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
