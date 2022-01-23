//nolint:lll
//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=./../../../mock/user/$GOPACKAGE/$GOFILE
package database

import (
	"context"
	"errors"
	"fmt"

	"firebase.google.com/go/v4/errorutils"
	"github.com/calmato/shs-web/api/internal/user/entity"
	"github.com/calmato/shs-web/api/pkg/database"
	"github.com/calmato/shs-web/api/pkg/firebase/authentication"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var (
	ErrInvalidArgument = errors.New("database: invalid argument")
	ErrNotFound        = errors.New("database: not found")
	ErrAlreadyExists   = errors.New("database: already exists")
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
type Student interface {
	List(ctx context.Context, p *ListStudentsParams, fields ...string) (entity.Students, error)
	MultiGet(ctx context.Context, ids []string, fields ...string) (entity.Students, error)
	Get(ctx context.Context, id string, fields ...string) (*entity.Student, error)
	Create(ctx context.Context, s *entity.Student) error
	Delete(ctx context.Context, studentID string) error
	Count(ctx context.Context) (int64, error)
}

type Teacher interface {
	List(ctx context.Context, p *ListTeachersParams, fields ...string) (entity.Teachers, error)
	MultiGet(ctx context.Context, ids []string, fields ...string) (entity.Teachers, error)
	Get(ctx context.Context, id string, fields ...string) (*entity.Teacher, error)
	Create(ctx context.Context, t *entity.Teacher) error
	UpdateMail(ctx context.Context, teacherID string, mail string) error
	UpdatePassword(ctx context.Context, teacherID string, password string) error
	UpdateRole(ctx context.Context, teacherID string, role entity.Role) error
	Delete(ctx context.Context, teacherID string) error
	Count(ctx context.Context) (int64, error)
}

/**
 * params
 */
type ListTeachersParams struct {
	Limit  int
	Offset int
}

type ListStudentsParams struct {
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
	switch err := err.(type) {
	case *mysql.MySQLError:
		if err.Number == 1062 {
			return fmt.Errorf("%w: %s", ErrAlreadyExists, err)
		}
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
		errors.Is(err, gorm.ErrPrimaryKeyRequired),
		errorutils.IsInvalidArgument(err),
		errorutils.IsOutOfRange(err),
		errorutils.IsUnauthenticated(err),
		errorutils.IsPermissionDenied(err),
		errorutils.IsFailedPrecondition(err):
		return fmt.Errorf("%w: %s", ErrInvalidArgument, err)
	case errors.Is(err, gorm.ErrRecordNotFound),
		errorutils.IsNotFound(err):
		return fmt.Errorf("%w: %s", ErrNotFound, err)
	case errorutils.IsConflict(err),
		errorutils.IsAlreadyExists(err):
		return fmt.Errorf("%w: %s", ErrAlreadyExists, err)
	case errors.Is(err, gorm.ErrNotImplemented),
		errorutils.IsUnavailable(err):
		return fmt.Errorf("%w: %s", ErrNotImplemented, err)
	case errors.Is(err, gorm.ErrDryRunModeUnsupported),
		errors.Is(err, gorm.ErrInvalidDB),
		errors.Is(err, gorm.ErrRegistered),
		errors.Is(err, gorm.ErrUnsupportedDriver),
		errors.Is(err, gorm.ErrUnsupportedRelation),
		errorutils.IsAborted(err),
		errorutils.IsCancelled(err),
		errorutils.IsDataLoss(err),
		errorutils.IsResourceExhausted(err),
		errorutils.IsInternal(err),
		errorutils.IsDeadlineExceeded(err):
		return fmt.Errorf("%w: %s", ErrInternal, err)
	default:
		return fmt.Errorf("%w: %s", ErrUnknown, err)
	}
}
