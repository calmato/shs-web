package validation

import "github.com/calmato/shs-web/api/proto/user"

func (v *requestValidation) Hello(req *user.HelloRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(user.HelloRequestValidationError)
		return validationError(validate.Error())
	}
	return nil
}
