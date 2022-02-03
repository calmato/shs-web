package database

import (
	"context"
	"errors"
	"time"

	"github.com/calmato/shs-web/api/internal/lesson/entity"
	"github.com/calmato/shs-web/api/pkg/database"
	"github.com/calmato/shs-web/api/pkg/jst"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const studentShiftTemplateTable = "student_shift_templates"

var studentShiftTemplateFields = []string{
	"student_id", "schedules", "suggested_lessons", "created_at", "updated_at",
}

type studentShiftTemplate struct {
	db  *database.Client
	now func() time.Time
}

func NewStudentShiftTemplate(db *database.Client) StudentShiftTemplate {
	return &studentShiftTemplate{
		db:  db,
		now: jst.Now,
	}
}

func (s *studentShiftTemplate) Get(ctx context.Context, studentID string) (*entity.StudentShiftTemplate, error) {
	var template *entity.StudentShiftTemplate

	stmt := s.db.DB.Table(studentShiftTemplateTable).
		Select(studentShiftTemplateFields).
		Where("student_id", studentID)

	err := stmt.First(&template).Error
	if err != nil {
		return nil, dbError(err)
	}
	err = template.Fill()
	if err != nil {
		return nil, dbError(err)
	}
	return template, nil
}

func (s *studentShiftTemplate) Upsert(
	ctx context.Context, studentID string, template *entity.StudentShiftTemplate,
) error {
	now := s.now()
	err := template.FillJSON()
	if err != nil {
		return dbError(err)
	}

	tx, err := s.db.Begin()
	if err != nil {
		return dbError(err)
	}
	defer s.db.Close(tx)

	var current *entity.StudentShiftTemplate
	err = tx.Table(studentShiftTemplateTable).
		Select([]string{"student_id", "created_at"}).
		Where("student_id = ?", studentID).
		First(&current).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return dbError(err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		template.CreatedAt = now
		template.UpdatedAt = now
	} else {
		template.CreatedAt = current.CreatedAt
		template.UpdatedAt = now
	}

	err = tx.Clauses(clause.OnConflict{UpdateAll: true}).
		Table(studentShiftTemplateTable).
		Create(&template).Error
	if err != nil {
		tx.Rollback()
		return dbError(err)
	}
	return dbError(tx.Commit().Error)
}
