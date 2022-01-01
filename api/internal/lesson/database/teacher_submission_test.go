package database

import (
	"time"

	"github.com/calmato/shs-web/api/internal/lesson/entity"
)

func testTeacherSubmission(teacherID string, summaryID int64, decided bool, now time.Time) *entity.TeacherSubmission {
	return &entity.TeacherSubmission{
		TeacherID:      teacherID,
		ShiftSummaryID: summaryID,
		Decided:        decided,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
}
