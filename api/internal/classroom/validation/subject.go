package validation

import "github.com/calmato/shs-web/api/proto/classroom"

func (v *requestValidation) ListSubjects(req *classroom.ListSubjectsRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(classroom.ListSubjectsRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}

func (v *requestValidation) MultiGetSubjects(req *classroom.MultiGetSubjectsRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(classroom.MultiGetSubjectsRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}

func (v *requestValidation) GetSubject(req *classroom.GetSubjectRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(classroom.GetSubjectRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}

func (v *requestValidation) CreateSubject(req *classroom.CreateSubjectRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(classroom.CreateSubjectRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}

func (v *requestValidation) UpdateSubject(req *classroom.UpdateSubjectRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(classroom.UpdateSubjectRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}

func (v *requestValidation) DeleteSubject(req *classroom.DeleteSubjectRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(classroom.DeleteSubjectRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}
