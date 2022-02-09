package entity

import (
	"fmt"

	"github.com/calmato/shs-web/api/proto/user"
)

type Teacher struct {
	*user.Teacher
}

type Teachers []*Teacher

func NewTeacher(t *user.Teacher) *Teacher {
	return &Teacher{
		Teacher: t,
	}
}

func (t *Teacher) Fullname() string {
	return fmt.Sprintf("%s %s", t.LastName, t.FirstName)
}

func NewTeachers(ts []*user.Teacher) Teachers {
	teachers := make(Teachers, len(ts))
	for i := range ts {
		teachers[i] = NewTeacher(ts[i])
	}
	return teachers
}
