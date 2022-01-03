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

func TestTeacherShift_ListByShiftSummaryID(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, teacherShiftTable, teacherSubmissionTable, shiftTable, shiftSummaryTable)

	now := jst.Date(2021, 12, 10, 12, 0, 0, 0)

	const teacherID = "teacherid"
	var openAt, endAt = jst.Date(2021, 12, 1, 0, 0, 0, 0), jst.Date(2021, 12, 15, 0, 0, 0, 0)
	summary := testShiftSummary(1, 2022, openAt, endAt, now)
	err = m.db.DB.Create(&summary).Error
	require.NoError(t, err)

	shifts := make(entity.Shifts, 5)
	shifts[0] = testShift(1, 1, jst.Date(2022, 1, 1, 0, 0, 0, 0), "1700", "1830", now)
	shifts[1] = testShift(2, 1, jst.Date(2022, 1, 1, 0, 0, 0, 0), "1830", "2000", now)
	shifts[2] = testShift(3, 1, jst.Date(2022, 1, 2, 0, 0, 0, 0), "1700", "1830", now)
	shifts[3] = testShift(4, 1, jst.Date(2022, 1, 2, 0, 0, 0, 0), "1830", "2000", now)
	shifts[4] = testShift(5, 1, jst.Date(2022, 1, 2, 0, 0, 0, 0), "2000", "2130", now)
	err = m.db.DB.Create(&shifts).Error
	require.NoError(t, err)

	submission := testTeacherSubmission(teacherID, 1, true, now)
	err = m.db.DB.Create(&submission).Error
	require.NoError(t, err)

	teacherShifts := make(entity.TeacherShifts, 3)
	teacherShifts[0] = testTeacherShift(teacherID, 1, 1, now)
	teacherShifts[1] = testTeacherShift(teacherID, 1, 2, now)
	teacherShifts[2] = testTeacherShift(teacherID, 1, 4, now)
	err = m.db.DB.Create(&teacherShifts).Error
	require.NoError(t, err)

	type args struct {
		teacherID string
		summaryID int64
	}
	type want struct {
		shifts entity.TeacherShifts
		isErr  bool
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
				teacherID: teacherID,
				summaryID: 1,
			},
			want: want{
				shifts: teacherShifts,
				isErr:  false,
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

			db := NewTeacherShift(m.db)
			actual, err := db.ListByShiftSummaryID(ctx, tt.args.teacherID, tt.args.summaryID)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			assert.Len(t, actual, len(tt.want.shifts))
			for i, shift := range tt.want.shifts {
				shift.CreatedAt, shift.UpdatedAt = actual[i].CreatedAt, actual[i].UpdatedAt
				assert.Contains(t, actual, shift)
			}
		})
	}
}

func TestTeacherShift_Replace(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, teacherShiftTable, teacherSubmissionTable, shiftTable, shiftSummaryTable)

	now := jst.Date(2021, 12, 10, 12, 0, 0, 0)

	const teacherID = "teacherid"
	var openAt, endAt = jst.Date(2021, 12, 1, 0, 0, 0, 0), jst.Date(2021, 12, 15, 0, 0, 0, 0)
	summary := testShiftSummary(1, 2022, openAt, endAt, now)
	err = m.db.DB.Create(&summary).Error
	require.NoError(t, err)

	shifts := make(entity.Shifts, 5)
	shifts[0] = testShift(1, 1, jst.Date(2022, 1, 1, 0, 0, 0, 0), "1700", "1830", now)
	shifts[1] = testShift(2, 1, jst.Date(2022, 1, 1, 0, 0, 0, 0), "1830", "2000", now)
	shifts[2] = testShift(3, 1, jst.Date(2022, 1, 2, 0, 0, 0, 0), "1700", "1830", now)
	shifts[3] = testShift(4, 1, jst.Date(2022, 1, 2, 0, 0, 0, 0), "1830", "2000", now)
	shifts[4] = testShift(5, 1, jst.Date(2022, 1, 2, 0, 0, 0, 0), "2000", "2130", now)
	err = m.db.DB.Create(&shifts).Error
	require.NoError(t, err)

	submission := testTeacherSubmission(teacherID, 1, true, now)

	teacherShifts := make(entity.TeacherShifts, 3)
	teacherShifts[0] = testTeacherShift(teacherID, 1, 1, now)
	teacherShifts[1] = testTeacherShift(teacherID, 1, 2, now)
	teacherShifts[2] = testTeacherShift(teacherID, 1, 4, now)

	otherShifts := make(entity.TeacherShifts, 1)
	otherShifts[0] = testTeacherShift(teacherID, 1, 0, now)

	type args struct {
		submission *entity.TeacherSubmission
		shifts     entity.TeacherShifts
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
			name:  "success insert",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {},
			args: args{
				submission: submission,
				shifts:     teacherShifts,
			},
			want: want{
				isErr: false,
			},
		},
		{
			name: "success update",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {
				err = m.db.DB.Create(&submission).Error
				require.NoError(t, err)
				err = m.db.DB.Create(&teacherShifts).Error
				require.NoError(t, err)
			},
			args: args{
				submission: testTeacherSubmission(teacherID, 1, false, now),
				shifts:     teacherShifts[1:],
			},
			want: want{
				isErr: false,
			},
		},
		{
			name: "success shift length 0",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {
				err = m.db.DB.Create(&submission).Error
				require.NoError(t, err)
				err = m.db.DB.Create(&teacherShifts).Error
				require.NoError(t, err)
			},
			args: args{
				submission: submission,
				shifts:     entity.TeacherShifts{},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name: "failed to replace teacher shifts",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {
				err = m.db.DB.Create(&submission).Error
				require.NoError(t, err)
				err = m.db.DB.Create(&teacherShifts).Error
				require.NoError(t, err)
			},
			args: args{
				submission: submission,
				shifts:     otherShifts,
			},
			want: want{
				isErr: true,
			},
		},
		{
			name: "failed to replace teacher submission",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {
				err = m.db.DB.Create(&submission).Error
				require.NoError(t, err)
				err = m.db.DB.Create(&teacherShifts).Error
				require.NoError(t, err)
			},
			args: args{
				submission: testTeacherSubmission(teacherID, 0, false, now),
				shifts:     entity.TeacherShifts{},
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

			_ = m.dbDelete(ctx, teacherShiftTable, teacherSubmissionTable)
			tt.setup(ctx, t, m)

			db := NewTeacherShift(m.db)
			err := db.Replace(ctx, tt.args.submission, tt.args.shifts)
			assert.Equal(t, tt.want.isErr, err != nil, err)
		})
	}
}

func testTeacherShift(teacherID string, summaryID, shiftID int64, now time.Time) *entity.TeacherShift {
	return &entity.TeacherShift{
		TeacherID:      teacherID,
		ShiftID:        shiftID,
		ShiftSummaryID: summaryID,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
}
