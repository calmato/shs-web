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
	ListShiftSummaries(req *lesson.ListShiftSummariesRequest) error
	CreateShifts(req *lesson.CreateShiftsRequest) error
}

type requestValidation struct{}

func NewRequestValidation() RequestValidation {
	return &requestValidation{}
}

func validationError(msg string) error {
	return fmt.Errorf("%w: %s", ErrRequestValidation, msg)
}
