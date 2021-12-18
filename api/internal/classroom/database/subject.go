package database

import (
	"context"
	"time"

	"github.com/calmato/shs-web/api/internal/classroom/entity"
	"github.com/calmato/shs-web/api/pkg/database"
	"github.com/calmato/shs-web/api/pkg/jst"
)

const subjectTable = "subjects"

var subjectFields = []string{
	"id", "name", "color", "created_at", "updated_at",
}

type subject struct {
	db  *database.Client
	now func() time.Time
}

func NewSubject(db *database.Client) Subject {
	return &subject{
		db:  db,
		now: jst.Now,
	}
}

func (s *subject) List(ctx context.Context, params *ListSubjectsParams, fields ...string) (entity.Subjects, error) {
	var subjects entity.Subjects
	if len(fields) == 0 {
		fields = subjectFields
	}

	stmt := s.db.DB.Table(subjectTable).Select(fields)
	if params.Limit > 0 {
		stmt.Limit(params.Limit)
	}
	if params.Offset > 0 {
		stmt.Offset(params.Offset)
	}

	err := stmt.Find(&subjects).Error
	return subjects, dbError(err)
}

func (s *subject) MultiGet(ctx context.Context, ids []int64, fields ...string) (entity.Subjects, error) {
	var subjects entity.Subjects
	if len(fields) == 0 {
		fields = subjectFields
	}

	stmt := s.db.DB.Table(subjectTable).Select(fields).
		Where("id IN (?)", ids)

	err := stmt.Find(&subjects).Error
	return subjects, dbError(err)
}

func (s *subject) Get(ctx context.Context, id int64, fields ...string) (*entity.Subject, error) {
	var subject *entity.Subject
	if len(fields) == 0 {
		fields = subjectFields
	}

	stmt := s.db.DB.Table(subjectTable).Select(fields).
		Where("id = ?", id)

	err := stmt.First(&subject).Error
	if err != nil {
		return nil, dbError(err)
	}
	return subject, nil
}

func (s *subject) Count(ctx context.Context) (int64, error) {
	var total int64
	err := s.db.DB.Table(subjectTable).Count(&total).Error
	return total, dbError(err)
}
