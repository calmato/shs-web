package database

import (
	"context"
	"time"

	"github.com/calmato/shs-web/api/internal/lesson/entity"
	"github.com/calmato/shs-web/api/pkg/database"
	"github.com/calmato/shs-web/api/pkg/jst"
)

const shiftTable = "shifts"

var shiftFields = []string{
	"id", "shift_summary_id", "date",
	"start_time", "end_time", "created_at", "updated_at",
}

type shift struct {
	db  *database.Client
	now func() time.Time
}

func NewShift(db *database.Client) Shift {
	return &shift{
		db:  db,
		now: jst.Now,
	}
}

func (s *shift) List(ctx context.Context, params *ListShiftsParams, fields ...string) (entity.Shifts, error) {
	var shifts entity.Shifts
	if len(fields) == 0 {
		fields = shiftFields
	}

	stmt := s.db.DB.Table(shiftTable).Select(fields)
	if params.ShiftID > 0 {
		stmt.Where("id = ?", params.ShiftID)
	}
	if params.ShiftSummaryID > 0 {
		stmt.Where("shift_summary_id = ?", params.ShiftSummaryID)
	}
	if params.Limit > 0 {
		stmt.Limit(params.Limit)
	}
	if params.Offset > 0 {
		stmt.Limit(params.Offset)
	}

	err := stmt.Find(&shifts).Error
	return shifts, dbError(err)
}

func (s *shift) ListByDuration(ctx context.Context, since, until time.Time, fields ...string) (entity.Shifts, error) {
	var shifts entity.Shifts
	if len(fields) == 0 {
		fields = shiftFields
	}

	stmt := s.db.DB.Table(shiftTable).Select(fields).
		Where("date >= ?", since).
		Where("date < ?", until)

	err := stmt.Find(&shifts).Error
	return shifts, dbError(err)
}

func (s *shift) ListBySummaryID(ctx context.Context, summaryID int64, fields ...string) (entity.Shifts, error) {
	var shifts entity.Shifts
	if len(fields) == 0 {
		fields = shiftFields
	}

	stmt := s.db.DB.Table(shiftTable).Select(fields).
		Where("shift_summary_id = ?", summaryID)

	err := stmt.Find(&shifts).Error
	return shifts, dbError(err)
}

func (s *shift) MultiGet(ctx context.Context, ids []int64, fields ...string) (entity.Shifts, error) {
	var shifts entity.Shifts
	if len(fields) == 0 {
		fields = shiftFields
	}

	stmt := s.db.DB.Table(shiftTable).Select(fields).
		Where("id IN (?)", ids)

	err := stmt.Find(&shifts).Error
	return shifts, dbError(err)
}

func (s *shift) Get(ctx context.Context, id int64, fields ...string) (*entity.Shift, error) {
	var shift *entity.Shift
	if len(fields) == 0 {
		fields = shiftFields
	}

	stmt := s.db.DB.Table(shiftTable).Select(fields).
		Where("id = ?", id)

	err := stmt.First(&shift).Error
	if err != nil {
		return nil, dbError(err)
	}
	return shift, nil
}

func (s *shift) MultipleCreate(ctx context.Context, summary *entity.ShiftSummary, shifts entity.Shifts) error {
	now := s.now()

	tx, err := s.db.Begin()
	if err != nil {
		return dbError(err)
	}
	defer s.db.Close(tx)

	summary.CreatedAt = now
	summary.UpdatedAt = now
	err = tx.Table(shiftSummaryTable).Create(&summary).Error
	if err != nil {
		tx.Rollback()
		return dbError(err)
	}

	for i := range shifts {
		shifts[i].ShiftSummaryID = summary.ID
		shifts[i].CreatedAt = now
		shifts[i].UpdatedAt = now
	}
	err = tx.Table(shiftTable).Create(&shifts).Error
	if err != nil {
		tx.Rollback()
		return dbError(err)
	}
	return dbError(tx.Commit().Error)
}
