package database

import (
	"context"
	"time"

	"github.com/calmato/shs-web/api/internal/lesson/entity"
	"github.com/calmato/shs-web/api/pkg/database"
	"github.com/calmato/shs-web/api/pkg/jst"
)

const teacherSubmissionTable = "teacher_submissions"

var teacherSubmissionFields = []string{
	"teacher_id", "shift_summary_id", "decided", "created_at", "updated_at",
}

type teacherSubmission struct {
	db  *database.Client
	now func() time.Time
}

func NewTeacherSubmission(db *database.Client) TeacherSubmission {
	return &teacherSubmission{
		db:  db,
		now: jst.Now,
	}
}

func (s *teacherSubmission) ListByShiftSummaryIDs(
	ctx context.Context, teacherID string, summaryIDs []int64, fields ...string,
) (entity.TeacherSubmissions, error) {
	var submissions entity.TeacherSubmissions
	if len(fields) == 0 {
		fields = teacherSubmissionFields
	}

	stmt := s.db.DB.Table(teacherSubmissionTable).Select(fields).
		Where("teacher_id = ?", teacherID).
		Where("shift_summary_id IN (?)", summaryIDs)

	err := stmt.Find(&submissions).Error
	return submissions, dbError(err)
}

func (s *teacherSubmission) ListByTeacherIDs(
	ctx context.Context, teacherIDs []string, summaryID int64, fields ...string,
) (entity.TeacherSubmissions, error) {
	var submissions entity.TeacherSubmissions
	if len(fields) == 0 {
		fields = teacherSubmissionFields
	}

	stmt := s.db.DB.Table(teacherSubmissionTable).Select(fields).
		Where("teacher_id IN (?)", teacherIDs).
		Where("shift_summary_id = ?", summaryID)

	err := stmt.Find(&submissions).Error
	return submissions, dbError(err)
}

func (s *teacherSubmission) Get(
	ctx context.Context, teacherID string, summaryID int64, fields ...string,
) (*entity.TeacherSubmission, error) {
	var submission *entity.TeacherSubmission
	if len(fields) == 0 {
		fields = teacherSubmissionFields
	}

	stmt := s.db.DB.Table(teacherSubmissionTable).Select(fields).
		Where("teacher_id = ?", teacherID).
		Where("shift_summary_id = ?", summaryID)

	err := stmt.First(&submission).Error
	return submission, dbError(err)
}
