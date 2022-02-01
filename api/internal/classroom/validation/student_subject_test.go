package validation

import (
	"testing"

	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/stretchr/testify/assert"
)

func TestMultiGetStudentSubjects(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *classroom.MultiGetStudentSubjectsRequest
		isErr bool
	}{
		{
			name: "success",
			req: &classroom.MultiGetStudentSubjectsRequest{
				StudentIds: []string{"studentid"},
			},
			isErr: false,
		},
		{
			name: "StudentIds is unique",
			req: &classroom.MultiGetStudentSubjectsRequest{
				StudentIds: []string{"studentid", "studentid"},
			},
			isErr: true,
		},
		{
			name: "StudentIds is items.min_len",
			req: &classroom.MultiGetStudentSubjectsRequest{
				StudentIds: []string{""},
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.MultiGetStudentSubjects(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestGetStudentSubject(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *classroom.GetStudentSubjectRequest
		isErr bool
	}{
		{
			name: "success",
			req: &classroom.GetStudentSubjectRequest{
				StudentId: "studentid",
			},
			isErr: false,
		},
		{
			name: "StudentId is items.min_len",
			req: &classroom.GetStudentSubjectRequest{
				StudentId: "",
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.GetStudentSubject(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestUpsertStudentSubject(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *classroom.UpsertStudentSubjectRequest
		isErr bool
	}{
		{
			name: "success",
			req: &classroom.UpsertStudentSubjectRequest{
				StudentId:  "studentid",
				SubjectIds: []int64{},
				SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
			},
			isErr: false,
		},
		{
			name: "StudentId is items.min_len",
			req: &classroom.UpsertStudentSubjectRequest{
				StudentId:  "",
				SubjectIds: []int64{1, 2},
				SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
			},
			isErr: true,
		},
		{
			name: "SubjectIds is unique",
			req: &classroom.UpsertStudentSubjectRequest{
				StudentId:  "",
				SubjectIds: []int64{1, 1},
				SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
			},
			isErr: true,
		},
		{
			name: "SubjectIds is items.gt",
			req: &classroom.UpsertStudentSubjectRequest{
				StudentId:  "",
				SubjectIds: []int64{0},
				SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
			},
			isErr: true,
		},
		{
			name: "SchoolType is defined_only",
			req: &classroom.UpsertStudentSubjectRequest{
				StudentId:  "",
				SubjectIds: []int64{1, 2},
				SchoolType: classroom.SchoolType_SCHOOL_TYPE_UNKNOWN,
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.UpsertStudentSubject(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}
