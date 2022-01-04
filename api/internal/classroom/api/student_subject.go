package api

import (
	"context"
	"fmt"

	"github.com/calmato/shs-web/api/internal/classroom/entity"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/calmato/shs-web/api/proto/user"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *classroomService) MultiGetStudentSubjects(
	ctx context.Context, req *classroom.MultiGetStudentSubjectsRequest,
) (*classroom.MultiGetStudentSubjectsResponse, error) {
	const prefixKey = "multiGetStudentSubjects"
	if err := s.validator.MultiGetStudentSubjects(req); err != nil {
		return nil, gRPCError(err)
	}

	sharedKey := fmt.Sprintf("%s:%v", prefixKey, req.StudentIds)
	res, err, _ := s.sharedGroup.Do(sharedKey, func() (interface{}, error) {
		studentSubjects, subjects, err := s.listStudentSubjects(ctx, req.StudentIds)
		if err != nil {
			return nil, err
		}
		res := &classroom.MultiGetStudentSubjectsResponse{
			StudentSubjects: studentSubjects.StudentsProto(),
			Subjects:        subjects.Proto(),
		}
		return res, nil
	})
	if err != nil {
		return nil, gRPCError(err)
	}

	return res.(*classroom.MultiGetStudentSubjectsResponse), nil
}

func (s *classroomService) listStudentSubjects(
	ctx context.Context, studentIDs []string,
) (entity.StudentSubjects, entity.Subjects, error) {
	studentSubjects, err := s.db.StudentSubject.ListByStudentIDs(ctx, studentIDs)
	if err != nil {
		return nil, nil, err
	}
	subjects, err := s.db.Subject.MultiGet(ctx, studentSubjects.SubjectIDs())
	if err != nil {
		return nil, nil, err
	}
	return studentSubjects, subjects, nil
}

func (s *classroomService) GetStudentSubject(
	ctx context.Context, req *classroom.GetStudentSubjectRequest,
) (*classroom.GetStudentSubjectResponse, error) {
	if err := s.validator.GetStudentSubject(req); err != nil {
		return nil, gRPCError(err)
	}

	studentSubjects, subjects, err := s.listStudentSubjects(ctx, []string{req.StudentId})
	if err != nil {
		return nil, gRPCError(err)
	}

	res := &classroom.GetStudentSubjectResponse{
		StudentSubject: studentSubjects.StudentProto(),
		Subjects:       subjects.Proto(),
	}
	return res, nil
}

func (s *classroomService) UpsertStudentSubject(
	ctx context.Context, req *classroom.UpsertStudentSubjectRequest,
) (*classroom.UpsertStudentSubjectResponse, error) {
	if err := s.validator.UpsertStudentSubject(req); err != nil {
		return nil, gRPCError(err)
	}
	schoolType, err := entity.NewSchoolType(req.SchoolType)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	eg, ectx := errgroup.WithContext(ctx)
	eg.Go(func() (err error) {
		in := &user.GetStudentRequest{Id: req.StudentId}
		_, err = s.user.GetStudent(ectx, in)
		return
	})
	eg.Go(func() error {
		subjects, err := s.db.Subject.MultiGet(ectx, req.SubjectIds)
		if err != nil {
			return err
		}
		if len(subjects) != len(req.SubjectIds) {
			return status.Error(codes.InvalidArgument, "api: subject ids length is unmatch")
		}
		return nil
	})
	if err := eg.Wait(); err != nil {
		return nil, gRPCError(err)
	}

	studentSubjects := entity.NewStudentSubjects(req.StudentId, req.SubjectIds)
	err = s.db.StudentSubject.Replace(ctx, schoolType, studentSubjects)
	if err != nil {
		return nil, gRPCError(err)
	}

	res := &classroom.UpsertStudentSubjectResponse{}
	return res, nil
}
