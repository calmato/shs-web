package entity

import "github.com/calmato/shs-web/api/proto/user"

type Student struct {
	*user.Student
}

type Students []*Student

func NewStudent(student *user.Student) *Student {
	return &Student{
		Student: student,
	}
}

func NewStudents(students []*user.Student) Students {
	ss := make(Students, len(students))
	for i := range students {
		ss[i] = NewStudent(students[i])
	}
	return ss
}

func (ss Students) IDs() []string {
	ids := make([]string, len(ss))
	for i := range ss {
		ids[i] = ss[i].Id
	}
	return ids
}
