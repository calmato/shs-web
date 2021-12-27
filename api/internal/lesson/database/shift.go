package database

import (
	"context"
	"time"

	"github.com/calmato/shs-web/api/internal/lesson/entity"
	"github.com/calmato/shs-web/api/pkg/database"
	"github.com/calmato/shs-web/api/pkg/jst"
)

const shiftTable = "shifts"

// var shiftFields = []string{
// 	"id", "shift_summary_id", "date",
// 	"start_time", "end_time", "created_at", "updated_at",
// }

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
