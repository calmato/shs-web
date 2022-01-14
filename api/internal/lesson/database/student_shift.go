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

const studentShiftTable = "student_shifts"

var studentShiftFields = []string{
	"student_id", "shift_id", "shift_summary_id", "created_at", "updated_at",
}

type studentShift struct {
	db  *database.Client
	now func() time.Time
}

func NewStudentShift(db *database.Client) StudentShift {
	return &studentShift{
		db:  db,
		now: jst.Now,
	}
}

func (s *studentShift) ListByShiftSummaryID(
	ctx context.Context, studentIDs []string, summaryID int64, fields ...string,
) (entity.StudentShifts, error) {
	var shifts entity.StudentShifts
	if len(fields) == 0 {
		fields = studentShiftFields
	}

	stmt := s.db.DB.Table(studentShiftTable).Select(fields).
		Where("student_id IN (?)", studentIDs).
		Where("shift_summary_id = ?", summaryID)

	err := stmt.Find(&shifts).Error
	return shifts, dbError(err)
}

func (s *studentShift) Replace(
	ctx context.Context, submission *entity.StudentSubmission, shifts entity.StudentShifts,
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

	var current *entity.StudentSubmission
	err = tx.Table(studentSubmissionTable).
		Select([]string{"student_id", "shift_summary_id", "created_at"}).
		Where("student_id = ?", submission.StudentID).
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

	err = tx.Table(studentShiftTable).
		Where("student_id = ?", submission.StudentID).
		Where("shift_summary_id = ?", submission.ShiftSummaryID).
		Delete(&entity.StudentShift{}).Error
	if err != nil {
		tx.Rollback()
		return dbError(err)
	}

	if len(shifts) > 0 {
		err = tx.Table(studentShiftTable).Create(&shifts).Error
		if err != nil {
			tx.Rollback()
			return dbError(err)
		}
	}

	err = tx.Clauses(clause.OnConflict{UpdateAll: true}).
		Table(studentSubmissionTable).
		Create(&submission).Error
	if err != nil {
		tx.Rollback()
		return dbError(err)
	}
	return dbError(tx.Commit().Error)
}
