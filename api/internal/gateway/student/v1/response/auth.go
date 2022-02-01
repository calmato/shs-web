package response

import "github.com/calmato/shs-web/api/internal/gateway/student/v1/entity"

type AuthResponse struct {
	*entity.Auth
	Subjects entity.Subjects `json:"subjects"` // 受講科目一覧
}
