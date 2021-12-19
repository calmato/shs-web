package entity

import (
	"strings"
	"time"

	"github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
)

type Subject struct {
	ID        int64     `json:"id"`        // 授業科目ID
	Name      string    `json:"name"`      // 授業科目名
	Color     string    `json:"color"`     // 表示色 (#rrggbb)
	CreatedAt time.Time `json:"createdAt"` // 登録日時
	UpdatedAt time.Time `json:"updatedAt"` // 更新日時
}

type Subjects []*Subject

func NewSubject(subject *entity.Subject) *Subject {
	return &Subject{
		ID:        subject.Id,
		Name:      subject.Name,
		Color:     strings.ToUpper(subject.Color),
		CreatedAt: jst.ParseFromUnix(subject.CreatedAt),
		UpdatedAt: jst.ParseFromUnix(subject.UpdatedAt),
	}
}

func NewSubjects(subjects entity.Subjects) Subjects {
	ss := make(Subjects, len(subjects))
	for i := range subjects {
		ss[i] = NewSubject(subjects[i])
	}
	return ss
}
