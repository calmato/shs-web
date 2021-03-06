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
	"github.com/gin-gonic/gin"
)

func (h *apiV1Handler) ListSubjects(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	typ, err := strconv.ParseInt(ctx.DefaultQuery("type", "0"), 10, 32)
	if err != nil {
		badRequest(ctx, err)
		return
	}
	schoolType, _ := entity.SchoolType(typ).ClassroomSchoolType()

	in := &classroom.ListSubjectsRequest{
		SchoolType: schoolType,
	}
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

func (h *apiV1Handler) CreateSubject(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	req := &request.CreateSubjectRequest{}
	if err := ctx.BindJSON(req); err != nil {
		badRequest(ctx, err)
		return
	}
	schoolType, err := req.SchoolType.ClassroomSchoolType()
	if err != nil {
		badRequest(ctx, err)
		return
	}

	in := &classroom.CreateSubjectRequest{
		Name:       req.Name,
		Color:      req.Color,
		SchoolType: schoolType,
	}
	out, err := h.classroom.CreateSubject(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}
	subject := gentity.NewSubject(out.Subject)

	res := &response.SubjectResponse{
		Subject: entity.NewSubject(subject),
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV1Handler) UpdateSubject(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	req := &request.UpdateSubjectRequest{}
	if err := ctx.BindJSON(req); err != nil {
		badRequest(ctx, err)
		return
	}
	subjectID, err := strconv.ParseInt(ctx.Param("subjectId"), 10, 64)
	if err != nil {
		badRequest(ctx, err)
		return
	}
	schoolType, err := req.SchoolType.ClassroomSchoolType()
	if err != nil {
		badRequest(ctx, err)
		return
	}

	in := &classroom.UpdateSubjectRequest{
		Id:         subjectID,
		Name:       req.Name,
		Color:      req.Color,
		SchoolType: schoolType,
	}
	_, err = h.classroom.UpdateSubject(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (h *apiV1Handler) DeleteSubject(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	subjectID, err := strconv.ParseInt(ctx.Param("subjectId"), 10, 64)
	if err != nil {
		badRequest(ctx, err)
		return
	}

	in := &classroom.DeleteSubjectRequest{Id: subjectID}
	_, err = h.classroom.DeleteSubject(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (h *apiV1Handler) multiGetTeacherSubjects(
	ctx context.Context, teacherIDs []string,
) (map[string]gentity.Subjects, error) {
	in := &classroom.MultiGetTeacherSubjectsRequest{
		TeacherIds: teacherIDs,
	}
	out, err := h.classroom.MultiGetTeacherSubjects(ctx, in)
	if err != nil {
		return nil, err
	}
	teacherSubjects := gentity.NewTeacherSubjects(out.TeacherSubjects)
	subjects := gentity.NewSubjects(out.Subjects)
	return subjects.GroupByTeacher(teacherSubjects), nil
}

func (h *apiV1Handler) getTeacherSubject(
	ctx context.Context, teacherID string,
) (gentity.Subjects, error) {
	in := &classroom.GetTeacherSubjectRequest{
		TeacherId: teacherID,
	}
	out, err := h.classroom.GetTeacherSubject(ctx, in)
	if err != nil {
		return nil, err
	}
	subjects := gentity.NewSubjects(out.Subjects)
	return subjects, nil
}

func (h *apiV1Handler) multiGetStudentSubjects(
	ctx context.Context, studentIDs []string,
) (map[string]gentity.Subjects, error) {
	in := &classroom.MultiGetStudentSubjectsRequest{
		StudentIds: studentIDs,
	}
	out, err := h.classroom.MultiGetStudentSubjects(ctx, in)
	if err != nil {
		return nil, err
	}
	studentSubjects := gentity.NewStudentSubjects(out.StudentSubjects)
	subjects := gentity.NewSubjects(out.Subjects)
	return subjects.GroupByStudent(studentSubjects), nil
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
