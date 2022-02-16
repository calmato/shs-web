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

func TestTeacherSubmission_ListByShiftSummaryIDs(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, teacherSubmissionTable, shiftSummaryTable)

	now := jst.Date(2021, 12, 10, 12, 0, 0, 0)

	const teacherID = "teacherid"

	summaries := make(entity.ShiftSummaries, 3)
	summaries[0] = testShiftSummary(1, 202202, jst.Date(2022, 1, 1, 0, 0, 0, 0), jst.Date(2022, 1, 15, 0, 0, 0, 0), now)
	summaries[1] = testShiftSummary(2, 202203, jst.Date(2022, 2, 1, 0, 0, 0, 0), jst.Date(2022, 2, 15, 0, 0, 0, 0), now)
	summaries[2] = testShiftSummary(3, 202204, jst.Date(2022, 3, 1, 0, 0, 0, 0), jst.Date(2022, 3, 15, 0, 0, 0, 0), now)
	err = m.db.DB.Create(&summaries).Error
	require.NoError(t, err)

	submissions := make(entity.TeacherSubmissions, 2)
	submissions[0] = testTeacherSubmission(teacherID, 1, true, now)
	submissions[1] = testTeacherSubmission(teacherID, 2, false, now)
	err = m.db.DB.Create(&submissions).Error
	require.NoError(t, err)

	type args struct {
		teacherID  string
		summaryIDs []int64
	}
	type want struct {
		submissions entity.TeacherSubmissions
		isErr       bool
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
				teacherID:  teacherID,
				summaryIDs: []int64{1, 2},
			},
			want: want{
				submissions: submissions,
				isErr:       false,
			},
		},
		{
			name:  "success is empty",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {},
			args: args{
				teacherID:  teacherID,
				summaryIDs: []int64{0},
			},
			want: want{
				submissions: entity.TeacherSubmissions{},
				isErr:       false,
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

			db := NewTeacherSubmission(m.db)
			actual, err := db.ListByShiftSummaryIDs(ctx, tt.args.teacherID, tt.args.summaryIDs)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			assert.Len(t, actual, len(tt.want.submissions))
			for i, expect := range tt.want.submissions {
				expect.CreatedAt, expect.UpdatedAt = actual[i].CreatedAt, actual[i].UpdatedAt
				assert.Contains(t, actual, expect)
			}
		})
	}
}

func TestTeacherSubmission_ListByTeacherIDs(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, teacherSubmissionTable, shiftSummaryTable)

	now := jst.Date(2021, 12, 10, 12, 0, 0, 0)

	summary := testShiftSummary(1, 202202, jst.Date(2022, 1, 1, 0, 0, 0, 0), jst.Date(2022, 1, 15, 0, 0, 0, 0), now)
	err = m.db.DB.Create(&summary).Error
	require.NoError(t, err)

	submissions := make(entity.TeacherSubmissions, 2)
	submissions[0] = testTeacherSubmission("teacherid01", 1, true, now)
	submissions[1] = testTeacherSubmission("teacherid02", 2, false, now)
	err = m.db.DB.Create(&submissions).Error
	require.NoError(t, err)

	type args struct {
		teacherIDs []string
		summaryID  int64
	}
	type want struct {
		submissions entity.TeacherSubmissions
		isErr       bool
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
				teacherIDs: []string{"teacherid01", "teacherid02"},
				summaryID:  1,
			},
			want: want{
				submissions: submissions,
				isErr:       false,
			},
		},
		{
			name:  "success is empty",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {},
			args: args{
				teacherIDs: []string{"teacherid00"},
				summaryID:  0,
			},
			want: want{
				submissions: entity.TeacherSubmissions{},
				isErr:       false,
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

			db := NewTeacherSubmission(m.db)
			actual, err := db.ListByTeacherIDs(ctx, tt.args.teacherIDs, tt.args.summaryID)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			assert.Len(t, actual, len(tt.want.submissions))
			for i, expect := range tt.want.submissions {
				expect.CreatedAt, expect.UpdatedAt = actual[i].CreatedAt, actual[i].UpdatedAt
				assert.Contains(t, actual, expect)
			}
		})
	}
}

func TestTeacherSubmission_Get(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, teacherSubmissionTable, shiftSummaryTable)

	now := jst.Date(2021, 12, 10, 12, 0, 0, 0)

	const teacherID = "teacherid"

	summary := testShiftSummary(1, 202202, jst.Date(2022, 1, 1, 0, 0, 0, 0), jst.Date(2022, 1, 15, 0, 0, 0, 0), now)
	err = m.db.DB.Create(&summary).Error
	require.NoError(t, err)

	submission := testTeacherSubmission(teacherID, 1, true, now)
	err = m.db.DB.Create(&submission).Error
	require.NoError(t, err)

	type args struct {
		teacherID string
		summaryID int64
	}
	type want struct {
		submission *entity.TeacherSubmission
		isErr      bool
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
				submission: submission,
				isErr:      false,
			},
		},
		{
			name:  "not found",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {},
			args: args{
				teacherID: teacherID,
				summaryID: 0,
			},
			want: want{
				submission: nil,
				isErr:      true,
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

			db := NewTeacherSubmission(m.db)
			actual, err := db.Get(ctx, tt.args.teacherID, tt.args.summaryID)
			if tt.want.isErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				tt.want.submission.CreatedAt = actual.CreatedAt
				tt.want.submission.UpdatedAt = actual.UpdatedAt
				assert.Equal(t, tt.want.submission, actual)
			}
		})
	}
}

func testTeacherSubmission(teacherID string, summaryID int64, decided bool, now time.Time) *entity.TeacherSubmission {
	return &entity.TeacherSubmission{
		TeacherID:      teacherID,
		ShiftSummaryID: summaryID,
		Decided:        decided,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
}
