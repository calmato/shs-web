package entity

import (
	"testing"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/stretchr/testify/assert"
)

func TestTeacherSubject_SubjectIDs(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name     string
		subjects TeacherSubjects
		expect   []int64
	}{
		{
			name: "success",
			subjects: TeacherSubjects{
				{
					TeacherID: "teacherid1",
					SubjectID: 1,
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					TeacherID: "teacherid1",
					SubjectID: 2,
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					TeacherID: "teacherid2",
					SubjectID: 1,
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
			expect: []int64{1, 2},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			subjectIDs := tt.subjects.SubjectIDs()
			for _, subjectID := range tt.expect {
				assert.Contains(t, subjectIDs, subjectID)
			}
		})
	}
}

func TestTeacherSubject_GroupByTeacherID(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name     string
		subjects TeacherSubjects
		expect   map[string]TeacherSubjects
	}{
		{
			name: "success",
			subjects: TeacherSubjects{
				{
					TeacherID: "teacherid1",
					SubjectID: 1,
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					TeacherID: "teacherid1",
					SubjectID: 2,
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					TeacherID: "teacherid2",
					SubjectID: 1,
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
			expect: map[string]TeacherSubjects{
				"teacherid1": {
					{
						TeacherID: "teacherid1",
						SubjectID: 1,
						CreatedAt: now,
						UpdatedAt: now,
					},
					{
						TeacherID: "teacherid1",
						SubjectID: 2,
						CreatedAt: now,
						UpdatedAt: now,
					},
				},
				"teacherid2": {
					{
						TeacherID: "teacherid2",
						SubjectID: 1,
						CreatedAt: now,
						UpdatedAt: now,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.subjects.GroupByTeacherID())
		})
	}
}

func TestTeacherSubject_TeacherProto(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name     string
		subjects TeacherSubjects
		expect   *classroom.TeacherSubject
	}{
		{
			name: "success",
			subjects: TeacherSubjects{
				{
					TeacherID: "teacherid1",
					SubjectID: 1,
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					TeacherID: "teacherid1",
					SubjectID: 2,
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
			expect: &classroom.TeacherSubject{
				TeacherId:  "teacherid1",
				SubjectIds: []int64{1, 2},
			},
		},
		{
			name:     "success empty",
			subjects: TeacherSubjects{},
			expect:   nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.subjects.TeacherProto())
		})
	}
}

func TestTeacherSubject_TeachersProto(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name     string
		subjects TeacherSubjects
		expect   []*classroom.TeacherSubject
	}{
		{
			name: "success",
			subjects: TeacherSubjects{
				{
					TeacherID: "teacherid1",
					SubjectID: 1,
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					TeacherID: "teacherid1",
					SubjectID: 2,
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					TeacherID: "teacherid2",
					SubjectID: 1,
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
			expect: []*classroom.TeacherSubject{
				{
					TeacherId:  "teacherid1",
					SubjectIds: []int64{1, 2},
				},
				{
					TeacherId:  "teacherid2",
					SubjectIds: []int64{1},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.subjects.TeachersProto())
		})
	}
}
