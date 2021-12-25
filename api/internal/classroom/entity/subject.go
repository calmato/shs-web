package entity

import (
	"errors"
	"time"

	"github.com/calmato/shs-web/api/proto/classroom"
)

var errInvalidSchoolType = errors.New("entity: invalid school type")

type SchoolType int32

const (
	SchoolTypeUnknown          SchoolType = 0
	SchoolTypeElementarySchool SchoolType = 1
	SchoolTypeJuniorHighSchool SchoolType = 2
	SchoolTypeHighSchool       SchoolType = 3
)

type Subject struct {
	ID         int64      `gorm:"primaryKey;autoIncrement;<-:create"`
	Name       string     `gorm:""`
	Color      string     `gorm:""`
	SchoolType SchoolType `gorm:""`
	CreatedAt  time.Time  `gorm:"<-:create"`
	UpdatedAt  time.Time  `gorm:""`
}

type Subjects []*Subject

func NewSubject(name, color string, schoolType classroom.SchoolType) *Subject {
	return &Subject{
		Name:       name,
		Color:      color,
		SchoolType: SchoolType(schoolType),
	}
}

func (s *Subject) Proto() *classroom.Subject {
	return &classroom.Subject{
		Id:         s.ID,
		Name:       s.Name,
		Color:      s.Color,
		SchoolType: classroom.SchoolType(s.SchoolType),
		CreatedAt:  s.CreatedAt.Unix(),
		UpdatedAt:  s.UpdatedAt.Unix(),
	}
}

func (ss Subjects) Proto() []*classroom.Subject {
	subjects := make([]*classroom.Subject, len(ss))
	for i := range ss {
		subjects[i] = ss[i].Proto()
	}
	return subjects
}

func NewSchoolType(schoolType classroom.SchoolType) (SchoolType, error) {
	switch schoolType {
	case classroom.SchoolType_SCHOOL_TYPE_ELEMENTARY_SCHOOL:
		return SchoolTypeElementarySchool, nil
	case classroom.SchoolType_SCHOOL_TYPE_JUNIOR_HIGH_SCHOOL:
		return SchoolTypeJuniorHighSchool, nil
	case classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL:
		return SchoolTypeHighSchool, nil
	default:
		return SchoolTypeUnknown, errInvalidSchoolType
	}
}
