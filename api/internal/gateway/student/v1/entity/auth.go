package entity

import (
	"github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/proto/user"
)

type Auth struct {
	ID            string     `json:"id"`            // 生徒ID
	LastName      string     `json:"lastName"`      // 姓
	FirstName     string     `json:"firstName"`     // 名
	LastNameKana  string     `json:"lastNameKana"`  // 姓(かな)
	FirstNameKana string     `json:"firstNameKana"` // 名(かな)
	Mail          string     `json:"mail"`          // メールアドレス
	SchoolType    SchoolType `json:"schoolType"`    // 校種
	Grade         int64      `json:"grade"`         // 学年
}

func NewAuth(student *entity.Student) *Auth {
	return &Auth{
		ID:            student.Id,
		LastName:      student.LastName,
		FirstName:     student.FirstName,
		LastNameKana:  student.LastNameKana,
		FirstNameKana: student.FirstNameKana,
		Mail:          student.Mail,
		SchoolType:    NewSchoolTypeFromUser(student.SchoolType),
		Grade:         student.Grade,
	}
}

func NewSchoolTypeFromUser(schoolType user.SchoolType) SchoolType {
	switch schoolType {
	case user.SchoolType_SCHOOL_TYPE_ELEMENTARY_SCHOOL:
		return SchoolTypeElementarySchool
	case user.SchoolType_SCHOOL_TYPE_JUNIOR_HIGH_SCHOOL:
		return SchoolTypeJuniorHighSchool
	case user.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL:
		return SchoolTypeHighSchool
	default:
		return SchoolTypeUnknown
	}
}
