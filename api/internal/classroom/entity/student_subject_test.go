package entity

import (
	"testing"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/stretchr/testify/assert"
)

func TestStudentSubjects(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		studentID  string
		subjectIDs []int64
		expect     StudentSubjects
	}{
		{
			name:       "success",
			studentID:  "studentid",
			subjectIDs: []int64{1},
			expect: StudentSubjects{
				{
					StudentID: "studentid",
					SubjectID: 1,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewStudentSubjects(tt.studentID, tt.subjectIDs))
		})
	}
}

func TestStudentSubject_SubjectIDs(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name     string
		subjects StudentSubjects
		expect   []int64
	}{
		{
			name: "success",
			subjects: StudentSubjects{
				{
					StudentID: "studentid1",
					SubjectID: 1,
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					StudentID: "studentid1",
					SubjectID: 2,
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					StudentID: "studentid2",
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

func TestStudentSubject_GroupByStudentID(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name     string
		subjects StudentSubjects
		expect   map[string]StudentSubjects
	}{
		{
			name: "success",
			subjects: StudentSubjects{
				{
					StudentID: "studentid1",
					SubjectID: 1,
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					StudentID: "studentid1",
					SubjectID: 2,
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					StudentID: "studentid2",
					SubjectID: 1,
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
			expect: map[string]StudentSubjects{
				"studentid1": {
					{
						StudentID: "studentid1",
						SubjectID: 1,
						CreatedAt: now,
						UpdatedAt: now,
					},
					{
						StudentID: "studentid1",
						SubjectID: 2,
						CreatedAt: now,
						UpdatedAt: now,
					},
				},
				"studentid2": {
					{
						StudentID: "studentid2",
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
			assert.Equal(t, tt.expect, tt.subjects.GroupByStudentID())
		})
	}
}

func TestStudentSubject_StudentProto(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name     string
		subjects StudentSubjects
		expect   *classroom.StudentSubject
	}{
		{
			name: "success",
			subjects: StudentSubjects{
				{
					StudentID: "studentid1",
					SubjectID: 1,
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					StudentID: "studentid1",
					SubjectID: 2,
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
			expect: &classroom.StudentSubject{
				StudentId:  "studentid1",
				SubjectIds: []int64{1, 2},
			},
		},
		{
			name:     "success empty",
			subjects: StudentSubjects{},
			expect:   nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.subjects.StudentProto())
		})
	}
}

func TestStudentSubject_StudentsProto(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name     string
		subjects StudentSubjects
		expect   []*classroom.StudentSubject
	}{
		{
			name: "success",
			subjects: StudentSubjects{
				{
					StudentID: "studentid1",
					SubjectID: 1,
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					StudentID: "studentid1",
					SubjectID: 2,
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					StudentID: "studentid2",
					SubjectID: 1,
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
			expect: []*classroom.StudentSubject{
				{
					StudentId:  "studentid1",
					SubjectIds: []int64{1, 2},
				},
				{
					StudentId:  "studentid2",
					SubjectIds: []int64{1},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := tt.subjects.StudentsProto()
			for i := range tt.expect {
				assert.Contains(t, actual, tt.expect[i])
			}
		})
	}
}
