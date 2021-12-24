package validation

import "github.com/calmato/shs-web/api/proto/classroom"

func (v *requestValidation) ListSchedules(req *classroom.ListSchedulesRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(classroom.ListSchedulesRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}

func (v *requestValidation) GetSchedule(req *classroom.GetScheduleRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(classroom.GetScheduleRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}
