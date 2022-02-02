package response

import "github.com/calmato/shs-web/api/internal/gateway/student/v1/entity"

type LessonsResponse struct {
	Lessons  entity.Lessons  `json:"lessons"`  // 授業一覧
	Teachers entity.Teachers `json:"teachers"` // 講師一覧
}
