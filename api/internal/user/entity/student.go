package entity

import (
	"time"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/pkg/uuid"
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
	SchoolType    SchoolType     `gorm:"-"`
	Grade         int64          `gorm:"-"`
	Password      string         `gorm:"-"`
	CreatedAt     time.Time      `gorm:"<-:create"`
	UpdatedAt     time.Time      `gorm:""`
	DeletedAt     gorm.DeletedAt `gorm:"default:null"`
}

type Students []*Student

func NewStudent(
	lastName, firstName, lastNameKana, firstNameKana, mail, password string,
	schoolType SchoolType, grade int64, now time.Time,
) *Student {
	uid := uuid.Base58Encode(uuid.New())
	return &Student{
		ID:            uid,
		LastName:      lastName,
		FirstName:     firstName,
		LastNameKana:  lastNameKana,
		FirstNameKana: firstNameKana,
		Mail:          mail,
		BirthYear:     newStudentBirthYear(schoolType, grade, now),
		SchoolType:    schoolType,
		Grade:         grade,
		Password:      password,
	}
}

func newStudentBirthYear(schoolType SchoolType, grade int64, now time.Time) int64 {
	year := int64(jst.FiscalYear(now))
	switch schoolType {
	case SchoolTypeElementarySchool:
		return year - (grade + 6)
	case SchoolTypeJuniorHighSchool:
		return year - (grade + 12)
	case SchoolTypeHighSchool:
		return year - (grade + 15)
	default:
		return 0
	}
}

func (s *Student) Fill(now time.Time) {
	year := int64(jst.FiscalYear(now))
	age := year - s.BirthYear
	switch {
	case age > 6 && age <= 12:
		s.SchoolType = SchoolTypeElementarySchool
		s.Grade = age - 6
	case age > 12 && age <= 15:
		s.SchoolType = SchoolTypeJuniorHighSchool
		s.Grade = age - 12
	case age > 15 && age <= 18:
		s.SchoolType = SchoolTypeHighSchool
		s.Grade = age - 15
	default:
		s.SchoolType = SchoolTypeUnknown
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
		SchoolType:    user.SchoolType(s.SchoolType),
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
