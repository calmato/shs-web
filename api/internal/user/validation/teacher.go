package validation

import (
	"fmt"

	"github.com/calmato/shs-web/api/proto/user"
)

func (v *requestValidation) ListTeachers(req *user.ListTeachersRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(user.ListTeachersRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}

func (v *requestValidation) CreateTeacher(req *user.CreateTeacherRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(user.CreateTeacherRequestValidationError)
		return validationError(validate.Error())
	}

	if req.Password != req.PasswordConfirmation {
		msg := fmt.Sprintf(eqFieldMessage, "PasswordConfirmation", "Password")
		return validationError(msg)
	}

	return nil
}
