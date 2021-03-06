package validation

import "github.com/calmato/shs-web/api/proto/lesson"

func (v *requestValidation) ListShiftSummaries(req *lesson.ListShiftSummariesRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(lesson.ListShiftSummariesRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}

func (v *requestValidation) GetShiftSummary(req *lesson.GetShiftSummaryRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(lesson.GetShiftSummaryRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}

func (v *requestValidation) UpdateShiftSummarySchedule(req *lesson.UpdateShiftSummaryScheduleRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(lesson.UpdateShiftSummaryScheduleRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}

func (v *requestValidation) UpdateShiftSummaryDecided(req *lesson.UpdateShiftSummaryDecidedRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(lesson.UpdateShiftSummaryDecidedRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}

func (v *requestValidation) DeleteShiftSummary(req *lesson.DeleteShiftSummaryRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(lesson.DeleteShiftSummaryRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}

func (v *requestValidation) ListShifts(req *lesson.ListShiftsRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(lesson.ListShiftsRequestValidationError)
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

func (v *requestValidation) ListSubmissions(req *lesson.ListSubmissionsRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(lesson.ListSubmissionsRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}
