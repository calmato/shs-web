package entity

import "github.com/calmato/shs-web/api/internal/gateway/entity"

type Auth struct {
	ID            string `json:"id"`            // 講師ID
	LastName      string `json:"lastName"`      // 姓
	FirstName     string `json:"firstName"`     // 名
	LastNameKana  string `json:"lastNameKana"`  // 姓(かな)
	FirstNameKana string `json:"firstNameKana"` // 名(かな)
	Mail          string `json:"mail"`          // メールアドレス
	Role          Role   `json:"role"`          // 権限
}

func NewAuth(teacher *entity.Teacher) *Auth {
	return &Auth{
		ID:            teacher.Id,
		LastName:      teacher.LastName,
		FirstName:     teacher.FirstName,
		LastNameKana:  teacher.LastNameKana,
		FirstNameKana: teacher.FirstNameKana,
		Mail:          teacher.Mail,
		Role:          NewRole(teacher.Role),
	}
}
