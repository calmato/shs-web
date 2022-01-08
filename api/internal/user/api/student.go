package api

import (
	"context"

	"github.com/calmato/shs-web/api/internal/user/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/user"
)

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
