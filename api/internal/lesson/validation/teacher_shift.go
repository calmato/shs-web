package validation

import "github.com/calmato/shs-web/api/proto/lesson"

func (v *requestValidation) ListTeacherSubmissionsByShiftSummaryIDs(
	req *lesson.ListTeacherSubmissionsByShiftSummaryIDsRequest,
) error {
	if err := req.Validate(); err != nil {
		validate := err.(lesson.ListTeacherSubmissionsByShiftSummaryIDsRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}

func (v *requestValidation) ListTeacherSubmissionsByTeacherIDs(
	req *lesson.ListTeacherSubmissionsByTeacherIDsRequest,
) error {
	if err := req.Validate(); err != nil {
		validate := err.(lesson.ListTeacherSubmissionsByTeacherIDsRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}

func (v *requestValidation) ListTeacherShifts(req *lesson.ListTeacherShiftsRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(lesson.ListTeacherShiftsRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}

func (v *requestValidation) GetTeacherShifts(req *lesson.GetTeacherShiftsRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(lesson.GetTeacherShiftsRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}

func (v *requestValidation) UpsertTeacherShifts(req *lesson.UpsertTeacherShiftsRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(lesson.UpsertTeacherShiftsRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}
