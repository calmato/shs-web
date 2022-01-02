package api

import (
	"context"

	"github.com/calmato/shs-web/api/internal/lesson/entity"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/calmato/shs-web/api/proto/user"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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

	submission := entity.NewTeacherSubmission(req.TeacherId, req.ShiftSummaryId, req.Desided)
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
