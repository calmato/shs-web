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
	Hello(req *user.HelloRequest) error
	ListTeachers(req *user.ListTeachersRequest) error
	CreateTeacher(req *user.CreateTeacherRequest) error
}

type requestValidation struct{}

func NewRequestValidation() RequestValidation {
	return &requestValidation{}
}

func validationError(msg string) error {
	return fmt.Errorf("%w: %s", ErrRequestValidation, msg)
}
