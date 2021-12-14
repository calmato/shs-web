package response

import "github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"

type TeacherResponse struct {
	*entity.Teacher
}

type TeachersResponse struct {
	Teachers entity.Teachers `json:"teachers"` // 講師一覧
}
