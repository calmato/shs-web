package api

import (
	"context"

	"github.com/calmato/shs-web/api/internal/user/database"
	"github.com/calmato/shs-web/api/internal/user/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/user"
	"golang.org/x/sync/errgroup"
)

func (s *userService) listStudents(ctx context.Context, limit, offset int64) (entity.Students, int64, error) {
	eg, ectx := errgroup.WithContext(ctx)
	var students entity.Students
	eg.Go(func() (err error) {
		params := &database.ListStudentsParams{
			Limit:  int(limit),
			Offset: int(offset),
		}
		students, err = s.db.Student.List(ectx, params)
		return
	})
	var total int64
	eg.Go(func() (err error) {
		total, err = s.db.Student.Count(ectx)
		return
	})
	if err := eg.Wait(); err != nil {
		return nil, 0, err
	}

	return students, total, nil
}

func (s *userService) ListStudents(
	ctx context.Context, req *user.ListStudentsRequest,
) (*user.ListStudentsResponse, error) {
	students, total, err := s.listStudents(ctx, req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}
	res := &user.ListStudentsResponse{
		Students: students.Proto(),
		Total:    total,
	}
	if err != nil {
		return nil, gRPCError(err)
	}
	return res, nil
}

func (s *userService) GetStudent(ctx context.Context, req *user.GetStudentRequest) (*user.GetStudentResponse, error) {
	if err := s.validator.GetStudent(req); err != nil {
		return nil, gRPCError(err)
	}

	student, err := s.db.Student.Get(ctx, req.Id)
	if err != nil {
		return nil, gRPCError(err)
	}

	res := &user.GetStudentResponse{
		Student: student.Proto(),
	}
	return res, nil
}

func (s *userService) CreateStudent(
	ctx context.Context, req *user.CreateStudentRequest,
) (*user.CreateStudentResponse, error) {
	if err := s.validator.CreateStudent(req); err != nil {
		return nil, gRPCError(err)
	}

	now := jst.Now()
	student := entity.NewStudent(
		req.LastName, req.FirstName, req.LastNameKana, req.FirstNameKana,
		req.Mail, req.Password, entity.SchoolType(req.SchoolType), req.Grade, now)

	err := s.db.Student.Create(ctx, student)
	if err != nil {
		return nil, gRPCError(err)
	}

	res := &user.CreateStudentResponse{
		Student: student.Proto(),
	}
	return res, nil
}
