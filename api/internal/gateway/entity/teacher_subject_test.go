package entity

import (
	"testing"

	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/stretchr/testify/assert"
)

func TestTeacherSubject(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		subject *classroom.TeacherSubject
		expect  *TeacherSubject
	}{
		{
			name: "success",
			subject: &classroom.TeacherSubject{
				TeacherId:  "teacherid",
				SubjectIds: []int64{1, 2},
			},
			expect: &TeacherSubject{
				TeacherSubject: &classroom.TeacherSubject{
					TeacherId:  "teacherid",
					SubjectIds: []int64{1, 2},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewTeacherSubject(tt.subject))
		})
	}
}

func TestTeacherSubjects(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		subjects []*classroom.TeacherSubject
		expect   TeacherSubjects
	}{
		{
			name: "success",
			subjects: []*classroom.TeacherSubject{
				{
					TeacherId:  "teacherid",
					SubjectIds: []int64{1, 2},
				},
			},
			expect: TeacherSubjects{
				{
					TeacherSubject: &classroom.TeacherSubject{
						TeacherId:  "teacherid",
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
			assert.Equal(t, tt.expect, NewTeacherSubjects(tt.subjects))
		})
	}
}
