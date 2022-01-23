package handler

import (
	"net/http"

	gentity "github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/internal/gateway/student/v1/response"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"
	"github.com/calmato/shs-web/api/internal/gateway/util"
	"github.com/gin-gonic/gin"
)

func (h *apiV1Handler) GetStudent(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	studentID := ctx.Param("studentId")

	var student *gentity.Student
	student, err := h.getStudent(c, studentID)
	if err != nil {
		httpError(ctx, err)
		return
	}

	res := &response.StudentResponse{
		Student: entity.NewStudent(student, nil),
	}
	ctx.JSON(http.StatusOK, res)
}
