package api

import (
	"context"

	"github.com/calmato/shs-web/api/internal/lesson/entity"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/calmato/shs-web/api/proto/lesson"
)

func (s *lessonService) CreateShifts(
	ctx context.Context, req *lesson.CreateShiftsRequest,
) (*lesson.CreateShiftsResponse, error) {
	if err := s.validator.CreateShifts(req); err != nil {
		return nil, gRPCError(err)
	}

	schedules, err := s.listClassroomSchedules(ctx)
	if err != nil {
		return nil, gRPCError(err)
	}
	schduleMap := schedules.Map()

	summary := entity.NewShiftSummary(req.YearMonth, req.OpenAt, req.EndAt)
	shifts, err := entity.NewShifts(summary, schduleMap, req.ClosedDates)
	if err != nil {
		return nil, gRPCError(err)
	}
	err = s.db.Shift.MultipleCreate(ctx, summary, shifts)
	if err != nil {
		return nil, gRPCError(err)
	}

	res := &lesson.CreateShiftsResponse{
		Summary: summary.Proto(),
		Shifts:  shifts.Proto(),
	}
	return res, nil
}

func (s *lessonService) listClassroomSchedules(ctx context.Context) (entity.Schedules, error) {
	in := &classroom.ListSchedulesRequest{}
	out, err := s.classroom.ListSchedules(ctx, in)
	if err != nil {
		return nil, err
	}
	schedules := entity.NewSchedules(out.Schedules)
	return schedules, nil
}
