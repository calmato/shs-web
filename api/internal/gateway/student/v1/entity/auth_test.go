package entity

import (
	"testing"

	"github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name    string
		student *entity.Student
		expect  *Auth
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
			expect: &Auth{
				ID:            "kSByoE6FetnPs5Byk3a9Zx",
				LastName:      "中村",
				FirstName:     "広大",
				LastNameKana:  "なかむら",
				FirstNameKana: "こうだい",
				Mail:          "student-test001@calmato.jp",
				SchoolType:    SchoolTypeHighSchool,
				Grade:         3,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewAuth(tt.student))
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
