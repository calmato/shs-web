package entity

import "github.com/calmato/shs-web/api/internal/gateway/entity"

type Teacher struct {
	ID            string `json:"id"`            // 講師ID
	LastName      string `json:"lastName"`      // 姓
	FirstName     string `json:"firstName"`     // 名
	LastNameKana  string `json:"lastNameKana"`  // 姓(かな)
	FirstNameKana string `json:"firstNameKana"` // 名(かな)
}

type Teachers []*Teacher

func NewTeacher(teacher *entity.Teacher) *Teacher {
	return &Teacher{
		ID:            teacher.Id,
		LastName:      teacher.LastName,
		FirstName:     teacher.FirstName,
		LastNameKana:  teacher.LastNameKana,
		FirstNameKana: teacher.FirstNameKana,
	}
}

func NewTeachers(teachers entity.Teachers) Teachers {
	ts := make(Teachers, len(teachers))
	for i := range teachers {
		ts[i] = NewTeacher(teachers[i])
	}
	return ts
}
