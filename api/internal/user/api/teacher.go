package api

import (
	"context"
	"fmt"

	"github.com/calmato/shs-web/api/internal/user/database"
	"github.com/calmato/shs-web/api/internal/user/entity"
	"github.com/calmato/shs-web/api/pkg/uuid"
	"github.com/calmato/shs-web/api/proto/user"
	"golang.org/x/sync/errgroup"
)

func (s *userService) ListTeachers(
	ctx context.Context, req *user.ListTeachersRequest,
) (*user.ListTeachersResponse, error) {
	const prefixKey = "listTeachers"
	if err := s.validator.ListTeachers(req); err != nil {
		return nil, gRPCError(err)
	}

	sharedKey := fmt.Sprintf("%s:%d:%d", prefixKey, req.Limit, req.Offset)
	res, err, _ := s.sharedGroup.Do(sharedKey, func() (interface{}, error) {
		teachers, total, err := s.listTeachers(ctx, req.Limit, req.Offset)
		if err != nil {
			return nil, err
		}
		res := &user.ListTeachersResponse{
			Teachers: teachers.Proto(),
			Total:    total,
		}
		return res, nil
	})
	if err != nil {
		return nil, gRPCError(err)
	}

	return res.(*user.ListTeachersResponse), nil
}

func (s *userService) listTeachers(ctx context.Context, limit, offset int64) (entity.Teachers, int64, error) {
	eg, ectx := errgroup.WithContext(ctx)
	var teachers entity.Teachers
	eg.Go(func() (err error) {
		params := &database.ListTeachersParams{
			Limit:  int(limit),
			Offset: int(offset),
		}
		teachers, err = s.db.Teacher.List(ectx, params)
		return
	})
	var total int64
	eg.Go(func() (err error) {
		total, err = s.db.Teacher.Count(ectx)
		return
	})
	if err := eg.Wait(); err != nil {
		return nil, 0, err
	}

	return teachers, total, nil
}

func (s *userService) GetTeacher(ctx context.Context, req *user.GetTeacherRequest) (*user.GetTeacherResponse, error) {
	if err := s.validator.GetTeacher(req); err != nil {
		return nil, gRPCError(err)
	}

	teacher, err := s.db.Teacher.Get(ctx, req.Id)
	if err != nil {
		return nil, gRPCError(err)
	}

	res := &user.GetTeacherResponse{
		Teacher: teacher.Proto(),
	}
	return res, nil
}

func (s *userService) CreateTeacher(
	ctx context.Context, req *user.CreateTeacherRequest,
) (*user.CreateTeacherResponse, error) {
	if err := s.validator.CreateTeacher(req); err != nil {
		return nil, gRPCError(err)
	}

	uid := uuid.Base58Encode(uuid.New())
	teacher := &entity.Teacher{
		ID:            uid,
		LastName:      req.LastName,
		FirstName:     req.FirstName,
		LastNameKana:  req.LastNameKana,
		FirstNameKana: req.FirstNameKana,
		Mail:          req.Mail,
		Role:          int32(req.Role),
		Password:      req.Password,
	}

	err := s.db.Teacher.Create(ctx, teacher)
	if err != nil {
		return nil, gRPCError(err)
	}

	res := &user.CreateTeacherResponse{
		Teacher: teacher.Proto(),
	}
	return res, nil
}
