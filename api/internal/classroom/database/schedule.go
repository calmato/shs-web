package database

import (
	"context"
	"fmt"
	"time"

	"github.com/calmato/shs-web/api/internal/classroom/entity"
	"github.com/calmato/shs-web/api/pkg/database"
	"github.com/calmato/shs-web/api/pkg/jst"
)

const scheduleTable = "schedules"

var scheduleFields = []string{
	"weekday", "is_closed", "lessons", "created_at", "updated_at",
}

type schedule struct {
	db  *database.Client
	now func() time.Time
}

func NewSchedule(db *database.Client) Schedule {
	return &schedule{
		db:  db,
		now: jst.Now,
	}
}

func (s *schedule) List(ctx context.Context, fields ...string) (entity.Schedules, error) {
	var schedules entity.Schedules
	if len(schedules) == 0 {
		fields = scheduleFields
	}

	stmt := s.db.DB.Table(scheduleTable).Select(fields)

	err := stmt.Find(&schedules).Error
	if err != nil {
		return nil, dbError(err)
	}
	err = schedules.Fill()
	if err != nil {
		return nil, dbError(err)
	}
	return schedules, nil
}

func (s *schedule) Get(ctx context.Context, weekday time.Weekday, fields ...string) (*entity.Schedule, error) {
	var schedule *entity.Schedule
	if len(fields) == 0 {
		fields = scheduleFields
	}

	stmt := s.db.DB.Table(scheduleTable).Select(fields).
		Where("weekday = ?", weekday)

	err := stmt.First(&schedule).Error
	if err != nil {
		return nil, dbError(err)
	}
	err = schedule.Fill()
	if err != nil {
		return nil, dbError(err)
	}
	return schedule, nil
}

func (s *schedule) MultipleUpdate(ctx context.Context, schedules entity.Schedules) error {
	now := s.now()
	err := schedules.FillJSON()
	if err != nil {
		return fmt.Errorf("database: failed to fill json: %w", ErrInvalidArgument)
	}

	tx, err := s.db.Begin()
	if err != nil {
		return dbError(err)
	}
	defer s.db.Close(tx)

	var currentSchedules entity.Schedules
	err = tx.Table(scheduleTable).Select([]string{"weekday", "created_at"}).
		Where("weekday IN (?)", schedules.Weekdays()).
		Find(&currentSchedules).Error
	if err != nil {
		return dbError(err)
	}
	scheduleMap := currentSchedules.Map()

	for _, schedule := range schedules {
		current, ok := scheduleMap[schedule.Weekday]
		if !ok {
			return fmt.Errorf("database: schedule not found: %w", ErrNotFound)
		}
		schedule.CreatedAt = current.CreatedAt
		schedule.UpdatedAt = now

		err = tx.Table(scheduleTable).Save(&schedule).Error
		if err != nil {
			tx.Rollback()
			return dbError(err)
		}
	}

	return dbError(tx.Commit().Error)
}
