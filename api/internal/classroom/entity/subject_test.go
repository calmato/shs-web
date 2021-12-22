package entity

import (
	"testing"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/stretchr/testify/assert"
)

func TestSubject_Proto(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name    string
		subject *Subject
		expect  *classroom.Subject
	}{
		{
			name: "success",
			subject: &Subject{
				ID:         1,
				Name:       "国語",
				Color:      "#F8BBD0",
				SchoolType: 1,
				CreatedAt:  now,
				UpdatedAt:  now,
			},
			expect: &classroom.Subject{
				Id:         1,
				Name:       "国語",
				Color:      "#F8BBD0",
				SchoolType: classroom.SchoolType_SCHOOL_TYPE_ELEMENTARY_SCHOOL,
				CreatedAt:  now.Unix(),
				UpdatedAt:  now.Unix(),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.subject.Proto())
		})
	}
}

func TestSubjects_Proto(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name     string
		subjects Subjects
		expect   []*classroom.Subject
	}{
		{
			name: "success",
			subjects: Subjects{
				{
					ID:         1,
					Name:       "国語",
					Color:      "#F8BBD0",
					SchoolType: 1,
					CreatedAt:  now,
					UpdatedAt:  now,
				},
			},
			expect: []*classroom.Subject{
				{
					Id:         1,
					Name:       "国語",
					Color:      "#F8BBD0",
					SchoolType: classroom.SchoolType_SCHOOL_TYPE_ELEMENTARY_SCHOOL,
					CreatedAt:  now.Unix(),
					UpdatedAt:  now.Unix(),
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.subjects.Proto())
		})
	}
}

func TestSchoolType(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		schoolType classroom.SchoolType
		expect     SchoolType
		isErr      bool
	}{
		{
			name:       "elementary school",
			schoolType: classroom.SchoolType_SCHOOL_TYPE_ELEMENTARY_SCHOOL,
			expect:     SchoolTypeElementarySchool,
			isErr:      false,
		},
		{
			name:       "junior high school",
			schoolType: classroom.SchoolType_SCHOOL_TYPE_JUNIOR_HIGH_SCHOOL,
			expect:     SchoolTypeJuniorHighSchool,
			isErr:      false,
		},
		{
			name:       "high school",
			schoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
			expect:     SchoolTypeHighSchool,
			isErr:      false,
		},
		{
			name:       "invalid school type",
			schoolType: classroom.SchoolType_SCHOOL_TYPE_UNKNOWN,
			expect:     SchoolTypeUnknown,
			isErr:      true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual, err := NewSchoolType(tt.schoolType)
			assert.Equal(t, tt.isErr, err != nil, err)
			assert.Equal(t, tt.expect, actual)
		})
	}
}
