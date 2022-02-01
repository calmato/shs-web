//nolint:lll
//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=./../../../mock/user/$GOPACKAGE/$GOFILE
package validation

import (
	"errors"
	"fmt"

	"github.com/calmato/shs-web/api/proto/user"
)

const eqFieldMessage = "%s must be a match %s"

var ErrRequestValidation = errors.New("validation: invalid argument")

type RequestValidation interface {
	ListTeachers(req *user.ListTeachersRequest) error
	MultiGetTeachers(req *user.MultiGetTeachersRequest) error
	GetTeacher(req *user.GetTeacherRequest) error
	CreateTeacher(req *user.CreateTeacherRequest) error
	UpdateTeacherMail(req *user.UpdateTeacherMailRequest) error
	UpdateTeacherPassword(req *user.UpdateTeacherPasswordRequest) error
	UpdateTeacherRole(req *user.UpdateTeacherRoleRequest) error
	DeleteTeacher(req *user.DeleteTeacherRequest) error
	ListStudents(req *user.ListStudentsRequest) error
	MultiGetStudents(req *user.MultiGetStudentsRequest) error
	GetStudent(req *user.GetStudentRequest) error
	CreateStudent(req *user.CreateStudentRequest) error
	UpdateStudentMail(req *user.UpdateStudentMailRequest) error
	UpdateStudentPassword(req *user.UpdateStudentPasswordRequest) error
	DeleteStudent(req *user.DeleteStudentRequest) error
}

type requestValidation struct{}

func NewRequestValidation() RequestValidation {
	return &requestValidation{}
}

func validationError(msg string) error {
	return fmt.Errorf("%w: %s", ErrRequestValidation, msg)
}
