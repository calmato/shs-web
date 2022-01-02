package response

import "github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"

type TeacherSubmissionsResponse struct {
	Summaries entity.TeacherSubmissions `json:"summaries"` // シフト募集概要一覧
}
