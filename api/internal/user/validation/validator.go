//nolint:lll
//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=./../../../mock/user/$GOPACKAGE/$GOFILE
package validation

import (
	"errors"
	"fmt"

	"github.com/calmato/shs-web/api/proto/user"
)

const eqFieldMessage = "%s must be a match %s"

var ErrRequestValidation = errors.New("validation: invalid argument")

type RequestValidation interface {
	ListTeachers(req *user.ListTeachersRequest) error
	GetTeacher(req *user.GetTeacherRequest) error
	CreateTeacher(req *user.CreateTeacherRequest) error
	UpdateTeacherMail(req *user.UpdateTeacherMailRequest) error
	UpdateTeacherPassword(req *user.UpdateTeacherPasswordRequest) error
	DeleteTeacher(req *user.DeleteTeacherRequest) error
	GetStudent(req *user.GetStudentRequest) error
}

type requestValidation struct{}

func NewRequestValidation() RequestValidation {
	return &requestValidation{}
}

func validationError(msg string) error {
	return fmt.Errorf("%w: %s", ErrRequestValidation, msg)
}
