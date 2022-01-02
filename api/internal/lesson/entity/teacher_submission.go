package entity

import (
	"time"

	"github.com/calmato/shs-web/api/proto/lesson"
)

type TeacherSubmission struct {
	TeacherID      string    `gorm:"primaryKey;<-:create"`
	ShiftSummaryID int64     `gorm:"primaryKey;<-:create"`
	Decided        bool      `gorm:""`
	CreatedAt      time.Time `gorm:"<-:create"`
	UpdatedAt      time.Time `gorm:""`
}

type TeacherSubmissions []*TeacherSubmission

func NewTeacherSubmission(teacherID string, summaryID int64, decided bool) *TeacherSubmission {
	return &TeacherSubmission{
		TeacherID:      teacherID,
		ShiftSummaryID: summaryID,
		Decided:        decided,
	}
}

func (s *TeacherSubmission) Proto() *lesson.TeacherSubmission {
	return &lesson.TeacherSubmission{
		TeacherId:      s.TeacherID,
		ShiftSummaryId: s.ShiftSummaryID,
		Decided:        s.Decided,
		CreatedAt:      s.CreatedAt.Unix(),
		UpdatedAt:      s.UpdatedAt.Unix(),
	}
}

func (ss TeacherSubmissions) Proto() []*lesson.TeacherSubmission {
	submissions := make([]*lesson.TeacherSubmission, len(ss))
	for i := range ss {
		submissions[i] = ss[i].Proto()
	}
	return submissions
}
