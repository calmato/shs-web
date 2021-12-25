package database

import (
	"context"
	"time"

	"github.com/calmato/shs-web/api/internal/classroom/entity"
	"github.com/calmato/shs-web/api/pkg/database"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/classroom"
)

const subjectTable = "subjects"

var subjectFields = []string{
	"id", "name", "color", "school_type", "created_at", "updated_at",
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
	if params.SchoolType != classroom.SchoolType_SCHOOL_TYPE_UNKNOWN {
		stmt.Where("school_type = ?", int32(params.SchoolType))
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

func (s *subject) Create(ctx context.Context, subject *entity.Subject) error {
	now := s.now()
	subject.CreatedAt = now
	subject.UpdatedAt = now

	tx, err := s.db.Begin()
	if err != nil {
		return dbError(err)
	}
	defer s.db.Close(tx)

	err = tx.Table(subjectTable).Create(&subject).Error
	if err != nil {
		tx.Rollback()
		return dbError(err)
	}
	return dbError(tx.Commit().Error)
}

func (s *subject) Update(ctx context.Context, subjectID int64, subject *entity.Subject) error {
	tx, err := s.db.Begin()
	if err != nil {
		return dbError(err)
	}
	defer s.db.Close(tx)

	var current *entity.Subject
	err = tx.Table(subjectTable).
		Select([]string{"id", "created_at"}).
		Where("id = ?", subjectID).
		First(&current).Error
	if err != nil {
		return dbError(err)
	}

	subject.ID = current.ID
	subject.CreatedAt = current.CreatedAt
	subject.UpdatedAt = s.now()

	err = tx.Table(subjectTable).Save(&subject).Error
	if err != nil {
		tx.Rollback()
		return dbError(err)
	}
	return dbError(tx.Commit().Error)
}

func (s *subject) Delete(ctx context.Context, subjectID int64) error {
	tx, err := s.db.Begin()
	if err != nil {
		return dbError(err)
	}
	defer s.db.Close(tx)

	err = tx.Table(subjectTable).
		Where("id = ?", subjectID).
		Delete(&entity.Subject{}).Error
	if err != nil {
		tx.Rollback()
		return dbError(err)
	}
	return dbError(tx.Commit().Error)
}

func (s *subject) Count(ctx context.Context) (int64, error) {
	var total int64
	err := s.db.DB.Table(subjectTable).Count(&total).Error
	return total, dbError(err)
}
