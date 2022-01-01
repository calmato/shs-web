package validation

import "github.com/calmato/shs-web/api/proto/lesson"

func (v *requestValidation) UpsertTeacherShifts(req *lesson.UpsertTeacherShiftsRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(lesson.UpsertTeacherShiftsRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}
