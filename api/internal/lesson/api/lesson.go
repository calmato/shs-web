package api

import (
	"context"

	"github.com/calmato/shs-web/api/internal/lesson/entity"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/calmato/shs-web/api/proto/user"
	"golang.org/x/sync/errgroup"
)

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
	eg.Go(func() error {
		_, err := s.db.Shift.Get(ectx, req.ShiftId, "id")
		return err
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
	}
	return res, nil
}
