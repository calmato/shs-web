package entity

import "github.com/calmato/shs-web/api/proto/user"

type Teacher struct {
	*user.Teacher
}

type Teachers []*Teacher

func NewTeacher(teacher *user.Teacher) *Teacher {
	return &Teacher{
		Teacher: teacher,
	}
}

func (t *Teacher) AdminRole() bool {
	return t.Role == user.Role_ROLE_ADMINISTRATOR
}

func NewTeachers(teachers []*user.Teacher) Teachers {
	ts := make(Teachers, len(teachers))
	for i := range teachers {
		ts[i] = NewTeacher(teachers[i])
	}
	return ts
}

func (ts Teachers) IDs() []string {
	ids := make([]string, len(ts))
	for i := range ts {
		ids[i] = ts[i].Id
	}
	return ids
}
