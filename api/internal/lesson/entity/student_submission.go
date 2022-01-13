package entity

import (
	"time"

	"github.com/calmato/shs-web/api/proto/lesson"
)

type StudentSubmission struct {
	StudentID      string    `gorm:"primaryKey;<-:create"`
	ShiftSummaryID int64     `gorm:"primaryKey;<-:create"`
	Decided        bool      `gorm:""`
	CreatedAt      time.Time `gorm:"<-:create"`
	UpdatedAt      time.Time `gorm:""`
}

type StudentSubmissions []*StudentSubmission

func NewStudentSubmission(studentID string, summaryID int64, decided bool) *StudentSubmission {
	return &StudentSubmission{
		StudentID:      studentID,
		ShiftSummaryID: summaryID,
		Decided:        decided,
	}
}

func (s *StudentSubmission) Proto() *lesson.StudentSubmission {
	return &lesson.StudentSubmission{
		StudentId:      s.StudentID,
		ShiftSummaryId: s.ShiftSummaryID,
		Decided:        s.Decided,
		CreatedAt:      s.CreatedAt.Unix(),
		UpdatedAt:      s.UpdatedAt.Unix(),
	}
}

func (ss StudentSubmissions) Proto() []*lesson.StudentSubmission {
	submissions := make([]*lesson.StudentSubmission, len(ss))
	for i := range ss {
		submissions[i] = ss[i].Proto()
	}
	return submissions
}
