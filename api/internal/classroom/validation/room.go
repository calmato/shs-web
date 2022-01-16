package validation

import "github.com/calmato/shs-web/api/proto/classroom"

func (v *requestValidation) GetRoom(req *classroom.GetRoomRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(classroom.GetRoomRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}

func (v *requestValidation) GetRoomsTotal(req *classroom.GetRoomsTotalRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(classroom.GetRoomsTotalRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}

func (v *requestValidation) UpdateRoomsTotal(req *classroom.UpdateRoomsTotalRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(classroom.UpdateRoomsTotalRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}
