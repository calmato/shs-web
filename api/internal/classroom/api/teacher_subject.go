package api

import (
	"context"
	"fmt"

	"github.com/calmato/shs-web/api/internal/classroom/entity"
	"github.com/calmato/shs-web/api/proto/classroom"
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
