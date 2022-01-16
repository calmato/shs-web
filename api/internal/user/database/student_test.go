package database

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/calmato/shs-web/api/internal/user/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStudents_List(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, studentTable)

	now := jst.Now()

	students := make(entity.Students, 3)
	students[0] = testStudent("cvcTyJFfgoDQrqC1KDHbRe", "student01@calmato.jp", now)
	students[1] = testStudent("jx2NB7t3xodUu53LYtYTf2", "student02@calmato.jp", now)
	students[2] = testStudent("kvnMftmwoVsCzZRKNTEZtg", "student03@calmato.jp", now)
	err = m.db.DB.Create(&students).Error
	require.NoError(t, err)

	type args struct {
		params *ListStudentsParams
	}
	type want struct {
		students entity.Students
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
				params: &ListStudentsParams{},
			},
			want: want{
				students: students,
				isErr:    false,
			},
		},
		{
			name:  "success with limit and offset",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {},
			args: args{
				params: &ListStudentsParams{
					Limit:  1,
					Offset: 1,
				},
			},
			want: want{
				students: students[1:1],
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

			db := NewStudent(m.db, m.auth)
			actual, err := db.List(ctx, tt.args.params)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			for i, student := range tt.want.students {
				student.CreatedAt = actual[i].CreatedAt
				student.UpdatedAt = actual[i].UpdatedAt
				assert.Contains(t, actual, student)
			}
		})
	}
}

func TestStudent_Get(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, studentTable)

	now := jst.Now()

	student := testStudent("cvcTyJFfgoDQnqC1KDHbKr", "student01@calmato.jp", now)
	err = m.db.DB.Create(&student).Error
	require.NoError(t, err)

	type args struct {
		id string
	}
	type want struct {
		student *entity.Student
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
				id: "cvcTyJFfgoDQnqC1KDHbKr",
			},
			want: want{
				student: student,
				isErr:   false,
			},
		},
		{
			name:  "not found",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {},
			args: args{
				id: "jx2NC7t3yodUu53LMtYLf1",
			},
			want: want{
				student: nil,
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

			db := NewStudent(m.db, m.auth)
			actual, err := db.Get(ctx, tt.args.id)
			if tt.want.isErr {
				assert.Error(t, err)
				assert.Nil(t, actual)
			} else {
				student.CreatedAt = actual.CreatedAt
				student.UpdatedAt = actual.UpdatedAt
				assert.NoError(t, err)
				assert.Equal(t, tt.want.student, actual)
			}
		})
	}
}

func TestStudent_Create(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	type args struct {
		student *entity.Student
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
				student: &entity.Student{
					ID:            idmock,
					LastName:      "浜田",
					FirstName:     "直志",
					LastNameKana:  "はまだ",
					FirstNameKana: "ただし",
					Mail:          "student-test001@calmato.jp",
					Password:      "12345678",
					BirthYear:     2005,
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
				student: &entity.Student{
					ID:       idmock,
					Mail:     "student-test001@calmato.jp",
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
				student: &entity.Student{
					ID:        idmock,
					Mail:      "student-test001@calmato.jp",
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
			_ = m.dbDelete(ctx, studentTable)
			tt.setup(ctx, t, m)

			db := NewStudent(m.db, m.auth)
			err := db.Create(ctx, tt.args.student)
			assert.Equal(t, tt.want.isErr, err != nil, err)
		})
	}
}

func TestStudent_Count(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, studentTable)

	now := jst.Now()

	Students := make(entity.Students, 3)
	Students[0] = testStudent("cvcTyJFfgoDQrqC1KDHbRe", "student01@calmato.jp", now)
	Students[1] = testStudent("jx2NB7t3xodUu53LYtYTf2", "student02@calmato.jp", now)
	Students[2] = testStudent("kvnMftmwoVsCzZRKNTEZtg", "student03@calmato.jp", now)
	err = m.db.DB.Create(&Students).Error
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

			db := NewStudent(m.db, m.auth)
			actual, err := db.Count(ctx)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			assert.Equal(t, tt.want.total, actual)
		})
	}
}

func testStudent(id string, mail string, now time.Time) *entity.Student {
	return &entity.Student{
		ID:            id,
		LastName:      "浜田",
		FirstName:     "直志",
		LastNameKana:  "はまだ",
		FirstNameKana: "ただし",
		Mail:          mail,
		BirthYear:     2005,
		CreatedAt:     now,
		UpdatedAt:     now,
	}
}
