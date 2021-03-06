package validation

import (
	"fmt"

	"github.com/calmato/shs-web/api/proto/user"
)

func (v *requestValidation) ListStudents(req *user.ListStudentsRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(user.ListStudentsRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}

func (v *requestValidation) MultiGetStudents(req *user.MultiGetStudentsRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(user.MultiGetStudentsRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}

func (v *requestValidation) GetStudent(req *user.GetStudentRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(user.GetStudentRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}

func (v *requestValidation) CreateStudent(req *user.CreateStudentRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(user.CreateStudentRequestValidationError)
		return validationError(validate.Error())
	}

	if req.Password != req.PasswordConfirmation {
		msg := fmt.Sprintf(eqFieldMessage, "PasswordConfirmation", "Password")
		return validationError(msg)
	}

	return nil
}

func (v *requestValidation) UpdateStudentMail(req *user.UpdateStudentMailRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(user.UpdateStudentMailRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}

func (v *requestValidation) UpdateStudentPassword(req *user.UpdateStudentPasswordRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(user.UpdateStudentPasswordRequestValidationError)
		return validationError(validate.Error())
	}

	if req.Password != req.PasswordConfirmation {
		msg := fmt.Sprintf(eqFieldMessage, "PasswordConfirmation", "Password")
		return validationError(msg)
	}

	return nil
}

func (v *requestValidation) DeleteStudent(req *user.DeleteStudentRequest) error {
	if err := req.Validate(); err != nil {
		validete := err.(user.DeleteStudentRequestValidationError)
		return validationError(validete.Error())
	}

	return nil
}
