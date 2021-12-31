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

func TestShiftSummary_List(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, shiftSummaryTable)

	now := jst.Now()

	summaries := make(entity.ShiftSummaries, 4)
	summaries[0] = testShiftSummary(1, 202201, now.AddDate(0, -1, 0), now.AddDate(0, -1, 1), now)
	summaries[1] = testShiftSummary(2, 202202, now.AddDate(0, 0, -1), now.AddDate(0, 0, 1), now)
	summaries[2] = testShiftSummary(3, 202203, now.AddDate(0, 1, 0), now.AddDate(0, 1, 1), now)
	summaries[3] = testShiftSummary(4, 202204, now.AddDate(0, 2, 0), now.AddDate(0, 2, 1), now)
	summaries.Fill(now)
	err = m.db.DB.Create(&summaries).Error
	require.NoError(t, err)

	type args struct {
		params *ListShiftSummariesParams
	}
	type want struct {
		summaries entity.ShiftSummaries
		isErr     bool
	}
	tests := []struct {
		name  string
		setup func(ctx context.Context, t *testing.T, m *mocks)
		args  args
		want  want
	}{
		{
			name:  "success order by asc",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {},
			args: args{
				params: &ListShiftSummariesParams{
					OrderBy: OrderByAsc,
				},
			},
			want: want{
				summaries: summaries,
				isErr:     false,
			},
		},
		{
			name:  "success order by desc",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {},
			args: args{
				params: &ListShiftSummariesParams{
					OrderBy: OrderByDesc,
				},
			},
			want: want{
				summaries: summaries,
				isErr:     false,
			},
		},
		{
			name:  "success with params",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {},
			args: args{
				params: &ListShiftSummariesParams{
					Limit:  10,
					Offset: 1,
					Status: entity.ShiftStatusWaiting,
				},
			},
			want: want{
				summaries: summaries[3:],
				isErr:     false,
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

			db := NewShiftSummary(m.db)
			actual, err := db.List(ctx, tt.args.params)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			assert.Len(t, actual, len(tt.want.summaries))

			actualMap := actual.Map()
			for _, s := range tt.want.summaries {
				summary := actualMap[s.ID]
				require.NotNil(t, summary)
				assert.Equal(t, s.Status, summary.Status)
			}
		})
	}
}

func TestShiftSummary_Get(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, shiftSummaryTable)

	now := jst.Now()

	summary := testShiftSummary(1, 202201, now.AddDate(0, -1, 0), now.AddDate(0, -1, 1), now)
	summary.Fill(now)
	err = m.db.DB.Create(&summary).Error
	require.NoError(t, err)

	type args struct {
		summaryID int64
	}
	type want struct {
		summary *entity.ShiftSummary
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
				summaryID: 1,
			},
			want: want{
				summary: summary,
				isErr:   false,
			},
		},
		{
			name:  "not found",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {},
			args: args{
				summaryID: 2,
			},
			want: want{
				summary: nil,
				isErr:   true,
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

			db := NewShiftSummary(m.db)
			actual, err := db.Get(ctx, tt.args.summaryID)
			if tt.want.isErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want.summary.Status, actual.Status)
			}
		})
	}
}

func TestShiftSummary_UpdateSchedule(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, shiftSummaryTable)

	now := jst.Now()
	summary := testShiftSummary(1, 202202, now.AddDate(0, 0, -1), now.AddDate(0, 0, 1), now)

	type args struct {
		summaryID int64
		openAt    time.Time
		endAt     time.Time
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
				err = m.db.DB.Create(&summary).Error
				require.NoError(t, err)
			},
			args: args{
				summaryID: 1,
				openAt:    jst.Now(),
				endAt:     jst.Now(),
			},
			want: want{
				isErr: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			_ = m.dbDelete(ctx, shiftSummaryTable)
			tt.setup(ctx, t, m)

			db := NewShiftSummary(m.db)
			err := db.UpdateSchedule(ctx, tt.args.summaryID, tt.args.openAt, tt.args.endAt)
			assert.Equal(t, tt.want.isErr, err != nil, err)
		})
	}
}

func TestShiftSummary_Delete(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, shiftSummaryTable)

	now := jst.Now()

	summary := testShiftSummary(1, 202202, now.AddDate(0, 0, -1), now.AddDate(0, 0, 1), now)
	shifts := make(entity.Shifts, 3)
	shifts[0] = testShift(1, 1, jst.Date(2022, 2, 1, 0, 0, 0, 0), "1700", "1830", now)
	shifts[1] = testShift(2, 1, jst.Date(2022, 2, 1, 0, 0, 0, 0), "1830", "2000", now)
	shifts[2] = testShift(3, 1, jst.Date(2022, 2, 2, 0, 0, 0, 0), "1700", "1830", now)

	type args struct {
		summaryID int64
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
				err = m.db.DB.Create(&summary).Error
				require.NoError(t, err)
				err = m.db.DB.Create(&shifts).Error
				require.NoError(t, err)
			},
			args: args{
				summaryID: 1,
			},
			want: want{
				isErr: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			_ = m.dbDelete(ctx, shiftSummaryTable)
			tt.setup(ctx, t, m)

			db := NewShiftSummary(m.db)
			err := db.Delete(ctx, tt.args.summaryID)
			assert.Equal(t, tt.want.isErr, err != nil, err)
		})
	}
}

func TestShiftSummary_Count(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, shiftSummaryTable)

	now := jst.Now()

	summaries := make(entity.ShiftSummaries, 4)
	summaries[0] = testShiftSummary(1, 202201, now.AddDate(0, -1, 0), now.AddDate(0, -1, 1), now)
	summaries[1] = testShiftSummary(2, 202202, now.AddDate(0, 0, -1), now.AddDate(0, 0, 1), now)
	summaries[2] = testShiftSummary(3, 202203, now.AddDate(0, 1, 0), now.AddDate(0, 1, 1), now)
	summaries[3] = testShiftSummary(4, 202204, now.AddDate(0, 2, 0), now.AddDate(0, 2, 1), now)
	summaries.Fill(now)
	err = m.db.DB.Create(&summaries).Error
	require.NoError(t, err)

	type args struct{}
	type want struct {
		total int64
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
			args:  args{},
			want: want{
				total: 4,
				isErr: false,
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

			db := NewShiftSummary(m.db)
			actual, err := db.Count(ctx)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			assert.Equal(t, tt.want.total, actual)
		})
	}
}

func testShiftSummary(id int64, yearMonth int32, openAt, endAt, now time.Time) *entity.ShiftSummary {
	return &entity.ShiftSummary{
		ID:        id,
		YearMonth: yearMonth,
		OpenAt:    openAt,
		EndAt:     endAt,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
