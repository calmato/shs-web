package validation

import (
	"fmt"
	"strings"
	"testing"

	"github.com/calmato/shs-web/api/proto/user"
	"github.com/stretchr/testify/assert"
)

func TestListTeachers(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *user.ListTeachersRequest
		isErr bool
	}{
		{
			name: "success",
			req: &user.ListTeachersRequest{
				Limit:  30,
				Offset: 0,
			},
			isErr: false,
		},
		{
			name: "Lismit is gte",
			req: &user.ListTeachersRequest{
				Limit:  -1,
				Offset: 0,
			},
			isErr: true,
		},
		{
			name: "Offset is gte",
			req: &user.ListTeachersRequest{
				Limit:  30,
				Offset: -1,
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.ListTeachers(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestGetTeacher(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *user.GetTeacherRequest
		isErr bool
	}{
		{
			name: "success",
			req: &user.GetTeacherRequest{
				Id: "cvcTyJFfgoDQrqC1KDHbRe",
			},
			isErr: false,
		},
		{
			name: "Id is min_len",
			req: &user.GetTeacherRequest{
				Id: "",
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.GetTeacher(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestCreateTeacher(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *user.CreateTeacherRequest
		isErr bool
	}{
		{
			name: "success",
			req: &user.CreateTeacherRequest{
				LastName:             "中村",
				FirstName:            "広大",
				LastNameKana:         "なかむら",
				FirstNameKana:        "こうだい",
				Mail:                 "teacher-test001@calmato.jp",
				Role:                 user.Role_ROLE_TEACHER,
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
			isErr: false,
		},
		{
			name: "LastName is min_len",
			req: &user.CreateTeacherRequest{
				LastName:             "",
				FirstName:            "広大",
				LastNameKana:         "なかむら",
				FirstNameKana:        "こうだい",
				Mail:                 "teacher-test001@calmato.jp",
				Role:                 user.Role_ROLE_TEACHER,
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
			isErr: true,
		},
		{
			name: "LastName is max_len",
			req: &user.CreateTeacherRequest{
				LastName:             strings.Repeat("x", 17),
				FirstName:            "広大",
				LastNameKana:         "なかむら",
				FirstNameKana:        "こうだい",
				Mail:                 "teacher-test001@calmato.jp",
				Role:                 user.Role_ROLE_TEACHER,
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
			isErr: true,
		},
		{
			name: "FirstName is min_len",
			req: &user.CreateTeacherRequest{
				LastName:             "中村",
				FirstName:            "",
				LastNameKana:         "なかむら",
				FirstNameKana:        "こうだい",
				Mail:                 "teacher-test001@calmato.jp",
				Role:                 user.Role_ROLE_TEACHER,
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
			isErr: true,
		},
		{
			name: "FirstName is max_len",
			req: &user.CreateTeacherRequest{
				LastName:             "中村",
				FirstName:            strings.Repeat("x", 17),
				LastNameKana:         "なかむら",
				FirstNameKana:        "こうだい",
				Mail:                 "teacher-test001@calmato.jp",
				Role:                 user.Role_ROLE_TEACHER,
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
			isErr: true,
		},
		{
			name: "LastNameKana is min_len",
			req: &user.CreateTeacherRequest{
				LastName:             "中村",
				FirstName:            "広大",
				LastNameKana:         "",
				FirstNameKana:        "こうだい",
				Mail:                 "teacher-test001@calmato.jp",
				Role:                 user.Role_ROLE_TEACHER,
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
			isErr: true,
		},
		{
			name: "LastNameKana is max_len",
			req: &user.CreateTeacherRequest{
				LastName:             "中村",
				FirstName:            "広大",
				LastNameKana:         strings.Repeat("x", 33),
				FirstNameKana:        "こうだい",
				Mail:                 "teacher-test001@calmato.jp",
				Role:                 user.Role_ROLE_TEACHER,
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
			isErr: true,
		},
		{
			name: "LastNameKana is pattern",
			req: &user.CreateTeacherRequest{
				LastName:             "中村",
				FirstName:            "広大",
				LastNameKana:         "ナカムラ",
				FirstNameKana:        "こうだい",
				Mail:                 "teacher-test001@calmato.jp",
				Role:                 user.Role_ROLE_TEACHER,
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
			isErr: true,
		},
		{
			name: "FirstNameKana is min_len",
			req: &user.CreateTeacherRequest{
				LastName:             "中村",
				FirstName:            "広大",
				LastNameKana:         "なかむら",
				FirstNameKana:        "",
				Mail:                 "teacher-test001@calmato.jp",
				Role:                 user.Role_ROLE_TEACHER,
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
			isErr: true,
		},
		{
			name: "FirstNameKana is max_len",
			req: &user.CreateTeacherRequest{
				LastName:             "中村",
				FirstName:            "広大",
				LastNameKana:         "なかむら",
				FirstNameKana:        strings.Repeat("x", 33),
				Mail:                 "teacher-test001@calmato.jp",
				Role:                 user.Role_ROLE_TEACHER,
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
			isErr: true,
		},
		{
			name: "FirstNameKana is pattern",
			req: &user.CreateTeacherRequest{
				LastName:             "中村",
				FirstName:            "広大",
				LastNameKana:         "なかむら",
				FirstNameKana:        "コウダイ",
				Mail:                 "teacher-test001@calmato.jp",
				Role:                 user.Role_ROLE_TEACHER,
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
			isErr: true,
		},
		{
			name: "Mail is min_len",
			req: &user.CreateTeacherRequest{
				LastName:             "中村",
				FirstName:            "広大",
				LastNameKana:         "なかむら",
				FirstNameKana:        "こうだい",
				Mail:                 "",
				Role:                 user.Role_ROLE_TEACHER,
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
			isErr: true,
		},
		{
			name: "Mail is max_len",
			req: &user.CreateTeacherRequest{
				LastName:             "中村",
				FirstName:            "広大",
				LastNameKana:         "なかむら",
				FirstNameKana:        "こうだい",
				Mail:                 fmt.Sprintf("%s@calmato.jp", strings.Repeat("x", 245)),
				Role:                 user.Role_ROLE_TEACHER,
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
			isErr: true,
		},
		{
			name: "Mail is email",
			req: &user.CreateTeacherRequest{
				LastName:             "中村",
				FirstName:            "広大",
				LastNameKana:         "なかむら",
				FirstNameKana:        "こうだい",
				Mail:                 "teacher-test01",
				Role:                 user.Role_ROLE_TEACHER,
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
			isErr: true,
		},
		{
			name: "Role is defined_only",
			req: &user.CreateTeacherRequest{
				LastName:             "中村",
				FirstName:            "広大",
				LastNameKana:         "なかむら",
				FirstNameKana:        "こうだい",
				Mail:                 "teacher-test01@calmato.jp",
				Role:                 user.Role(-1),
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
			isErr: true,
		},
		{
			name: "Password is min_len",
			req: &user.CreateTeacherRequest{
				LastName:             "中村",
				FirstName:            "広大",
				LastNameKana:         "なかむら",
				FirstNameKana:        "こうだい",
				Mail:                 "teacher-test01@calmato.jp",
				Role:                 user.Role_ROLE_TEACHER,
				Password:             strings.Repeat("x", 5),
				PasswordConfirmation: "12345678",
			},
			isErr: true,
		},
		{
			name: "Password is max_len",
			req: &user.CreateTeacherRequest{
				LastName:             "中村",
				FirstName:            "広大",
				LastNameKana:         "なかむら",
				FirstNameKana:        "こうだい",
				Mail:                 "teacher-test01@calmato.jp",
				Role:                 user.Role_ROLE_TEACHER,
				Password:             strings.Repeat("x", 33),
				PasswordConfirmation: "12345678",
			},
			isErr: true,
		},
		{
			name: "Password is pattern",
			req: &user.CreateTeacherRequest{
				LastName:             "中村",
				FirstName:            "広大",
				LastNameKana:         "なかむら",
				FirstNameKana:        "こうだい",
				Mail:                 "teacher-test01@calmato.jp",
				Role:                 user.Role_ROLE_TEACHER,
				Password:             "あいうえおあ",
				PasswordConfirmation: "12345678",
			},
			isErr: true,
		},
		{
			name: "PasswordConfirmation is min_len",
			req: &user.CreateTeacherRequest{
				LastName:             "中村",
				FirstName:            "広大",
				LastNameKana:         "なかむら",
				FirstNameKana:        "こうだい",
				Mail:                 "teacher-test01@calmato.jp",
				Role:                 user.Role_ROLE_TEACHER,
				Password:             "12345678",
				PasswordConfirmation: "",
			},
			isErr: true,
		},
		{
			name: "Custom validation is password match",
			req: &user.CreateTeacherRequest{
				LastName:             "中村",
				FirstName:            "広大",
				LastNameKana:         "なかむら",
				FirstNameKana:        "こうだい",
				Mail:                 "teacher-test01@calmato.jp",
				Role:                 user.Role_ROLE_TEACHER,
				Password:             "12345678",
				PasswordConfirmation: "87654321",
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.CreateTeacher(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestUpdateTeacherMail(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *user.UpdateTeacherMailRequest
		isErr bool
	}{
		{
			name: "success",
			req: &user.UpdateTeacherMailRequest{
				Id:   "cvcTyJFfgoDQrqC1KDHbRe",
				Mail: "teacher-test01@calmato.jp",
			},
			isErr: false,
		},
		{
			name: "Id is min_len",
			req: &user.UpdateTeacherMailRequest{
				Id:   "",
				Mail: "teacher-test01@calmato.jp",
			},
			isErr: true,
		},
		{
			name: "Mail is min_len",
			req: &user.UpdateTeacherMailRequest{
				Id:   "cvcTyJFfgoDQrqC1KDHbRe",
				Mail: "",
			},
			isErr: true,
		},
		{
			name: "Mail is max_len",
			req: &user.UpdateTeacherMailRequest{
				Id:   "cvcTyJFfgoDQrqC1KDHbRe",
				Mail: fmt.Sprintf("%s@calmato.jp", strings.Repeat("x", 245)),
			},
			isErr: true,
		},
		{
			name: "Mail is email",
			req: &user.UpdateTeacherMailRequest{
				Id:   "cvcTyJFfgoDQrqC1KDHbRe",
				Mail: "teacher-test01",
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.UpdateTeacherMail(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestUpdateTeacherPassword(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *user.UpdateTeacherPasswordRequest
		isErr bool
	}{
		{
			name: "success",
			req: &user.UpdateTeacherPasswordRequest{
				Id:                   "cvcTyJFfgoDQrqC1KDHbRe",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
			isErr: false,
		},
		{
			name: "Id is min_len",
			req: &user.UpdateTeacherPasswordRequest{
				Id:                   "",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
			isErr: true,
		},
		{
			name: "Password is min_len",
			req: &user.UpdateTeacherPasswordRequest{
				Id:                   "cvcTyJFfgoDQrqC1KDHbRe",
				Password:             strings.Repeat("x", 5),
				PasswordConfirmation: "12345678",
			},
			isErr: true,
		},
		{
			name: "Password is max_len",
			req: &user.UpdateTeacherPasswordRequest{
				Id:                   "cvcTyJFfgoDQrqC1KDHbRe",
				Password:             strings.Repeat("x", 33),
				PasswordConfirmation: "12345678",
			},
			isErr: true,
		},
		{
			name: "Password is pattern",
			req: &user.UpdateTeacherPasswordRequest{
				Id:                   "cvcTyJFfgoDQrqC1KDHbRe",
				Password:             "あいうえおあ",
				PasswordConfirmation: "12345678",
			},
			isErr: true,
		},
		{
			name: "Custom validation is password match",
			req: &user.UpdateTeacherPasswordRequest{
				Id:                   "cvcTyJFfgoDQrqC1KDHbRe",
				Password:             "12345678",
				PasswordConfirmation: "87654321",
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.UpdateTeacherPassword(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestUpdateTeacherRole(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *user.UpdateTeacherRoleRequest
		isErr bool
	}{
		{
			name: "success",
			req: &user.UpdateTeacherRoleRequest{
				Id:   "cvcTyJFfgoDQrqC1KDHbRe",
				Role: user.Role_ROLE_TEACHER,
			},
			isErr: false,
		},
		{
			name: "Id is min_len",
			req: &user.UpdateTeacherRoleRequest{
				Id:   "",
				Role: user.Role_ROLE_TEACHER,
			},
			isErr: true,
		},
		{
			name: "Role is defined_only",
			req: &user.UpdateTeacherRoleRequest{
				Id:   "cvcTyJFfgoDQrqC1KDHbRe",
				Role: user.Role(-1),
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.UpdateTeacherRole(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestDeleteTeacher(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *user.DeleteTeacherRequest
		isErr bool
	}{
		{
			name: "success",
			req: &user.DeleteTeacherRequest{
				Id: "cvcTyJFfgoDQrqC1KDHbRe",
			},
			isErr: false,
		},
		{
			name: "Id is min_len",
			req: &user.DeleteTeacherRequest{
				Id: "",
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.DeleteTeacher(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}
