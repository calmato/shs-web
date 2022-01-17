package response

import "github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"

type LessonResponse struct {
	*entity.Lesson
}
