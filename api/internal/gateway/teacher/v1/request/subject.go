package request

import "github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"

type CreateSubjectRequest struct {
	Name       string            `json:"name,omitempty"`       // 授業科目名
	Color      string            `json:"color,omitempty"`      // 表示色
	SchoolType entity.SchoolType `json:"schoolType,omitempty"` // 校種
}

type UpdateSubjectRequest struct {
	Name       string            `json:"name,omitempty"`       // 授業科目名
	Color      string            `json:"color,omitempty"`      // 表示色
	SchoolType entity.SchoolType `json:"schoolType,omitempty"` // 校種
}
