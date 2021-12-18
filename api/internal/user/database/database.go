//nolint:lll
//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=./../../../mock/user/$GOPACKAGE/$GOFILE
package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/calmato/shs-web/api/internal/user/entity"
	"github.com/calmato/shs-web/api/pkg/database"
	"github.com/calmato/shs-web/api/pkg/firebase/authentication"
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
	Auth     authentication.Client
}

type Database struct {
	Student Student
	Teacher Teacher
}

func NewDatabase(params *Params) *Database {
	return &Database{
		Student: NewStudent(params.Database, params.Auth),
		Teacher: NewTeacher(params.Database, params.Auth),
	}
}

/**
 * interface
 */
type Student interface{}

type Teacher interface {
	List(ctx context.Context, p *ListTeachersParams, fields ...string) (entity.Teachers, error)
	Get(ctx context.Context, id string, fields ...string) (*entity.Teacher, error)
	Create(ctx context.Context, t *entity.Teacher) error
	Count(ctx context.Context) (int64, error)
}

/**
 * params
 */
type ListTeachersParams struct {
	Limit  int
	Offset int
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
