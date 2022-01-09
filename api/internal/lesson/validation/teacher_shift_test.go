package validation

import (
	"testing"

	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/stretchr/testify/assert"
)

func TestListTeacherSubmissionsByShiftSummaryIDs(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *lesson.ListTeacherSubmissionsByShiftSummaryIDsRequest
		isErr bool
	}{
		{
			name: "success",
			req: &lesson.ListTeacherSubmissionsByShiftSummaryIDsRequest{
				TeacherId:       "teacherid",
				ShiftSummaryIds: []int64{1, 2},
			},
			isErr: false,
		},
		{
			name: "TeacherId is min_len",
			req: &lesson.ListTeacherSubmissionsByShiftSummaryIDsRequest{
				TeacherId:       "",
				ShiftSummaryIds: []int64{1, 2},
			},
			isErr: true,
		},
		{
			name: "ShiftSummaryIds is unique",
			req: &lesson.ListTeacherSubmissionsByShiftSummaryIDsRequest{
				TeacherId:       "teacherid",
				ShiftSummaryIds: []int64{1, 1},
			},
			isErr: true,
		},
		{
			name: "ShiftSummaryIds is items.gt",
			req: &lesson.ListTeacherSubmissionsByShiftSummaryIDsRequest{
				TeacherId:       "teacherid",
				ShiftSummaryIds: []int64{0},
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.ListTeacherSubmissionsByShiftSummaryIDs(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestListTeacherShifts(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *lesson.ListTeacherShiftsRequest
		isErr bool
	}{
		{
			name: "success",
			req: &lesson.ListTeacherShiftsRequest{
				TeacherIds:     []string{"teacherid"},
				ShiftSummaryId: 1,
			},
			isErr: false,
		},
		{
			name: "TeacherIds is unique",
			req: &lesson.ListTeacherShiftsRequest{
				TeacherIds:     []string{"teacherid", "teacherid"},
				ShiftSummaryId: 1,
			},
			isErr: true,
		},
		{
			name: "TeacherIds is items.min_len",
			req: &lesson.ListTeacherShiftsRequest{
				TeacherIds:     []string{""},
				ShiftSummaryId: 1,
			},
			isErr: true,
		},
		{
			name: "ShiftSummaryId is gt",
			req: &lesson.ListTeacherShiftsRequest{
				TeacherIds:     []string{"teacherid"},
				ShiftSummaryId: 0,
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.ListTeacherShifts(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestGetTeacherShifts(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *lesson.GetTeacherShiftsRequest
		isErr bool
	}{
		{
			name: "success",
			req: &lesson.GetTeacherShiftsRequest{
				TeacherId:      "teacherid",
				ShiftSummaryId: 1,
			},
			isErr: false,
		},
		{
			name: "TeacherId is min_len",
			req: &lesson.GetTeacherShiftsRequest{
				TeacherId:      "",
				ShiftSummaryId: 1,
			},
			isErr: true,
		},
		{
			name: "ShiftSummaryId is gt",
			req: &lesson.GetTeacherShiftsRequest{
				TeacherId:      "teacherid",
				ShiftSummaryId: 0,
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.GetTeacherShifts(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

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
				Decided:        true,
			},
			isErr: false,
		},
		{
			name: "TeacherId is min_len",
			req: &lesson.UpsertTeacherShiftsRequest{
				TeacherId:      "",
				ShiftSummaryId: 1,
				ShiftIds:       []int64{1, 2},
				Decided:        true,
			},
			isErr: true,
		},
		{
			name: "ShiftSummaryId is gt",
			req: &lesson.UpsertTeacherShiftsRequest{
				TeacherId:      "teacherid",
				ShiftSummaryId: 0,
				ShiftIds:       []int64{1, 2},
				Decided:        true,
			},
			isErr: true,
		},
		{
			name: "ShiftIds is unique",
			req: &lesson.UpsertTeacherShiftsRequest{
				TeacherId:      "teacherid",
				ShiftSummaryId: 1,
				ShiftIds:       []int64{1, 1},
				Decided:        true,
			},
			isErr: true,
		},
		{
			name: "ShiftIds is items.gt",
			req: &lesson.UpsertTeacherShiftsRequest{
				TeacherId:      "teacherid",
				ShiftSummaryId: 1,
				ShiftIds:       []int64{0},
				Decided:        true,
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
