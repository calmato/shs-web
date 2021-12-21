package validation

import (
	"testing"

	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/stretchr/testify/assert"
)

func TestMultiGetTeacherSubjects(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *classroom.MultiGetTeacherSubjectsRequest
		isErr bool
	}{
		{
			name: "success",
			req: &classroom.MultiGetTeacherSubjectsRequest{
				TeacherIds: []string{"teacherid"},
			},
			isErr: false,
		},
		{
			name: "TeacherIds is min_items",
			req: &classroom.MultiGetTeacherSubjectsRequest{
				TeacherIds: []string{},
			},
			isErr: true,
		},
		{
			name: "TeacherIds is unique",
			req: &classroom.MultiGetTeacherSubjectsRequest{
				TeacherIds: []string{"teacherid", "teacherid"},
			},
			isErr: true,
		},
		{
			name: "TeacherIds is items.min_len",
			req: &classroom.MultiGetTeacherSubjectsRequest{
				TeacherIds: []string{""},
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.MultiGetTeacherSubjects(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestGetTeacherSubject(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *classroom.GetTeacherSubjectRequest
		isErr bool
	}{
		{
			name: "success",
			req: &classroom.GetTeacherSubjectRequest{
				TeacherId: "teacherid",
			},
			isErr: false,
		},
		{
			name: "TeacherIds is items.min_len",
			req: &classroom.GetTeacherSubjectRequest{
				TeacherId: "",
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.GetTeacherSubject(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}
