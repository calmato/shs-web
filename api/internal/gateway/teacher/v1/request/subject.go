package request

type CreateSubjectRequest struct {
	Name       string `json:"name,omitempty"`       // 授業科目名
	Color      string `json:"color,omitempty"`      // 表示色
	SchoolType int32  `json:"schoolType,omitempty"` // 校種
}

type UpdateSubjectRequest struct {
	Name       string `json:"name,omitempty"`       // 授業科目名
	Color      string `json:"color,omitempty"`      // 表示色
	SchoolType int32  `json:"schoolType,omitempty"` // 校種
}
