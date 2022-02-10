package api

import (
	"context"
	"fmt"

	"github.com/calmato/shs-web/api/internal/lesson/database"
	"github.com/calmato/shs-web/api/internal/lesson/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/calmato/shs-web/api/proto/messenger"
	"go.uber.org/zap"
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

func (s *lessonService) UpdateShiftSummaryDecided(
	ctx context.Context, req *lesson.UpdateShiftSummaryDecidedRequest,
) (*lesson.UpdateShiftSummaryDecidedResponse, error) {
	if err := s.validator.UpdateShiftSummaryDecided(req); err != nil {
		return nil, gRPCError(err)
	}

	err := s.db.ShiftSummary.UpdateDecided(ctx, req.Id, req.Decided)
	if err != nil {
		return nil, gRPCError(err)
	}

	s.waitGroup.Add(1)
	go func() {
		defer s.waitGroup.Done()
		if !req.Decided {
			return
		}
		bgctx := context.Background()
		in := &messenger.NotifyLessonDecidedRequest{ShiftSummaryId: req.Id}
		if _, err := s.messenger.NotifyLessonDecided(bgctx, in); err != nil {
			s.logger.Error("Failed to notify lesson decided", zap.Int64("shiftSummaryId", req.Id))
		}
	}()
	return &lesson.UpdateShiftSummaryDecidedResponse{}, nil
}

func (s *lessonService) DeleteShiftSummary(
	ctx context.Context, req *lesson.DeleteShiftSummaryRequest,
) (*lesson.DeleteShiftSummaryResponse, error) {
	if err := s.validator.DeleteShiftSummary(req); err != nil {
		return nil, gRPCError(err)
	}

	err := s.db.ShiftSummary.Delete(ctx, req.Id)
	if err != nil {
		return nil, gRPCError(err)
	}

	return &lesson.DeleteShiftSummaryResponse{}, nil
}

func (s *lessonService) ListShifts(
	ctx context.Context, req *lesson.ListShiftsRequest,
) (*lesson.ListShiftsResponse, error) {
	const prefixKey = "listShifts"
	if err := s.validator.ListShifts(req); err != nil {
		return nil, gRPCError(err)
	}

	sharedKey := fmt.Sprintf("%s:%d:%d", prefixKey, req.ShiftSummaryId, req.ShiftId)
	res, err, _ := s.sharedGroup.Do(sharedKey, func() (interface{}, error) {
		params := &database.ListShiftsParams{
			ShiftSummaryID: req.ShiftSummaryId,
			ShiftID:        req.ShiftId,
		}
		shifts, err := s.db.Shift.List(ctx, params)
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

	summary.Fill(s.now())
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

func (s *lessonService) ListSubmissions(
	ctx context.Context, req *lesson.ListSubmissionsRequest,
) (*lesson.ListSubmissionsResponse, error) {
	const prefixKey = "listSubmissions"
	if err := s.validator.ListSubmissions(req); err != nil {
		return nil, gRPCError(err)
	}
	sharedKey := fmt.Sprintf("%s:%d", prefixKey, req.ShiftId)
	res, err, _ := s.sharedGroup.Do(sharedKey, func() (interface{}, error) {
		teachers, students, err := s.listSubmissions(ctx, req.ShiftId)
		if err != nil {
			return nil, err
		}
		res := &lesson.ListSubmissionsResponse{
			TeacherShifts: teachers.Proto(),
			StudentShifts: students.Proto(),
		}
		return res, nil
	})
	if err != nil {
		return nil, gRPCError(err)
	}
	return res.(*lesson.ListSubmissionsResponse), nil
}

func (s *lessonService) listSubmissions(
	ctx context.Context, shiftID int64,
) (entity.TeacherShifts, entity.StudentShifts, error) {
	eg, ectx := errgroup.WithContext(ctx)
	var teachers entity.TeacherShifts
	eg.Go(func() (err error) {
		teachers, err = s.db.TeacherShift.ListByShiftID(ectx, shiftID)
		return
	})
	var students entity.StudentShifts
	eg.Go(func() (err error) {
		students, err = s.db.StudentShift.ListByShiftID(ectx, shiftID)
		return
	})
	if err := eg.Wait(); err != nil {
		return nil, nil, err
	}
	return teachers, students, nil
}
