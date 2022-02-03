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

func TestStudentShiftTemplate_Get(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, studentShiftTemplateTable)

	now := jst.Date(2021, 12, 10, 12, 0, 0, 0)

	const studentID = "studentid"
	template := testStudentShiftTemplate(studentID, now)
	err = m.db.DB.Create(&template).Error
	require.NoError(t, err)

	type args struct {
		studentID string
	}
	type want struct {
		template *entity.StudentShiftTemplate
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
				studentID: studentID,
			},
			want: want{
				template: template,
				isErr:    false,
			},
		},
		{
			name:  "not found",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {},
			args: args{
				studentID: "",
			},
			want: want{
				template: nil,
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

			db := NewStudentShiftTemplate(m.db)
			actual, err := db.Get(ctx, tt.args.studentID)
			if tt.want.isErr {
				assert.Error(t, err)
				return
			}
			tt.want.template.CreatedAt, tt.want.template.UpdatedAt = actual.CreatedAt, actual.UpdatedAt
			tt.want.template.SchedulesJSON = actual.SchedulesJSON
			tt.want.template.SuggestedLessonsJSON = actual.SuggestedLessonsJSON
			assert.NoError(t, err)
			assert.Equal(t, actual, tt.want.template)
		})
	}
}

func TestStudentShiftTemplate_Upsert(t *testing.T) {
	m, err := newMock()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = m.dbDelete(ctx, studentShiftTemplateTable)

	now := jst.Date(2021, 12, 10, 12, 0, 0, 0)

	const studentID = "studentid"
	template := testStudentShiftTemplate(studentID, now)

	type args struct {
		studentID string
		template  *entity.StudentShiftTemplate
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
			name:  "success to create",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {},
			args: args{
				studentID: studentID,
				template:  template,
			},
			want: want{
				isErr: false,
			},
		},
		{
			name: "success to update",
			setup: func(ctx context.Context, t *testing.T, m *mocks) {
				err := m.db.DB.Create(&template).Error
				require.NoError(t, err)
			},
			args: args{
				studentID: studentID,
				template:  template,
			},
			want: want{
				isErr: false,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			m.dbDelete(ctx, studentShiftTemplateTable)
			tt.setup(ctx, t, m)

			db := NewStudentShiftTemplate(m.db)
			err := db.Upsert(ctx, tt.args.studentID, tt.args.template)
			assert.Equal(t, tt.want.isErr, err != nil, err)
		})
	}
}

func testStudentShiftTemplate(studentID string, now time.Time) *entity.StudentShiftTemplate {
	template := &entity.StudentShiftTemplate{
		StudentID: studentID,
		Schedules: entity.ShiftSchedules{
			{
				Weekday: time.Sunday,
				Lessons: entity.LessonSchedules{
					{StartTime: "1700", EndTime: "1830"},
					{StartTime: "1830", EndTime: "2000"},
				},
			},
		},
		SuggestedLessons: entity.SuggestedLessons{
			{SubjectID: 1, Total: 4},
		},
	}
	_ = template.FillJSON()
	return template
}
