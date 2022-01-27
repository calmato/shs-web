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
	ErrAlreadyExists   = errors.New("database: already exists")
	ErrNotImplemented  = errors.New("database: not implemented")
	ErrInternal        = errors.New("database: internal")
	ErrUnknown         = errors.New("database: unknown")
)

type Params struct {
	Database *database.Client
}

type Database struct {
	Lesson            Lesson
	ShiftSummary      ShiftSummary
	Shift             Shift
	TeacherSubmission TeacherSubmission
	TeacherShift      TeacherShift
	StudentSubmission StudentSubmission
	StudentShift      StudentShift
}

func NewDatabase(params *Params) *Database {
	return &Database{
		Lesson:            NewLesson(params.Database),
		ShiftSummary:      NewShiftSummary(params.Database),
		Shift:             NewShift(params.Database),
		TeacherSubmission: NewTeacherSubmission(params.Database),
		TeacherShift:      NewTeacherShift(params.Database),
		StudentSubmission: NewStudentSubmission(params.Database),
		StudentShift:      NewStudentShift(params.Database),
	}
}

/**
 * interface
 */
type Lesson interface {
	List(ctx context.Context, p *ListLessonsParams, fields ...string) (entity.Lessons, error)
	Create(ctx context.Context, lesson *entity.Lesson) error
	Update(ctx context.Context, lessonID int64, lesson *entity.Lesson) error
	Delete(ctx context.Context, lessonID int64) error
	Count(ctx context.Context, p *ListLessonsParams) (int64, error)
}

type ShiftSummary interface {
	List(ctx context.Context, p *ListShiftSummariesParams, fields ...string) (entity.ShiftSummaries, error)
	Get(ctx context.Context, id int64, fields ...string) (*entity.ShiftSummary, error)
	UpdateSchedule(ctx context.Context, id int64, openAt, endAt time.Time) error
	Delete(ctx context.Context, id int64) error
	Count(ctx context.Context) (int64, error)
}

type Shift interface {
	List(ctx context.Context, p *ListShiftsParams, fields ...string) (entity.Shifts, error)
	MultiGet(ctx context.Context, ids []int64, fields ...string) (entity.Shifts, error)
	Get(ctx context.Context, id int64, fields ...string) (*entity.Shift, error)
	MultipleCreate(ctx context.Context, summary *entity.ShiftSummary, shifts entity.Shifts) error
}

type TeacherSubmission interface {
	ListByShiftSummaryIDs(
		ctx context.Context, teacherID string, summaryIDs []int64, fields ...string,
	) (entity.TeacherSubmissions, error)
	Get(ctx context.Context, teacherID string, summaryID int64, fields ...string) (*entity.TeacherSubmission, error)
}

type TeacherShift interface {
	ListByShiftSummaryID(
		ctx context.Context, teacherIDs []string, summaryID int64, fields ...string,
	) (entity.TeacherShifts, error)
	ListByShiftID(ctx context.Context, shiftID int64, fields ...string) (entity.TeacherShifts, error)
	Replace(ctx context.Context, submission *entity.TeacherSubmission, shifts entity.TeacherShifts) error
}

type StudentSubmission interface {
	ListByShiftSummaryIDs(
		ctx context.Context, studentID string, summaryIDs []int64, fields ...string,
	) (entity.StudentSubmissions, error)
	ListByStudentIDs(
		ctx context.Context, studentIDs []string, summaryID int64, fields ...string,
	) (entity.StudentSubmissions, error)
	Get(ctx context.Context, studentID string, summaryID int64, fields ...string) (*entity.StudentSubmission, error)
}

type StudentShift interface {
	ListByShiftSummaryID(
		ctx context.Context, studentIDs []string, summaryID int64, fields ...string,
	) (entity.StudentShifts, error)
	ListByShiftID(ctx context.Context, shiftID int64, fields ...string) (entity.StudentShifts, error)
	Replace(ctx context.Context, submission *entity.StudentSubmission, shifts entity.StudentShifts) error
}

/**
 * params
 */
type OrderBy int32

const (
	OrderByNone OrderBy = iota
	OrderByAsc
	OrderByDesc
)

type ListLessonsParams struct {
	Limit          int
	Offset         int
	ShiftSummaryID int64
	ShiftID        int64
	TeacherID      string
	StudentID      string
}

type ListShiftSummariesParams struct {
	Limit   int
	Offset  int
	Status  entity.ShiftStatus
	OrderBy OrderBy
}

type ListShiftsParams struct {
	Limit          int
	Offset         int
	ShiftSummaryID int64
	ShiftID        int64
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
