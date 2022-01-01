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

const teacherShiftTable = "teacher_shifts"

// var teacherShiftFields = []string{
// 	"teacher_id", "shift_id", "shift_summary_id", "created_at", "updated_at",
// }

type teacherShift struct {
	db  *database.Client
	now func() time.Time
}

func NewTeacherShift(db *database.Client) TeacherShift {
	return &teacherShift{
		db:  db,
		now: jst.Now,
	}
}

func (s *teacherShift) Replace(
	ctx context.Context, submission *entity.TeacherSubmission, shifts entity.TeacherShifts,
) error {
	now := s.now()
	for i := range shifts {
		shifts[i].CreatedAt = now
		shifts[i].UpdatedAt = now
	}

	tx, err := s.db.Begin()
	if err != nil {
		return dbError(err)
	}
	defer s.db.Close(tx)

	var current *entity.TeacherSubmission
	err = tx.Table(teacherSubmissionTable).
		Select([]string{"teacher_id", "shift_summary_id", "created_at"}).
		Where("teacher_id = ?", submission.TeacherID).
		Where("shift_summary_id = ?", submission.ShiftSummaryID).
		First(&current).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return dbError(err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		submission.CreatedAt = now
		submission.UpdatedAt = now
	} else {
		submission.CreatedAt = current.CreatedAt
		submission.UpdatedAt = now
	}

	err = tx.Table(teacherShiftTable).
		Where("teacher_id = ?", submission.TeacherID).
		Where("shift_summary_id = ?", submission.ShiftSummaryID).
		Delete(&entity.TeacherShift{}).Error
	if err != nil {
		tx.Rollback()
		return dbError(err)
	}

	if len(shifts) > 0 {
		err = tx.Table(teacherShiftTable).Create(&shifts).Error
		if err != nil {
			tx.Rollback()
			return dbError(err)
		}
	}

	err = tx.Clauses(clause.OnConflict{UpdateAll: true}).
		Table(teacherSubmissionTable).
		Create(&submission).Error
	if err != nil {
		tx.Rollback()
		return dbError(err)
	}
	return dbError(tx.Commit().Error)
}
