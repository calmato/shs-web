//nolint:lll
//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=./../../../mock/classroom/$GOPACKAGE/$GOFILE
package validation

import (
	"errors"
	"fmt"

	"github.com/calmato/shs-web/api/proto/classroom"
)

var ErrRequestValidation = errors.New("validation: invalid argument")

type RequestValidation interface {
	ListSubjects(req *classroom.ListSubjectsRequest) error
	MultiGetSubjects(req *classroom.MultiGetSubjectsRequest) error
	GetSubject(req *classroom.GetSubjectRequest) error
	CreateSubject(req *classroom.CreateSubjectRequest) error
	UpdateSubject(req *classroom.UpdateSubjectRequest) error
	DeleteSubject(req *classroom.DeleteSubjectRequest) error
	MultiGetTeacherSubjects(req *classroom.MultiGetTeacherSubjectsRequest) error
	GetTeacherSubject(req *classroom.GetTeacherSubjectRequest) error
	UpsertTeacherSubject(req *classroom.UpsertTeacherSubjectRequest) error
	MultiGetStudentSubjects(req *classroom.MultiGetStudentSubjectsRequest) error
	GetStudentSubject(req *classroom.GetStudentSubjectRequest) error
	UpsertStudentSubject(req *classroom.UpsertStudentSubjectRequest) error
	ListSchedules(req *classroom.ListSchedulesRequest) error
	GetSchedule(req *classroom.GetScheduleRequest) error
	UpdateSchedules(req *classroom.UpdateSchedulesRequest) error
	GetRoom(req *classroom.GetRoomRequest) error
	GetRoomsTotal(req *classroom.GetRoomsTotalRequest) error
	UpdateRoomsTotal(req *classroom.UpdateRoomsTotalRequest) error
}

type requestValidation struct{}

func NewRequestValidation() RequestValidation {
	return &requestValidation{}
}

func validationError(msg string) error {
	return fmt.Errorf("%w: %s", ErrRequestValidation, msg)
}
