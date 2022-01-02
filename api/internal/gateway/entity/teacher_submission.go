package entity

import "github.com/calmato/shs-web/api/proto/lesson"

type TeacherSubmission struct {
	*lesson.TeacherSubmission
}

type TeacherSubmissions []*TeacherSubmission

func NewTeacherSubmission(submission *lesson.TeacherSubmission) *TeacherSubmission {
	return &TeacherSubmission{
		TeacherSubmission: submission,
	}
}

func NewTeacherSubmissions(submissions []*lesson.TeacherSubmission) TeacherSubmissions {
	ss := make(TeacherSubmissions, len(submissions))
	for i := range submissions {
		ss[i] = NewTeacherSubmission(submissions[i])
	}
	return ss
}

func (ss TeacherSubmissions) MapByShiftSummaryID() map[int64]*TeacherSubmission {
	res := make(map[int64]*TeacherSubmission, len(ss))
	for _, s := range ss {
		res[s.ShiftSummaryId] = s
	}
	return res
}
