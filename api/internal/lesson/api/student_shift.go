package api

import (
	"context"
	"errors"

	"github.com/calmato/shs-web/api/internal/lesson/database"
	"github.com/calmato/shs-web/api/internal/lesson/entity"
	"github.com/calmato/shs-web/api/pkg/set"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/calmato/shs-web/api/proto/user"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *lessonService) ListStudentSubmissionsByShiftSummaryIDs(
	ctx context.Context, req *lesson.ListStudentSubmissionsByShiftSummaryIDsRequest,
) (*lesson.ListStudentSubmissionsByShiftSummaryIDsResponse, error) {
	if err := s.validator.ListStudentSubmissionsByShiftSummaryIDs(req); err != nil {
		return nil, gRPCError(err)
	}

	submissions, err := s.db.StudentSubmission.ListByShiftSummaryIDs(ctx, req.StudentId, req.ShiftSummaryIds)
	if err != nil {
		return nil, gRPCError(err)
	}

	res := &lesson.ListStudentSubmissionsByShiftSummaryIDsResponse{
		Submissions: submissions.Proto(),
	}
	return res, nil
}

func (s *lessonService) ListStudentSubmissionsByStudentIDs(
	ctx context.Context, req *lesson.ListStudentSubmissionsByStudentIDsRequest,
) (*lesson.ListStudentSubmissionsByStudentIDsResponse, error) {
	if err := s.validator.ListStudentSubmissionsByStudentIDs(req); err != nil {
		return nil, gRPCError(err)
	}

	submissions, err := s.db.StudentSubmission.ListByStudentIDs(ctx, req.StudentIds, req.ShiftSummaryId)
	if err != nil {
		return nil, gRPCError(err)
	}

	res := &lesson.ListStudentSubmissionsByStudentIDsResponse{
		Submissions: submissions.Proto(),
	}
	return res, nil
}

func (s *lessonService) ListStudentShifts(
	ctx context.Context, req *lesson.ListStudentShiftsRequest,
) (*lesson.ListStudentShiftsResponse, error) {
	if err := s.validator.ListStudentShifts(req); err != nil {
		return nil, gRPCError(err)
	}

	shifts, err := s.db.StudentShift.ListByShiftSummaryID(ctx, req.StudentIds, req.ShiftSummaryId)
	if err != nil {
		return nil, gRPCError(err)
	}

	res := &lesson.ListStudentShiftsResponse{
		Shifts: shifts.Proto(),
	}
	return res, nil
}

func (s *lessonService) GetStudentShifts(
	ctx context.Context, req *lesson.GetStudentShiftsRequest,
) (*lesson.GetStudentShiftsResponse, error) {
	if err := s.validator.GetStudentShifts(req); err != nil {
		return nil, gRPCError(err)
	}

	eg, ectx := errgroup.WithContext(ctx)
	var submission *entity.StudentSubmission
	eg.Go(func() (err error) {
		submission, err = s.db.StudentSubmission.Get(ectx, req.StudentId, req.ShiftSummaryId)
		return
	})
	var shifts entity.StudentShifts
	eg.Go(func() (err error) {
		shifts, err = s.db.StudentShift.ListByShiftSummaryID(ectx, []string{req.StudentId}, req.ShiftSummaryId)
		return
	})
	if err := eg.Wait(); err != nil && !errors.Is(err, database.ErrNotFound) {
		return nil, gRPCError(err)
	}

	res := &lesson.GetStudentShiftsResponse{
		Shifts: shifts.Proto(),
	}
	if submission != nil {
		res.Submission = submission.Proto()
	}
	return res, nil
}

func (s *lessonService) UpsertStudentShifts(
	ctx context.Context, req *lesson.UpsertStudentShiftsRequest,
) (*lesson.UpsertStudentShiftsResponse, error) {
	if err := s.validator.UpsertStudentShifts(req); err != nil {
		return nil, gRPCError(err)
	}

	eg, ectx := errgroup.WithContext(ctx)
	eg.Go(func() (err error) {
		in := &user.GetStudentRequest{Id: req.StudentId}
		_, err = s.user.GetStudent(ectx, in)
		return
	})
	var subject *classroom.StudentSubject
	eg.Go(func() error {
		in := &classroom.GetStudentSubjectRequest{
			StudentId: req.StudentId,
		}
		out, err := s.classroom.GetStudentSubject(ctx, in)
		if err != nil {
			return err
		}
		subject = out.StudentSubject
		return nil
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

	lessons := entity.NewSuggestedLessons(req.Lessons)
	if !s.isContainsStudentSubjects(subject, lessons) {
		return nil, status.Error(codes.InvalidArgument, "api: contains invalid subject id")
	}

	submission := entity.NewStudentSubmission(req.StudentId, req.ShiftSummaryId, req.Decided, lessons)
	err := submission.FillJSON()
	if err != nil {
		return nil, gRPCError(err)
	}
	shifts := entity.NewStudentShifts(req.StudentId, req.ShiftSummaryId, req.ShiftIds)
	err = s.db.StudentShift.Replace(ctx, submission, shifts)
	if err != nil {
		return nil, gRPCError(err)
	}

	res := &lesson.UpsertStudentShiftsResponse{
		Submission: submission.Proto(),
		Shifts:     shifts.Proto(),
	}
	return res, nil
}

func (s *lessonService) isContainsStudentSubjects(
	subject *classroom.StudentSubject, lessons entity.SuggestedLessons,
) bool {
	set := set.New(len(subject.SubjectIds))
	set.AddInt64s(subject.SubjectIds...)
	for i := range lessons {
		if !set.Contains(lessons[i].SubjectID) {
			return false
		}
	}
	return true
}
