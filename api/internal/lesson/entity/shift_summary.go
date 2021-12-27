package entity

import (
	"time"

	"github.com/calmato/shs-web/api/proto/lesson"
)

type ShiftStatus int32

const (
	ShiftStatusUnknown   ShiftStatus = 0
	ShiftStatusWaiting   ShiftStatus = 1
	ShiftStatusAccepting ShiftStatus = 2
	ShiftStatusFinished  ShiftStatus = 3
)

type ShiftSummary struct {
	ID        int64       `gorm:"primaryKey;autoIncrement;<-:create"`
	YearMonth int32       `gorm:""`
	Status    ShiftStatus `gorm:"-"`
	OpenAt    time.Time   `gorm:""`
	EndAt     time.Time   `gorm:""`
	CreatedAt time.Time   `gorm:"<-:create"`
	UpdatedAt time.Time   `gorm:""`
}

type ShiftSummaries []*ShiftSummary

func (s *ShiftSummary) Fill(now time.Time) {
	if now.Before(s.OpenAt) { // 開始前: now < OpenAt
		s.Status = ShiftStatusWaiting
	} else if now.After(s.EndAt) { // 締切後: CloseAt < now
		s.Status = ShiftStatusFinished
	} else {
		s.Status = ShiftStatusAccepting
	}
}

func (s *ShiftSummary) Proto() *lesson.ShiftSummary {
	return &lesson.ShiftSummary{
		Id:        s.ID,
		YearMonth: s.YearMonth,
		Status:    lesson.ShiftStatus(s.Status),
		OpenAt:    s.OpenAt.Unix(),
		EndAt:     s.EndAt.Unix(),
		CreatedAt: s.CreatedAt.Unix(),
		UpdatedAt: s.UpdatedAt.Unix(),
	}
}

func (ss ShiftSummaries) Fill(now time.Time) {
	for i := range ss {
		ss[i].Fill(now)
	}
}

func (ss ShiftSummaries) Proto() []*lesson.ShiftSummary {
	summaries := make([]*lesson.ShiftSummary, len(ss))
	for i := range ss {
		summaries[i] = ss[i].Proto()
	}
	return summaries
}
