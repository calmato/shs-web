package entity

import (
	"strings"
	"time"

	"github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/classroom"
)

type Subject struct {
	ID         int64      `json:"id"`        // 授業科目ID
	Name       string     `json:"name"`      // 授業科目名
	Color      string     `json:"color"`     // 表示色 (#rrggbb)
	SchoolType SchoolType `json:"shoolType"` // 校種
	CreatedAt  time.Time  `json:"createdAt"` // 登録日時
	UpdatedAt  time.Time  `json:"updatedAt"` // 更新日時
}

type Subjects []*Subject

type SchoolType int32

const (
	SchoolTypeUnknown          SchoolType = 0
	SchoolTypeElementarySchool SchoolType = 1
	SchoolTypeJuniorHighSchool SchoolType = 2
	SchoolTypeHighSchool       SchoolType = 3
)

func NewSubject(subject *entity.Subject) *Subject {
	return &Subject{
		ID:         subject.Id,
		Name:       subject.Name,
		Color:      strings.ToUpper(subject.Color),
		SchoolType: NewSchoolType(subject.SchoolType),
		CreatedAt:  jst.ParseFromUnix(subject.CreatedAt),
		UpdatedAt:  jst.ParseFromUnix(subject.UpdatedAt),
	}
}

func NewSubjects(subjects entity.Subjects) Subjects {
	ss := make(Subjects, len(subjects))
	for i := range subjects {
		ss[i] = NewSubject(subjects[i])
	}
	return ss
}

func NewSchoolType(schoolType classroom.SchoolType) SchoolType {
	switch schoolType {
	case classroom.SchoolType_SCHOOL_TYPE_ELEMENTARY_SCHOOL:
		return SchoolTypeElementarySchool
	case classroom.SchoolType_SCHOOL_TYPE_JUNIOR_HIGH_SCHOOL:
		return SchoolTypeJuniorHighSchool
	case classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL:
		return SchoolTypeHighSchool
	default:
		return SchoolTypeUnknown
	}
}
