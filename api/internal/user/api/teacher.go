package api

import (
	"context"

	"github.com/calmato/shs-web/api/internal/user/database"
	"github.com/calmato/shs-web/api/internal/user/entity"
	"github.com/calmato/shs-web/api/pkg/uuid"
	"github.com/calmato/shs-web/api/proto/user"
)

func (s *userService) ListTeachers(
	ctx context.Context, req *user.ListTeachersRequest,
) (*user.ListTeachersResponse, error) {
	if err := s.validator.ListTeachers(req); err != nil {
		return nil, gRPCError(err)
	}

	params := &database.ListTeachersParams{
		Limit:  int(req.Limit),
		Offset: int(req.Offset),
	}
	teachers, err := s.db.Teacher.List(ctx, params)
	if err != nil {
		return nil, gRPCError(err)
	}

	res := &user.ListTeachersResponse{
		Teachers: teachers.Proto(),
	}
	return res, nil
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
