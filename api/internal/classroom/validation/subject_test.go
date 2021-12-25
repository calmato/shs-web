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
			name: "success",
			req: &classroom.ListSubjectsRequest{
				SchoolType: classroom.SchoolType_SCHOOL_TYPE_UNKNOWN,
			},
			isErr: false,
		},
		{
			name: "SchoolType is defined_only",
			req: &classroom.ListSubjectsRequest{
				SchoolType: classroom.SchoolType(-1),
			},
			isErr: true,
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

func TestCreateSubject(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *classroom.CreateSubjectRequest
		isErr bool
	}{
		{
			name: "success",
			req: &classroom.CreateSubjectRequest{
				Name:       "国語",
				Color:      "#F8BBD0",
				SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
			},
			isErr: false,
		},
		{
			name: "Name is min_len",
			req: &classroom.CreateSubjectRequest{
				Name:       "",
				Color:      "#F8BBD0",
				SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
			},
			isErr: true,
		},
		{
			name: "Color is len",
			req: &classroom.CreateSubjectRequest{
				Name:       "国語",
				Color:      "#00000",
				SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
			},
			isErr: true,
		},
		{
			name: "Color is pattern",
			req: &classroom.CreateSubjectRequest{
				Name:       "国語",
				Color:      "#00000G",
				SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
			},
			isErr: true,
		},
		{
			name: "SchoolType is defined_only",
			req: &classroom.CreateSubjectRequest{
				Name:       "国語",
				Color:      "#F8BBD0",
				SchoolType: classroom.SchoolType(-1),
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.CreateSubject(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestUpdateSubject(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *classroom.UpdateSubjectRequest
		isErr bool
	}{
		{
			name: "success",
			req: &classroom.UpdateSubjectRequest{
				Id:         1,
				Name:       "国語",
				Color:      "#F8BBD0",
				SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
			},
			isErr: false,
		},
		{
			name: "Id is gt",
			req: &classroom.UpdateSubjectRequest{
				Id:         0,
				Name:       "国語",
				Color:      "#F8BBD0",
				SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
			},
			isErr: true,
		},
		{
			name: "Name is min_len",
			req: &classroom.UpdateSubjectRequest{
				Id:         1,
				Name:       "",
				Color:      "#F8BBD0",
				SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
			},
			isErr: true,
		},
		{
			name: "Color is len",
			req: &classroom.UpdateSubjectRequest{
				Id:         1,
				Name:       "国語",
				Color:      "#00000",
				SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
			},
			isErr: true,
		},
		{
			name: "Color is pattern",
			req: &classroom.UpdateSubjectRequest{
				Id:         1,
				Name:       "国語",
				Color:      "#00000G",
				SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
			},
			isErr: true,
		},
		{
			name: "SchoolType is defined_only",
			req: &classroom.UpdateSubjectRequest{
				Id:         1,
				Name:       "国語",
				Color:      "#F8BBD0",
				SchoolType: classroom.SchoolType(-1),
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.UpdateSubject(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestDeleteSubject(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *classroom.DeleteSubjectRequest
		isErr bool
	}{
		{
			name: "success",
			req: &classroom.DeleteSubjectRequest{
				Id: 1,
			},
			isErr: false,
		},
		{
			name: "Id is gt",
			req: &classroom.DeleteSubjectRequest{
				Id: 0,
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.DeleteSubject(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}
