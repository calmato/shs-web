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
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	var gtemplate *gentity.StudentShiftTemplate
	eg.Go(func() error {
		in := &lesson.GetStudentShiftTemplateRequest{StudentId: studentID}
		out, err := h.lesson.GetStudentShiftTemplate(ectx, in)
		if st, ok := status.FromError(err); ok && st.Code() == codes.NotFound {
			gtemplate = &gentity.StudentShiftTemplate{
				Lessons:   gentity.StudentSuggestedLessons{},
				Schedules: gentity.StudentShiftSchedules{},
			}
			return nil
		}
		if err != nil {
			return err
		}
		gtemplate = gentity.NewStudentShiftTemplate(out.Template)
		return nil
	})
	var gsubmission *gentity.StudentSubmission
	var gstudentShifts gentity.StudentShifts
	eg.Go(func() error {
		in := &lesson.GetStudentShiftsRequest{
			StudentId:      studentID,
			ShiftSummaryId: shiftSummaryID,
		}
		out, err := h.lesson.GetStudentShifts(ectx, in)
		if st, ok := status.FromError(err); ok && st.Code() == codes.NotFound {
			return nil
		}
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
	shiftsMap, err := gshifts.GroupByDate()
	if err != nil {
		httpError(ctx, err)
		return
	}

	var suggestedLessons entity.StudentSuggestedLessons
	var shifts entity.StudentShiftDetails
	if gsubmission == nil {
		// 未提出 -> テンプレート使用
		suggestedLessons = entity.NewStudentSuggestedLessons(gtemplate.Lessons)
		shifts = entity.NewStudentShiftDetailsWithTemplate(summary, shiftsMap, gtemplate.Schedules)
	} else {
		// 提出済み -> 提出データ使用
		suggestedLessons = entity.NewStudentSuggestedLessons(gsubmission.SuggestedLessons)
		shifts = entity.NewStudentShiftDetailsForMonth(summary, shiftsMap, gstudentShifts.MapByShiftID())
	}

	res := &response.SubmissionResponse{
		Summary:          entity.NewStudentSubmission(gsummary, gsubmission),
		Shifts:           shifts,
		SuggestedLessons: suggestedLessons,
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

	var suggestedLessons entity.StudentSuggestedLessons
	if gsubmission != nil {
		suggestedLessons = entity.NewStudentSuggestedLessons(gsubmission.SuggestedLessons)
	}
	res := &response.SubmissionResponse{
		Summary:          entity.NewStudentSubmission(gsummary, gsubmission),
		Shifts:           entity.NewStudentShiftDetailsForMonth(summary, shifts, gstudentShifts.MapByShiftID()),
		SuggestedLessons: suggestedLessons,
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV1Handler) GetSubmissionTemplate(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	studentID := getStudentID(ctx)

	eg, ectx := errgroup.WithContext(c)
	var gschedules gentity.Schedules
	eg.Go(func() error {
		in := &classroom.ListSchedulesRequest{}
		out, err := h.classroom.ListSchedules(ectx, in)
		if err != nil {
			return err
		}
		gschedules = gentity.NewSchedules(out.Schedules)
		return nil
	})
	var gtemplate *gentity.StudentShiftTemplate
	eg.Go(func() error {
		in := &lesson.GetStudentShiftTemplateRequest{StudentId: studentID}
		out, err := h.lesson.GetStudentShiftTemplate(ectx, in)
		if err != nil && status.Code(err) != codes.NotFound {
			return err
		}
		gtemplate = gentity.NewStudentShiftTemplate(out.Template)
		return nil
	})
	if err := eg.Wait(); err != nil {
		httpError(ctx, err)
		return
	}

	res := response.SubmissionTemplateResponse{
		Schedules:        entity.NewStudentSchedules(gschedules, gtemplate.Schedules),
		SuggestedLessons: entity.NewStudentSuggestedLessons(gtemplate.Lessons),
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV1Handler) UpsertSubmissionTemplate(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	studentID := getStudentID(ctx)

	req := &request.UpsertSubmissionTemplateRequest{}
	if err := ctx.BindJSON(req); err != nil {
		badRequest(ctx, err)
		return
	}

	in := &lesson.UpsertStudentShiftTemplateRequest{
		StudentId: studentID,
		Template: &lesson.StudentShiftTemplateToUpsert{
			Schedules:        h.newStudentShiftSchedulesToUpsert(req.Schedules),
			SuggestedLessons: h.newSuggestedLessonsToUpsert(req.SuggestedLessons),
		},
	}
	_, err := h.lesson.UpsertStudentShiftTemplate(c, in)
	if err != nil {
		httpError(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (h *apiV1Handler) newStudentShiftSchedulesToUpsert(
	schedules entity.StudentSchedules,
) []*lesson.StudentShiftTemplateToUpsert_Schedule {
	req := make([]*lesson.StudentShiftTemplateToUpsert_Schedule, len(schedules))
	for i := range schedules {
		lessons := make([]*lesson.StudentShiftTemplateToUpsert_Lesson, 0, len(schedules[i].Lessons))
		for _, l := range schedules[i].Lessons {
			if !l.Enabled {
				continue
			}
			lesson := &lesson.StudentShiftTemplateToUpsert_Lesson{
				StartTime: l.StartTime,
				EndTime:   l.EndTime,
			}
			lessons = append(lessons, lesson)
		}
		req[i] = &lesson.StudentShiftTemplateToUpsert_Schedule{
			Weekday: int32(schedules[i].Weekday),
			Lessons: lessons,
		}
	}
	return req
}

func (h *apiV1Handler) newSuggestedLessonsToUpsert(
	lessons entity.StudentSuggestedLessons,
) []*lesson.StudentSuggestedLesson {
	res := make([]*lesson.StudentSuggestedLesson, len(lessons))
	for i := range lessons {
		res[i] = &lesson.StudentSuggestedLesson{
			SubjectId: lessons[i].SubjectID,
			Total:     lessons[i].Total,
		}
	}
	return res
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
