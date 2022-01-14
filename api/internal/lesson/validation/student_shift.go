package validation

import "github.com/calmato/shs-web/api/proto/lesson"

func (v *requestValidation) ListStudentSubmissionsByShiftSummaryIDs(
	req *lesson.ListStudentSubmissionsByShiftSummaryIDsRequest,
) error {
	if err := req.Validate(); err != nil {
		validate := err.(lesson.ListStudentSubmissionsByShiftSummaryIDsRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}

func (v *requestValidation) ListStudentShifts(req *lesson.ListStudentShiftsRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(lesson.ListStudentShiftsRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}

func (v *requestValidation) GetStudentShifts(req *lesson.GetStudentShiftsRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(lesson.GetStudentShiftsRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}

func (v *requestValidation) UpsertStudentShifts(req *lesson.UpsertStudentShiftsRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(lesson.UpsertStudentShiftsRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}
