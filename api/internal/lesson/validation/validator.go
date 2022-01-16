//nolint:lll
//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=./../../../mock/lesson/$GOPACKAGE/$GOFILE
package validation

import (
	"errors"
	"fmt"

	"github.com/calmato/shs-web/api/proto/lesson"
)

var ErrRequestValidation = errors.New("validation: invalid argument")

type RequestValidation interface {
	CreateLesson(req *lesson.CreateLessonRequest) error
	ListShiftSummaries(req *lesson.ListShiftSummariesRequest) error
	GetShiftSummary(req *lesson.GetShiftSummaryRequest) error
	UpdateShiftSummarySchedule(req *lesson.UpdateShiftSummaryScheduleRequest) error
	DeleteShiftSummary(req *lesson.DeleteShiftSummaryRequest) error
	ListShifts(req *lesson.ListShiftsRequest) error
	CreateShifts(req *lesson.CreateShiftsRequest) error
	ListTeacherSubmissionsByShiftSummaryIDs(req *lesson.ListTeacherSubmissionsByShiftSummaryIDsRequest) error
	ListTeacherShifts(req *lesson.ListTeacherShiftsRequest) error
	GetTeacherShifts(req *lesson.GetTeacherShiftsRequest) error
	UpsertTeacherShifts(req *lesson.UpsertTeacherShiftsRequest) error
	ListStudentSubmissionsByShiftSummaryIDs(req *lesson.ListStudentSubmissionsByShiftSummaryIDsRequest) error
	ListStudentSubmissionsByStudentIDs(req *lesson.ListStudentSubmissionsByStudentIDsRequest) error
	ListStudentShifts(req *lesson.ListStudentShiftsRequest) error
	GetStudentShifts(req *lesson.GetStudentShiftsRequest) error
	UpsertStudentShifts(req *lesson.UpsertStudentShiftsRequest) error
}

type requestValidation struct{}

func NewRequestValidation() RequestValidation {
	return &requestValidation{}
}

func validationError(msg string) error {
	return fmt.Errorf("%w: %s", ErrRequestValidation, msg)
}
