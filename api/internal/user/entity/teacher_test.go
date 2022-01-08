package entity

import (
	"testing"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/stretchr/testify/assert"
)

func TestTeacher_Proto(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name    string
		teacher *Teacher
		expect  *user.Teacher
	}{
		{
			name: "success",
			teacher: &Teacher{
				ID:            "kSByoE6FetnPs5Byk3a9Zx",
				LastName:      "中村",
				FirstName:     "広大",
				LastNameKana:  "なかむら",
				FirstNameKana: "こうだい",
				Mail:          "teacher-test001@calmato.jp",
				Password:      "12345678",
				Role:          0,
				CreatedAt:     now,
				UpdatedAt:     now,
			},
			expect: &user.Teacher{
				Id:            "kSByoE6FetnPs5Byk3a9Zx",
				LastName:      "中村",
				FirstName:     "広大",
				LastNameKana:  "なかむら",
				FirstNameKana: "こうだい",
				Mail:          "teacher-test001@calmato.jp",
				Role:          user.Role(0),
				CreatedAt:     now.Unix(),
				UpdatedAt:     now.Unix(),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.teacher.Proto())
		})
	}
}

func TestTeachers_Proto(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name     string
		teachers Teachers
		expect   []*user.Teacher
	}{
		{
			name: "success",
			teachers: Teachers{
				{
					ID:            "kSByoE6FetnPs5Byk3a9Zx",
					LastName:      "中村",
					FirstName:     "広大",
					LastNameKana:  "なかむら",
					FirstNameKana: "こうだい",
					Mail:          "teacher-test001@calmato.jp",
					Password:      "12345678",
					Role:          0,
					CreatedAt:     now,
					UpdatedAt:     now,
				},
			},
			expect: []*user.Teacher{
				{
					Id:            "kSByoE6FetnPs5Byk3a9Zx",
					LastName:      "中村",
					FirstName:     "広大",
					LastNameKana:  "なかむら",
					FirstNameKana: "こうだい",
					Mail:          "teacher-test001@calmato.jp",
					Role:          user.Role(0),
					CreatedAt:     now.Unix(),
					UpdatedAt:     now.Unix(),
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.teachers.Proto())
		})
	}
}

func TestRole(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		role   user.Role
		expect Role
		isErr  bool
	}{
		{
			name:   "success teacher role",
			role:   user.Role_ROLE_TEACHER,
			expect: RoleTeacher,
			isErr:  false,
		},
		{
			name:   "success administrator role",
			role:   user.Role_ROLE_ADMINISTRATOR,
			expect: RoleAdministrator,
			isErr:  false,
		},
		{
			name:   "success invalid role",
			role:   user.Role(-1),
			expect: RoleTeacher,
			isErr:  true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual, err := NewRole(tt.role)
			assert.Equal(t, tt.isErr, err != nil, err)
			assert.Equal(t, tt.expect, actual)
		})
	}
}
