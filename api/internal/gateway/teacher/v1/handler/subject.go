package handler

import (
	"context"
	"net/http"

	gentity "github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/response"
	"github.com/calmato/shs-web/api/internal/gateway/util"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/gin-gonic/gin"
)

func (h *apiV1Handler) ListSubjects(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	in := &classroom.ListSubjectsRequest{}
	out, err := h.classroom.ListSubjects(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}
	subjects := gentity.NewSubjects(out.Subjects)

	res := &response.SubjectsResponse{
		Subjects: entity.NewSubjects(subjects),
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV1Handler) multiGetSubjects(ctx context.Context, subjectIDs []int64) (gentity.Subjects, error) {
	in := &classroom.MultiGetSubjectsRequest{
		Ids: subjectIDs,
	}
	out, err := h.classroom.MultiGetSubjects(ctx, in)
	if err != nil {
		return nil, err
	}
	subjects := gentity.NewSubjects(out.Subjects)
	return subjects, nil
}

func (h *apiV1Handler) getTeacherSubject(
	ctx context.Context, teacherID string,
) (*gentity.TeacherSubject, gentity.Subjects, error) {
	in := &classroom.GetTeacherSubjectRequest{
		TeacherId: teacherID,
	}
	out, err := h.classroom.GetTeacherSubject(ctx, in)
	if err != nil {
		return nil, nil, err
	}
	teacherSubject := gentity.NewTeacherSubject(out.TeacherSubject)
	subjects := gentity.NewSubjects(out.Subjects)
	return teacherSubject, subjects, nil
}
