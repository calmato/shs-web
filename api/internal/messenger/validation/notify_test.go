package validation

import (
	"testing"

	"github.com/calmato/shs-web/api/proto/messenger"
	"github.com/stretchr/testify/assert"
)

func TestLessonDecided(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *messenger.NotifyLessonDecidedRequest
		isErr bool
	}{
		{
			name: "success",
			req: &messenger.NotifyLessonDecidedRequest{
				ShiftSummaryId: 1,
			},
			isErr: false,
		},
		{
			name: "ShiftSummaryId is gt",
			req: &messenger.NotifyLessonDecidedRequest{
				ShiftSummaryId: 0,
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.NotifyLessonDecided(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}
