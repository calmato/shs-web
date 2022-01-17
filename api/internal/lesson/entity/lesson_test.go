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
