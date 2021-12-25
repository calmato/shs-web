package database

import (
	"context"
	"testing"
	"time"

	"github.com/calmato/shs-web/api/internal/classroom/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSchedule_List(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, scheduleTable)

	now := jst.Now()

	lessons := entity.Lessons{
		{StartTime: "1700", EndTime: "1830"},
		{StartTime: "1830", EndTime: "2000"},
		{StartTime: "2000", EndTime: "2100"},
	}
	schedules := make(entity.Schedules, 2)
	schedules[0] = testSchedule(time.Sunday, true, nil, now)
	schedules[1] = testSchedule(time.Monday, false, lessons, now)
	err = m.db.DB.Create(&schedules).Error
	require.NoError(t, err)

	type args struct{}
	type want struct {
		schedules entity.Schedules
		isErr     bool
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
				schedules: schedules,
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

			db := NewSchedule(m.db)
			actual, err := db.List(ctx)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			for i, schedule := range tt.want.schedules {
				schedule.LessonsJSON = actual[i].LessonsJSON
				schedule.CreatedAt = actual[i].CreatedAt
				schedule.UpdatedAt = actual[i].UpdatedAt
				assert.Contains(t, actual, schedule)
			}
		})
	}
}

func TestSchedule_Get(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, scheduleTable)

	now := jst.Now()

	lessons := entity.Lessons{
		{StartTime: "1700", EndTime: "1830"},
		{StartTime: "1830", EndTime: "2000"},
		{StartTime: "2000", EndTime: "2100"},
	}
	schedule := testSchedule(time.Monday, false, lessons, now)
	err = m.db.DB.Create(&schedule).Error
	require.NoError(t, err)

	type args struct {
		weekday time.Weekday
	}
	type want struct {
		schedule *entity.Schedule
		isErr    bool
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
				weekday: time.Monday,
			},
			want: want{
				schedule: schedule,
				isErr:    false,
			},
		},
		{
			name:  "not found",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {},
			args: args{
				weekday: time.Friday,
			},
			want: want{
				schedule: nil,
				isErr:    true,
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

			db := NewSchedule(m.db)
			actual, err := db.Get(ctx, tt.args.weekday)
			if tt.want.isErr {
				assert.Error(t, err)
				assert.Nil(t, actual)
			} else {
				schedule.LessonsJSON = actual.LessonsJSON // ignore
				schedule.CreatedAt = actual.CreatedAt     // ignore
				schedule.UpdatedAt = actual.UpdatedAt     // ignore
				assert.NoError(t, err)
				assert.Equal(t, tt.want.schedule, actual)
			}
		})
	}
}

func TestSchedule_MultipleUpdate(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, scheduleTable)

	now := jst.Now()

	lessons := entity.Lessons{
		{StartTime: "1700", EndTime: "1830"},
		{StartTime: "1830", EndTime: "2000"},
		{StartTime: "2000", EndTime: "2100"},
	}
	schedules := make(entity.Schedules, 1)
	schedules[0] = testSchedule(time.Monday, false, lessons, now)

	type args struct {
		schedules entity.Schedules
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
				err := m.db.DB.Create(&schedules).Error
				require.NoError(t, err)
			},
			args: args{
				schedules: entity.Schedules{
					{
						Weekday:  time.Monday,
						IsClosed: true,
						Lessons:  nil,
					},
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
				schedules: schedules,
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

			_ = m.dbDelete(ctx, scheduleTable)
			tt.setup(ctx, t, m)

			db := NewSchedule(m.db)
			err := db.MultipleUpdate(ctx, tt.args.schedules)
			assert.Equal(t, tt.want.isErr, err != nil, err)
		})
	}
}

func testSchedule(weekday time.Weekday, closed bool, lessons entity.Lessons, now time.Time) *entity.Schedule {
	schedule := &entity.Schedule{
		Weekday:   weekday,
		IsClosed:  false,
		Lessons:   lessons,
		CreatedAt: now,
		UpdatedAt: now,
	}
	_ = schedule.FillJSON()
	return schedule
}
