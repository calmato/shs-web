package api

import (
	"context"
	"fmt"

	"github.com/calmato/shs-web/api/internal/lesson/database"
	"github.com/calmato/shs-web/api/internal/lesson/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/calmato/shs-web/api/proto/lesson"
	"golang.org/x/sync/errgroup"
)

func (s *lessonService) ListShiftSummaries(
	ctx context.Context, req *lesson.ListShiftSummariesRequest,
) (*lesson.ListShiftSummariesResponse, error) {
	const prefixKey = "listShiftSummaries"
	if err := s.validator.ListShiftSummaries(req); err != nil {
		return nil, gRPCError(err)
	}

	sharedKey := fmt.Sprintf("%s:%d:%d:%d", prefixKey, req.Status, req.Limit, req.Offset)
	res, err, _ := s.sharedGroup.Do(sharedKey, func() (interface{}, error) {
		status, _ := entity.NewShiftStatus(req.Status)
		orderBy := s.getShiftsummariesOrderBy(req.OrderBy)
		summaries, total, err := s.listShiftSummaries(ctx, status, req.Limit, req.Offset, orderBy)
		if err != nil {
			return nil, err
		}
		res := &lesson.ListShiftSummariesResponse{
			Summaries: summaries.Proto(),
			Total:     total,
		}
		return res, nil
	})
	if err != nil {
		return nil, gRPCError(err)
	}

	return res.(*lesson.ListShiftSummariesResponse), nil
}

func (s *lessonService) getShiftsummariesOrderBy(orderBy lesson.ListShiftSummariesRequest_OrderBy) database.OrderBy {
	switch orderBy {
	case lesson.ListShiftSummariesRequest_ORDER_BY_YEAR_MONTH_ASC:
		return database.OrderByAsc
	case lesson.ListShiftSummariesRequest_ORDER_BY_YEAR_MONTH_DESC:
		return database.OrderByDesc
	default:
		return database.OrderByNone
	}
}

func (s *lessonService) listShiftSummaries(
	ctx context.Context, status entity.ShiftStatus, limit, offset int64, orderBy database.OrderBy,
) (entity.ShiftSummaries, int64, error) {
	eg, ectx := errgroup.WithContext(ctx)
	var summaries entity.ShiftSummaries
	eg.Go(func() (err error) {
		params := &database.ListShiftSummariesParams{
			Status:  status,
			Limit:   int(limit),
			Offset:  int(offset),
			OrderBy: orderBy,
		}
		summaries, err = s.db.ShiftSummary.List(ectx, params)
		return
	})
	var total int64
	eg.Go(func() (err error) {
		total, err = s.db.ShiftSummary.Count(ectx)
		return
	})
	if err := eg.Wait(); err != nil {
		return nil, 0, err
	}

	return summaries, total, nil
}

func (s *lessonService) GetShiftSummary(
	ctx context.Context, req *lesson.GetShiftSummaryRequest,
) (*lesson.GetShiftSummaryResponse, error) {
	if err := s.validator.GetShiftSummary(req); err != nil {
		return nil, gRPCError(err)
	}

	summary, err := s.db.ShiftSummary.Get(ctx, req.Id)
	if err != nil {
		return nil, gRPCError(err)
	}

	res := &lesson.GetShiftSummaryResponse{
		Summary: summary.Proto(),
	}
	return res, nil
}

func (s *lessonService) UpdateShiftSummarySchedule(
	ctx context.Context, req *lesson.UpdateShiftSummaryScheduleRequest,
) (*lesson.UpdateShiftSummaryShceduleResponse, error) {
	if err := s.validator.UpdateShiftSummarySchedule(req); err != nil {
		return nil, gRPCError(err)
	}

	openAt := jst.ParseFromUnix(req.OpenAt)
	endAt := jst.ParseFromUnix(req.EndAt)
	err := s.db.ShiftSummary.UpdateSchedule(ctx, req.Id, openAt, endAt)
	if err != nil {
		return nil, gRPCError(err)
	}
	return &lesson.UpdateShiftSummaryShceduleResponse{}, nil
}

func (s *lessonService) ListShifts(
	ctx context.Context, req *lesson.ListShiftsRequest,
) (*lesson.ListShiftsResponse, error) {
	const prefixKey = "listShifts"
	if err := s.validator.ListShifts(req); err != nil {
		return nil, gRPCError(err)
	}

	sharedKey := fmt.Sprintf("%s:%d", prefixKey, req.ShiftSummaryId)
	res, err, _ := s.sharedGroup.Do(sharedKey, func() (interface{}, error) {
		shifts, err := s.db.Shift.ListBySummaryID(ctx, req.ShiftSummaryId)
		if err != nil {
			return nil, err
		}
		res := &lesson.ListShiftsResponse{
			Shifts: shifts.Proto(),
		}
		return res, nil
	})
	if err != nil {
		return nil, gRPCError(err)
	}

	return res.(*lesson.ListShiftsResponse), nil
}

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
