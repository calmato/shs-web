package entity

import "github.com/calmato/shs-web/api/proto/classroom"

type TeacherSubject struct {
	*classroom.TeacherSubject
}

type TeacherSubjects []*TeacherSubject

func NewTeacherSubject(subject *classroom.TeacherSubject) *TeacherSubject {
	return &TeacherSubject{
		TeacherSubject: subject,
	}
}

func NewTeacherSubjects(subjects []*classroom.TeacherSubject) TeacherSubjects {
	ss := make(TeacherSubjects, len(subjects))
	for i := range subjects {
		ss[i] = NewTeacherSubject(subjects[i])
	}
	return ss
}
