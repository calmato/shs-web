package entity

import (
	"testing"

	"github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/stretchr/testify/assert"
)

func TestLesson(t *testing.T) {
	t.Parallel()
	now := jst.Date(2022, 1, 15, 12, 30, 0, 0)
	tests := []struct {
		name   string
		lesson *entity.Lesson
		shift  *entity.Shift
		expect *Lesson
		isErr  bool
	}{
		{
			name: "success",
			lesson: &entity.Lesson{
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
			shift: &entity.Shift{
				Shift: &lesson.Shift{
					Id:             1,
					ShiftSummaryId: 1,
					Date:           "20220202",
					StartTime:      "1700",
					EndTime:        "1830",
					CreatedAt:      now.Unix(),
					UpdatedAt:      now.Unix(),
				},
			},
			expect: &Lesson{
				ID:        1,
				ShiftID:   1,
				SubjectID: 1,
				Room:      1,
				TeacherID: "teacherid",
				StudentID: "studentid",
				Notes:     "感想",
				StartAt:   jst.Date(2022, 2, 2, 17, 0, 0, 0),
				EndAt:     jst.Date(2022, 2, 2, 18, 30, 0, 0),
				CreatedAt: now,
				UpdatedAt: now,
			},
			isErr: false,
		},
		{
			name: "failed to parse start time",
			lesson: &entity.Lesson{
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
			shift: &entity.Shift{
				Shift: &lesson.Shift{
					Id:             1,
					ShiftSummaryId: 1,
					Date:           "20220202",
					StartTime:      "",
					EndTime:        "1830",
					CreatedAt:      now.Unix(),
					UpdatedAt:      now.Unix(),
				},
			},
			expect: nil,
			isErr:  true,
		},
		{
			name: "failed to parse ent time",
			lesson: &entity.Lesson{
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
			shift: &entity.Shift{
				Shift: &lesson.Shift{
					Id:             1,
					ShiftSummaryId: 1,
					Date:           "20220202",
					StartTime:      "1700",
					EndTime:        "",
					CreatedAt:      now.Unix(),
					UpdatedAt:      now.Unix(),
				},
			},
			expect: nil,
			isErr:  true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual, err := NewLesson(tt.lesson, tt.shift)
			assert.Equal(t, tt.isErr, err != nil, err)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestLessons(t *testing.T) {
	t.Parallel()
	now := jst.Date(2022, 1, 15, 12, 30, 0, 0)
	tests := []struct {
		name    string
		lessons entity.Lessons
		shifts  map[int64]*entity.Shift
		expect  Lessons
		isErr   bool
	}{
		{
			name: "success",
			lessons: entity.Lessons{
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
				{
					Lesson: &lesson.Lesson{
						Id:             2,
						ShiftSummaryId: 1,
						ShiftId:        2,
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
			shifts: map[int64]*entity.Shift{
				1: {
					Shift: &lesson.Shift{
						Id:             1,
						ShiftSummaryId: 1,
						Date:           "20220202",
						StartTime:      "1700",
						EndTime:        "1830",
						CreatedAt:      now.Unix(),
						UpdatedAt:      now.Unix(),
					},
				},
				2: {
					Shift: &lesson.Shift{
						Id:             2,
						ShiftSummaryId: 1,
						Date:           "20220202",
						StartTime:      "1830",
						EndTime:        "2000",
						CreatedAt:      now.Unix(),
						UpdatedAt:      now.Unix(),
					},
				},
			},
			expect: Lessons{
				{
					ID:        1,
					ShiftID:   1,
					SubjectID: 1,
					Room:      1,
					TeacherID: "teacherid",
					StudentID: "studentid",
					Notes:     "感想",
					StartAt:   jst.Date(2022, 2, 2, 17, 0, 0, 0),
					EndAt:     jst.Date(2022, 2, 2, 18, 30, 0, 0),
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					ID:        2,
					ShiftID:   2,
					SubjectID: 1,
					Room:      1,
					TeacherID: "teacherid",
					StudentID: "studentid",
					Notes:     "感想",
					StartAt:   jst.Date(2022, 2, 2, 18, 30, 0, 0),
					EndAt:     jst.Date(2022, 2, 2, 20, 0, 0, 0),
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
			isErr: false,
		},
		{
			name: "failed to shift not found",
			lessons: entity.Lessons{
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
				{
					Lesson: &lesson.Lesson{
						Id:             2,
						ShiftSummaryId: 1,
						ShiftId:        2,
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
			shifts: map[int64]*entity.Shift{},
			expect: nil,
			isErr:  true,
		},
		{
			name: "failed to new lesson",
			lessons: entity.Lessons{
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
				{
					Lesson: &lesson.Lesson{
						Id:             2,
						ShiftSummaryId: 1,
						ShiftId:        2,
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
			shifts: map[int64]*entity.Shift{
				1: {
					Shift: &lesson.Shift{
						Id:             1,
						ShiftSummaryId: 1,
						Date:           "20220202",
						StartTime:      "",
						EndTime:        "",
						CreatedAt:      now.Unix(),
						UpdatedAt:      now.Unix(),
					},
				},
			},
			expect: nil,
			isErr:  true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual, err := NewLessons(tt.lessons, tt.shifts)
			assert.Equal(t, tt.isErr, err != nil, err)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestLessons_SortByStartAt(t *testing.T) {
	t.Parallel()
	now := jst.Date(2022, 1, 15, 12, 30, 0, 0)
	tests := []struct {
		name    string
		lessons Lessons
		expect  Lessons
	}{
		{
			name: "success",
			lessons: Lessons{
				{
					ID:        1,
					ShiftID:   1,
					SubjectID: 1,
					Room:      1,
					TeacherID: "teacherid",
					StudentID: "studentid",
					Notes:     "感想",
					StartAt:   jst.Date(2022, 2, 2, 17, 0, 0, 0),
					EndAt:     jst.Date(2022, 2, 2, 18, 30, 0, 0),
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					ID:        3,
					ShiftID:   3,
					SubjectID: 1,
					Room:      1,
					TeacherID: "teacherid",
					StudentID: "studentid",
					Notes:     "感想",
					StartAt:   jst.Date(2022, 2, 2, 20, 0, 0, 0),
					EndAt:     jst.Date(2022, 2, 2, 21, 30, 0, 0),
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					ID:        2,
					ShiftID:   2,
					SubjectID: 1,
					Room:      1,
					TeacherID: "teacherid",
					StudentID: "studentid",
					Notes:     "感想",
					StartAt:   jst.Date(2022, 2, 2, 18, 30, 0, 0),
					EndAt:     jst.Date(2022, 2, 2, 20, 0, 0, 0),
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
			expect: Lessons{
				{
					ID:        1,
					ShiftID:   1,
					SubjectID: 1,
					Room:      1,
					TeacherID: "teacherid",
					StudentID: "studentid",
					Notes:     "感想",
					StartAt:   jst.Date(2022, 2, 2, 17, 0, 0, 0),
					EndAt:     jst.Date(2022, 2, 2, 18, 30, 0, 0),
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					ID:        2,
					ShiftID:   2,
					SubjectID: 1,
					Room:      1,
					TeacherID: "teacherid",
					StudentID: "studentid",
					Notes:     "感想",
					StartAt:   jst.Date(2022, 2, 2, 18, 30, 0, 0),
					EndAt:     jst.Date(2022, 2, 2, 20, 0, 0, 0),
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					ID:        3,
					ShiftID:   3,
					SubjectID: 1,
					Room:      1,
					TeacherID: "teacherid",
					StudentID: "studentid",
					Notes:     "感想",
					StartAt:   jst.Date(2022, 2, 2, 20, 0, 0, 0),
					EndAt:     jst.Date(2022, 2, 2, 21, 30, 0, 0),
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.lessons.SortByStartAt()
			assert.Equal(t, tt.expect, tt.lessons)
		})
	}
}
