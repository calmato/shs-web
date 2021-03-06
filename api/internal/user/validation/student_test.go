package validation

import (
	"fmt"
	"strings"
	"testing"

	"github.com/calmato/shs-web/api/proto/user"
	"github.com/stretchr/testify/assert"
)

func TestListStudents(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *user.ListStudentsRequest
		isErr bool
	}{
		{
			name: "success",
			req: &user.ListStudentsRequest{
				Limit:  30,
				Offset: 0,
			},
			isErr: false,
		},
		{
			name: "Lismit is gte",
			req: &user.ListStudentsRequest{
				Limit:  -1,
				Offset: 0,
			},
			isErr: true,
		},
		{
			name: "Offset is gte",
			req: &user.ListStudentsRequest{
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
			err := validator.ListStudents(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestMultiGetStudents(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *user.MultiGetStudentsRequest
		isErr bool
	}{
		{
			name: "success",
			req: &user.MultiGetStudentsRequest{
				Ids: []string{"cvcTyJFfgoDQrqC1KDHbRe"},
			},
			isErr: false,
		},
		{
			name: "Ids is unique",
			req: &user.MultiGetStudentsRequest{
				Ids: []string{"studentid", "studentid"},
			},
			isErr: true,
		},
		{
			name: "Ids is min_len",
			req: &user.MultiGetStudentsRequest{
				Ids: []string{""},
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.MultiGetStudents(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestGetStudent(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *user.GetStudentRequest
		isErr bool
	}{
		{
			name: "success",
			req: &user.GetStudentRequest{
				Id: "MvcNyJFfdoDTrqC1KDHbRw",
			},
			isErr: false,
		},
		{
			name: "Id is min_len",
			req: &user.GetStudentRequest{
				Id: "",
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.GetStudent(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestCreateStudent(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *user.CreateStudentRequest
		isErr bool
	}{
		{
			name: "success",
			req: &user.CreateStudentRequest{
				LastName:             "??????",
				FirstName:            "?????????",
				LastNameKana:         "?????????",
				FirstNameKana:        "?????????",
				Mail:                 "student-test001@calmato.jp",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
				SchoolType:           user.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
				Grade:                1,
			},
			isErr: false,
		},
		{
			name: "LastName is min_len",
			req: &user.CreateStudentRequest{
				LastName:             "",
				FirstName:            "?????????",
				LastNameKana:         "?????????",
				FirstNameKana:        "?????????",
				Mail:                 "student-test001@calmato.jp",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
				SchoolType:           user.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
				Grade:                1,
			},
			isErr: true,
		},
		{
			name: "LastName is max_len",
			req: &user.CreateStudentRequest{
				LastName:             strings.Repeat("x", 17),
				FirstName:            "??????",
				LastNameKana:         "?????????",
				FirstNameKana:        "?????????",
				Mail:                 "student-test001@calmato.jp",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
				SchoolType:           user.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
				Grade:                1,
			},
			isErr: true,
		},
		{
			name: "FirstName is min_len",
			req: &user.CreateStudentRequest{
				LastName:             "??????",
				FirstName:            "",
				LastNameKana:         "?????????",
				FirstNameKana:        "?????????",
				Mail:                 "student-test001@calmato.jp",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
				SchoolType:           user.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
				Grade:                1,
			},
			isErr: true,
		},
		{
			name: "FirstName is max_len",
			req: &user.CreateStudentRequest{
				LastName:             "??????",
				FirstName:            strings.Repeat("x", 17),
				LastNameKana:         "?????????",
				FirstNameKana:        "?????????",
				Mail:                 "student-test001@calmato.jp",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
				SchoolType:           user.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
				Grade:                1,
			},
			isErr: true,
		},
		{
			name: "LastNameKana is min_len",
			req: &user.CreateStudentRequest{
				LastName:             "??????",
				FirstName:            "??????",
				LastNameKana:         "",
				FirstNameKana:        "?????????",
				Mail:                 "student-test001@calmato.jp",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
				SchoolType:           user.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
				Grade:                1,
			},
			isErr: true,
		},
		{
			name: "LastNameKana is max_len",
			req: &user.CreateStudentRequest{
				LastName:             "??????",
				FirstName:            "??????",
				LastNameKana:         strings.Repeat("x", 33),
				FirstNameKana:        "?????????",
				Mail:                 "student-test001@calmato.jp",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
				SchoolType:           user.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
				Grade:                1,
			},
			isErr: true,
		},
		{
			name: "LastNameKana is pattern",
			req: &user.CreateStudentRequest{
				LastName:             "??????",
				FirstName:            "??????",
				LastNameKana:         "?????????",
				FirstNameKana:        "?????????",
				Mail:                 "student-test001@calmato.jp",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
				SchoolType:           user.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
				Grade:                1,
			},
			isErr: true,
		},
		{
			name: "FirstNameKana is min_len",
			req: &user.CreateStudentRequest{
				LastName:             "??????",
				FirstName:            "??????",
				LastNameKana:         "?????????",
				FirstNameKana:        "",
				Mail:                 "student-test001@calmato.jp",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
				SchoolType:           user.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
				Grade:                1,
			},
			isErr: true,
		},
		{
			name: "FirstNameKana is max_len",
			req: &user.CreateStudentRequest{
				LastName:             "??????",
				FirstName:            "??????",
				LastNameKana:         "?????????",
				FirstNameKana:        strings.Repeat("x", 33),
				Mail:                 "student-test001@calmato.jp",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
				SchoolType:           user.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
				Grade:                1,
			},
			isErr: true,
		},
		{
			name: "FirstNameKana is pattern",
			req: &user.CreateStudentRequest{
				LastName:             "??????",
				FirstName:            "??????",
				LastNameKana:         "?????????",
				FirstNameKana:        "?????????",
				Mail:                 "student-test001@calmato.jp",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
				SchoolType:           user.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
				Grade:                1,
			},
			isErr: true,
		},
		{
			name: "Mail is min_len",
			req: &user.CreateStudentRequest{
				LastName:             "??????",
				FirstName:            "??????",
				LastNameKana:         "?????????",
				FirstNameKana:        "?????????",
				Mail:                 "",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
				SchoolType:           user.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
				Grade:                1,
			},
			isErr: true,
		},
		{
			name: "Mail is max_len",
			req: &user.CreateStudentRequest{
				LastName:             "??????",
				FirstName:            "??????",
				LastNameKana:         "?????????",
				FirstNameKana:        "?????????",
				Mail:                 fmt.Sprintf("%s@calmato.jp", strings.Repeat("x", 245)),
				Password:             "12345678",
				PasswordConfirmation: "12345678",
				SchoolType:           user.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
				Grade:                1,
			},
			isErr: true,
		},
		{
			name: "Mail is email",
			req: &user.CreateStudentRequest{
				LastName:             "??????",
				FirstName:            "??????",
				LastNameKana:         "?????????",
				FirstNameKana:        "?????????",
				Mail:                 "student-test01",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
				SchoolType:           user.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
				Grade:                1,
			},
			isErr: true,
		},
		{
			name: "Password is min_len",
			req: &user.CreateStudentRequest{
				LastName:             "??????",
				FirstName:            "??????",
				LastNameKana:         "?????????",
				FirstNameKana:        "?????????",
				Mail:                 "student-test01@calmato.jp",
				Password:             strings.Repeat("x", 5),
				PasswordConfirmation: "12345678",
				SchoolType:           user.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
				Grade:                1,
			},
			isErr: true,
		},
		{
			name: "Password is max_len",
			req: &user.CreateStudentRequest{
				LastName:             "??????",
				FirstName:            "??????",
				LastNameKana:         "?????????",
				FirstNameKana:        "?????????",
				Mail:                 "student-test01@calmato.jp",
				Password:             strings.Repeat("x", 33),
				PasswordConfirmation: "12345678",
				SchoolType:           user.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
				Grade:                1,
			},
			isErr: true,
		},
		{
			name: "Password is pattern",
			req: &user.CreateStudentRequest{
				LastName:             "??????",
				FirstName:            "??????",
				LastNameKana:         "?????????",
				FirstNameKana:        "?????????",
				Mail:                 "student-test01@calmato.jp",
				Password:             "??????????????????",
				PasswordConfirmation: "12345678",
				SchoolType:           user.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
				Grade:                1,
			},
			isErr: true,
		},
		{
			name: "PasswordConfirmation is min_len",
			req: &user.CreateStudentRequest{
				LastName:             "??????",
				FirstName:            "??????",
				LastNameKana:         "?????????",
				FirstNameKana:        "?????????",
				Mail:                 "student-test01@calmato.jp",
				Password:             "12345678",
				PasswordConfirmation: "",
				SchoolType:           user.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
				Grade:                1,
			},
			isErr: true,
		},
		{
			name: "Custom validation is password match",
			req: &user.CreateStudentRequest{
				LastName:             "??????",
				FirstName:            "??????",
				LastNameKana:         "?????????",
				FirstNameKana:        "?????????",
				Mail:                 "student-test01@calmato.jp",
				Password:             "12345678",
				PasswordConfirmation: "87654321",
				SchoolType:           user.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
				Grade:                1,
			},
			isErr: true,
		},
		{
			name: "SchoolType is defined_only",
			req: &user.CreateStudentRequest{
				LastName:             "??????",
				FirstName:            "??????",
				LastNameKana:         "?????????",
				FirstNameKana:        "?????????",
				Mail:                 "student-test01@calmato.jp",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
				SchoolType:           user.SchoolType(-1),
				Grade:                1,
			},
			isErr: true,
		},
		{
			name: "Grade is min",
			req: &user.CreateStudentRequest{
				LastName:             "??????",
				FirstName:            "??????",
				LastNameKana:         "?????????",
				FirstNameKana:        "?????????",
				Mail:                 "student-test01@calmato.jp",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
				SchoolType:           user.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
				Grade:                -1,
			},
			isErr: true,
		},
		{
			name: "Grade is max",
			req: &user.CreateStudentRequest{
				LastName:             "??????",
				FirstName:            "??????",
				LastNameKana:         "?????????",
				FirstNameKana:        "?????????",
				Mail:                 "student-test01@calmato.jp",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
				SchoolType:           user.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
				Grade:                8,
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.CreateStudent(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestUpdateStudentMail(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *user.UpdateStudentMailRequest
		isErr bool
	}{
		{
			name: "success",
			req: &user.UpdateStudentMailRequest{
				Id:   "cvcTyJFfgoDQrqC1KDHbRe",
				Mail: "student-test01@calmato.jp",
			},
			isErr: false,
		},
		{
			name: "Id is min_len",
			req: &user.UpdateStudentMailRequest{
				Id:   "",
				Mail: "student-test01@calmato.jp",
			},
			isErr: true,
		},
		{
			name: "Mail is min_len",
			req: &user.UpdateStudentMailRequest{
				Id:   "cvcTyJFfgoDQrqC1KDHbRe",
				Mail: "",
			},
			isErr: true,
		},
		{
			name: "Mail is max_len",
			req: &user.UpdateStudentMailRequest{
				Id:   "cvcTyJFfgoDQrqC1KDHbRe",
				Mail: fmt.Sprintf("%s@calmato.jp", strings.Repeat("x", 245)),
			},
			isErr: true,
		},
		{
			name: "Mail is email",
			req: &user.UpdateStudentMailRequest{
				Id:   "cvcTyJFfgoDQrqC1KDHbRe",
				Mail: "student-test01",
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.UpdateStudentMail(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestUpdateStudentPassword(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *user.UpdateStudentPasswordRequest
		isErr bool
	}{
		{
			name: "success",
			req: &user.UpdateStudentPasswordRequest{
				Id:                   "cvcTyJFfgoDQrqC1KDHbRe",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
			isErr: false,
		},
		{
			name: "Id is min_len",
			req: &user.UpdateStudentPasswordRequest{
				Id:                   "",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
			isErr: true,
		},
		{
			name: "Password is min_len",
			req: &user.UpdateStudentPasswordRequest{
				Id:                   "cvcTyJFfgoDQrqC1KDHbRe",
				Password:             strings.Repeat("x", 5),
				PasswordConfirmation: "12345678",
			},
			isErr: true,
		},
		{
			name: "Password is max_len",
			req: &user.UpdateStudentPasswordRequest{
				Id:                   "cvcTyJFfgoDQrqC1KDHbRe",
				Password:             strings.Repeat("x", 33),
				PasswordConfirmation: "12345678",
			},
			isErr: true,
		},
		{
			name: "Password is pattern",
			req: &user.UpdateStudentPasswordRequest{
				Id:                   "cvcTyJFfgoDQrqC1KDHbRe",
				Password:             "??????????????????",
				PasswordConfirmation: "12345678",
			},
			isErr: true,
		},
		{
			name: "Custom validation is password match",
			req: &user.UpdateStudentPasswordRequest{
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
			err := validator.UpdateStudentPassword(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestDeleteStudent(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()
	tests := []struct {
		name  string
		req   *user.DeleteStudentRequest
		isErr bool
	}{
		{
			name: "success",
			req: &user.DeleteStudentRequest{
				Id: "cvcTyJFfgoDQrqC1KDHbRe",
			},
			isErr: false,
		},
		{
			name: "Id is min_len",
			req: &user.DeleteStudentRequest{
				Id: "",
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.DeleteStudent(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}
