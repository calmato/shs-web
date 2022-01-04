package validation

import "github.com/calmato/shs-web/api/proto/classroom"

func (v *requestValidation) MultiGetStudentSubjects(req *classroom.MultiGetStudentSubjectsRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(classroom.MultiGetStudentSubjectsRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}

func (v *requestValidation) GetStudentSubject(req *classroom.GetStudentSubjectRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(classroom.GetStudentSubjectRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}

func (v *requestValidation) UpsertStudentSubject(req *classroom.UpsertStudentSubjectRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(classroom.UpsertStudentSubjectRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}
