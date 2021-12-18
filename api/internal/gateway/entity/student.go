package entity

import (
	"time"

	"github.com/calmato/shs-web/api/proto/user"
	"gorm.io/gorm"
)

type Student struct {
	ID            string         `gorm:"primaryKey;<-:create"`
	LastName      string         `gorm:""`
	FirstName     string         `gorm:""`
	LastNameKana  string         `gorm:""`
	FirstNameKana string         `gorm:""`
	Mail          string         `gorm:""`
	BirthYear     int64          `gorm:""`
	Password      string         `gorm:"-"`
	CreatedAt     time.Time      `gorm:"<-:create"`
	UpdatedAt     time.Time      `gorm:""`
	DeletedAt     gorm.DeletedAt `gorm:"default:null"`
}

type Students []*Student

func (s *Student) Proto() *user.Student {
	return &user.Student{
		Id:            s.ID,
		LastName:      s.LastName,
		FirstName:     s.FirstName,
		LastNameKana:  s.LastNameKana,
		FirstNameKana: s.FirstNameKana,
		Mail:          s.Mail,
		BirthYear:     s.BirthYear,
		CreatedAt:     s.CreatedAt.Unix(),
		UpdatedAt:     s.CreatedAt.Unix(),
	}
}

func (ss Students) Proto() []*user.Student {
	students := make([]*user.Student, len(ss))
	for i := range ss {
		students[i] = ss[i].Proto()
	}
	return students
}
