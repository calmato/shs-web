//nolint:lll
//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=./../../../mock/lesson/$GOPACKAGE/$GOFILE
package database

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/calmato/shs-web/api/internal/lesson/entity"
	"github.com/calmato/shs-web/api/pkg/database"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var (
	ErrInvalidArgument = errors.New("database: invalid argument")
	ErrNotFound        = errors.New("database: not found")
	ErrNotImplemented  = errors.New("database: not implemented")
	ErrInternal        = errors.New("database: internal")
	ErrUnknown         = errors.New("database: unknown")
)

type Params struct {
	Database *database.Client
}

type Database struct {
	ShiftSummary ShiftSummary
	Shift        Shift
}

func NewDatabase(params *Params) *Database {
	return &Database{
		ShiftSummary: NewShiftSummary(params.Database),
		Shift:        NewShift(params.Database),
	}
}

/**
 * interface
 */
type ShiftSummary interface {
	List(ctx context.Context, p *ListShiftSummariesParams, fields ...string) (entity.ShiftSummaries, error)
	Get(ctx context.Context, id int64, fields ...string) (*entity.ShiftSummary, error)
	UpdateSchedule(ctx context.Context, id int64, openAt, endAt time.Time) error
	Count(ctx context.Context) (int64, error)
}

type Shift interface {
	ListBySummaryID(ctx context.Context, summaryID int64, fields ...string) (entity.Shifts, error)
	MultipleCreate(ctx context.Context, summary *entity.ShiftSummary, shifts entity.Shifts) error
}

/**
 * params
 */
type ListShiftSummariesParams struct {
	Limit  int
	Offset int
	Status entity.ShiftStatus
}

/**
 * private methods
 */
func dbError(err error) error {
	if err == nil {
		return nil
	}

	//nolint:gocritic
	switch err.(type) {
	case *mysql.MySQLError:
		return fmt.Errorf("%w: %s", ErrUnknown, err)
	}

	switch {
	case errors.Is(err, gorm.ErrEmptySlice),
		errors.Is(err, gorm.ErrInvalidData),
		errors.Is(err, gorm.ErrInvalidField),
		errors.Is(err, gorm.ErrInvalidTransaction),
		errors.Is(err, gorm.ErrInvalidValue),
		errors.Is(err, gorm.ErrInvalidValueOfLength),
		errors.Is(err, gorm.ErrMissingWhereClause),
		errors.Is(err, gorm.ErrModelValueRequired),
		errors.Is(err, gorm.ErrPrimaryKeyRequired):
		return fmt.Errorf("%w: %s", ErrInvalidArgument, err)
	case errors.Is(err, gorm.ErrRecordNotFound):
		return fmt.Errorf("%w: %s", ErrNotFound, err)
	case errors.Is(err, gorm.ErrNotImplemented):
		return fmt.Errorf("%w: %s", ErrNotImplemented, err)
	case errors.Is(err, gorm.ErrDryRunModeUnsupported),
		errors.Is(err, gorm.ErrInvalidDB),
		errors.Is(err, gorm.ErrRegistered),
		errors.Is(err, gorm.ErrUnsupportedDriver),
		errors.Is(err, gorm.ErrUnsupportedRelation):
		return fmt.Errorf("%w: %s", ErrInternal, err)
	default:
		return fmt.Errorf("%w: %s", ErrUnknown, err)
	}
}
