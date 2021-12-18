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
}

type requestValidation struct{}

func NewRequestValidation() RequestValidation {
	return &requestValidation{}
}

func validationError(msg string) error {
	return fmt.Errorf("%w: %s", ErrRequestValidation, msg)
}