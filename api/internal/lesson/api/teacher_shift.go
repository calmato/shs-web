package api

import (
	"context"
	"errors"

	"github.com/calmato/shs-web/api/internal/lesson/database"
	"github.com/calmato/shs-web/api/internal/lesson/entity"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/calmato/shs-web/api/proto/user"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *lessonService) ListTeacherSubmissionsByShiftSummaryIDs(
	ctx context.Context, req *lesson.ListTeacherSubmissionsByShiftSummaryIDsRequest,
) (*lesson.ListTeacherSubmissionsByShiftSummaryIDsResponse, error) {
	if err := s.validator.ListTeacherSubmissionsByShiftSummaryIDs(req); err != nil {
		return nil, gRPCError(err)
	}

	submissions, err := s.db.TeacherSubmission.ListByShiftSummaryIDs(ctx, req.TeacherId, req.ShiftSummaryIds)
	if err != nil {
		return nil, gRPCError(err)
	}

	res := &lesson.ListTeacherSubmissionsByShiftSummaryIDsResponse{
		Submissions: submissions.Proto(),
	}
	return res, nil
}

func (s *lessonService) ListTeacherShifts(
	ctx context.Context, req *lesson.ListTeacherShiftsRequest,
) (*lesson.ListTeacherShiftsResponse, error) {
	if err := s.validator.ListTeacherShifts(req); err != nil {
		return nil, gRPCError(err)
	}

	shifts, err := s.db.TeacherShift.ListByShiftSummaryID(ctx, req.TeacherIds, req.ShiftSummaryId)
	if err != nil {
		return nil, gRPCError(err)
	}

	res := &lesson.ListTeacherShiftsResponse{
		Shifts: shifts.Proto(),
	}
	return res, nil
}

func (s *lessonService) GetTeacherShifts(
	ctx context.Context, req *lesson.GetTeacherShiftsRequest,
) (*lesson.GetTeacherShiftsResponse, error) {
	if err := s.validator.GetTeacherShifts(req); err != nil {
		return nil, gRPCError(err)
	}

	eg, ectx := errgroup.WithContext(ctx)
	var submission *entity.TeacherSubmission
	eg.Go(func() (err error) {
		submission, err = s.db.TeacherSubmission.Get(ectx, req.TeacherId, req.ShiftSummaryId)
		return
	})
	var shifts entity.TeacherShifts
	eg.Go(func() (err error) {
		shifts, err = s.db.TeacherShift.ListByShiftSummaryID(ectx, []string{req.TeacherId}, req.ShiftSummaryId)
		return
	})
	if err := eg.Wait(); err != nil && !errors.Is(err, database.ErrNotFound) {
		return nil, gRPCError(err)
	}

	res := &lesson.GetTeacherShiftsResponse{
		Submission: submission.Proto(),
		Shifts:     shifts.Proto(),
	}
	return res, nil
}

func (s *lessonService) UpsertTeacherShifts(
	ctx context.Context, req *lesson.UpsertTeacherShiftsRequest,
) (*lesson.UpsertTeacherShiftsResponse, error) {
	if err := s.validator.UpsertTeacherShifts(req); err != nil {
		return nil, gRPCError(err)
	}

	eg, ectx := errgroup.WithContext(ctx)
	eg.Go(func() (err error) {
		in := &user.GetTeacherRequest{Id: req.TeacherId}
		_, err = s.user.GetTeacher(ectx, in)
		return
	})
	var summary *entity.ShiftSummary
	eg.Go(func() (err error) {
		summary, err = s.db.ShiftSummary.Get(ectx, req.ShiftSummaryId)
		return
	})
	eg.Go(func() error {
		shifts, err := s.db.Shift.MultiGet(ectx, req.ShiftIds, "id", "shift_summary_id")
		if err != nil {
			return err
		}
		shifts = shifts.GroupByShiftSummaryID()[req.ShiftSummaryId]
		if len(req.ShiftIds) != len(shifts) {
			return status.Error(codes.InvalidArgument, "api: shift ids length is unmatch")
		}
		return nil
	})
	if err := eg.Wait(); err != nil {
		return nil, gRPCError(err)
	}

	if !summary.EnabledSubmit() {
		return nil, status.Error(codes.FailedPrecondition, "api: outside of shift submission")
	}

	submission := entity.NewTeacherSubmission(req.TeacherId, req.ShiftSummaryId, req.Decided)
	shifts := entity.NewTeacherShifts(req.TeacherId, req.ShiftSummaryId, req.ShiftIds)
	err := s.db.TeacherShift.Replace(ctx, submission, shifts)
	if err != nil {
		return nil, gRPCError(err)
	}

	res := &lesson.UpsertTeacherShiftsResponse{
		Submission: submission.Proto(),
		Shifts:     shifts.Proto(),
	}
	return res, nil
}
