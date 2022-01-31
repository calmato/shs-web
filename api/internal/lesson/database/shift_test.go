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

func TestShift_List(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, shiftTable, shiftSummaryTable)

	now := jst.Now()

	openAt := jst.Date(2022, 1, 1, 0, 0, 0, 0)
	endAt := jst.Date(2021, 1, 14, 23, 59, 59, 0)
	summary := testShiftSummary(1, 202202, openAt, endAt, now)
	err = m.db.DB.Create(&summary).Error
	require.NoError(t, err)

	shifts := make(entity.Shifts, 3)
	shifts[0] = testShift(1, 1, jst.Date(2022, 2, 1, 0, 0, 0, 0), "1700", "1830", now)
	shifts[1] = testShift(2, 1, jst.Date(2022, 2, 1, 0, 0, 0, 0), "1830", "2000", now)
	shifts[2] = testShift(3, 1, jst.Date(2022, 2, 2, 0, 0, 0, 0), "1700", "1830", now)
	err = m.db.DB.Create(&shifts).Error
	require.NoError(t, err)

	type args struct {
		params *ListShiftsParams
	}
	type want struct {
		shifts entity.Shifts
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
				params: &ListShiftsParams{
					ShiftSummaryID: 1,
				},
			},
			want: want{
				shifts: shifts,
				isErr:  false,
			},
		},
		{
			name:  "success to length is 0",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {},
			args: args{
				params: &ListShiftsParams{
					ShiftSummaryID: 100,
					ShiftID:        100,
					Limit:          100,
					Offset:         100,
				},
			},
			want: want{
				shifts: entity.Shifts{},
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

			db := NewShift(m.db)
			actual, err := db.List(ctx, tt.args.params)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			assert.Len(t, actual, len(tt.want.shifts))
			for i, shift := range tt.want.shifts {
				shift.Date = actual[i].Date
				shift.CreatedAt = actual[i].CreatedAt
				shift.UpdatedAt = actual[i].UpdatedAt
				assert.Contains(t, actual, shift)
			}
		})
	}
}

func TestShift_ListByDuration(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, shiftTable, shiftSummaryTable)

	now := jst.Now()

	openAt := jst.Date(2022, 1, 1, 0, 0, 0, 0)
	endAt := jst.Date(2021, 1, 14, 23, 59, 59, 0)
	summary := testShiftSummary(1, 202202, openAt, endAt, now)
	err = m.db.DB.Create(&summary).Error
	require.NoError(t, err)

	shifts := make(entity.Shifts, 3)
	shifts[0] = testShift(1, 1, jst.Date(2022, 2, 1, 0, 0, 0, 0), "1700", "1830", now)
	shifts[1] = testShift(2, 1, jst.Date(2022, 2, 1, 0, 0, 0, 0), "1830", "2000", now)
	shifts[2] = testShift(3, 1, jst.Date(2022, 2, 2, 0, 0, 0, 0), "1700", "1830", now)
	err = m.db.DB.Create(&shifts).Error
	require.NoError(t, err)

	type args struct {
		since time.Time
		until time.Time
	}
	type want struct {
		shifts entity.Shifts
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
				since: jst.BeginningOfDay(jst.Date(2022, 2, 1, 0, 0, 0, 0)),
				until: jst.EndOfDay(jst.Date(2022, 2, 1, 0, 0, 0, 0)),
			},
			want: want{
				shifts: shifts[:2],
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

			db := NewShift(m.db)
			actual, err := db.ListByDuration(ctx, tt.args.since, tt.args.until)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			assert.Len(t, actual, len(tt.want.shifts))
			for i, shift := range tt.want.shifts {
				shift.Date = actual[i].Date
				shift.CreatedAt = actual[i].CreatedAt
				shift.UpdatedAt = actual[i].UpdatedAt
				assert.Contains(t, actual, shift)
			}
		})
	}
}

func TestShift_MultiGet(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, shiftTable, shiftSummaryTable)

	now := jst.Now()

	openAt := jst.Date(2022, 1, 1, 0, 0, 0, 0)
	endAt := jst.Date(2021, 1, 14, 23, 59, 59, 0)
	summary := testShiftSummary(1, 202202, openAt, endAt, now)
	err = m.db.DB.Create(&summary).Error
	require.NoError(t, err)

	shifts := make(entity.Shifts, 3)
	shifts[0] = testShift(1, 1, jst.Date(2022, 2, 1, 0, 0, 0, 0), "1700", "1830", now)
	shifts[1] = testShift(2, 1, jst.Date(2022, 2, 1, 0, 0, 0, 0), "1830", "2000", now)
	shifts[2] = testShift(3, 1, jst.Date(2022, 2, 2, 0, 0, 0, 0), "1700", "1830", now)
	err = m.db.DB.Create(&shifts).Error
	require.NoError(t, err)

	type args struct {
		shiftIDs []int64
	}
	type want struct {
		shifts entity.Shifts
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
				shiftIDs: []int64{1, 2, 3},
			},
			want: want{
				shifts: shifts,
				isErr:  false,
			},
		},
		{
			name:  "success is empty",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {},
			args: args{
				shiftIDs: []int64{},
			},
			want: want{
				shifts: entity.Shifts{},
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

			db := NewShift(m.db)
			actual, err := db.MultiGet(ctx, tt.args.shiftIDs)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			assert.Len(t, actual, len(tt.want.shifts))
			for i, shift := range tt.want.shifts {
				shift.Date = actual[i].Date
				shift.CreatedAt = actual[i].CreatedAt
				shift.UpdatedAt = actual[i].UpdatedAt
				assert.Contains(t, actual, shift)
			}
		})
	}
}

func TestShift_Get(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, shiftTable, shiftSummaryTable)

	now := jst.Now()

	openAt := jst.Date(2022, 1, 1, 0, 0, 0, 0)
	endAt := jst.Date(2021, 1, 14, 23, 59, 59, 0)
	summary := testShiftSummary(1, 202202, openAt, endAt, now)
	err = m.db.DB.Create(&summary).Error
	require.NoError(t, err)

	shift := testShift(1, 1, jst.Date(2022, 2, 1, 0, 0, 0, 0), "1700", "1830", now)
	err = m.db.DB.Create(&shift).Error
	require.NoError(t, err)

	type args struct {
		shiftID int64
	}
	type want struct {
		shift *entity.Shift
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
				shiftID: 1,
			},
			want: want{
				shift: shift,
				isErr: false,
			},
		},
		{
			name:  "not found",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {},
			args: args{
				shiftID: 0,
			},
			want: want{
				shift: nil,
				isErr: true,
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

			db := NewShift(m.db)
			actual, err := db.Get(ctx, tt.args.shiftID)
			if tt.want.isErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			tt.want.shift.Date = actual.Date
			tt.want.shift.CreatedAt = actual.CreatedAt
			tt.want.shift.UpdatedAt = actual.UpdatedAt
			assert.Equal(t, tt.want.shift, actual)
		})
	}
}

func TestShift_MultipleCreate(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, shiftTable, shiftSummaryTable)

	now := jst.Now()

	openAt := jst.Date(2022, 1, 1, 0, 0, 0, 0)
	endAt := jst.Date(2021, 1, 14, 23, 59, 59, 0)
	summary := testShiftSummary(1, 202202, openAt, endAt, now)
	shifts := make(entity.Shifts, 3)
	shifts[0] = testShift(1, 1, jst.Date(2022, 2, 1, 0, 0, 0, 0), "1700", "1830", now)
	shifts[1] = testShift(2, 1, jst.Date(2022, 2, 1, 0, 0, 0, 0), "1830", "2000", now)
	shifts[2] = testShift(3, 1, jst.Date(2022, 2, 2, 0, 0, 0, 0), "1700", "1830", now)

	type args struct {
		summary *entity.ShiftSummary
		shifts  entity.Shifts
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
				summary: &entity.ShiftSummary{
					YearMonth: 202202,
					OpenAt:    openAt,
					EndAt:     endAt,
				},
				shifts: entity.Shifts{
					{
						Date:      jst.Date(2022, 2, 1, 0, 0, 0, 0),
						StartTime: "1700",
						EndTime:   "1830",
					},
					{
						Date:      jst.Date(2022, 2, 1, 0, 0, 0, 0),
						StartTime: "1830",
						EndTime:   "2000",
					},
					{
						Date:      jst.Date(2022, 2, 2, 0, 0, 0, 0),
						StartTime: "1700",
						EndTime:   "1830",
					},
				},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name: "already exists",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {
				err := m.db.DB.Create(&summary).Error
				require.NoError(t, err)
				err = m.db.DB.Create(&shifts).Error
				require.NoError(t, err)
			},
			args: args{
				summary: summary,
				shifts:  shifts,
			},
			want: want{
				isErr: true,
			},
		},
		{
			name:  "failed to create shifts",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {},
			args: args{
				summary: summary,
				shifts:  entity.Shifts{{}},
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

			_ = m.dbDelete(ctx, shiftTable, shiftSummaryTable)
			tt.setup(ctx, t, m)

			db := NewShift(m.db)
			err := db.MultipleCreate(ctx, tt.args.summary, tt.args.shifts)
			assert.Equal(t, tt.want.isErr, err != nil, err)
		})
	}
}

func testShift(id, shiftSummaryID int64, date time.Time, startTime, endTime string, now time.Time) *entity.Shift {
	return &entity.Shift{
		ID:             id,
		ShiftSummaryID: shiftSummaryID,
		Date:           date,
		StartTime:      startTime,
		EndTime:        endTime,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
}
