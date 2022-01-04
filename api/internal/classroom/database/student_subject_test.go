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

func TestStudentSubject_List(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, studentSubjectTable, subjectTable)

	now := jst.Now()

	const studentID1, studentID2 = "studentid1", "studentid2"

	subjects := make(entity.Subjects, 3)
	subjects[0] = testSubject(1, "国語", entity.SchoolTypeElementarySchool, now)
	subjects[1] = testSubject(2, "数学", entity.SchoolTypeJuniorHighSchool, now)
	subjects[2] = testSubject(3, "英語", entity.SchoolTypeHighSchool, now)
	err = m.db.DB.Create(&subjects).Error
	require.NoError(t, err)

	studentSubjects := make(entity.StudentSubjects, 4)
	studentSubjects[0] = testStudentSubject(studentID1, 1, now)
	studentSubjects[1] = testStudentSubject(studentID1, 2, now)
	studentSubjects[2] = testStudentSubject(studentID1, 3, now)
	studentSubjects[3] = testStudentSubject(studentID2, 1, now)
	err = m.db.DB.Create(&studentSubjects).Error
	require.NoError(t, err)

	type args struct {
		studentIDs []string
	}
	type want struct {
		subjects entity.StudentSubjects
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
				studentIDs: []string{studentID1, studentID2},
			},
			want: want{
				subjects: studentSubjects,
				isErr:    false,
			},
		},
		{
			name:  "success empty",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {},
			args: args{
				studentIDs: []string{},
			},
			want: want{
				subjects: entity.StudentSubjects{},
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

			db := NewStudentSubject(m.db)
			actual, err := db.ListByStudentIDs(ctx, tt.args.studentIDs)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			for i, subject := range tt.want.subjects {
				subject.CreatedAt = actual[i].CreatedAt
				subject.UpdatedAt = actual[i].UpdatedAt
				assert.Contains(t, actual, subject)
			}
		})
	}
}

func TestStudentSubject_Replace(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, studentSubjectTable, subjectTable)

	now := jst.Now()

	const studentID1, studentID2 = "studentid1", "studentid2"

	subjects := make(entity.Subjects, 5)
	subjects[0] = testSubject(1, "国語", entity.SchoolTypeElementarySchool, now)
	subjects[1] = testSubject(2, "数学", entity.SchoolTypeJuniorHighSchool, now)
	subjects[2] = testSubject(3, "社会", entity.SchoolTypeHighSchool, now)
	subjects[3] = testSubject(4, "理科", entity.SchoolTypeHighSchool, now)
	subjects[4] = testSubject(5, "英語", entity.SchoolTypeHighSchool, now)
	err = m.db.DB.Create(&subjects).Error
	require.NoError(t, err)

	studentSubjects := make(entity.StudentSubjects, 4)
	studentSubjects[0] = testStudentSubject(studentID1, 1, now)
	studentSubjects[1] = testStudentSubject(studentID1, 2, now)
	studentSubjects[2] = testStudentSubject(studentID1, 3, now)
	studentSubjects[2] = testStudentSubject(studentID1, 4, now)
	studentSubjects[3] = testStudentSubject(studentID2, 1, now)

	type args struct {
		schoolType entity.SchoolType
		subjects   entity.StudentSubjects
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
				subjects: entity.StudentSubjects{
					testStudentSubject(studentID1, 3, now),
					testStudentSubject(studentID1, 4, now),
					testStudentSubject(studentID1, 5, now),
				},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name: "success when exists",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {
				err := m.db.DB.Create(studentSubjects).Error
				require.NoError(t, err)
			},
			args: args{
				schoolType: entity.SchoolTypeHighSchool,
				subjects: entity.StudentSubjects{
					testStudentSubject(studentID1, 4, now),
					testStudentSubject(studentID1, 5, now),
				},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name: "success to subject length is 0",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {
				err := m.db.DB.Create(studentSubjects).Error
				require.NoError(t, err)
			},
			args: args{
				schoolType: entity.SchoolTypeHighSchool,
				subjects:   entity.StudentSubjects{},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name:  "failed to insert subject",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {},
			args: args{
				schoolType: entity.SchoolTypeHighSchool,
				subjects: entity.StudentSubjects{
					testStudentSubject(studentID1, 0, now),
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

			_ = m.dbDelete(ctx, studentSubjectTable)
			tt.setup(ctx, t, m)

			db := NewStudentSubject(m.db)
			err := db.Replace(ctx, tt.args.schoolType, tt.args.subjects)
			assert.Equal(t, tt.want.isErr, err != nil, err)
		})
	}
}

func testStudentSubject(studentID string, subjectID int64, now time.Time) *entity.StudentSubject {
	return &entity.StudentSubject{
		StudentID: studentID,
		SubjectID: subjectID,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
