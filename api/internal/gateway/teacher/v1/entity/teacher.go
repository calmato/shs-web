package entity

import (
	"errors"
	"time"

	"github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/user"
)

var errInvalidRole = errors.New("entity: invalid role")

type Teacher struct {
	ID            string    `json:"id"`            // 講師ID
	LastName      string    `json:"lastName"`      // 姓
	FirstName     string    `json:"firstName"`     // 名
	LastNameKana  string    `json:"lastNameKana"`  // 姓(かな)
	FirstNameKana string    `json:"firstNameKana"` // 名(かな)
	Mail          string    `json:"mail"`          // メールアドレス
	Role          Role      `json:"role"`          // 権限
	CreatedAt     time.Time `json:"createdAt"`     // 登録日時
	UpdatedAt     time.Time `json:"updatedAt"`     // 更新日時
}

type Teachers []*Teacher

type Role int32

const (
	RoleUnknown       Role = 0
	RoleTeacher       Role = 1
	RoleAdministrator Role = 2
)

func NewTeacher(teacher *entity.Teacher) *Teacher {
	return &Teacher{
		ID:            teacher.Id,
		LastName:      teacher.LastName,
		FirstName:     teacher.FirstName,
		LastNameKana:  teacher.LastNameKana,
		FirstNameKana: teacher.FirstNameKana,
		Mail:          teacher.Mail,
		Role:          NewRole(teacher.Role),
		CreatedAt:     jst.ParseFromUnix(teacher.CreatedAt),
		UpdatedAt:     jst.ParseFromUnix(teacher.UpdatedAt),
	}
}

func NewTeachers(teachers entity.Teachers) Teachers {
	ts := make(Teachers, len(teachers))
	for i := range teachers {
		ts[i] = NewTeacher(teachers[i])
	}
	return ts
}

func NewRole(role user.Role) Role {
	switch role {
	case user.Role_ROLE_TEACHER:
		return RoleTeacher
	case user.Role_ROLE_ADMINISTRATOR:
		return RoleAdministrator
	default:
		return RoleUnknown
	}
}

func (r Role) UserRole() (user.Role, error) {
	switch r {
	case RoleAdministrator:
		return user.Role_ROLE_ADMINISTRATOR, nil
	case RoleTeacher:
		return user.Role_ROLE_TEACHER, nil
	default:
		return user.Role_ROLE_TEACHER, errInvalidRole
	}
}
