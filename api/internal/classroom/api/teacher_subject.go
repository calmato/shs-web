package api

import (
	"context"
	"fmt"

	"github.com/calmato/shs-web/api/internal/classroom/entity"
	"github.com/calmato/shs-web/api/proto/classroom"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *classroomService) MultiGetTeacherSubjects(
	ctx context.Context, req *classroom.MultiGetTeacherSubjectsRequest,
) (*classroom.MultiGetTeacherSubjectsResponse, error) {
	const prefixKey = "multiGetTeacherSubjects"
	if err := s.validator.MultiGetTeacherSubjects(req); err != nil {
		return nil, gRPCError(err)
	}

	sharedKey := fmt.Sprintf("%s:%v", prefixKey, req.TeacherIds)
	res, err, _ := s.sharedGroup.Do(sharedKey, func() (interface{}, error) {
		teacherSubjects, subjects, err := s.listTeacherSubjects(ctx, req.TeacherIds)
		if err != nil {
			return nil, err
		}
		res := &classroom.MultiGetTeacherSubjectsResponse{
			TeacherSubjects: teacherSubjects.TeachersProto(),
			Subjects:        subjects.Proto(),
		}
		return res, nil
	})
	if err != nil {
		return nil, gRPCError(err)
	}

	return res.(*classroom.MultiGetTeacherSubjectsResponse), nil
}

func (s *classroomService) listTeacherSubjects(
	ctx context.Context, teacherIDs []string,
) (entity.TeacherSubjects, entity.Subjects, error) {
	teacherSubjects, err := s.db.TeacherSubject.ListByTeacherIDs(ctx, teacherIDs)
	if err != nil {
		return nil, nil, err
	}
	subjects, err := s.db.Subject.MultiGet(ctx, teacherSubjects.SubjectIDs())
	if err != nil {
		return nil, nil, err
	}
	return teacherSubjects, subjects, nil
}

func (s *classroomService) GetTeacherSubject(
	ctx context.Context, req *classroom.GetTeacherSubjectRequest,
) (*classroom.GetTeacherSubjectResponse, error) {
	if err := s.validator.GetTeacherSubject(req); err != nil {
		return nil, gRPCError(err)
	}

	teacherSubjects, subjects, err := s.listTeacherSubjects(ctx, []string{req.TeacherId})
	if err != nil {
		return nil, gRPCError(err)
	}

	res := &classroom.GetTeacherSubjectResponse{
		TeacherSubject: teacherSubjects.TeacherProto(),
		Subjects:       subjects.Proto(),
	}
	return res, nil
}

func (s *classroomService) UpsertTeacherSubject(
	ctx context.Context, req *classroom.UpsertTeacherSubjectRequest,
) (*classroom.UpsertTeacherSubjectResponse, error) {
	if err := s.validator.UpsertTeacherSubject(req); err != nil {
		return nil, gRPCError(err)
	}

	schoolType, err := entity.NewSchoolType(req.SchoolType)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	subjects, err := s.db.Subject.MultiGet(ctx, req.SubjectIds)
	if err != nil {
		return nil, gRPCError(err)
	}
	if len(subjects) != len(req.SubjectIds) {
		return nil, status.Error(codes.InvalidArgument, "api: subject ids length is unmatch")
	}

	teacherSubjects := entity.NewTeacherSubjects(req.TeacherId, req.SubjectIds)
	err = s.db.TeacherSubject.Replace(ctx, schoolType, teacherSubjects)
	if err != nil {
		return nil, gRPCError(err)
	}

	res := &classroom.UpsertTeacherSubjectResponse{}
	return res, nil
}
