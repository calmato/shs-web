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

func (v *requestValidation) UpsertTeacherShifts(req *lesson.UpsertTeacherShiftsRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(lesson.UpsertTeacherShiftsRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}
