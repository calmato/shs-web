package database

import (
	"context"
	"time"

	"github.com/calmato/shs-web/api/internal/lesson/entity"
	"github.com/calmato/shs-web/api/pkg/database"
	"github.com/calmato/shs-web/api/pkg/jst"
)

const studentSubmissionTable = "student_submissions"

var studentSubmissionFields = []string{
	"student_id", "shift_summary_id", "decided",
	"suggested_lessons", "created_at", "updated_at",
}

type studentSubmission struct {
	db  *database.Client
	now func() time.Time
}

func NewStudentSubmission(db *database.Client) StudentSubmission {
	return &studentSubmission{
		db:  db,
		now: jst.Now,
	}
}

func (s *studentSubmission) ListByShiftSummaryIDs(
	ctx context.Context, studentID string, summaryIDs []int64, fields ...string,
) (entity.StudentSubmissions, error) {
	var submissions entity.StudentSubmissions
	if len(fields) == 0 {
		fields = studentSubmissionFields
	}

	stmt := s.db.DB.Table(studentSubmissionTable).Select(fields).
		Where("student_id = ?", studentID).
		Where("shift_summary_id IN (?)", summaryIDs)

	err := stmt.Find(&submissions).Error
	if err != nil {
		return nil, dbError(err)
	}
	err = submissions.Fill()
	if err != nil {
		return nil, dbError(err)
	}
	return submissions, nil
}

func (s *studentSubmission) ListByStudentIDs(
	ctx context.Context, studentIDs []string, summaryID int64, fields ...string,
) (entity.StudentSubmissions, error) {
	var submissions entity.StudentSubmissions
	if len(fields) == 0 {
		fields = studentSubmissionFields
	}

	stmt := s.db.DB.Table(studentSubmissionTable).Select(fields).
		Where("student_id IN (?)", studentIDs).
		Where("shift_summary_id = ?", summaryID)

	err := stmt.Find(&submissions).Error
	if err != nil {
		return nil, dbError(err)
	}
	err = submissions.Fill()
	if err != nil {
		return nil, dbError(err)
	}
	return submissions, nil
}

func (s *studentSubmission) Get(
	ctx context.Context, studentID string, summaryID int64, fields ...string,
) (*entity.StudentSubmission, error) {
	var submission *entity.StudentSubmission
	if len(fields) == 0 {
		fields = studentSubmissionFields
	}

	stmt := s.db.DB.Table(studentSubmissionTable).Select(fields).
		Where("student_id = ?", studentID).
		Where("shift_summary_id = ?", summaryID)

	err := stmt.First(&submission).Error
	if err != nil {
		return nil, dbError(err)
	}
	err = submission.Fill()
	if err != nil {
		return nil, dbError(err)
	}
	return submission, nil
}
