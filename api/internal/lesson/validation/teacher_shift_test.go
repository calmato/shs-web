package validation

import (
	"testing"

	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/stretchr/testify/assert"
)

func TestUpsertTeacherShifts(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *lesson.UpsertTeacherShiftsRequest
		isErr bool
	}{
		{
			name: "success",
			req: &lesson.UpsertTeacherShiftsRequest{
				TeacherId:      "teacherid",
				ShiftSummaryId: 1,
				ShiftIds:       []int64{1, 2},
				Desided:        true,
			},
			isErr: false,
		},
		{
			name: "TeacherId is min_len",
			req: &lesson.UpsertTeacherShiftsRequest{
				TeacherId:      "",
				ShiftSummaryId: 1,
				ShiftIds:       []int64{1, 2},
				Desided:        true,
			},
			isErr: true,
		},
		{
			name: "ShiftSummaryId is gt",
			req: &lesson.UpsertTeacherShiftsRequest{
				TeacherId:      "teacherid",
				ShiftSummaryId: 0,
				ShiftIds:       []int64{1, 2},
				Desided:        true,
			},
			isErr: true,
		},
		{
			name: "ShiftIds is unique",
			req: &lesson.UpsertTeacherShiftsRequest{
				TeacherId:      "teacherid",
				ShiftSummaryId: 1,
				ShiftIds:       []int64{1, 1},
				Desided:        true,
			},
			isErr: true,
		},
		{
			name: "ShiftIds is items.gt",
			req: &lesson.UpsertTeacherShiftsRequest{
				TeacherId:      "teacherid",
				ShiftSummaryId: 1,
				ShiftIds:       []int64{0},
				Desided:        true,
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.UpsertTeacherShifts(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}
