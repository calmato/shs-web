package entity

import (
	"time"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
)

type StudentSubmission struct {
	StudentID        string
	ShiftSummaryID   int64
	Decided          bool
	SuggestedLessons StudentSuggestedLessons
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type StudentSubmissions []*StudentSubmission

func NewStudentSubmission(submission *lesson.StudentSubmission) *StudentSubmission {
	return &StudentSubmission{
		StudentID:        submission.StudentId,
		ShiftSummaryID:   submission.ShiftSummaryId,
		Decided:          submission.Decided,
		SuggestedLessons: NewStudentSuggestedLessons(submission.SuggestedLessons),
		CreatedAt:        jst.ParseFromUnix(submission.CreatedAt),
		UpdatedAt:        jst.ParseFromUnix(submission.UpdatedAt),
	}
}

func NewStudentSubmissions(submissions []*lesson.StudentSubmission) StudentSubmissions {
	ss := make(StudentSubmissions, len(submissions))
	for i := range submissions {
		ss[i] = NewStudentSubmission(submissions[i])
	}
	return ss
}

func (ss StudentSubmissions) MapByShiftSummaryID() map[int64]*StudentSubmission {
	res := make(map[int64]*StudentSubmission, len(ss))
	for _, s := range ss {
		res[s.ShiftSummaryID] = s
	}
	return res
}

func (ss StudentSubmissions) MapByStudentID() map[string]*StudentSubmission {
	res := make(map[string]*StudentSubmission, len(ss))
	for _, s := range ss {
		res[s.StudentID] = s
	}
	return res
}
