package entity

import "github.com/calmato/shs-web/api/proto/classroom"

type StudentSubject struct {
	*classroom.StudentSubject
}

type StudentSubjects []*StudentSubject

func NewStudentSubject(subject *classroom.StudentSubject) *StudentSubject {
	return &StudentSubject{
		StudentSubject: subject,
	}
}

func NewStudentSubjects(subjects []*classroom.StudentSubject) StudentSubjects {
	ss := make(StudentSubjects, len(subjects))
	for i := range subjects {
		ss[i] = NewStudentSubject(subjects[i])
	}
	return ss
}
