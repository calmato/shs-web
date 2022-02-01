package response

import "github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"

type LessonResponse struct {
	*entity.Lesson
}

type LessonsResponse struct {
	Lessons  entity.Lessons  `json:"lessons"`  // 授業一覧
	Teachers entity.Teachers `json:"teachers"` // 講師一覧
	Students entity.Students `json:"students"` // 生徒一覧
}
