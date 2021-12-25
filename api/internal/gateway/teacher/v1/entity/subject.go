package entity

import (
	"errors"
	"time"

	"github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/classroom"
)

var errInvalidSchoolType = errors.New("entity: invalid school type")

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
		Color:      subject.Color,
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

func (ss Subjects) GroupBySchoolType() map[SchoolType]Subjects {
	res := map[SchoolType]Subjects{
		SchoolTypeElementarySchool: make(Subjects, 0),
		SchoolTypeJuniorHighSchool: make(Subjects, 0),
		SchoolTypeHighSchool:       make(Subjects, 0),
	}

	for _, s := range ss {
		if _, ok := res[s.SchoolType]; !ok {
			continue
		}
		res[s.SchoolType] = append(res[s.SchoolType], s)
	}
	return res
}

func (t SchoolType) ClassroomSchoolType() (classroom.SchoolType, error) {
	switch t {
	case SchoolTypeElementarySchool:
		return classroom.SchoolType_SCHOOL_TYPE_ELEMENTARY_SCHOOL, nil
	case SchoolTypeJuniorHighSchool:
		return classroom.SchoolType_SCHOOL_TYPE_JUNIOR_HIGH_SCHOOL, nil
	case SchoolTypeHighSchool:
		return classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL, nil
	default:
		return classroom.SchoolType_SCHOOL_TYPE_UNKNOWN, errInvalidSchoolType
	}
}
