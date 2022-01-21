package entity

import (
	"testing"

	"github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/stretchr/testify/assert"
)

func TestStudent(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 8, 2, 18, 30, 0, 0)
	tests := []struct {
		name     string
		student  *entity.Student
		subjects entity.Subjects
		expect   *Student
	}{
		{
			name: "success",
			student: &entity.Student{
				Student: &user.Student{
					Id:            "kSByoE6FetnPs5Byk3a9Zx",
					LastName:      "中村",
					FirstName:     "広大",
					LastNameKana:  "なかむら",
					FirstNameKana: "こうだい",
					Mail:          "student-test001@calmato.jp",
					SchoolType:    user.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
					Grade:         3,
					CreatedAt:     now.Unix(),
					UpdatedAt:     now.Unix(),
				},
			},
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
			expect: &Student{
				ID:            "kSByoE6FetnPs5Byk3a9Zx",
				LastName:      "中村",
				FirstName:     "広大",
				LastNameKana:  "なかむら",
				FirstNameKana: "こうだい",
				Mail:          "student-test001@calmato.jp",
				SchoolType:    SchoolTypeHighSchool,
				Grade:         3,
				Subjects: Subjects{
					{
						ID:         1,
						Name:       "国語",
						Color:      "#F8BBD0",
						SchoolType: SchoolTypeHighSchool,
						CreatedAt:  now,
						UpdatedAt:  now,
					},
				},
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewStudent(tt.student, tt.subjects))
		})
	}
}

func TestStudents(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 8, 2, 18, 30, 0, 0)
	tests := []struct {
		name     string
		students entity.Students
		subjects map[string]entity.Subjects
		expect   Students
	}{
		{
			name: "success",
			students: entity.Students{
				{
					Student: &user.Student{
						Id:            "kSByoE6FetnPs5Byk3a9Zx",
						LastName:      "中村",
						FirstName:     "広大",
						LastNameKana:  "なかむら",
						FirstNameKana: "こうだい",
						Mail:          "student-test001@calmato.jp",
						SchoolType:    user.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
						Grade:         3,
						CreatedAt:     now.Unix(),
						UpdatedAt:     now.Unix(),
					},
				},
			},
			subjects: map[string]entity.Subjects{
				"kSByoE6FetnPs5Byk3a9Zx": {
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
			},
			expect: Students{
				{
					ID:            "kSByoE6FetnPs5Byk3a9Zx",
					LastName:      "中村",
					FirstName:     "広大",
					LastNameKana:  "なかむら",
					FirstNameKana: "こうだい",
					Mail:          "student-test001@calmato.jp",
					SchoolType:    SchoolTypeHighSchool,
					Grade:         3,
					Subjects: Subjects{
						{
							ID:         1,
							Name:       "国語",
							Color:      "#F8BBD0",
							SchoolType: SchoolTypeHighSchool,
							CreatedAt:  now,
							UpdatedAt:  now,
						},
					},
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewStudents(tt.students, tt.subjects))
		})
	}
}

func TestSchoolTypeFromUser(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		schoolType user.SchoolType
		expect     SchoolType
	}{
		{
			name:       "successl elementary school",
			schoolType: user.SchoolType_SCHOOL_TYPE_ELEMENTARY_SCHOOL,
			expect:     SchoolTypeElementarySchool,
		},
		{
			name:       "successl junior high school",
			schoolType: user.SchoolType_SCHOOL_TYPE_JUNIOR_HIGH_SCHOOL,
			expect:     SchoolTypeJuniorHighSchool,
		},
		{
			name:       "successl high school",
			schoolType: user.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
			expect:     SchoolTypeHighSchool,
		},
		{
			name:       "successl invalid school type",
			schoolType: user.SchoolType_SCHOOL_TYPE_UNKNOWN,
			expect:     SchoolTypeUnknown,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewSchoolTypeFromUser(tt.schoolType))
		})
	}
}
