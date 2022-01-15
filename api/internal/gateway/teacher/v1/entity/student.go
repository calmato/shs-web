package entity

import (
	"time"

	"github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/user"
)

type Student struct {
	ID            string     `json:"id"`            // 講師ID
	LastName      string     `json:"lastName"`      // 姓
	FirstName     string     `json:"firstName"`     // 名
	LastNameKana  string     `json:"lastNameKana"`  // 姓(かな)
	FirstNameKana string     `json:"firstNameKana"` // 名(かな)
	Mail          string     `json:"mail"`          // メールアドレス
	SchoolType    SchoolType `json:"schoolType"`    // 校種
	Grade         int64      `json:"grade"`         // 学年
	CreatedAt     time.Time  `json:"createdAt"`     // 登録日時
	UpdatedAt     time.Time  `json:"updatedAt"`     // 更新日時
}

type Students []*Student

func NewStudent(student *entity.Student) *Student {
	return &Student{
		ID:            student.Id,
		LastName:      student.LastName,
		FirstName:     student.FirstName,
		LastNameKana:  student.LastNameKana,
		FirstNameKana: student.FirstNameKana,
		Mail:          student.Mail,
		SchoolType:    NewSchoolTypeFromUser(student.SchoolType),
		Grade:         student.Grade,
		CreatedAt:     jst.ParseFromUnix(student.CreatedAt),
		UpdatedAt:     jst.ParseFromUnix(student.UpdatedAt),
	}
}

func NewStudents(students entity.Students) Students {
	ts := make(Students, len(students))
	for i := range students {
		ts[i] = NewStudent(students[i])
	}
	return ts
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