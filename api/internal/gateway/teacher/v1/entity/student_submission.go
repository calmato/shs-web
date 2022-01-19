package entity

import (
	"github.com/calmato/shs-web/api/internal/gateway/entity"
)

type StudentSubmissionDetail struct {
	Student               *Student `json:"student"`               // 生徒情報
	LessonTotal           int64    `json:"lessonTotal"`           // 登録受講授業数
	SuggestedClassesTotal int64    `json:"suggestedClassesTotal"` // 希望受講授業数
}

type StudentSubmissionDetails []*StudentSubmissionDetail

func NewStudentSubmissionDetail(
	student *entity.Student, submission *entity.StudentSubmission, lessons entity.Lessons,
) *StudentSubmissionDetail {
	var suggestedClassesTotal int64
	if submission != nil {
		suggestedClassesTotal = submission.SuggestedClasses
	}
	return &StudentSubmissionDetail{
		Student:               NewStudent(student),
		LessonTotal:           int64(len(lessons)),
		SuggestedClassesTotal: suggestedClassesTotal,
	}
}

func NewStudentSubmissionDetails(
	students entity.Students,
	submissionMap map[string]*entity.StudentSubmission,
	lessonsMap map[string]entity.Lessons,
) StudentSubmissionDetails {
	ss := make(StudentSubmissionDetails, len(students))
	for i, student := range students {
		submission := submissionMap[student.Id]
		lessons := lessonsMap[student.Id]
		ss[i] = NewStudentSubmissionDetail(student, submission, lessons)
	}
	return ss
}
