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

func TestTeacherSubject_List(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, teacherSubjectTable, subjectTable)

	now := jst.Now()

	const teacherID1, teacherID2 = "teacherid1", "teacherid2"

	subjects := make(entity.Subjects, 3)
	subjects[0] = testSubject(1, "国語", entity.SchoolTypeElementarySchool, now)
	subjects[1] = testSubject(2, "数学", entity.SchoolTypeJuniorHighSchool, now)
	subjects[2] = testSubject(3, "英語", entity.SchoolTypeHighSchool, now)
	err = m.db.DB.Create(&subjects).Error
	require.NoError(t, err)

	teacherSubjects := make(entity.TeacherSubjects, 4)
	teacherSubjects[0] = testTeacherSubject(teacherID1, 1, now)
	teacherSubjects[1] = testTeacherSubject(teacherID1, 2, now)
	teacherSubjects[2] = testTeacherSubject(teacherID1, 3, now)
	teacherSubjects[3] = testTeacherSubject(teacherID2, 1, now)
	err = m.db.DB.Create(&teacherSubjects).Error
	require.NoError(t, err)

	type args struct {
		teacherIDs []string
	}
	type want struct {
		subjects entity.TeacherSubjects
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
				teacherIDs: []string{teacherID1, teacherID2},
			},
			want: want{
				subjects: teacherSubjects,
				isErr:    false,
			},
		},
		{
			name:  "success empty",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {},
			args: args{
				teacherIDs: []string{},
			},
			want: want{
				subjects: entity.TeacherSubjects{},
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

			db := NewTeacherSubject(m.db)
			actual, err := db.ListByTeacherIDs(ctx, tt.args.teacherIDs)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			for i, subject := range tt.want.subjects {
				subject.CreatedAt = actual[i].CreatedAt
				subject.UpdatedAt = actual[i].UpdatedAt
				assert.Contains(t, actual, subject)
			}
		})
	}
}

func TestTeacherSubject_Replace(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, teacherSubjectTable, subjectTable)

	now := jst.Now()

	const teacherID1, teacherID2 = "teacherid1", "teacherid2"

	subjects := make(entity.Subjects, 5)
	subjects[0] = testSubject(1, "国語", entity.SchoolTypeElementarySchool, now)
	subjects[1] = testSubject(2, "数学", entity.SchoolTypeJuniorHighSchool, now)
	subjects[2] = testSubject(3, "社会", entity.SchoolTypeHighSchool, now)
	subjects[3] = testSubject(4, "理科", entity.SchoolTypeHighSchool, now)
	subjects[4] = testSubject(5, "英語", entity.SchoolTypeHighSchool, now)
	err = m.db.DB.Create(&subjects).Error
	require.NoError(t, err)

	teacherSubjects := make(entity.TeacherSubjects, 4)
	teacherSubjects[0] = testTeacherSubject(teacherID1, 1, now)
	teacherSubjects[1] = testTeacherSubject(teacherID1, 2, now)
	teacherSubjects[2] = testTeacherSubject(teacherID1, 3, now)
	teacherSubjects[2] = testTeacherSubject(teacherID1, 4, now)
	teacherSubjects[3] = testTeacherSubject(teacherID2, 1, now)

	type args struct {
		schoolType entity.SchoolType
		subjects   entity.TeacherSubjects
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
			name:  "success when empty",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {},
			args: args{
				schoolType: entity.SchoolTypeHighSchool,
				subjects: entity.TeacherSubjects{
					testTeacherSubject(teacherID1, 3, now),
					testTeacherSubject(teacherID1, 4, now),
					testTeacherSubject(teacherID1, 5, now),
				},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name: "success when exists",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {
				err := m.db.DB.Create(teacherSubjects).Error
				require.NoError(t, err)
			},
			args: args{
				schoolType: entity.SchoolTypeHighSchool,
				subjects: entity.TeacherSubjects{
					testTeacherSubject(teacherID1, 4, now),
					testTeacherSubject(teacherID1, 5, now),
				},
			},
			want: want{
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

			_ = m.dbDelete(ctx, teacherSubjectTable)
			tt.setup(ctx, t, m)

			db := NewTeacherSubject(m.db)
			err := db.Replace(ctx, tt.args.schoolType, tt.args.subjects)
			assert.Equal(t, tt.want.isErr, err != nil, err)
		})
	}
}

func testTeacherSubject(teacherID string, subjectID int64, now time.Time) *entity.TeacherSubject {
	return &entity.TeacherSubject{
		TeacherID: teacherID,
		SubjectID: subjectID,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
