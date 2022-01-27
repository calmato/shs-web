package response

import "github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"

type LessonResponse struct {
	*entity.Lesson
}

type LessonsResponse struct {
	Lessons entity.Lessons `json:"lessons"` // 授業一覧
	Total   int64          `json:"total"`   // 授業合計数
}
