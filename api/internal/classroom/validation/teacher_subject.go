package validation

import "github.com/calmato/shs-web/api/proto/classroom"

func (v *requestValidation) MultiGetTeacherSubjects(req *classroom.MultiGetTeacherSubjectsRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(classroom.MultiGetTeacherSubjectsRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}

func (v *requestValidation) GetTeacherSubject(req *classroom.GetTeacherSubjectRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(classroom.GetTeacherSubjectRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}

func (v *requestValidation) UpdateTeacherSubject(req *classroom.UpdateTeacherSubjectRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(classroom.UpdateTeacherSubjectRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}
