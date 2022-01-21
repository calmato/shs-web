package entity

import (
	"testing"

	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/stretchr/testify/assert"
)

func TestStudentSubject(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		subject *classroom.StudentSubject
		expect  *StudentSubject
	}{
		{
			name: "success",
			subject: &classroom.StudentSubject{
				StudentId:  "studentid",
				SubjectIds: []int64{1, 2},
			},
			expect: &StudentSubject{
				StudentSubject: &classroom.StudentSubject{
					StudentId:  "studentid",
					SubjectIds: []int64{1, 2},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewStudentSubject(tt.subject))
		})
	}
}

func TestStudentSubjects(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		subjects []*classroom.StudentSubject
		expect   StudentSubjects
	}{
		{
			name: "success",
			subjects: []*classroom.StudentSubject{
				{
					StudentId:  "studentid",
					SubjectIds: []int64{1, 2},
				},
			},
			expect: StudentSubjects{
				{
					StudentSubject: &classroom.StudentSubject{
						StudentId:  "studentid",
						SubjectIds: []int64{1, 2},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewStudentSubjects(tt.subjects))
		})
	}
}
