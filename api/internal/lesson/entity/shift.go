package entity

import (
	"time"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
)

type Shift struct {
	ID             int64     `gorm:"primaryKey;autoIncrement;<-:create"`
	ShiftSummaryID int64     `gorm:""`
	Date           time.Time `gorm:""`
	StartTime      string    `gorm:""`
	EndTime        string    `gorm:""`
	CreatedAt      time.Time `gorm:"<-:create"`
	UpdatedAt      time.Time `gorm:""`
}

type Shifts []*Shift

func (s *Shift) Proto() *lesson.Shift {
	return &lesson.Shift{
		Id:             s.ID,
		ShiftSummaryId: s.ShiftSummaryID,
		Date:           jst.FormatYYYYMMDD(s.Date),
		StartTime:      s.StartTime,
		EndTime:        s.EndTime,
		CreatedAt:      s.CreatedAt.Unix(),
		UpdatedAt:      s.UpdatedAt.Unix(),
	}
}

func (ss Shifts) Proto() []*lesson.Shift {
	shifts := make([]*lesson.Shift, len(ss))
	for i := range ss {
		shifts[i] = ss[i].Proto()
	}
	return shifts
}
