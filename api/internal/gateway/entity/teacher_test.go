package entity

import (
	"testing"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/stretchr/testify/assert"
)

func TestTeacher(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name    string
		teacher *user.Teacher
		expect  *Teacher
	}{
		{
			name: "success",
			teacher: &user.Teacher{
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
			expect: &Teacher{
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
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewTeacher(tt.teacher))
		})
	}
}

func TestTeacher_AdminRole(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		teacher *Teacher
		expect  bool
	}{
		{
			name: "success admin role",
			teacher: &Teacher{
				Teacher: &user.Teacher{
					Role: user.Role_ROLE_ADMINISTRATOR,
				},
			},
			expect: true,
		},
		{
			name: "success non admin role",
			teacher: &Teacher{
				Teacher: &user.Teacher{
					Role: user.Role_ROLE_TEACHER,
				},
			},
			expect: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.teacher.AdminRole())
		})
	}
}

func TestTeachers(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name     string
		teachers []*user.Teacher
		expect   Teachers
	}{
		{
			name: "success",
			teachers: []*user.Teacher{
				{
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
			expect: Teachers{
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

func TestTeachers_IDs(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name     string
		teachers Teachers
		expect   []string
	}{
		{
			name: "success",
			teachers: Teachers{
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
				{
					Teacher: &user.Teacher{
						Id:            "teacherid",
						LastName:      "テスト",
						FirstName:     "講師",
						LastNameKana:  "てすと",
						FirstNameKana: "こうし",
						Mail:          "teacher-test002@calmato.jp",
						Role:          user.Role_ROLE_ADMINISTRATOR,
						CreatedAt:     now.Unix(),
						UpdatedAt:     now.Unix(),
					},
				},
			},
			expect: []string{
				"kSByoE6FetnPs5Byk3a9Zx",
				"teacherid",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.teachers.IDs())
		})
	}
}
