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

func (v *requestValidation) GetTeacher(req *user.GetTeacherRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(user.GetTeacherRequestValidationError)
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

func (v *requestValidation) UpdateTeacherMail(req *user.UpdateTeacherMailRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(user.UpdateTeacherMailRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}

func (v *requestValidation) UpdateTeacherPassword(req *user.UpdateTeacherPasswordRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(user.UpdateTeacherPasswordRequestValidationError)
		return validationError(validate.Error())
	}

	if req.Password != req.PasswordConfirmation {
		msg := fmt.Sprintf(eqFieldMessage, "PasswordConfirmation", "Password")
		return validationError(msg)
	}

	return nil
}

func (v *requestValidation) UpdateTeacherRole(req *user.UpdateTeacherRoleRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(user.UpdateTeacherRoleRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}

func (v *requestValidation) DeleteTeacher(req *user.DeleteTeacherRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(user.DeleteTeacherRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}
