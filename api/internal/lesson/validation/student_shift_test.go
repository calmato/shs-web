package validation

import (
	"testing"

	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/stretchr/testify/assert"
)

func TestListStudentSubmissionsByShiftSummaryIDs(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *lesson.ListStudentSubmissionsByShiftSummaryIDsRequest
		isErr bool
	}{
		{
			name: "success",
			req: &lesson.ListStudentSubmissionsByShiftSummaryIDsRequest{
				StudentId:       "studentid",
				ShiftSummaryIds: []int64{1, 2},
			},
			isErr: false,
		},
		{
			name: "StudentId is min_len",
			req: &lesson.ListStudentSubmissionsByShiftSummaryIDsRequest{
				StudentId:       "",
				ShiftSummaryIds: []int64{1, 2},
			},
			isErr: true,
		},
		{
			name: "ShiftSummaryIds is unique",
			req: &lesson.ListStudentSubmissionsByShiftSummaryIDsRequest{
				StudentId:       "studentid",
				ShiftSummaryIds: []int64{1, 1},
			},
			isErr: true,
		},
		{
			name: "ShiftSummaryIds is items.gt",
			req: &lesson.ListStudentSubmissionsByShiftSummaryIDsRequest{
				StudentId:       "studentid",
				ShiftSummaryIds: []int64{0},
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.ListStudentSubmissionsByShiftSummaryIDs(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestListStudentShifts(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *lesson.ListStudentShiftsRequest
		isErr bool
	}{
		{
			name: "success",
			req: &lesson.ListStudentShiftsRequest{
				StudentIds:     []string{"studentid"},
				ShiftSummaryId: 1,
			},
			isErr: false,
		},
		{
			name: "StudentIds is unique",
			req: &lesson.ListStudentShiftsRequest{
				StudentIds:     []string{"studentid", "studentid"},
				ShiftSummaryId: 1,
			},
			isErr: true,
		},
		{
			name: "StudentIds is items.min_len",
			req: &lesson.ListStudentShiftsRequest{
				StudentIds:     []string{""},
				ShiftSummaryId: 1,
			},
			isErr: true,
		},
		{
			name: "ShiftSummaryId is gt",
			req: &lesson.ListStudentShiftsRequest{
				StudentIds:     []string{"studentid"},
				ShiftSummaryId: 0,
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.ListStudentShifts(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestGetStudentShifts(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *lesson.GetStudentShiftsRequest
		isErr bool
	}{
		{
			name: "success",
			req: &lesson.GetStudentShiftsRequest{
				StudentId:      "studentid",
				ShiftSummaryId: 1,
			},
			isErr: false,
		},
		{
			name: "StudentId is min_len",
			req: &lesson.GetStudentShiftsRequest{
				StudentId:      "",
				ShiftSummaryId: 1,
			},
			isErr: true,
		},
		{
			name: "ShiftSummaryId is gt",
			req: &lesson.GetStudentShiftsRequest{
				StudentId:      "studentid",
				ShiftSummaryId: 0,
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.GetStudentShifts(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestUpsertStudentShifts(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *lesson.UpsertStudentShiftsRequest
		isErr bool
	}{
		{
			name: "success",
			req: &lesson.UpsertStudentShiftsRequest{
				StudentId:      "studentid",
				ShiftSummaryId: 1,
				ShiftIds:       []int64{1, 2},
				Decided:        true,
			},
			isErr: false,
		},
		{
			name: "StudentId is min_len",
			req: &lesson.UpsertStudentShiftsRequest{
				StudentId:      "",
				ShiftSummaryId: 1,
				ShiftIds:       []int64{1, 2},
				Decided:        true,
			},
			isErr: true,
		},
		{
			name: "ShiftSummaryId is gt",
			req: &lesson.UpsertStudentShiftsRequest{
				StudentId:      "studentid",
				ShiftSummaryId: 0,
				ShiftIds:       []int64{1, 2},
				Decided:        true,
			},
			isErr: true,
		},
		{
			name: "ShiftIds is unique",
			req: &lesson.UpsertStudentShiftsRequest{
				StudentId:      "studentid",
				ShiftSummaryId: 1,
				ShiftIds:       []int64{1, 1},
				Decided:        true,
			},
			isErr: true,
		},
		{
			name: "ShiftIds is items.gt",
			req: &lesson.UpsertStudentShiftsRequest{
				StudentId:      "studentid",
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
			err := validator.UpsertStudentShifts(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}
