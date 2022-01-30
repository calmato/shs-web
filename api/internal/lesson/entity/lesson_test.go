package entity

import (
	"testing"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/stretchr/testify/assert"
)

func TestLesson_Proto(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name   string
		lesson *Lesson
		expect *lesson.Lesson
	}{
		{
			name: "success",
			lesson: &Lesson{
				ID:             1,
				ShiftSummaryID: 1,
				ShiftID:        1,
				SubjectID:      1,
				RoomID:         1,
				TeacherID:      "teacherid",
				StudentID:      "studentid",
				Notes:          "",
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			expect: &lesson.Lesson{
				Id:             1,
				ShiftSummaryId: 1,
				ShiftId:        1,
				SubjectId:      1,
				RoomId:         1,
				TeacherId:      "teacherid",
				StudentId:      "studentid",
				Notes:          "",
				CreatedAt:      now.Unix(),
				UpdatedAt:      now.Unix(),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.lesson.Proto())
		})
	}
}

func TestLessons_ShiftSummaryIDs(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name    string
		lessons Lessons
		expect  []int64
	}{
		{
			name: "success",
			lessons: Lessons{
				{
					ID:             1,
					ShiftSummaryID: 1,
					ShiftID:        1,
					SubjectID:      1,
					RoomID:         1,
					TeacherID:      "teacherid",
					StudentID:      "studentid",
					Notes:          "",
					CreatedAt:      now,
					UpdatedAt:      now,
				},
			},
			expect: []int64{1},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.lessons.ShiftSummaryIDs())
		})
	}
}

func TestLessons_ShiftIDs(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name    string
		lessons Lessons
		expect  []int64
	}{
		{
			name: "success",
			lessons: Lessons{
				{
					ID:             1,
					ShiftSummaryID: 1,
					ShiftID:        1,
					SubjectID:      1,
					RoomID:         1,
					TeacherID:      "teacherid",
					StudentID:      "studentid",
					Notes:          "",
					CreatedAt:      now,
					UpdatedAt:      now,
				},
			},
			expect: []int64{1},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.lessons.ShiftIDs())
		})
	}
}

func TestLessons_Decided(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name      string
		lessons   Lessons
		summaries map[int64]*ShiftSummary
		expect    Lessons
		isErr     bool
	}{
		{
			name: "success",
			lessons: Lessons{
				{
					ID:             1,
					ShiftSummaryID: 1,
					ShiftID:        1,
					SubjectID:      1,
					RoomID:         1,
					TeacherID:      "teacherid",
					StudentID:      "studentid",
					Notes:          "",
					CreatedAt:      now,
					UpdatedAt:      now,
				},
			},
			summaries: map[int64]*ShiftSummary{
				1: {Decided: true},
			},
			expect: Lessons{
				{
					ID:             1,
					ShiftSummaryID: 1,
					ShiftID:        1,
					SubjectID:      1,
					RoomID:         1,
					TeacherID:      "teacherid",
					StudentID:      "studentid",
					Notes:          "",
					CreatedAt:      now,
					UpdatedAt:      now,
				},
			},
			isErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual, err := tt.lessons.Decided(tt.summaries)
			assert.Equal(t, tt.isErr, err != nil, err)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestLessons_Proto(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name    string
		lessons Lessons
		expect  []*lesson.Lesson
	}{
		{
			name: "success",
			lessons: Lessons{
				{
					ID:             1,
					ShiftSummaryID: 1,
					ShiftID:        1,
					SubjectID:      1,
					RoomID:         1,
					TeacherID:      "teacherid",
					StudentID:      "studentid",
					Notes:          "",
					CreatedAt:      now,
					UpdatedAt:      now,
				},
			},
			expect: []*lesson.Lesson{
				{
					Id:             1,
					ShiftSummaryId: 1,
					ShiftId:        1,
					SubjectId:      1,
					RoomId:         1,
					TeacherId:      "teacherid",
					StudentId:      "studentid",
					Notes:          "",
					CreatedAt:      now.Unix(),
					UpdatedAt:      now.Unix(),
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.lessons.Proto())
		})
	}
}
