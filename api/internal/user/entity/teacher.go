package entity

import (
	"time"

	"github.com/calmato/shs-web/api/proto/user"
	"gorm.io/gorm"
)

type Teacher struct {
	ID            string         `gorm:"primaryKey;<-:create"`
	LastName      string         `gorm:""`
	FirstName     string         `gorm:""`
	LastNameKana  string         `gorm:""`
	FirstNameKana string         `gorm:""`
	Mail          string         `gorm:""`
	Role          int32          `gorm:""`
	Password      string         `gorm:"-"`
	CreatedAt     time.Time      `gorm:"<-:create"`
	UpdatedAt     time.Time      `gorm:""`
	DeletedAt     gorm.DeletedAt `gorm:"default:null"`
}

type Teachers []*Teacher

func (t *Teacher) Proto() *user.Teacher {
	return &user.Teacher{
		Id:            t.ID,
		LastName:      t.LastName,
		FirstName:     t.FirstName,
		LastNameKana:  t.LastNameKana,
		FirstNameKana: t.FirstNameKana,
		Mail:          t.Mail,
		Role:          user.Role(t.Role),
		CreatedAt:     t.CreatedAt.Unix(),
		UpdatedAt:     t.CreatedAt.Unix(),
	}
}

func (ts Teachers) Proto() []*user.Teacher {
	teachers := make([]*user.Teacher, len(ts))
	for i := range ts {
		teachers[i] = ts[i].Proto()
	}
	return teachers
}
