package api

import (
	"context"

	"github.com/calmato/shs-web/api/internal/lesson/database"
	"github.com/calmato/shs-web/api/internal/lesson/entity"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/calmato/shs-web/api/proto/user"
	"golang.org/x/sync/errgroup"
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

	shifts, err := s.db.Shift.MultiGet(ctx, lessons.ShiftIDs())
	if err != nil {
		return nil, gRPCError(err)
	}

	res := &lesson.ListLessonsResponse{
		Lessons: lessons.Proto(),
		Shifts:  shifts.Proto(),
		Total:   total,
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
