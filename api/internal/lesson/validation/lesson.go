package validation

import "github.com/calmato/shs-web/api/proto/lesson"

func (v *requestValidation) ListLessons(req *lesson.ListLessonsRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(lesson.ListLessonsRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}

func (v *requestValidation) ListLessonsByDuration(req *lesson.ListLessonsByDurationRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(lesson.ListLessonsByDurationRequestValidationError)
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

func (v *requestValidation) UpdateLesson(req *lesson.UpdateLessonRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(lesson.UpdateLessonRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}

func (v *requestValidation) DeleteLesson(req *lesson.DeleteLessonRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(lesson.DeleteLessonRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}
