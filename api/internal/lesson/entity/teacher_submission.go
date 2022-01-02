package entity

import (
	"time"

	"github.com/calmato/shs-web/api/proto/lesson"
)

type TeacherSubmissionStatus int32

const (
	TeacherSubmissionStatusUnknown   TeacherSubmissionStatus = 0
	TeacherSubmissionStatusWaiting   TeacherSubmissionStatus = 1
	TeacherSubmissionStatusEditing   TeacherSubmissionStatus = 2
	TeacherSubmissionStatusSubmitted TeacherSubmissionStatus = 3
)

type TeacherSubmission struct {
	TeacherID      string                  `gorm:"primaryKey;<-:create"`
	ShiftSummaryID int64                   `gorm:"primaryKey;<-:create"`
	Decided        bool                    `gorm:""`
	Status         TeacherSubmissionStatus `gorm:"-"`
	CreatedAt      time.Time               `gorm:"<-:create"`
	UpdatedAt      time.Time               `gorm:""`
}

type TeacherSubmissions []*TeacherSubmission

func NewTeacherSubmission(teacherID string, summaryID int64, decided bool) *TeacherSubmission {
	return &TeacherSubmission{
		TeacherID:      teacherID,
		ShiftSummaryID: summaryID,
		Decided:        decided,
	}
}

func (s *TeacherSubmission) Fill(shifts TeacherShifts) {
	if s.Decided {
		s.Status = TeacherSubmissionStatusSubmitted
		return
	}
	if len(shifts) > 0 {
		s.Status = TeacherSubmissionStatusEditing
		return
	}
	s.Status = TeacherSubmissionStatusWaiting
}

func (s *TeacherSubmission) Proto() *lesson.TeacherSubmission {
	return &lesson.TeacherSubmission{
		TeacherId:      s.TeacherID,
		ShiftSummaryId: s.ShiftSummaryID,
		Decided:        s.Decided,
		Status:         lesson.TeacherSubmissionStatus(s.Status),
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
