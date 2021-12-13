package entity

import "github.com/calmato/shs-web/api/proto/user"

type Teacher struct {
	*user.Teacher
}

func NewTeacher(teacher *user.Teacher) *Teacher {
	return &Teacher{
		Teacher: teacher,
	}
}
