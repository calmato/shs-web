package validation

import (
	"fmt"

	"github.com/calmato/shs-web/api/proto/user"
)

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
