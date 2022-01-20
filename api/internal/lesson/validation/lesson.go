package validation

import "github.com/calmato/shs-web/api/proto/lesson"

func (v *requestValidation) ListLessonsByShiftSummaryID(req *lesson.ListLessonsByShiftSummaryIDRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(lesson.ListLessonsByShiftSummaryIDRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}

func (v *requestValidation) ListLessonsByTeacherID(req *lesson.ListLessonsByTeacherIDRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(lesson.ListLessonsByTeacherIDRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}

func (v *requestValidation) ListLessonsByStudentID(req *lesson.ListLessonsByStudentIDRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(lesson.ListLessonsByStudentIDRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}

func (v *requestValidation) CreateLesson(req *lesson.CreateLessonRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(lesson.CreateLessonRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}
