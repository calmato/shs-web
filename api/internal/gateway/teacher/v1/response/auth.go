package response

import "github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"

type AuthResponse struct {
	*entity.Auth
	Subjects map[entity.SchoolType]entity.Subjects `json:"subjects`
}
