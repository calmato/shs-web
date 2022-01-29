package response

import "github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"

type StudentResponse struct {
	*entity.Student
}

type StudentsResponse struct {
	Students entity.Students `json:"students"` // 生徒一覧
	Total    int64           `json:"total"`    // 生徒合計数
}
