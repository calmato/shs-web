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
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

func (h *apiV1Handler) ListShiftSummaries(ctx *gin.Context) {
	c := util.SetMetadata(ctx)
	const (
		defaultLimit  = "30"
		defaultOffset = "0"
		defaultStatus = "0"
	)

	limit, err := strconv.ParseInt(ctx.DefaultQuery("limit", defaultLimit), 10, 64)
	if err != nil {
		badRequest(ctx, err)
		return
	}
	offset, err := strconv.ParseInt(ctx.DefaultQuery("offset", defaultOffset), 10, 64)
	if err != nil {
		badRequest(ctx, err)
		return
	}
	status, err := strconv.ParseInt(ctx.DefaultQuery("status", defaultStatus), 10, 32)
	if err != nil {
		badRequest(ctx, err)
		return
	}
	shiftStatus, _ := entity.ShiftStatus(status).LessonShiftStatus()

	in := &lesson.ListShiftSummariesRequest{
		Limit:   limit,
		Offset:  offset,
		Status:  shiftStatus,
		OrderBy: lesson.ListShiftSummariesRequest_ORDER_BY_YEAR_MONTH_DESC,
	}
	out, err := h.lesson.ListShiftSummaries(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}
	summaries := gentity.NewShiftSummaries(out.Summaries)

	res := &response.ShiftSummariesResponse{
		Summaries: entity.NewShiftSummaries(summaries),
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV1Handler) UpdateShiftSummarySchedule(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	req := &request.UpdateShiftSummaryScheduleRequest{}
	if err := ctx.BindJSON(req); err != nil {
		badRequest(ctx, err)
		return
	}
	shiftSummaryID, err := strconv.ParseInt(ctx.Param("shiftId"), 10, 64)
	if err != nil {
		badRequest(ctx, err)
		return
	}
	openAt, err := jst.ParseFromYYYYMMDD(req.OpenDate)
	if err != nil {
		badRequest(ctx, err)
		return
	}
	endAt, err := jst.ParseFromYYYYMMDD(req.EndDate)
	if err != nil {
		badRequest(ctx, err)
		return
	}

	in := &lesson.UpdateShiftSummaryScheduleRequest{
		Id:     shiftSummaryID,
		OpenAt: jst.BeginningOfDay(openAt).Unix(),
		EndAt:  jst.EndOfDay(endAt).Unix(),
	}
	_, err = h.lesson.UpdateShiftSummarySchedule(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (h *apiV1Handler) DeleteShiftSummary(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	shiftSummaryID, err := strconv.ParseInt(ctx.Param("shiftId"), 10, 64)
	if err != nil {
		badRequest(ctx, err)
		return
	}

	in := &lesson.DeleteShiftSummaryRequest{
		Id: shiftSummaryID,
	}
	_, err = h.lesson.DeleteShiftSummary(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (h *apiV1Handler) ListShifts(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	shiftSummaryID, err := strconv.ParseInt(ctx.Param("shiftId"), 10, 64)
	if err != nil {
		badRequest(ctx, err)
		return
	}

	eg, ectx := errgroup.WithContext(c)
	var gteachers gentity.Teachers
	eg.Go(func() error {
		in := &user.ListTeachersRequest{
			Limit:  0,
			Offset: 0,
		}
		out, err := h.user.ListTeachers(ectx, in)
		if err != nil {
			return err
		}
		gteachers = gentity.NewTeachers(out.Teachers)
		return nil
	})
	var gstudents gentity.Students
	eg.Go(func() error {
		in := &user.ListStudentRequest{
			Limit:  0,
			Offset: 0,
		}
		out, err := h.user.ListStudents(ectx, in)
		if err != nil {
			return err
		}
		gstudents = gentity.NewStudents(out.Students)
		return nil
	})
	if err := eg.Wait(); err != nil {
		httpError(ctx, err)
		return
	}

	eg, ectx = errgroup.WithContext(c)
	var rooms int64
	eg.Go(func() error {
		in := &classroom.GetRoomsTotalRequest{}
		out, err := h.classroom.GetRoomsTotal(ectx, in)
		if err != nil {
			return err
		}
		rooms = out.Total
		return nil
	})
	var gsummary *gentity.ShiftSummary
	eg.Go(func() (err error) {
		gsummary, err = h.getShiftSummary(ectx, shiftSummaryID)
		return
	})
	var gshifts gentity.Shifts
	eg.Go(func() (err error) {
		gshifts, err = h.listShiftsBySummaryID(ectx, shiftSummaryID)
		return
	})
	var gteacherShifts gentity.TeacherShifts
	eg.Go(func() error {
		in := &lesson.ListTeacherShiftsRequest{
			TeacherIds:     gteachers.IDs(),
			ShiftSummaryId: shiftSummaryID,
		}
		out, err := h.lesson.ListTeacherShifts(ectx, in)
		if err != nil {
			return err
		}
		gteacherShifts = gentity.NewTeacherShifts(out.Shifts)
		return nil
	})
	var gstudentSubmissions gentity.StudentSubmissions
	eg.Go((func() error {
		in := &lesson.ListStudentSubmissionsByStudentIDsRequest{
			StudentIds:     gstudents.IDs(),
			ShiftSummaryId: shiftSummaryID,
		}
		out, err := h.lesson.ListStudentSubmissionsByStudentIDs(ectx, in)
		if err != nil {
			return err
		}
		gstudentSubmissions = gentity.NewStudentSubmissions(out.Submissions)
		return nil
	}))
	var gstudentShifts gentity.StudentShifts
	eg.Go(func() error {
		in := &lesson.ListStudentShiftsRequest{
			StudentIds:     gstudents.IDs(),
			ShiftSummaryId: shiftSummaryID,
		}
		out, err := h.lesson.ListStudentShifts(ectx, in)
		if err != nil {
			return err
		}
		gstudentShifts = gentity.NewStudentShifts(out.Shifts)
		return nil
	})
	if err := eg.Wait(); err != nil {
		httpError(ctx, err)
		return
	}

	shiftsMap, err := gshifts.GroupByDate()
	if err != nil {
		httpError(ctx, err)
		return
	}
	studentSubmissionMap := gstudentSubmissions.MapByStudentID()
	teacherShiftsMap := gteacherShifts.GroupByTeacherID()
	studentShiftsMap := gstudentShifts.GroupByStudentID()
	summary := entity.NewShiftSummary(gsummary)
	res := &response.ShiftsResponse{
		Summary:  summary,
		Shifts:   entity.NewShiftDetailsForMonth(summary, shiftsMap),
		Rooms:    rooms,
		Teachers: entity.NewTeacherSubmissionDetails(gteachers, teacherShiftsMap),
		Students: entity.NewStudentSubmissionDetails(gstudents, studentSubmissionMap, studentShiftsMap),
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV1Handler) CreateShifts(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	req := &request.CreateShiftsRequest{}
	if err := ctx.BindJSON(req); err != nil {
		badRequest(ctx, err)
		return
	}
	yearMonth, err := strconv.ParseInt(req.YearMonth, 10, 64)
	if err != nil {
		badRequest(ctx, err)
		return
	}
	openAt, err := jst.ParseFromYYYYMMDD(req.OpenDate)
	if err != nil {
		badRequest(ctx, err)
		return
	}
	endAt, err := jst.ParseFromYYYYMMDD(req.EndDate)
	if err != nil {
		badRequest(ctx, err)
		return
	}

	in := &lesson.CreateShiftsRequest{
		YearMonth:   int32(yearMonth),
		OpenAt:      jst.BeginningOfDay(openAt).Unix(),
		EndAt:       jst.EndOfDay(endAt).Unix(),
		ClosedDates: req.ClosedDates,
	}
	out, err := h.lesson.CreateShifts(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}
	gsummary := gentity.NewShiftSummary(out.Summary)
	gshifts := gentity.NewShifts(out.Shifts)

	summary := entity.NewShiftSummary(gsummary)
	shiftsMap, err := gshifts.GroupByDate()
	if err != nil {
		httpError(ctx, err)
		return
	}
	res := &response.ShiftsResponse{
		Summary: summary,
		Shifts:  entity.NewShiftDetailsForMonth(summary, shiftsMap),
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV1Handler) getShiftSummary(ctx context.Context, summaryID int64) (*gentity.ShiftSummary, error) {
	in := &lesson.GetShiftSummaryRequest{
		Id: summaryID,
	}
	out, err := h.lesson.GetShiftSummary(ctx, in)
	if err != nil {
		return nil, err
	}
	summary := gentity.NewShiftSummary(out.Summary)
	return summary, nil
}

func (h *apiV1Handler) listShiftsBySummaryID(ctx context.Context, summaryID int64) (gentity.Shifts, error) {
	in := &lesson.ListShiftsRequest{
		ShiftSummaryId: summaryID,
	}
	out, err := h.lesson.ListShifts(ctx, in)
	if err != nil {
		return nil, err
	}
	shifts := gentity.NewShifts(out.Shifts)
	return shifts, nil
}
