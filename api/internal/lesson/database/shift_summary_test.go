package database

import (
	"time"

	"github.com/calmato/shs-web/api/internal/lesson/entity"
)

func testShiftSummary(id int64, yearMonth int32, openAt, endAt, now time.Time) *entity.ShiftSummary {
	return &entity.ShiftSummary{
		ID:        id,
		YearMonth: yearMonth,
		OpenAt:    openAt,
		EndAt:     endAt,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
