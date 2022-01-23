package handler

import (
	"context"

	gentity "github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/proto/user"
)

func (h *apiV1Handler) multiGetStudents(ctx context.Context, studentIDs []string) (gentity.Students, error) {
	in := &user.MultiGetStudentsRequest{
		Ids: studentIDs,
	}
	out, err := h.user.MultiGetStudents(ctx, in)
	if err != nil {
		return nil, err
	}
	return gentity.NewStudents(out.Students), nil
}
