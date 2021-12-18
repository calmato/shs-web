package database

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/calmato/shs-web/api/internal/user/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTeacher_List(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, teacherTable)

	now := jst.Now()

	teachers := make(entity.Teachers, 3)
	teachers[0] = testTeacher("cvcTyJFfgoDQrqC1KDHbRe", "teacher01@calmato.jp", now)
	teachers[1] = testTeacher("jx2NB7t3xodUu53LYtYTf2", "teacher02@calmato.jp", now)
	teachers[2] = testTeacher("kvnMftmwoVsCzZRKNTEZtg", "teacher03@calmato.jp", now)
	err = m.db.DB.Create(&teachers).Error
	require.NoError(t, err)

	type args struct {
		params *ListTeachersParams
	}
	type want struct {
		teachers entity.Teachers
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
				params: &ListTeachersParams{},
			},
			want: want{
				teachers: teachers,
				isErr:    false,
			},
		},
		{
			name:  "success with limit and offset",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {},
			args: args{
				params: &ListTeachersParams{
					Limit:  1,
					Offset: 1,
				},
			},
			want: want{
				teachers: teachers[1:1],
				isErr:    false,
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

			db := NewTeacher(m.db, m.auth)
			actual, err := db.List(ctx, tt.args.params)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			for i, teacher := range tt.want.teachers {
				teacher.CreatedAt = actual[i].CreatedAt // ignore
				teacher.UpdatedAt = actual[i].UpdatedAt // ignore
				assert.Contains(t, actual, teacher)
			}
		})
	}
}

func TestTeacher_Get(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, teacherTable)

	now := jst.Now()

	teacher := testTeacher("cvcTyJFfgoDQrqC1KDHbRe", "teacher01@calmato.jp", now)
	err = m.db.DB.Create(&teacher).Error
	require.NoError(t, err)

	type args struct {
		id string
	}
	type want struct {
		teacher *entity.Teacher
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
				id: "cvcTyJFfgoDQrqC1KDHbRe",
			},
			want: want{
				teacher: teacher,
				isErr:   false,
			},
		},
		{
			name:  "not found",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {},
			args: args{
				id: "jx2NB7t3xodUu53LYtYTf2",
			},
			want: want{
				teacher: nil,
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

			db := NewTeacher(m.db, m.auth)
			actual, err := db.Get(ctx, tt.args.id)
			if tt.want.isErr {
				assert.Error(t, err)
				assert.Nil(t, actual)
			} else {
				teacher.CreatedAt = actual.CreatedAt // ignore
				teacher.UpdatedAt = actual.UpdatedAt // ignore
				assert.NoError(t, err)
				assert.Equal(t, tt.want.teacher, actual)
			}
		})
	}
}

func TestTeacher_Create(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	type args struct {
		teacher *entity.Teacher
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
				teacher: &entity.Teacher{
					ID:            idmock,
					LastName:      "中村",
					FirstName:     "広大",
					LastNameKana:  "なかむら",
					FirstNameKana: "こうだい",
					Mail:          "teacher-test001@calmato.jp",
					Password:      "12345678",
					Role:          0,
				},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name:  "failed to create in firebase authentication",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {},
			args: args{
				teacher: &entity.Teacher{
					ID:       idmock,
					Mail:     "teacher-test001@calmato.jp",
					Password: "",
				},
			},
			want: want{
				isErr: true,
			},
		},
		{
			name:  "failed to create in mysql",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {},
			args: args{
				teacher: &entity.Teacher{
					ID:        idmock,
					Mail:      "teacher-test001@calmato.jp",
					Password:  "12345678",
					LastName:  strings.Repeat("x", 17),
					FirstName: strings.Repeat("x", 17),
				},
			},
			want: want{
				isErr: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			_ = m.authDelete(ctx, idmock)
			_ = m.dbDelete(ctx, teacherTable)
			tt.setup(ctx, t, m)

			db := NewTeacher(m.db, m.auth)
			err := db.Create(ctx, tt.args.teacher)
			assert.Equal(t, tt.want.isErr, err != nil, err)
		})
	}
}

func TestTeacher_Count(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, teacherTable)

	now := jst.Now()

	teachers := make(entity.Teachers, 3)
	teachers[0] = testTeacher("cvcTyJFfgoDQrqC1KDHbRe", "teacher01@calmato.jp", now)
	teachers[1] = testTeacher("jx2NB7t3xodUu53LYtYTf2", "teacher02@calmato.jp", now)
	teachers[2] = testTeacher("kvnMftmwoVsCzZRKNTEZtg", "teacher03@calmato.jp", now)
	err = m.db.DB.Create(&teachers).Error
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

			db := NewTeacher(m.db, m.auth)
			actual, err := db.Count(ctx)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			assert.Equal(t, tt.want.total, actual)
		})
	}
}

func testTeacher(id string, mail string, now time.Time) *entity.Teacher {
	return &entity.Teacher{
		ID:            id,
		LastName:      "中村",
		FirstName:     "広大",
		LastNameKana:  "なかむら",
		FirstNameKana: "こうだい",
		Mail:          mail,
		Role:          int32(user.Role_ROLE_TEACHER),
		CreatedAt:     now,
		UpdatedAt:     now,
	}
}
