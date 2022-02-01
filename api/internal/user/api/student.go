package api

import (
	"context"
	"fmt"

	"github.com/calmato/shs-web/api/internal/user/database"
	"github.com/calmato/shs-web/api/internal/user/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/user"
	"golang.org/x/sync/errgroup"
)

func (s *userService) ListStudents(
	ctx context.Context, req *user.ListStudentsRequest,
) (*user.ListStudentsResponse, error) {
	const prefixKey = "listStudents"
	if err := s.validator.ListStudents(req); err != nil {
		return nil, gRPCError(err)
	}

	sharedKey := fmt.Sprintf("%s:%d:%d", prefixKey, req.Limit, req.Offset)
	res, err, _ := s.sharedGroup.Do(sharedKey, func() (interface{}, error) {
		students, total, err := s.listStudents(ctx, req.Limit, req.Offset)
		if err != nil {
			return nil, err
		}
		res := &user.ListStudentsResponse{
			Students: students.Proto(),
			Total:    total,
		}
		return res, nil
	})
	if err != nil {
		return nil, gRPCError(err)
	}

	return res.(*user.ListStudentsResponse), nil
}

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

func (s *userService) MultiGetStudents(
	ctx context.Context, req *user.MultiGetStudentsRequest,
) (*user.MultiGetStudentsResponse, error) {
	if err := s.validator.MultiGetStudents(req); err != nil {
		return nil, gRPCError(err)
	}

	students, err := s.db.Student.MultiGet(ctx, req.Ids)
	if err != nil {
		return nil, gRPCError(err)
	}

	res := &user.MultiGetStudentsResponse{
		Students: students.Proto(),
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

func (s *userService) UpdateStudentMail(
	ctx context.Context, req *user.UpdateStudentMailRequest,
) (*user.UpdateStudentMailResponse, error) {
	if err := s.validator.UpdateStudentMail(req); err != nil {
		return nil, gRPCError(err)
	}

	err := s.db.Student.UpdateMail(ctx, req.Id, req.Mail)
	if err != nil {
		return nil, gRPCError(err)
	}
	return &user.UpdateStudentMailResponse{}, nil
}

func (s *userService) UpdateStudentPassword(
	ctx context.Context, req *user.UpdateStudentPasswordRequest,
) (*user.UpdateStudentPasswordResponse, error) {
	if err := s.validator.UpdateStudentPassword(req); err != nil {
		return nil, gRPCError(err)
	}

	err := s.db.Student.UpdatePassword(ctx, req.Id, req.Password)
	if err != nil {
		return nil, gRPCError(err)
	}
	return &user.UpdateStudentPasswordResponse{}, nil
}

func (s *userService) DeleteStudent(
	ctx context.Context, req *user.DeleteStudentRequest,
) (*user.DeleteStudentResponse, error) {
	if err := s.validator.DeleteStudent(req); err != nil {
		return nil, gRPCError(err)
	}

	err := s.db.Student.Delete(ctx, req.Id)
	if err != nil {
		return nil, gRPCError(err)
	}
	return &user.DeleteStudentResponse{}, nil
}
