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
	student *entity.Student, submission *entity.StudentSubmission, shifts entity.StudentShifts,
) *StudentSubmissionDetail {
	var suggestedClassesTotal int64
	if submission != nil {
		suggestedClassesTotal = submission.SuggestedClasses
	}
	return &StudentSubmissionDetail{
		Student:               NewStudent(student),
		LessonTotal:           int64(len(shifts)),
		SuggestedClassesTotal: suggestedClassesTotal,
	}
}

func NewStudentSubmissionDetails(
	students entity.Students,
	submissionMap map[string]*entity.StudentSubmission,
	shiftsMap map[string]entity.StudentShifts,
) StudentSubmissionDetails {
	ss := make(StudentSubmissionDetails, len(students))
	for i, student := range students {
		submission := submissionMap[student.Id]
		shifts := shiftsMap[student.Id]
		ss[i] = NewStudentSubmissionDetail(student, submission, shifts)
	}
	return ss
}
