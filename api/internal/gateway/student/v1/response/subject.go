package response

import "github.com/calmato/shs-web/api/internal/gateway/student/v1/entity"

type SubjectsResponse struct {
	Subjects entity.Subjects `json:"subjects"` // 授業科目一覧
}
