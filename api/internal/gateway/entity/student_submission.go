package entity

import "github.com/calmato/shs-web/api/proto/lesson"

type StudentSubmission struct {
	*lesson.StudentSubmission
}

type StudentSubmissions []*StudentSubmission

func NewStudentSubmission(submission *lesson.StudentSubmission) *StudentSubmission {
	return &StudentSubmission{
		StudentSubmission: submission,
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
		res[s.ShiftSummaryId] = s
	}
	return res
}

func (ss StudentSubmissions) MapByStudentID() map[string]*StudentSubmission {
	res := make(map[string]*StudentSubmission, len(ss))
	for _, s := range ss {
		res[s.StudentId] = s
	}
	return res
}
