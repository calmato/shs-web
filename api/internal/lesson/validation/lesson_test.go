package validation

import (
	"strings"
	"testing"

	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/stretchr/testify/assert"
)

func TestListLessonsByShiftSummaryID(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *lesson.ListLessonsByShiftSummaryIDRequest
		isErr bool
	}{
		{
			name: "success",
			req: &lesson.ListLessonsByShiftSummaryIDRequest{
				ShiftSummaryId: 1,
			},
			isErr: false,
		},
		{
			name: "ShiftSummaryId is gt",
			req: &lesson.ListLessonsByShiftSummaryIDRequest{
				ShiftSummaryId: 0,
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.ListLessonsByShiftSummaryID(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestCreateLesson(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *lesson.CreateLessonRequest
		isErr bool
	}{
		{
			name: "success",
			req: &lesson.CreateLessonRequest{
				ShiftSummaryId: 1,
				ShiftId:        1,
				SubjectId:      1,
				RoomId:         1,
				TeacherId:      "teacherid",
				StudentId:      "studentid",
				Notes:          "",
			},
			isErr: false,
		},
		{
			name: "ShiftSummaryId is gt",
			req: &lesson.CreateLessonRequest{
				ShiftSummaryId: 0,
				ShiftId:        1,
				SubjectId:      1,
				RoomId:         1,
				TeacherId:      "teacherid",
				StudentId:      "studentid",
				Notes:          "",
			},
			isErr: true,
		},
		{
			name: "ShiftId is gt",
			req: &lesson.CreateLessonRequest{
				ShiftSummaryId: 1,
				ShiftId:        0,
				SubjectId:      1,
				RoomId:         1,
				TeacherId:      "teacherid",
				StudentId:      "studentid",
				Notes:          "",
			},
			isErr: true,
		},
		{
			name: "SubjectId is gt",
			req: &lesson.CreateLessonRequest{
				ShiftSummaryId: 1,
				ShiftId:        1,
				SubjectId:      0,
				RoomId:         1,
				TeacherId:      "teacherid",
				StudentId:      "studentid",
				Notes:          "",
			},
			isErr: true,
		},
		{
			name: "RoomId is gt",
			req: &lesson.CreateLessonRequest{
				ShiftSummaryId: 1,
				ShiftId:        1,
				SubjectId:      1,
				RoomId:         0,
				TeacherId:      "teacherid",
				StudentId:      "studentid",
				Notes:          "",
			},
			isErr: true,
		},
		{
			name: "TeacherId is min_len",
			req: &lesson.CreateLessonRequest{
				ShiftSummaryId: 1,
				ShiftId:        1,
				SubjectId:      1,
				RoomId:         1,
				TeacherId:      "",
				StudentId:      "studentid",
				Notes:          "",
			},
			isErr: true,
		},
		{
			name: "StudentId is min_len",
			req: &lesson.CreateLessonRequest{
				ShiftSummaryId: 1,
				ShiftId:        1,
				SubjectId:      1,
				RoomId:         1,
				TeacherId:      "teacherid",
				StudentId:      "",
				Notes:          "",
			},
			isErr: true,
		},
		{
			name: "Notes is max_len",
			req: &lesson.CreateLessonRequest{
				ShiftSummaryId: 1,
				ShiftId:        1,
				SubjectId:      1,
				RoomId:         1,
				TeacherId:      "teacherid",
				StudentId:      "studentid",
				Notes:          strings.Repeat("x", 201),
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.CreateLesson(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}
