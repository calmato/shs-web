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

func TestRoom_Replace(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, roomTable)

	now := jst.Now()

	rooms := make(entity.Rooms, 3)
	rooms[0] = testRoom(1, now)
	rooms[1] = testRoom(2, now)
	rooms[2] = testRoom(3, now)

	type args struct {
		rooms entity.Rooms
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
				err = m.db.DB.Create(&rooms).Error
				require.NoError(t, err)
			},
			args: args{
				rooms: entity.Rooms{
					testRoom(1, now),
					testRoom(2, now),
					testRoom(3, now),
					testRoom(4, now),
				},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name:  "failed to duplicate primary key",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {},
			args: args{
				rooms: entity.Rooms{
					testRoom(1, now),
					testRoom(1, now),
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

			_ = m.dbDelete(ctx, teacherSubjectTable)
			tt.setup(ctx, t, m)

			db := NewRoom(m.db)
			err := db.Replace(ctx, tt.args.rooms)
			assert.Equal(t, tt.want.isErr, err != nil, err)
		})
	}
}

func TestRoom_Count(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, roomTable)

	now := jst.Now()

	rooms := make(entity.Rooms, 3)
	rooms[0] = testRoom(1, now)
	rooms[1] = testRoom(2, now)
	rooms[2] = testRoom(3, now)
	err = m.db.DB.Create(&rooms).Error
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
				total: 3,
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

			db := NewRoom(m.db)
			actual, err := db.Count(ctx)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			assert.Equal(t, tt.want.total, actual)
		})
	}
}

func testRoom(id int32, now time.Time) *entity.Room {
	return &entity.Room{
		ID:        id,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
