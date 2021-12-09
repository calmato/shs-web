package validation

import (
	"fmt"
	"strings"
	"testing"

	"github.com/calmato/shs-web/api/proto/user"
	"github.com/stretchr/testify/assert"
)

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
				Role:                 0,
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
				Role:                 0,
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
				Role:                 0,
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
				Role:                 0,
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
				Role:                 0,
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
				Role:                 0,
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
				Role:                 0,
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
				Role:                 0,
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
				Role:                 0,
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
				Role:                 0,
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
				Role:                 0,
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
				Role:                 0,
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
				Role:                 0,
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
				Role:                 0,
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
				Role:                 -1,
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
				Role:                 0,
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
				Role:                 0,
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
				Role:                 0,
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
				Role:                 0,
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
				Role:                 0,
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
