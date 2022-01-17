package entity

import (
	"testing"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/stretchr/testify/assert"
)

func TestLesson(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name   string
		lesson *lesson.Lesson
		expect *Lesson
	}{
		{
			name: "success",
			lesson: &lesson.Lesson{
				Id:             1,
				ShiftSummaryId: 1,
				ShiftId:        1,
				SubjectId:      1,
				RoomId:         1,
				TeacherId:      "teacherid",
				StudentId:      "studentid",
				Notes:          "感想",
				CreatedAt:      now.Unix(),
				UpdatedAt:      now.Unix(),
			},
			expect: &Lesson{
				Lesson: &lesson.Lesson{
					Id:             1,
					ShiftSummaryId: 1,
					ShiftId:        1,
					SubjectId:      1,
					RoomId:         1,
					TeacherId:      "teacherid",
					StudentId:      "studentid",
					Notes:          "感想",
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
			assert.Equal(t, tt.expect, NewLesson(tt.lesson))
		})
	}
}

func TestLessons(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name    string
		lessons []*lesson.Lesson
		expect  Lessons
	}{
		{
			name: "success",
			lessons: []*lesson.Lesson{
				{
					Id:             1,
					ShiftSummaryId: 1,
					ShiftId:        1,
					SubjectId:      1,
					RoomId:         1,
					TeacherId:      "teacherid",
					StudentId:      "studentid",
					Notes:          "感想",
					CreatedAt:      now.Unix(),
					UpdatedAt:      now.Unix(),
				},
			},
			expect: Lessons{
				{
					Lesson: &lesson.Lesson{
						Id:             1,
						ShiftSummaryId: 1,
						ShiftId:        1,
						SubjectId:      1,
						RoomId:         1,
						TeacherId:      "teacherid",
						StudentId:      "studentid",
						Notes:          "感想",
						CreatedAt:      now.Unix(),
						UpdatedAt:      now.Unix(),
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewLessons(tt.lessons))
		})
	}
}
