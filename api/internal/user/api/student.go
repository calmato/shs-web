package api

import (
	"context"

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
