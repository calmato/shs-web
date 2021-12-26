package database

import (
	"context"
	"testing"
	"time"

	"github.com/calmato/shs-web/api/internal/user/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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

func testStudent(id string, mail string, now time.Time) *entity.Student {
	return &entity.Student{
		ID:            id,
		LastName:      "浜田",
		FirstName:     "直志",
		LastNameKana:  "はまだ",
		FirstNameKana: "ただし",
		Mail:          mail,
		BirthYear:     2021,
		CreatedAt:     now,
		UpdatedAt:     now,
	}
}
