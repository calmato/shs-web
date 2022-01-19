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

func TestListLessonsByTeacherID(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *lesson.ListLessonsByTeacherIDRequest
		isErr bool
	}{
		{
			name: "success",
			req: &lesson.ListLessonsByTeacherIDRequest{
				ShiftSummaryId: 1,
				TeacherId:      "teacherid",
			},
			isErr: false,
		},
		{
			name: "ShiftSummaryId is gt",
			req: &lesson.ListLessonsByTeacherIDRequest{
				ShiftSummaryId: 0,
				TeacherId:      "teacherid",
			},
			isErr: true,
		},
		{
			name: "TeacherId is min_len",
			req: &lesson.ListLessonsByTeacherIDRequest{
				ShiftSummaryId: 1,
				TeacherId:      "",
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.ListLessonsByTeacherID(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestListLessonsByStudentID(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *lesson.ListLessonsByStudentIDRequest
		isErr bool
	}{
		{
			name: "success",
			req: &lesson.ListLessonsByStudentIDRequest{
				ShiftSummaryId: 1,
				StudentId:      "studentid",
			},
			isErr: false,
		},
		{
			name: "ShiftSummaryId is gt",
			req: &lesson.ListLessonsByStudentIDRequest{
				ShiftSummaryId: 0,
				StudentId:      "studentid",
			},
			isErr: true,
		},
		{
			name: "StudentId is min_len",
			req: &lesson.ListLessonsByStudentIDRequest{
				ShiftSummaryId: 1,
				StudentId:      "",
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.ListLessonsByStudentID(tt.req)
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
