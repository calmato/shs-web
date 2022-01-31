package api

import (
	"context"
	"fmt"

	"github.com/calmato/shs-web/api/internal/lesson/database"
	"github.com/calmato/shs-web/api/internal/lesson/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/calmato/shs-web/api/proto/user"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *lessonService) ListLessons(
	ctx context.Context, req *lesson.ListLessonsRequest,
) (*lesson.ListLessonsResponse, error) {
	if err := s.validator.ListLessons(req); err != nil {
		return nil, gRPCError(err)
	}

	params := &database.ListLessonsParams{
		ShiftSummaryID: req.ShiftSummaryId,
		ShiftID:        req.ShiftId,
		TeacherID:      req.TeacherId,
		StudentID:      req.StudentId,
	}
	eg, ectx := errgroup.WithContext(ctx)
	var lessons entity.Lessons
	eg.Go(func() (err error) {
		lessons, err = s.db.Lesson.List(ectx, params)
		return
	})
	var total int64
	eg.Go(func() (err error) {
		total, err = s.db.Lesson.Count(ectx, params)
		return
	})
	if err := eg.Wait(); err != nil {
		return nil, gRPCError(err)
	}

	eg, ectx = errgroup.WithContext(ctx)
	var summaries entity.ShiftSummaries
	eg.Go(func() (err error) {
		if !req.OnlyDecided {
			return
		}
		summaries, err = s.db.ShiftSummary.MultiGet(ectx, lessons.ShiftSummaryIDs())
		return
	})
	var shifts entity.Shifts
	eg.Go(func() (err error) {
		shifts, err = s.db.Shift.MultiGet(ctx, lessons.ShiftIDs())
		return
	})
	if err := eg.Wait(); err != nil {
		return nil, gRPCError(err)
	}

	if req.OnlyDecided {
		var err error
		lessons, err = lessons.Decided(summaries.Map())
		if err != nil {
			return nil, gRPCError(err)
		}
	}

	res := &lesson.ListLessonsResponse{
		Lessons: lessons.Proto(),
		Shifts:  shifts.Proto(),
		Total:   total,
	}
	return res, nil
}

func (s *lessonService) ListLessonsByDuration(
	ctx context.Context, req *lesson.ListLessonsByDurationRequest,
) (*lesson.ListLessonsByDurationResponse, error) {
	if err := s.validator.ListLessonsByDuration(req); err != nil {
		return nil, gRPCError(err)
	}
	since, err := jst.ParseFromYYYYMMDD(req.Since)
	if err != nil {
		err = fmt.Errorf("failed to parse error: %w", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	until, err := jst.ParseFromYYYYMMDD(req.Until)
	if err != nil {
		err = fmt.Errorf("failed to parse error: %w", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	shifts, err := s.db.Shift.ListByDuration(ctx, jst.BeginningOfDay(since), jst.EndOfDay(until))
	if err != nil {
		return nil, gRPCError(err)
	}
	if len(shifts) == 0 {
		return &lesson.ListLessonsByDurationResponse{}, nil
	}

	params := &database.ListLessonsParams{
		ShiftIDs:  shifts.IDs(),
		TeacherID: req.TeacherId,
		StudentID: req.StudentId,
	}
	lessons, err := s.db.Lesson.List(ctx, params)
	if err != nil {
		return nil, gRPCError(err)
	}

	// 確定しているスケジュール飲みに絞り込み
	summaries, err := s.db.ShiftSummary.MultiGet(ctx, lessons.ShiftSummaryIDs())
	if err != nil {
		return nil, gRPCError(err)
	}
	lessons, err = lessons.Decided(summaries.Map())
	if err != nil {
		return nil, gRPCError(err)
	}

	res := &lesson.ListLessonsByDurationResponse{
		Lessons: lessons.Proto(),
		Shifts:  shifts.Proto(),
	}
	return res, nil
}

func (s *lessonService) CreateLesson(
	ctx context.Context, req *lesson.CreateLessonRequest,
) (*lesson.CreateLessonResponse, error) {
	if err := s.validator.CreateLesson(req); err != nil {
		return nil, gRPCError(err)
	}

	eg, ectx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		in := &classroom.GetSubjectRequest{Id: req.SubjectId}
		_, err := s.classroom.GetSubject(ectx, in)
		return err
	})
	eg.Go(func() error {
		in := &classroom.GetRoomRequest{Id: req.RoomId}
		_, err := s.classroom.GetRoom(ectx, in)
		return err
	})
	eg.Go(func() error {
		in := &user.GetTeacherRequest{Id: req.TeacherId}
		_, err := s.user.GetTeacher(ectx, in)
		return err
	})
	eg.Go(func() error {
		in := &user.GetStudentRequest{Id: req.StudentId}
		_, err := s.user.GetStudent(ectx, in)
		return err
	})
	eg.Go(func() error {
		_, err := s.db.ShiftSummary.Get(ectx, req.ShiftSummaryId, "id")
		return err
	})
	var shift *entity.Shift
	eg.Go(func() (err error) {
		shift, err = s.db.Shift.Get(ectx, req.ShiftId)
		return
	})
	if err := eg.Wait(); err != nil {
		return nil, gRPCError(err)
	}

	l := &entity.Lesson{
		ShiftSummaryID: req.ShiftSummaryId,
		ShiftID:        req.ShiftId,
		SubjectID:      req.SubjectId,
		RoomID:         req.RoomId,
		TeacherID:      req.TeacherId,
		StudentID:      req.StudentId,
		Notes:          req.Notes,
	}
	err := s.db.Lesson.Create(ctx, l)
	if err != nil {
		return nil, gRPCError(err)
	}

	res := &lesson.CreateLessonResponse{
		Lesson: l.Proto(),
		Shift:  shift.Proto(),
	}
	return res, nil
}

func (s *lessonService) UpdateLesson(
	ctx context.Context, req *lesson.UpdateLessonRequest,
) (*lesson.UpdateLessonResponse, error) {
	if err := s.validator.UpdateLesson(req); err != nil {
		return nil, gRPCError(err)
	}

	eg, ectx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		in := &classroom.GetSubjectRequest{Id: req.SubjectId}
		_, err := s.classroom.GetSubject(ectx, in)
		return err
	})
	eg.Go(func() error {
		in := &classroom.GetRoomRequest{Id: req.RoomId}
		_, err := s.classroom.GetRoom(ectx, in)
		return err
	})
	eg.Go(func() error {
		in := &user.GetTeacherRequest{Id: req.TeacherId}
		_, err := s.user.GetTeacher(ectx, in)
		return err
	})
	eg.Go(func() error {
		in := &user.GetStudentRequest{Id: req.StudentId}
		_, err := s.user.GetStudent(ectx, in)
		return err
	})
	eg.Go(func() error {
		_, err := s.db.ShiftSummary.Get(ectx, req.ShiftSummaryId, "id")
		return err
	})
	eg.Go(func() error {
		_, err := s.db.Shift.Get(ectx, req.ShiftId)
		return err
	})
	if err := eg.Wait(); err != nil {
		return nil, gRPCError(err)
	}

	l := &entity.Lesson{
		ID:             req.LessonId,
		ShiftSummaryID: req.ShiftSummaryId,
		ShiftID:        req.ShiftId,
		SubjectID:      req.SubjectId,
		RoomID:         req.RoomId,
		TeacherID:      req.TeacherId,
		StudentID:      req.StudentId,
		Notes:          req.Notes,
	}
	err := s.db.Lesson.Update(ctx, req.LessonId, l)
	if err != nil {
		return nil, gRPCError(err)
	}

	res := &lesson.UpdateLessonResponse{}
	return res, nil
}

func (s *lessonService) DeleteLesson(
	ctx context.Context, req *lesson.DeleteLessonRequest,
) (*lesson.DeleteLessonResponse, error) {
	if err := s.validator.DeleteLesson(req); err != nil {
		return nil, gRPCError(err)
	}

	err := s.db.Lesson.Delete(ctx, req.LessonId)
	if err != nil {
		return nil, gRPCError(err)
	}

	res := &lesson.DeleteLessonResponse{}
	return res, nil
}
