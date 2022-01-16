package database

import (
	"context"
	"testing"
	"time"

	"github.com/calmato/shs-web/api/internal/lesson/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLesson_List(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, lessonTable, shiftTable, shiftSummaryTable)

	now := jst.Now()

	summary := testShiftSummary(1, 202201, now.AddDate(0, -1, 0), now.AddDate(0, -1, 1), now)
	err = m.db.DB.Create(&summary).Error
	require.NoError(t, err)

	shifts := make(entity.Shifts, 3)
	shifts[0] = testShift(1, 1, jst.Date(2022, 2, 1, 0, 0, 0, 0), "1700", "1830", now)
	shifts[1] = testShift(2, 1, jst.Date(2022, 2, 1, 0, 0, 0, 0), "1830", "2000", now)
	shifts[2] = testShift(3, 1, jst.Date(2022, 2, 2, 0, 0, 0, 0), "1700", "1830", now)
	err = m.db.DB.Create(&shifts).Error
	require.NoError(t, err)

	lessons := make(entity.Lessons, 3)
	lessons[0] = testLesson(1, 1, 1, 1, "teacherid1", "studentid1", now)
	lessons[1] = testLesson(2, 1, 1, 2, "teacherid2", "studentid2", now)
	lessons[2] = testLesson(3, 1, 2, 1, "teacherid1", "studentid2", now)
	err = m.db.DB.Create(&lessons).Error
	require.NoError(t, err)

	type args struct {
		params *ListLessonsParams
	}
	type want struct {
		lessons entity.Lessons
		isErr   bool
	}
	tests := []struct {
		name  string
		setup func(ctx context.Context, t *testing.T, m *mocks)
		args  args
		want  want
	}{
		{
			name:  "success",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {},
			args: args{
				params: &ListLessonsParams{
					ShiftSummaryID: 1,
				},
			},
			want: want{
				lessons: lessons,
				isErr:   false,
			},
		},
		{
			name:  "success",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {},
			args: args{
				params: &ListLessonsParams{
					ShiftSummaryID: 100,
					Limit:          1,
					Offset:         100,
				},
			},
			want: want{
				lessons: entity.Lessons{},
				isErr:   false,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			tt.setup(ctx, t, m)

			db := NewLesson(m.db)
			actual, err := db.List(ctx, tt.args.params)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			assert.Len(t, actual, len(tt.want.lessons))
			for i, lesson := range tt.want.lessons {
				lesson.CreatedAt = actual[i].CreatedAt
				lesson.UpdatedAt = actual[i].UpdatedAt
				assert.Contains(t, actual, lesson)
			}
		})
	}
}

func TestLesson_Create(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, lessonTable, shiftTable, shiftSummaryTable)

	now := jst.Now()

	summary := testShiftSummary(1, 202201, now.AddDate(0, -1, 0), now.AddDate(0, -1, 1), now)
	err = m.db.DB.Create(&summary).Error
	require.NoError(t, err)

	shift := testShift(1, 1, jst.Date(2022, 2, 1, 0, 0, 0, 0), "1700", "1830", now)
	err = m.db.DB.Create(&shift).Error
	require.NoError(t, err)

	lesson := testLesson(1, 1, 1, 1, "teacherid", "studentid", now)

	type args struct {
		lesson *entity.Lesson
	}
	type want struct {
		isErr bool
	}
	tests := []struct {
		name  string
		setup func(ctx context.Context, t *testing.T, m *mocks)
		args  args
		want  want
	}{
		{
			name:  "success",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {},
			args: args{
				lesson: &entity.Lesson{
					ShiftSummaryID: 1,
					ShiftID:        1,
					SubjectID:      1,
					RoomID:         1,
					TeacherID:      "teacherid",
					StudentID:      "studentid",
				},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name: "already exists",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {
				err = m.db.DB.Create(&lesson).Error
				require.NoError(t, err)
			},
			args: args{
				lesson: &entity.Lesson{
					ShiftSummaryID: 1,
					ShiftID:        1,
					SubjectID:      1,
					RoomID:         1,
					TeacherID:      "teacherid",
					StudentID:      "studentid",
				},
			},
			want: want{
				isErr: true,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			_ = m.dbDelete(ctx, lessonTable)
			tt.setup(ctx, t, m)

			db := NewLesson(m.db)
			err := db.Create(ctx, tt.args.lesson)
			assert.Equal(t, tt.want.isErr, err != nil, err)
		})
	}
}

func TestLesson_Update(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, lessonTable, shiftTable, shiftSummaryTable)

	now := jst.Now()

	summary := testShiftSummary(1, 202201, now.AddDate(0, -1, 0), now.AddDate(0, -1, 1), now)
	err = m.db.DB.Create(&summary).Error
	require.NoError(t, err)

	shift := testShift(1, 1, jst.Date(2022, 2, 1, 0, 0, 0, 0), "1700", "1830", now)
	err = m.db.DB.Create(&shift).Error
	require.NoError(t, err)

	lesson := testLesson(1, 1, 1, 1, "teacherid", "studentid", now)

	type args struct {
		lessonID int64
		lesson   *entity.Lesson
	}
	type want struct {
		isErr bool
	}
	tests := []struct {
		name  string
		setup func(ctx context.Context, t *testing.T, m *mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {
				err := m.db.DB.Create(&lesson).Error
				require.NoError(t, err)
			},
			args: args{
				lessonID: 1,
				lesson: &entity.Lesson{
					ShiftSummaryID: 1,
					ShiftID:        1,
					SubjectID:      2,
					RoomID:         1,
					TeacherID:      "teacherid",
					StudentID:      "studentid",
				},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name:  "not found",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {},
			args: args{
				lessonID: 1,
				lesson: &entity.Lesson{
					ShiftSummaryID: 1,
					ShiftID:        1,
					SubjectID:      1,
					RoomID:         1,
					TeacherID:      "teacherid",
					StudentID:      "studentid",
				},
			},
			want: want{
				isErr: true,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			_ = m.dbDelete(ctx, lessonTable)
			tt.setup(ctx, t, m)

			db := NewLesson(m.db)
			err := db.Update(ctx, tt.args.lessonID, tt.args.lesson)
			assert.Equal(t, tt.want.isErr, err != nil, err)
		})
	}
}

func TestLesson_Delete(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, lessonTable, shiftTable, shiftSummaryTable)

	now := jst.Now()

	summary := testShiftSummary(1, 202201, now.AddDate(0, -1, 0), now.AddDate(0, -1, 1), now)
	err = m.db.DB.Create(&summary).Error
	require.NoError(t, err)

	shift := testShift(1, 1, jst.Date(2022, 2, 1, 0, 0, 0, 0), "1700", "1830", now)
	err = m.db.DB.Create(&shift).Error
	require.NoError(t, err)

	lesson := testLesson(1, 1, 1, 1, "teacherid", "studentid", now)

	type args struct {
		lessonID int64
	}
	type want struct {
		isErr bool
	}
	tests := []struct {
		name  string
		setup func(ctx context.Context, t *testing.T, m *mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {
				err := m.db.DB.Create(&lesson).Error
				require.NoError(t, err)
			},
			args: args{
				lessonID: 1,
			},
			want: want{
				isErr: false,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			_ = m.dbDelete(ctx, lessonTable)
			tt.setup(ctx, t, m)

			db := NewLesson(m.db)
			err := db.Delete(ctx, tt.args.lessonID)
			assert.Equal(t, tt.want.isErr, err != nil, err)
		})
	}
}

func testLesson(id, summaryID, shiftID int64, roomID int32, teacherID, studentID string, now time.Time) *entity.Lesson {
	return &entity.Lesson{
		ID:             id,
		ShiftSummaryID: summaryID,
		ShiftID:        shiftID,
		SubjectID:      1,
		RoomID:         roomID,
		TeacherID:      teacherID,
		StudentID:      studentID,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
}
