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

func TestStudentSubmission_ListByShiftSummaryIDs(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, studentSubmissionTable, shiftSummaryTable)

	now := jst.Date(2021, 12, 10, 12, 0, 0, 0)

	const studentID = "studentid"

	summaries := make(entity.ShiftSummaries, 3)
	summaries[0] = testShiftSummary(1, 202202, jst.Date(2022, 1, 1, 0, 0, 0, 0), jst.Date(2022, 1, 15, 0, 0, 0, 0), now)
	summaries[1] = testShiftSummary(2, 202203, jst.Date(2022, 2, 1, 0, 0, 0, 0), jst.Date(2022, 2, 15, 0, 0, 0, 0), now)
	summaries[2] = testShiftSummary(3, 202204, jst.Date(2022, 3, 1, 0, 0, 0, 0), jst.Date(2022, 3, 15, 0, 0, 0, 0), now)
	err = m.db.DB.Create(&summaries).Error
	require.NoError(t, err)

	submissions := make(entity.StudentSubmissions, 2)
	submissions[0] = testStudentSubmission(studentID, 1, true, now)
	submissions[1] = testStudentSubmission(studentID, 2, false, now)
	err = m.db.DB.Create(&submissions).Error
	require.NoError(t, err)

	type args struct {
		studentID  string
		summaryIDs []int64
	}
	type want struct {
		submissions entity.StudentSubmissions
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
				studentID:  studentID,
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
				studentID:  studentID,
				summaryIDs: []int64{0},
			},
			want: want{
				submissions: entity.StudentSubmissions{},
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

			db := NewStudentSubmission(m.db)
			actual, err := db.ListByShiftSummaryIDs(ctx, tt.args.studentID, tt.args.summaryIDs)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			assert.Len(t, actual, len(tt.want.submissions))
			for i, expect := range tt.want.submissions {
				expect.CreatedAt, expect.UpdatedAt = actual[i].CreatedAt, actual[i].UpdatedAt
				expect.SuggestedLessonsJSON = actual[i].SuggestedLessonsJSON
				assert.Contains(t, actual, expect)
			}
		})
	}
}

func TestStudentSubmission_ListByStudentIDs(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, studentSubmissionTable, shiftSummaryTable)

	now := jst.Date(2021, 12, 10, 12, 0, 0, 0)

	summary := testShiftSummary(1, 202202, jst.Date(2022, 1, 1, 0, 0, 0, 0), jst.Date(2022, 1, 15, 0, 0, 0, 0), now)
	err = m.db.DB.Create(&summary).Error
	require.NoError(t, err)

	submissions := make(entity.StudentSubmissions, 2)
	submissions[0] = testStudentSubmission("studentid1", 1, true, now)
	submissions[1] = testStudentSubmission("studentid2", 1, false, now)
	err = m.db.DB.Create(&submissions).Error
	require.NoError(t, err)

	type args struct {
		studentIDs []string
		summaryID  int64
	}
	type want struct {
		submissions entity.StudentSubmissions
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
				studentIDs: []string{"studentid1", "studentid2"},
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
				studentIDs: []string{"student"},
				summaryID:  1,
			},
			want: want{
				submissions: entity.StudentSubmissions{},
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

			db := NewStudentSubmission(m.db)
			actual, err := db.ListByStudentIDs(ctx, tt.args.studentIDs, tt.args.summaryID)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			assert.Len(t, actual, len(tt.want.submissions))
			for i, expect := range tt.want.submissions {
				expect.CreatedAt, expect.UpdatedAt = actual[i].CreatedAt, actual[i].UpdatedAt
				expect.SuggestedLessonsJSON = actual[i].SuggestedLessonsJSON
				assert.Contains(t, actual, expect)
			}
		})
	}
}

func TestStudentSubmission_Get(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, studentSubmissionTable, shiftSummaryTable)

	now := jst.Date(2021, 12, 10, 12, 0, 0, 0)

	const studentID = "studentid"

	summary := testShiftSummary(1, 202202, jst.Date(2022, 1, 1, 0, 0, 0, 0), jst.Date(2022, 1, 15, 0, 0, 0, 0), now)
	err = m.db.DB.Create(&summary).Error
	require.NoError(t, err)

	submission := testStudentSubmission(studentID, 1, true, now)
	err = m.db.DB.Create(&submission).Error
	require.NoError(t, err)

	type args struct {
		studentID string
		summaryID int64
	}
	type want struct {
		submission *entity.StudentSubmission
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
				studentID: studentID,
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
				studentID: studentID,
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

			db := NewStudentSubmission(m.db)
			actual, err := db.Get(ctx, tt.args.studentID, tt.args.summaryID)
			if tt.want.isErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				tt.want.submission.SuggestedLessonsJSON = actual.SuggestedLessonsJSON
				tt.want.submission.CreatedAt = actual.CreatedAt
				tt.want.submission.UpdatedAt = actual.UpdatedAt
				assert.Equal(t, tt.want.submission, actual)
			}
		})
	}
}

func testStudentSubmission(studentID string, summaryID int64, decided bool, now time.Time) *entity.StudentSubmission {
	submission := &entity.StudentSubmission{
		StudentID:      studentID,
		ShiftSummaryID: summaryID,
		Decided:        decided,
		SuggestedLessons: entity.SuggestedLessons{
			{SubjectID: 1, Total: 4},
			{SubjectID: 2, Total: 4},
		},
		CreatedAt: now,
		UpdatedAt: now,
	}
	_ = submission.FillJSON()
	return submission
}
