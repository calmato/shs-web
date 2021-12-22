package entity

import (
	"testing"

	"github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/stretchr/testify/assert"
)

func TestSubject(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 8, 2, 18, 30, 0, 0)
	tests := []struct {
		name    string
		subject *entity.Subject
		expect  *Subject
	}{
		{
			name: "success",
			subject: &entity.Subject{
				Subject: &classroom.Subject{
					Id:         1,
					Name:       "国語",
					Color:      "#f8bbd0",
					SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
					CreatedAt:  now.Unix(),
					UpdatedAt:  now.Unix(),
				},
			},
			expect: &Subject{
				ID:         1,
				Name:       "国語",
				Color:      "#F8BBD0",
				SchoolType: SchoolTypeHighSchool,
				CreatedAt:  now,
				UpdatedAt:  now,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewSubject(tt.subject))
		})
	}
}

func TestSubjects(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 8, 2, 18, 30, 0, 0)
	tests := []struct {
		name     string
		subjects entity.Subjects
		expect   Subjects
	}{
		{
			name: "success",
			subjects: entity.Subjects{
				{
					Subject: &classroom.Subject{
						Id:         1,
						Name:       "国語",
						Color:      "#f8bbd0",
						SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
						CreatedAt:  now.Unix(),
						UpdatedAt:  now.Unix(),
					},
				},
			},
			expect: []*Subject{
				{
					ID:         1,
					Name:       "国語",
					Color:      "#F8BBD0",
					SchoolType: SchoolTypeHighSchool,
					CreatedAt:  now,
					UpdatedAt:  now,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewSubjects(tt.subjects))
		})
	}
}

func TestSubjects_GroupBySchoolType(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 8, 2, 18, 30, 0, 0)
	tests := []struct {
		name     string
		subjects Subjects
		expect   map[SchoolType]Subjects
	}{
		{
			name: "success",
			subjects: Subjects{
				{
					ID:         1,
					Name:       "国語",
					Color:      "#F8BBD0",
					SchoolType: SchoolTypeHighSchool,
					CreatedAt:  now,
					UpdatedAt:  now,
				},
			},
			expect: map[SchoolType]Subjects{
				SchoolTypeElementarySchool: {},
				SchoolTypeJuniorHighSchool: {},
				SchoolTypeHighSchool: {
					{
						ID:         1,
						Name:       "国語",
						Color:      "#F8BBD0",
						SchoolType: SchoolTypeHighSchool,
						CreatedAt:  now,
						UpdatedAt:  now,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.subjects.GroupBySchoolType())
		})
	}
}
