package entity

import (
	"fmt"

	"github.com/calmato/shs-web/api/proto/user"
)

type Student struct {
	*user.Student
}

type Students []*Student

func NewStudent(t *user.Student) *Student {
	return &Student{
		Student: t,
	}
}

func (s *Student) Fullname() string {
	return fmt.Sprintf("%s %s", s.LastName, s.FirstName)
}

func NewStudents(ss []*user.Student) Students {
	students := make(Students, len(ss))
	for i := range ss {
		students[i] = NewStudent(ss[i])
	}
	return students
}
