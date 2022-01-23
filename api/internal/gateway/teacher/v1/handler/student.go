package handler

import (
	"context"

	gentity "github.com/calmato/shs-web/api/internal/gateway/entity"
)

func (h *apiV1Handler) multiGetStudents(ctx context.Context, studentIDs []string) (gentity.Students, error) {
	// TODO: 実装
	return gentity.Students{}, nil
}
