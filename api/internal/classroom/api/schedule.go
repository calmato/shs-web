package api

import (
	"context"
	"time"

	"github.com/calmato/shs-web/api/proto/classroom"
)

func (s *classroomService) ListSchedules(
	ctx context.Context, req *classroom.ListSchedulesRequest,
) (*classroom.ListSchedulesResponse, error) {
	if err := s.validator.ListSchedules(req); err != nil {
		return nil, gRPCError(err)
	}

	const sharedKey = "listSchedules"
	res, err, _ := s.sharedGroup.Do(sharedKey, func() (interface{}, error) {
		schedules, err := s.db.Schedule.List(ctx)
		if err != nil {
			return nil, err
		}
		res := &classroom.ListSchedulesResponse{
			Schedules: schedules.Proto(),
		}
		return res, nil
	})
	if err != nil {
		return nil, gRPCError(err)
	}

	return res.(*classroom.ListSchedulesResponse), nil
}

func (s *classroomService) GetSchedule(
	ctx context.Context, req *classroom.GetScheduleRequest,
) (*classroom.GetScheduleResponse, error) {
	if err := s.validator.GetSchedule(req); err != nil {
		return nil, gRPCError(err)
	}

	schedule, err := s.db.Schedule.Get(ctx, time.Weekday(req.Weekday))
	if err != nil {
		return nil, gRPCError(err)
	}

	res := &classroom.GetScheduleResponse{
		Schedule: schedule.Proto(),
	}
	return res, nil
}
