package handler

import (
	"context"
	"net/http"

	gentity "github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/internal/gateway/student/v1/entity"
	"github.com/calmato/shs-web/api/internal/gateway/student/v1/response"
	"github.com/calmato/shs-web/api/internal/gateway/util"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/gin-gonic/gin"
)

func (h *apiV1Handler) ListLessons(ctx *gin.Context) {
	c := util.SetMetadata(ctx)
	now := h.now()
	var (
		defaultSince = jst.FormatYYYYMMDD(jst.BeginningOfMonth(now.Year(), int(now.Month())))
		defaultUntil = jst.FormatYYYYMMDD(jst.EndOfMonth(now.Year(), int(now.Month())))
	)

	in := &lesson.ListLessonsByDurationRequest{
		StudentId: getStudentID(ctx),
		Since:     ctx.DefaultQuery("since", defaultSince),
		Until:     ctx.DefaultQuery("until", defaultUntil),
	}
	out, err := h.lesson.ListLessonsByDuration(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}
	glessons := gentity.NewLessons(out.Lessons)
	gShifts := gentity.NewShifts(out.Shifts)

	gteachers, err := h.multiGetTeachers(c, glessons.TeacherIDs())
	if err != nil {
		httpError(ctx, err)
		return
	}

	lessons, err := entity.NewLessons(glessons, gShifts.Map())
	if err != nil {
		httpError(ctx, err)
		return
	}
	res := &response.LessonsResponse{
		Lessons:  lessons,
		Teachers: entity.NewTeachers(gteachers),
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV1Handler) multiGetTeachers(ctx context.Context, teacherIDs []string) (gentity.Teachers, error) {
	in := &user.MultiGetTeachersRequest{
		Ids: teacherIDs,
	}
	out, err := h.user.MultiGetTeachers(ctx, in)
	if err != nil {
		return nil, err
	}
	return gentity.NewTeachers(out.Teachers), nil
}
