package handler

import (
	"fmt"
	"net/http"
	"strconv"

	gentity "github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/request"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/response"
	"github.com/calmato/shs-web/api/internal/gateway/util"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

func (h *apiV1Handler) ListTeacherSubmissions(ctx *gin.Context) {
	c := util.SetMetadata(ctx)
	const (
		defaultLimit  = "30"
		defaultOffset = "0"
	)

	teacherID := ctx.Param("teacherId")
	if teacherID != getTeacherID(ctx) {
		err := fmt.Errorf("api: %s is not your teacher id", teacherID)
		forbidden(ctx, err)
		return
	}
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

	shiftsIn := &lesson.ListShiftSummariesRequest{
		Status:  lesson.ShiftStatus_SHIFT_STATUS_ACCEPTING,
		OrderBy: lesson.ListShiftSummariesRequest_ORDER_BY_YEAR_MONTH_DESC,
		Limit:   limit,
		Offset:  offset,
	}
	shiftsOut, err := h.lesson.ListShiftSummaries(c, shiftsIn)
	if err != nil {
		httpError(ctx, err)
		return
	}
	summaries := gentity.NewShiftSummaries(shiftsOut.Summaries)

	submissionsIn := &lesson.ListTeacherSubmissionsByShiftSummaryIDsRequest{
		TeacherId:       teacherID,
		ShiftSummaryIds: summaries.IDs(),
	}
	submissionsOut, err := h.lesson.ListTeacherSubmissionsByShiftSummaryIDs(c, submissionsIn)
	if err != nil {
		httpError(ctx, err)
		return
	}
	submissions := gentity.NewTeacherSubmissions(submissionsOut.Submissions)

	res := &response.TeacherSubmissionsResponse{
		Summaries: entity.NewTeacherSubmissions(summaries, submissions.MapByShiftSummaryID()),
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV1Handler) ListTeacherShifts(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	teacherID := ctx.Param("teacherId")
	if teacherID != getTeacherID(ctx) {
		err := fmt.Errorf("api: %s is not your teacher id", teacherID)
		forbidden(ctx, err)
		return
	}
	shiftSummaryID, err := strconv.ParseInt(ctx.Param("summaryId"), 10, 64)
	if err != nil {
		badRequest(ctx, err)
		return
	}

	eg, ectx := errgroup.WithContext(c)
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
	var submission *gentity.TeacherSubmission
	var teacherShifts gentity.TeacherShifts
	eg.Go(func() error {
		in := &lesson.GetTeacherShiftsRequest{
			TeacherId:      teacherID,
			ShiftSummaryId: shiftSummaryID,
		}
		out, err := h.lesson.GetTeacherShifts(ectx, in)
		if err != nil {
			return err
		}
		submission = gentity.NewTeacherSubmission(out.Submission)
		teacherShifts = gentity.NewTeacherShifts(out.Shifts)
		return nil
	})
	if err := eg.Wait(); err != nil {
		httpError(ctx, err)
		return
	}
	summary := entity.NewShiftSummary(gsummary)
	shifts, err := gshifts.GroupByDate()
	if err != nil {
		httpError(ctx, err)
		return
	}

	res := &response.TeacherShiftsResponse{
		Summary: entity.NewTeacherSubmission(gsummary, submission),
		Shifts:  entity.NewTeacherShiftDetailsForMonth(summary, shifts, teacherShifts.MapByShiftID()),
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV1Handler) ListEnabledTeacherShifts(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	teacherID := ctx.Param("teacherId")
	shiftSummaryID, err := strconv.ParseInt(ctx.Param("shiftId"), 10, 64)
	if err != nil {
		badRequest(ctx, err)
		return
	}

	eg, ectx := errgroup.WithContext(c)
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
	var submission *gentity.TeacherSubmission
	var teacherShifts gentity.TeacherShifts
	eg.Go(func() error {
		in := &lesson.GetTeacherShiftsRequest{
			TeacherId:      teacherID,
			ShiftSummaryId: shiftSummaryID,
		}
		out, err := h.lesson.GetTeacherShifts(ectx, in)
		if err != nil {
			return err
		}
		submission = gentity.NewTeacherSubmission(out.Submission)
		teacherShifts = gentity.NewTeacherShifts(out.Shifts)
		return nil
	})
	if err := eg.Wait(); err != nil {
		httpError(ctx, err)
		return
	}
	summary := entity.NewShiftSummary(gsummary)
	shifts, err := gshifts.GroupByDate()
	if err != nil {
		httpError(ctx, err)
		return
	}

	res := &response.TeacherShiftsResponse{
		Summary: entity.NewTeacherSubmission(gsummary, submission),
		Shifts:  entity.NewEnabledTeacherShiftDetails(summary, shifts, teacherShifts.MapByShiftID()),
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV1Handler) UpsertTeacherShifts(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	teacherID := ctx.Param("teacherId")
	if teacherID != getTeacherID(ctx) {
		err := fmt.Errorf("api: %s is not your teacher id", teacherID)
		forbidden(ctx, err)
		return
	}
	req := &request.UpsertTeacherShiftsRequest{}
	if err := ctx.BindJSON(req); err != nil {
		badRequest(ctx, err)
		return
	}
	shiftSummaryID, err := strconv.ParseInt(ctx.Param("summaryId"), 10, 64)
	if err != nil {
		badRequest(ctx, err)
		return
	}

	eg, ectx := errgroup.WithContext(c)
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
	if err := eg.Wait(); err != nil {
		httpError(ctx, err)
		return
	}
	summary := entity.NewShiftSummary(gsummary)
	shifts, err := gshifts.GroupByDate()
	if err != nil {
		httpError(ctx, err)
		return
	}

	in := &lesson.UpsertTeacherShiftsRequest{
		TeacherId:      teacherID,
		ShiftSummaryId: shiftSummaryID,
		ShiftIds:       req.ShiftIDs,
		Decided:        true,
	}
	out, err := h.lesson.UpsertTeacherShifts(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}
	submission := gentity.NewTeacherSubmission(out.Submission)
	teacherShifts := gentity.NewTeacherShifts(out.Shifts)

	res := &response.TeacherShiftsResponse{
		Summary: entity.NewTeacherSubmission(gsummary, submission),
		Shifts:  entity.NewTeacherShiftDetailsForMonth(summary, shifts, teacherShifts.MapByShiftID()),
	}
	ctx.JSON(http.StatusOK, res)
}
