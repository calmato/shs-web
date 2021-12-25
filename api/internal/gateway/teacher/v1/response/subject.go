package response

import "github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"

type SubjectResponse struct {
	*entity.Subject
}

type SubjectsResponse struct {
	Subjects entity.Subjects `json:"subjects"` // 授業科目一覧
}
