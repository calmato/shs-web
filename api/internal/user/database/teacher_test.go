package database

import (
	"context"
	"strings"
	"testing"

	"github.com/calmato/shs-web/api/internal/user/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
