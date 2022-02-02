package handler

import (
	"context"
	"net/http"
	"strconv"

	gentity "github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/internal/gateway/student/v1/entity"
	"github.com/calmato/shs-web/api/internal/gateway/student/v1/request"
	"github.com/calmato/shs-web/api/internal/gateway/student/v1/response"
	"github.com/calmato/shs-web/api/internal/gateway/util"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

func (h *apiV1Handler) ListSubmissions(ctx *gin.Context) {
	c := util.SetMetadata(ctx)
	const (
		defaultLimit  = "30"
		defaultOffset = "0"
	)

	studentID := getStudentID(ctx)
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

	submissionsIn := &lesson.ListStudentSubmissionsByShiftSummaryIDsRequest{
		StudentId:       studentID,
		ShiftSummaryIds: summaries.IDs(),
	}
	submissionsOut, err := h.lesson.ListStudentSubmissionsByShiftSummaryIDs(c, submissionsIn)
	if err != nil {
		httpError(ctx, err)
		return
	}
	submissions := gentity.NewStudentSubmissions(submissionsOut.Submissions)

	res := &response.SubmissionsResponse{
		Summaries: entity.NewStudentSubmissions(summaries, submissions.MapByShiftSummaryID()),
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV1Handler) GetSubmission(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	studentID := getStudentID(ctx)
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
	var gsubmission *gentity.StudentSubmission
	var gstudentShifts gentity.StudentShifts
	eg.Go(func() error {
		in := &lesson.GetStudentShiftsRequest{
			StudentId:      studentID,
			ShiftSummaryId: shiftSummaryID,
		}
		out, err := h.lesson.GetStudentShifts(ectx, in)
		if err != nil {
			return err
		}
		gsubmission = gentity.NewStudentSubmission(out.Submission)
		gstudentShifts = gentity.NewStudentShifts(out.Shifts)
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

	res := &response.SubmissionResponse{
		Summary:          entity.NewStudentSubmission(gsummary, gsubmission),
		Shifts:           entity.NewStudentShiftDetailsForMonth(summary, shifts, gstudentShifts.MapByShiftID()),
		SuggestedLessons: entity.NewStudentSuggestedLessons(gsubmission),
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV1Handler) UpsertSubmission(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	studentID := getStudentID(ctx)
	req := &request.UpsertSubmissionRequest{}
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

	lessons := make([]*lesson.StudentSuggestedLesson, len(req.SuggestedLessons))
	for i := range req.SuggestedLessons {
		lessons[i] = &lesson.StudentSuggestedLesson{
			SubjectId: req.SuggestedLessons[i].SubjectID,
			Total:     req.SuggestedLessons[i].Total,
		}
	}
	in := &lesson.UpsertStudentShiftsRequest{
		StudentId:      studentID,
		ShiftSummaryId: shiftSummaryID,
		ShiftIds:       req.ShiftIDs,
		Decided:        true,
		Lessons:        lessons,
	}
	out, err := h.lesson.UpsertStudentShifts(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}
	gsubmission := gentity.NewStudentSubmission(out.Submission)
	gstudentShifts := gentity.NewStudentShifts(out.Shifts)

	res := &response.SubmissionResponse{
		Summary:          entity.NewStudentSubmission(gsummary, gsubmission),
		Shifts:           entity.NewStudentShiftDetailsForMonth(summary, shifts, gstudentShifts.MapByShiftID()),
		SuggestedLessons: entity.NewStudentSuggestedLessons(gsubmission),
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
