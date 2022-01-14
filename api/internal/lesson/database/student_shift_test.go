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

func TestStudentShift_ListByShiftSummaryID(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, studentShiftTable, studentSubmissionTable, shiftTable, shiftSummaryTable)

	now := jst.Date(2021, 12, 10, 12, 0, 0, 0)

	const studentID = "studentid"
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

	submission := testStudentSubmission(studentID, 1, true, now)
	err = m.db.DB.Create(&submission).Error
	require.NoError(t, err)

	studentShifts := make(entity.StudentShifts, 3)
	studentShifts[0] = testStudentShift(studentID, 1, 1, now)
	studentShifts[1] = testStudentShift(studentID, 1, 2, now)
	studentShifts[2] = testStudentShift(studentID, 1, 4, now)
	err = m.db.DB.Create(&studentShifts).Error
	require.NoError(t, err)

	type args struct {
		studentIDs []string
		summaryID  int64
	}
	type want struct {
		shifts entity.StudentShifts
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
				studentIDs: []string{studentID},
				summaryID:  1,
			},
			want: want{
				shifts: studentShifts,
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

			db := NewStudentShift(m.db)
			actual, err := db.ListByShiftSummaryID(ctx, tt.args.studentIDs, tt.args.summaryID)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			assert.Len(t, actual, len(tt.want.shifts))
			for i, shift := range tt.want.shifts {
				shift.CreatedAt, shift.UpdatedAt = actual[i].CreatedAt, actual[i].UpdatedAt
				assert.Contains(t, actual, shift)
			}
		})
	}
}

func TestStudentShift_Replace(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, studentShiftTable, studentSubmissionTable, shiftTable, shiftSummaryTable)

	now := jst.Date(2021, 12, 10, 12, 0, 0, 0)

	const studentID = "studentid"
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

	submission := testStudentSubmission(studentID, 1, true, now)

	studentShifts := make(entity.StudentShifts, 3)
	studentShifts[0] = testStudentShift(studentID, 1, 1, now)
	studentShifts[1] = testStudentShift(studentID, 1, 2, now)
	studentShifts[2] = testStudentShift(studentID, 1, 4, now)

	otherShifts := make(entity.StudentShifts, 1)
	otherShifts[0] = testStudentShift(studentID, 1, 0, now)

	type args struct {
		submission *entity.StudentSubmission
		shifts     entity.StudentShifts
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
				shifts:     studentShifts,
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
				err = m.db.DB.Create(&studentShifts).Error
				require.NoError(t, err)
			},
			args: args{
				submission: testStudentSubmission(studentID, 1, false, now),
				shifts:     studentShifts[1:],
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
				err = m.db.DB.Create(&studentShifts).Error
				require.NoError(t, err)
			},
			args: args{
				submission: submission,
				shifts:     entity.StudentShifts{},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name: "failed to replace student shifts",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {
				err = m.db.DB.Create(&submission).Error
				require.NoError(t, err)
				err = m.db.DB.Create(&studentShifts).Error
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
			name: "failed to replace student submission",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {
				err = m.db.DB.Create(&submission).Error
				require.NoError(t, err)
				err = m.db.DB.Create(&studentShifts).Error
				require.NoError(t, err)
			},
			args: args{
				submission: testStudentSubmission(studentID, 0, false, now),
				shifts:     entity.StudentShifts{},
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

			_ = m.dbDelete(ctx, studentShiftTable, studentSubmissionTable)
			tt.setup(ctx, t, m)

			db := NewStudentShift(m.db)
			err := db.Replace(ctx, tt.args.submission, tt.args.shifts)
			assert.Equal(t, tt.want.isErr, err != nil, err)
		})
	}
}

func testStudentShift(studentID string, summaryID, shiftID int64, now time.Time) *entity.StudentShift {
	return &entity.StudentShift{
		StudentID:      studentID,
		ShiftID:        shiftID,
		ShiftSummaryID: summaryID,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
}
