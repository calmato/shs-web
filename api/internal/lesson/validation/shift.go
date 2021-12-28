package validation

import "github.com/calmato/shs-web/api/proto/lesson"

func (v *requestValidation) ListShiftSummaries(req *lesson.ListShiftSummariesRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(lesson.ListShiftSummariesRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}

func (v *requestValidation) CreateShifts(req *lesson.CreateShiftsRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(lesson.CreateShiftsRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}
