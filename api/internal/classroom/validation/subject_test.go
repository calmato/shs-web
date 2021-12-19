package validation

import (
	"testing"

	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/stretchr/testify/assert"
)

func TestListSubjects(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *classroom.ListSubjectsRequest
		isErr bool
	}{
		{
			name:  "success",
			req:   &classroom.ListSubjectsRequest{},
			isErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.ListSubjects(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestMultiGetSubjects(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *classroom.MultiGetSubjectsRequest
		isErr bool
	}{
		{
			name: "success",
			req: &classroom.MultiGetSubjectsRequest{
				Ids: []int64{1, 2, 3},
			},
			isErr: false,
		},
		{
			name: "Ids is min_items",
			req: &classroom.MultiGetSubjectsRequest{
				Ids: []int64{},
			},
			isErr: true,
		},
		{
			name: "Ids is unique",
			req: &classroom.MultiGetSubjectsRequest{
				Ids: []int64{1, 1},
			},
			isErr: true,
		},
		{
			name: "Ids is items.gt",
			req: &classroom.MultiGetSubjectsRequest{
				Ids: []int64{0},
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.MultiGetSubjects(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestGetSubject(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *classroom.GetSubjectRequest
		isErr bool
	}{
		{
			name: "success",
			req: &classroom.GetSubjectRequest{
				Id: 1,
			},
			isErr: false,
		},
		{
			name: "Id is gte",
			req: &classroom.GetSubjectRequest{
				Id: 0,
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.GetSubject(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}
