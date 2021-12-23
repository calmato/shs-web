package handler

import (
	"net/http"

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

func (h *apiV1Handler) GetAuth(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	teacherID := getTeacherID(ctx)

	eg, ectx := errgroup.WithContext(c)
	var teacher *gentity.Teacher
	eg.Go(func() (err error) {
		teacher, err = h.getTeacher(ectx, teacherID)
		return
	})
	var subjects gentity.Subjects
	eg.Go(func() (err error) {
		_, subjects, err = h.getTeacherSubject(ectx, teacherID)
		return
	})
	if err := eg.Wait(); err != nil {
		httpError(ctx, err)
		return
	}

	res := &response.AuthResponse{
		Auth:     entity.NewAuth(teacher),
		Subjects: entity.NewSubjects(subjects).GroupBySchoolType(),
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV1Handler) UpdateMySubjects(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	teacherID := getTeacherID(ctx)

	req := &request.UpdateMySubjectRequest{}
	if err := ctx.BindJSON(req); err != nil {
		badRequest(ctx, err)
		return
	}
	schoolType, err := entity.SchoolType(req.SchoolType).ClassroomSchoolType()
	if err != nil {
		badRequest(ctx, err)
		return
	}

	in := &classroom.UpdateTeacherSubjectRequest{
		TeacherId:  teacherID,
		SubjectIds: req.SubjectIDs,
		SchoolType: schoolType,
	}
	_, err = h.classroom.UpdateTeacherSubject(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (h *apiV1Handler) UpdateMyMail(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	teacherID := getTeacherID(ctx)

	req := &request.UpdateMyMailRequest{}
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

func (h *apiV1Handler) UpdateMyPassword(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	teacherID := getTeacherID(ctx)

	req := &request.UpdateMyPasswordRequest{}
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
