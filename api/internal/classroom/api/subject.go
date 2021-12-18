package api

import (
	"context"
	"fmt"

	"github.com/calmato/shs-web/api/internal/classroom/database"
	"github.com/calmato/shs-web/api/internal/classroom/entity"
	"github.com/calmato/shs-web/api/proto/classroom"
	"golang.org/x/sync/errgroup"
)

func (s *classroomService) ListSubjects(
	ctx context.Context, req *classroom.ListSubjectsRequest,
) (*classroom.ListSubjectsResponse, error) {
	const prefixKey = "listSubjects"
	if err := s.validator.ListSubjects(req); err != nil {
		return nil, gRPCError(err)
	}

	sharedKey := fmt.Sprintf("%s:%d:%d", prefixKey, req.Limit, req.Offset)
	res, err, _ := s.sharedGroup.Do(sharedKey, func() (interface{}, error) {
		subjects, total, err := s.listSubjects(ctx, req.Limit, req.Offset)
		if err != nil {
			return nil, err
		}
		res := &classroom.ListSubjectsResponse{
			Subjects: subjects.Proto(),
			Total:    total,
		}
		return res, nil
	})
	if err != nil {
		return nil, gRPCError(err)
	}

	return res.(*classroom.ListSubjectsResponse), nil
}

func (s *classroomService) listSubjects(ctx context.Context, limit, offset int64) (entity.Subjects, int64, error) {
	eg, ectx := errgroup.WithContext(ctx)
	var subjects entity.Subjects
	eg.Go(func() (err error) {
		params := &database.ListSubjectsParams{
			Limit:  int(limit),
			Offset: int(offset),
		}
		subjects, err = s.db.Subject.List(ectx, params)
		return
	})
	var total int64
	eg.Go(func() (err error) {
		total, err = s.db.Subject.Count(ectx)
		return
	})
	if err := eg.Wait(); err != nil {
		return nil, 0, err
	}

	return subjects, total, nil
}

func (s *classroomService) MultiGetSubjects(
	ctx context.Context, req *classroom.MultiGetSubjectsRequest,
) (*classroom.MultiGetSubjectsResponse, error) {
	const prefixKey = "multiGetSubjects"
	if err := s.validator.MultiGetSubjects(req); err != nil {
		return nil, gRPCError(err)
	}

	sharedKey := fmt.Sprintf("%s:%v", prefixKey, req.Ids)
	res, err, _ := s.sharedGroup.Do(sharedKey, func() (interface{}, error) {
		subjects, err := s.db.Subject.MultiGet(ctx, req.Ids)
		if err != nil {
			return nil, err
		}
		res := &classroom.MultiGetSubjectsResponse{
			Subjects: subjects.Proto(),
		}
		return res, nil
	})
	if err != nil {
		return nil, gRPCError(err)
	}

	return res.(*classroom.MultiGetSubjectsResponse), nil
}

func (s *classroomService) GetSubject(
	ctx context.Context, req *classroom.GetSubjectRequest,
) (*classroom.GetSubjectResponse, error) {
	if err := s.validator.GetSubject(req); err != nil {
		return nil, gRPCError(err)
	}

	subject, err := s.db.Subject.Get(ctx, req.Id)
	if err != nil {
		return nil, gRPCError(err)
	}

	res := &classroom.GetSubjectResponse{
		Subject: subject.Proto(),
	}
	return res, nil
}
