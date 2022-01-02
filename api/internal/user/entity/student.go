package entity

import (
	"jst"
	"time"

	"github.com/calmato/shs-web/api/proto/user"
	"gorm.io/gorm"
)

type SchoolType int32

const (
	SchoolTypeUnknown          SchoolType = 0
	SchoolTypeElementarySchool SchoolType = 1
	SchoolTypeJuniorHighSchool SchoolType = 2
	SchoolTypeHighSchool       SchoolType = 3
)

type Student struct {
	ID            string         `gorm:"primaryKey;<-:create"`
	LastName      string         `gorm:""`
	FirstName     string         `gorm:""`
	LastNameKana  string         `gorm:""`
	FirstNameKana string         `gorm:""`
	Mail          string         `gorm:""`
	BirthYear     int64          `gorm:""`
	Schooltype    SchoolType     `gorm:"-"`
	Grade         int64          `gorm:"-"`
	Password      string         `gorm:"-"`
	CreatedAt     time.Time      `gorm:"<-:create"`
	UpdatedAt     time.Time      `gorm:""`
	DeletedAt     gorm.DeletedAt `gorm:"default:null"`
}

type Students []*Student

func (s *Student) Fill(now time.Time) {
	year := jst.FisicalYear(now)
	age := year - s.BirthYear
	switch {
	case age >= 7 && age <= 12:
		s.Schooltype = SchoolTypeElementarySchool
		s.Grade = age - 6
	case age >= 13 && age <= 15:
		s.Schooltype = SchoolTypeJuniorHighSchool
		s.Grade = age - 12
	case age >= 16 && age <= 18:
		s.Schooltype = SchoolTypeHighSchool
		s.Grade = age - 15
	default:
		s.Schooltype = SchoolTypeUnknown
		s.Grade = 0
	}
}

func (s *Student) Proto() *user.Student {
	return &user.Student{
		Id:            s.ID,
		LastName:      s.LastName,
		FirstName:     s.FirstName,
		LastNameKana:  s.LastNameKana,
		FirstNameKana: s.FirstNameKana,
		Mail:          s.Mail,
		BirthYear:     s.BirthYear,
		SchoolType:    user.SchoolType(s.Schooltype),
		Grade:         s.Grade,
		CreatedAt:     s.CreatedAt.Unix(),
		UpdatedAt:     s.CreatedAt.Unix(),
	}
}

func (ss Students) Fill(now time.Time) {
	for i := range ss {
		ss[i].Fill(now)
	}
}

func (ss Students) Proto() []*user.Student {
	students := make([]*user.Student, len(ss))
	for i := range ss {
		students[i] = ss[i].Proto()
	}
	return students
}
