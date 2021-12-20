package handler

import (
	"net/http"

	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/request"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/response"
	"github.com/calmato/shs-web/api/internal/gateway/util"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *apiV1Handler) GetAuth(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	teacherID := getTeacherID(ctx)
	teacher, err := h.getTeacher(c, teacherID)
	if err != nil {
		httpError(ctx, err)
		return
	}

	res := &response.AuthResponse{
		Auth: entity.NewAuth(teacher),
	}
	ctx.JSON(http.StatusOK, res)
}

// mock
func (h *apiV1Handler) UpdateMySubjects(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	req := &request.UpdateMySubjectRequest{}
	if err := ctx.BindJSON(req); err != nil {
		badRequest(ctx, err)
		return
	}
	_, err := entity.SchoolType(req.SchoolType).ClassroomSchoolType()
	if err != nil {
		badRequest(ctx, err)
		return
	}

	subjects, err := h.multiGetSubjects(c, req.SubjectIDs)
	if err != nil {
		httpError(ctx, err)
		return
	}
	if len(req.SubjectIDs) != len(subjects) {
		err := status.Error(codes.InvalidArgument, "handler: unmatch subjects length")
		httpError(ctx, err)
		return
	}

	// TODO: update my subjects

	res := &response.AuthResponse{}
	ctx.JSON(http.StatusOK, res)
}
