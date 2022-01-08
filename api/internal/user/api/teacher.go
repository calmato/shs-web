package api

import (
	"context"
	"fmt"

	"github.com/calmato/shs-web/api/internal/user/database"
	"github.com/calmato/shs-web/api/internal/user/entity"
	"github.com/calmato/shs-web/api/pkg/uuid"
	"github.com/calmato/shs-web/api/proto/user"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	role, err := entity.NewRole(req.Role)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "api: invalid role")
	}

	uid := uuid.Base58Encode(uuid.New())
	teacher := &entity.Teacher{
		ID:            uid,
		LastName:      req.LastName,
		FirstName:     req.FirstName,
		LastNameKana:  req.LastNameKana,
		FirstNameKana: req.FirstNameKana,
		Mail:          req.Mail,
		Role:          role,
		Password:      req.Password,
	}

	err = s.db.Teacher.Create(ctx, teacher)
	if err != nil {
		return nil, gRPCError(err)
	}

	res := &user.CreateTeacherResponse{
		Teacher: teacher.Proto(),
	}
	return res, nil
}

func (s *userService) UpdateTeacherMail(
	ctx context.Context, req *user.UpdateTeacherMailRequest,
) (*user.UpdateTeacherMailResponse, error) {
	if err := s.validator.UpdateTeacherMail(req); err != nil {
		return nil, gRPCError(err)
	}

	err := s.db.Teacher.UpdateMail(ctx, req.Id, req.Mail)
	if err != nil {
		return nil, gRPCError(err)
	}
	return &user.UpdateTeacherMailResponse{}, nil
}

func (s *userService) UpdateTeacherPassword(
	ctx context.Context, req *user.UpdateTeacherPasswordRequest,
) (*user.UpdateTeacherPasswordResponse, error) {
	if err := s.validator.UpdateTeacherPassword(req); err != nil {
		return nil, gRPCError(err)
	}

	err := s.db.Teacher.UpdatePassword(ctx, req.Id, req.Password)
	if err != nil {
		return nil, gRPCError(err)
	}
	return &user.UpdateTeacherPasswordResponse{}, nil
}

func (s *userService) UpdateTeacherRole(
	ctx context.Context, req *user.UpdateTeacherRoleRequest,
) (*user.UpdateTeacherRoleResponse, error) {
	if err := s.validator.UpdateTeacherRole(req); err != nil {
		return nil, gRPCError(err)
	}
	role, err := entity.NewRole(req.Role)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "api: invalid role")
	}

	err = s.db.Teacher.UpdateRole(ctx, req.Id, role)
	if err != nil {
		return nil, gRPCError(err)
	}
	return &user.UpdateTeacherRoleResponse{}, nil
}

func (s *userService) DeleteTeacher(
	ctx context.Context, req *user.DeleteTeacherRequest,
) (*user.DeleteTeacherResponse, error) {
	if err := s.validator.DeleteTeacher(req); err != nil {
		return nil, gRPCError(err)
	}

	err := s.db.Teacher.Delete(ctx, req.Id)
	if err != nil {
		return nil, gRPCError(err)
	}
	return &user.DeleteTeacherResponse{}, nil
}
