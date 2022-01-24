package handler

import (
	"net/http"
	"strconv"

	gentity "github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"
	"github.com/calmato/shs-web/api/internal/gateway/teacher/v1/response"

	"github.com/calmato/shs-web/api/internal/gateway/util"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

func (h *apiV1Handler) ListEnabledStudentShifts(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	studentID := ctx.Param("studentId")
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
	var submission *gentity.StudentSubmission
	var studentShifts gentity.StudentShifts
	eg.Go(func() error {
		in := &lesson.GetStudentShiftsRequest{
			StudentId:      studentID,
			ShiftSummaryId: shiftSummaryID,
		}
		out, err := h.lesson.GetStudentShifts(ectx, in)
		if err != nil {
			return err
		}
		submission = gentity.NewStudentSubmission(out.Submission)
		studentShifts = gentity.NewStudentShifts(out.Shifts)
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

	res := &response.StudentShiftsResponse{
		Summary:          entity.NewStudentSubmission(gsummary, submission),
		Shifts:           entity.NewEnabledStudentShiftDetails(summary, shifts, studentShifts.MapByShiftID()),
		SuggestedLessons: entity.NewStudentSuggestedLessons(submission),
	}
	ctx.JSON(http.StatusOK, res)
}
