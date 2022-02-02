package handler

import (
	"net/http"

	gentity "github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/request"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/response"
	"github.com/calmato/shs-web/api/internal/gateway/util"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/gin-gonic/gin"
)

func (h *apiV1Handler) ListSchedules(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	in := &classroom.ListSchedulesRequest{}
	out, err := h.classroom.ListSchedules(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}
	schedules := gentity.NewSchedules(out.Schedules)

	res := &response.SchedulesResponse{
		Schedules: entity.NewSchedules(schedules),
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV1Handler) UpdateSchedules(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	req := &request.UpdateSchedulesRequest{}
	if err := ctx.BindJSON(req); err != nil {
		badRequest(ctx, err)
		return
	}

	in := &classroom.UpdateSchedulesRequest{
		Schedules: h.newSchedulesToUpdate(req.Schedules),
	}
	_, err := h.classroom.UpdateSchedules(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (h *apiV1Handler) newScheduleLessonsToUpdate(
	lessons []*request.ScheduleLesson, isClosed bool,
) []*classroom.ScheduleToUpdate_Lesson {
	if isClosed {
		return nil
	}
	res := make([]*classroom.ScheduleToUpdate_Lesson, len(lessons))
	for i := range lessons {
		res[i] = &classroom.ScheduleToUpdate_Lesson{
			StartTime: lessons[i].StartTime,
			EndTime:   lessons[i].EndTime,
		}
	}
	return res
}

func (h *apiV1Handler) newSchedulesToUpdate(schedules []*request.ScheduleToUpdate) []*classroom.ScheduleToUpdate {
	res := make([]*classroom.ScheduleToUpdate, len(schedules))
	for i := range schedules {
		res[i] = &classroom.ScheduleToUpdate{
			Weekday:  int32(schedules[i].Weekday),
			IsClosed: schedules[i].IsClosed,
			Lessons:  h.newScheduleLessonsToUpdate(schedules[i].Lessons, schedules[i].IsClosed),
		}
	}
	return res
}
