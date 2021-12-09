package api

import (
	"context"

	"github.com/calmato/shs-web/api/internal/user/entity"
	"github.com/calmato/shs-web/api/pkg/uuid"
	"github.com/calmato/shs-web/api/proto/user"
)

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
