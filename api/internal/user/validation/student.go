package validation

import (
	"github.com/calmato/shs-web/api/proto/user"
)

func (v *requestValidation) GetStudent(req *user.GetStudentRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(user.GetStudentRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}
