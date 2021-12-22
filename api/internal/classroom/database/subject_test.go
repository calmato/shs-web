package database

import (
	"context"
	"testing"
	"time"

	"github.com/calmato/shs-web/api/internal/classroom/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubject_List(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, subjectTable)

	now := jst.Now()

	subjects := make(entity.Subjects, 3)
	subjects[0] = testSubject(1, "国語", entity.SchoolTypeElementarySchool, now)
	subjects[1] = testSubject(2, "数学", entity.SchoolTypeElementarySchool, now)
	subjects[2] = testSubject(3, "英語", entity.SchoolTypeElementarySchool, now)
	err = m.db.DB.Create(&subjects).Error
	require.NoError(t, err)

	type args struct {
		params *ListSubjectsParams
	}
	type want struct {
		subjects entity.Subjects
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
				params: &ListSubjectsParams{},
			},
			want: want{
				subjects: subjects,
				isErr:    false,
			},
		},
		{
			name:  "success with school_type",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {},
			args: args{
				params: &ListSubjectsParams{
					SchoolType: classroom.SchoolType_SCHOOL_TYPE_JUNIOR_HIGH_SCHOOL,
				},
			},
			want: want{
				subjects: subjects[1:1],
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

			db := NewSubject(m.db)
			actual, err := db.List(ctx, tt.args.params)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			for i, subject := range tt.want.subjects {
				subject.CreatedAt = actual[i].CreatedAt
				subject.UpdatedAt = actual[i].UpdatedAt
				assert.Contains(t, actual, subject)
			}
		})
	}
}

func TestSubject_MultiGet(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, subjectTable)

	now := jst.Now()

	subjects := make(entity.Subjects, 3)
	subjects[0] = testSubject(1, "国語", entity.SchoolTypeElementarySchool, now)
	subjects[1] = testSubject(2, "数学", entity.SchoolTypeElementarySchool, now)
	subjects[2] = testSubject(3, "英語", entity.SchoolTypeElementarySchool, now)
	err = m.db.DB.Create(&subjects).Error
	require.NoError(t, err)

	type args struct {
		ids []int64
	}
	type want struct {
		subjects entity.Subjects
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
				ids: []int64{1, 2, 3},
			},
			want: want{
				subjects: subjects,
				isErr:    false,
			},
		},
		{
			name:  "success to empty",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {},
			args: args{
				ids: []int64{0},
			},
			want: want{
				subjects: entity.Subjects{},
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

			db := NewSubject(m.db)
			actual, err := db.MultiGet(ctx, tt.args.ids)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			assert.Len(t, actual, len(tt.want.subjects))
			for i, subject := range tt.want.subjects {
				subject.CreatedAt = actual[i].CreatedAt
				subject.UpdatedAt = actual[i].UpdatedAt
				assert.Contains(t, actual, subject)
			}
		})
	}
}

func TestSubject_Get(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, subjectTable)

	now := jst.Now()

	subject := testSubject(1, "国語", entity.SchoolTypeElementarySchool, now)
	err = m.db.DB.Create(&subject).Error
	require.NoError(t, err)

	type args struct {
		id int64
	}
	type want struct {
		subject *entity.Subject
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
				id: 1,
			},
			want: want{
				subject: subject,
				isErr:   false,
			},
		},
		{
			name:  "not found",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {},
			args: args{
				id: 0,
			},
			want: want{
				subject: nil,
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

			db := NewSubject(m.db)
			actual, err := db.Get(ctx, tt.args.id)
			if tt.want.isErr {
				assert.Error(t, err)
				assert.Nil(t, actual)
			} else {
				subject.CreatedAt = actual.CreatedAt // ignore
				subject.UpdatedAt = actual.UpdatedAt // ignore
				assert.NoError(t, err)
				assert.Equal(t, tt.want.subject, actual)
			}
		})
	}
}

func TestSubject_Count(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, subjectTable)

	now := jst.Now()

	subjects := make(entity.Subjects, 3)
	subjects[0] = testSubject(1, "国語", entity.SchoolTypeElementarySchool, now)
	subjects[1] = testSubject(2, "数学", entity.SchoolTypeElementarySchool, now)
	subjects[2] = testSubject(3, "英語", entity.SchoolTypeElementarySchool, now)
	err = m.db.DB.Create(&subjects).Error
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

			db := NewSubject(m.db)
			actual, err := db.Count(ctx)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			assert.Equal(t, tt.want.total, actual)
		})
	}
}

func testSubject(id int64, name string, schoolType entity.SchoolType, now time.Time) *entity.Subject {
	return &entity.Subject{
		ID:         id,
		Name:       name,
		Color:      "#F8BBD0",
		SchoolType: schoolType,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}
