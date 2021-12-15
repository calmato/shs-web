package entity

import (
	"testing"

	"github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/stretchr/testify/assert"
)

func TestTeacher(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 8, 2, 18, 30, 0, 0)
	tests := []struct {
		name    string
		teacher *entity.Teacher
		expect  *Teacher
	}{
		{
			name: "success",
			teacher: &entity.Teacher{
				Teacher: &user.Teacher{
					Id:            "kSByoE6FetnPs5Byk3a9Zx",
					LastName:      "中村",
					FirstName:     "広大",
					LastNameKana:  "なかむら",
					FirstNameKana: "こうだい",
					Mail:          "teacher-test001@calmato.jp",
					Role:          user.Role_ROLE_TEACHER,
					CreatedAt:     now.Unix(),
					UpdatedAt:     now.Unix(),
				},
			},
			expect: &Teacher{
				ID:            "kSByoE6FetnPs5Byk3a9Zx",
				LastName:      "中村",
				FirstName:     "広大",
				LastNameKana:  "なかむら",
				FirstNameKana: "こうだい",
				Mail:          "teacher-test001@calmato.jp",
				Role:          RoleTeacher,
				CreatedAt:     now,
				UpdatedAt:     now,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewTeacher(tt.teacher))
		})
	}
}

func TestTeachers(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 8, 2, 18, 30, 0, 0)
	tests := []struct {
		name     string
		teachers entity.Teachers
		expect   Teachers
	}{
		{
			name: "success",
			teachers: entity.Teachers{
				{
					Teacher: &user.Teacher{
						Id:            "kSByoE6FetnPs5Byk3a9Zx",
						LastName:      "中村",
						FirstName:     "広大",
						LastNameKana:  "なかむら",
						FirstNameKana: "こうだい",
						Mail:          "teacher-test001@calmato.jp",
						Role:          user.Role_ROLE_TEACHER,
						CreatedAt:     now.Unix(),
						UpdatedAt:     now.Unix(),
					},
				},
			},
			expect: Teachers{
				{
					ID:            "kSByoE6FetnPs5Byk3a9Zx",
					LastName:      "中村",
					FirstName:     "広大",
					LastNameKana:  "なかむら",
					FirstNameKana: "こうだい",
					Mail:          "teacher-test001@calmato.jp",
					Role:          RoleTeacher,
					CreatedAt:     now,
					UpdatedAt:     now,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewTeachers(tt.teachers))
		})
	}
}

func TestRole(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		role   user.Role
		expect Role
	}{
		{
			name:   "success teacher role",
			role:   user.Role_ROLE_TEACHER,
			expect: RoleTeacher,
		},
		{
			name:   "success administrator role",
			role:   user.Role_ROLE_ADMINISTRATOR,
			expect: RoleAdministrator,
		},
		{
			name:   "success invalid role",
			role:   -1,
			expect: RoleUnknown,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewRole(tt.role))
		})
	}
}

func TestRole_UserRole(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		role   Role
		expect user.Role
		isErr  bool
	}{
		{
			name:   "success administrator role",
			role:   RoleAdministrator,
			expect: user.Role_ROLE_ADMINISTRATOR,
			isErr:  false,
		},
		{
			name:   "success teacher role",
			role:   RoleTeacher,
			expect: user.Role_ROLE_TEACHER,
			isErr:  false,
		},
		{
			name:   "success invalid role",
			role:   RoleUnknown,
			expect: user.Role_ROLE_TEACHER,
			isErr:  true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual, err := tt.role.UserRole()
			assert.Equal(t, tt.isErr, err != nil, err)
			assert.Equal(t, tt.expect, actual)
		})
	}
}
