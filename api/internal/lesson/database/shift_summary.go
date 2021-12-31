package database

import (
	"context"
	"time"

	"github.com/calmato/shs-web/api/internal/lesson/entity"
	"github.com/calmato/shs-web/api/pkg/database"
	"github.com/calmato/shs-web/api/pkg/jst"
)

const shiftSummaryTable = "shift_summaries"

var shiftSummaryFields = []string{
	"id", "year_month", "open_at", "end_at", "created_at", "updated_at",
}

type shiftSummary struct {
	db  *database.Client
	now func() time.Time
}

func NewShiftSummary(db *database.Client) ShiftSummary {
	return &shiftSummary{
		db:  db,
		now: jst.Now,
	}
}

func (s *shiftSummary) List(
	ctx context.Context, params *ListShiftSummariesParams, fields ...string,
) (entity.ShiftSummaries, error) {
	now := s.now()
	var summaries entity.ShiftSummaries
	if len(fields) == 0 {
		fields = shiftSummaryFields
	}

	stmt := s.db.DB.Table(shiftSummaryTable).Select(fields)
	switch params.Status {
	case entity.ShiftStatusWaiting:
		stmt.Where("open_at > ?", now)
	case entity.ShiftStatusAccepting:
		now := s.now()
		stmt.Where("open_at <= ?", now).Where("end_at >= ?", now)
	case entity.ShiftStatusFinished:
		stmt.Where("end_at < ?", now)
	}

	switch params.OrderBy {
	case OrderByAsc:
		stmt.Order("`year_month` ASC")
	case OrderByDesc:
		stmt.Order("`year_month` DESC")
	}

	if params.Limit > 0 {
		stmt.Limit(params.Limit)
	}
	if params.Offset > 0 {
		stmt.Offset(params.Offset)
	}

	err := stmt.Find(&summaries).Error
	if err != nil {
		return nil, dbError(err)
	}
	summaries.Fill(now)
	return summaries, nil
}

func (s *shiftSummary) Get(ctx context.Context, id int64, fields ...string) (*entity.ShiftSummary, error) {
	var summary *entity.ShiftSummary
	if len(fields) == 0 {
		fields = shiftSummaryFields
	}

	stmt := s.db.DB.Table(shiftSummaryTable).Select(fields).
		Where("id = ?", id)

	err := stmt.First(&summary).Error
	if err != nil {
		return nil, dbError(err)
	}
	summary.Fill(s.now())
	return summary, nil
}

func (s *shiftSummary) UpdateSchedule(ctx context.Context, id int64, openAt, endAt time.Time) error {
	tx, err := s.db.Begin()
	if err != nil {
		return dbError(err)
	}
	defer s.db.Close(tx)

	params := map[string]interface{}{
		"id":         id,
		"open_at":    openAt,
		"end_at":     endAt,
		"updated_at": s.now(),
	}

	err = tx.Table(shiftSummaryTable).Where("id = ?", id).Updates(params).Error
	if err != nil {
		tx.Rollback()
		return dbError(err)
	}
	return dbError(tx.Commit().Error)
}

func (s *shiftSummary) Count(ctx context.Context) (int64, error) {
	var total int64
	err := s.db.DB.Table(shiftSummaryTable).Count(&total).Error
	return total, dbError(err)
}
