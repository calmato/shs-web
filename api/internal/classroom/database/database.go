//nolint:lll
//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=./../../../mock/classroom/$GOPACKAGE/$GOFILE
package database

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/calmato/shs-web/api/internal/classroom/entity"
	"github.com/calmato/shs-web/api/pkg/database"
	"github.com/calmato/shs-web/api/proto/classroom"
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
}

type Database struct {
	Subject        Subject
	TeacherSubject TeacherSubject
	StudentSubject StudentSubject
	Schedule       Schedule
	Room           Room
}

func NewDatabase(params *Params) *Database {
	return &Database{
		Subject:        NewSubject(params.Database),
		TeacherSubject: NewTeacherSubject(params.Database),
		StudentSubject: NewStudentSubject(params.Database),
		Schedule:       NewSchedule(params.Database),
		Room:           NewRoom(params.Database),
	}
}

/**
 * interface
 */
type Subject interface {
	List(ctx context.Context, p *ListSubjectsParams, fields ...string) (entity.Subjects, error)
	MultiGet(ctx context.Context, ids []int64, fields ...string) (entity.Subjects, error)
	Get(ctx context.Context, id int64, fields ...string) (*entity.Subject, error)
	Create(ctx context.Context, subject *entity.Subject) error
	Update(ctx context.Context, subjectID int64, subject *entity.Subject) error
	Delete(ctx context.Context, subjectID int64) error
	Count(ctx context.Context) (int64, error)
}

type TeacherSubject interface {
	ListByTeacherIDs(ctx context.Context, teacherIDs []string, fields ...string) (entity.TeacherSubjects, error)
	Replace(ctx context.Context, schoolType entity.SchoolType, subjects entity.TeacherSubjects) error
}

type StudentSubject interface {
	ListByStudentIDs(ctx context.Context, studentIDs []string, fields ...string) (entity.StudentSubjects, error)
	Replace(ctx context.Context, schoolType entity.SchoolType, subjects entity.StudentSubjects) error
}

type Schedule interface {
	List(ctx context.Context, fields ...string) (entity.Schedules, error)
	Get(ctx context.Context, weekday time.Weekday, fields ...string) (*entity.Schedule, error)
	MultipleUpdate(ctx context.Context, schedules entity.Schedules) error
}

type Room interface {
	Get(ctx context.Context, id int32, fields ...string) (*entity.Room, error)
	Replace(ctx context.Context, rooms entity.Rooms) error
	Count(ctx context.Context) (int64, error)
}

/**
 * params
 */
type ListSubjectsParams struct {
	SchoolType classroom.SchoolType
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
