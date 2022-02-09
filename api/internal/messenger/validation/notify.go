package validation

import "github.com/calmato/shs-web/api/proto/messenger"

func (v *requestValidation) NotifyLessonDecided(req *messenger.NotifyLessonDecidedRequest) error {
	if err := req.Validate(); err != nil {
		validate := err.(messenger.NotifyLessonDecidedRequestValidationError)
		return validationError(validate.Error())
	}

	return nil
}
