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
					Color:      "#F8BBD0",
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
						Color:      "#F8BBD0",
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

func TestSchoolTypeFromClassroom(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		schoolType classroom.SchoolType
		expect     SchoolType
	}{
		{
			name:       "successl elementary school",
			schoolType: classroom.SchoolType_SCHOOL_TYPE_ELEMENTARY_SCHOOL,
			expect:     SchoolTypeElementarySchool,
		},
		{
			name:       "successl junior high school",
			schoolType: classroom.SchoolType_SCHOOL_TYPE_JUNIOR_HIGH_SCHOOL,
			expect:     SchoolTypeJuniorHighSchool,
		},
		{
			name:       "successl high school",
			schoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
			expect:     SchoolTypeHighSchool,
		},
		{
			name:       "successl invalid school type",
			schoolType: classroom.SchoolType_SCHOOL_TYPE_UNKNOWN,
			expect:     SchoolTypeUnknown,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewSchoolTypeFromClassroom(tt.schoolType))
		})
	}
}

func TestSubjects_SortByColor(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		subjects Subjects
		expect   Subjects
	}{
		{
			name: "success",
			subjects: Subjects{
				{
					ID:         1,
					Name:       "現代文",
					Color:      "#F8BBD0",
					SchoolType: SchoolTypeHighSchool,
				},
				{
					ID:         3,
					Name:       "数学Ⅰ",
					Color:      "#BBDEFB",
					SchoolType: SchoolTypeHighSchool,
				},
				{
					ID:         2,
					Name:       "古典",
					Color:      "#F8BBD0",
					SchoolType: SchoolTypeHighSchool,
				},
				{
					ID:         4,
					Name:       "数学Ⅱ",
					Color:      "#BBDEFB",
					SchoolType: SchoolTypeHighSchool,
				},
			},
			expect: Subjects{
				{
					ID:         3,
					Name:       "数学Ⅰ",
					Color:      "#BBDEFB",
					SchoolType: SchoolTypeHighSchool,
				},
				{
					ID:         4,
					Name:       "数学Ⅱ",
					Color:      "#BBDEFB",
					SchoolType: SchoolTypeHighSchool,
				},
				{
					ID:         1,
					Name:       "現代文",
					Color:      "#F8BBD0",
					SchoolType: SchoolTypeHighSchool,
				},
				{
					ID:         2,
					Name:       "古典",
					Color:      "#F8BBD0",
					SchoolType: SchoolTypeHighSchool,
				},
			},
		},
		{
			name:     "success when subjects is length 0",
			subjects: Subjects{},
			expect:   Subjects{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.subjects.SortByColor()
			assert.Equal(t, tt.expect, tt.subjects)
		})
	}
}

func TestSubjects_SortBySchoolType(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		subjects Subjects
		expect   Subjects
	}{
		{
			name: "success",
			subjects: Subjects{
				{
					ID:         3,
					Name:       "国語",
					Color:      "#F8BBD0",
					SchoolType: SchoolTypeHighSchool,
				},
				{
					ID:         1,
					Name:       "国語",
					Color:      "#F8BBD0",
					SchoolType: SchoolTypeElementarySchool,
				},
				{
					ID:         2,
					Name:       "国語",
					Color:      "#F8BBD0",
					SchoolType: SchoolTypeJuniorHighSchool,
				},
			},
			expect: Subjects{
				{
					ID:         1,
					Name:       "国語",
					Color:      "#F8BBD0",
					SchoolType: SchoolTypeElementarySchool,
				},
				{
					ID:         2,
					Name:       "国語",
					Color:      "#F8BBD0",
					SchoolType: SchoolTypeJuniorHighSchool,
				},
				{
					ID:         3,
					Name:       "国語",
					Color:      "#F8BBD0",
					SchoolType: SchoolTypeHighSchool,
				},
			},
		},
		{
			name:     "success when subjects is length 0",
			subjects: Subjects{},
			expect:   Subjects{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.subjects.SortBySchoolType()
			assert.Equal(t, tt.expect, tt.subjects)
		})
	}
}

func TestSchoolType_ClassroomSchoolType(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		schoolType SchoolType
		expect     classroom.SchoolType
		isErr      bool
	}{
		{
			name:       "successl elementary school",
			schoolType: SchoolTypeElementarySchool,
			expect:     classroom.SchoolType_SCHOOL_TYPE_ELEMENTARY_SCHOOL,
			isErr:      false,
		},
		{
			name:       "successl junior high school",
			schoolType: SchoolTypeJuniorHighSchool,
			expect:     classroom.SchoolType_SCHOOL_TYPE_JUNIOR_HIGH_SCHOOL,
			isErr:      false,
		},
		{
			name:       "successl high school",
			schoolType: SchoolTypeHighSchool,
			expect:     classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
			isErr:      false,
		},
		{
			name:       "failed to invalid school type",
			schoolType: SchoolTypeUnknown,
			expect:     classroom.SchoolType_SCHOOL_TYPE_UNKNOWN,
			isErr:      true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual, err := tt.schoolType.ClassroomSchoolType()
			assert.Equal(t, tt.isErr, err != nil, err)
			assert.Equal(t, tt.expect, actual)
		})
	}
}
