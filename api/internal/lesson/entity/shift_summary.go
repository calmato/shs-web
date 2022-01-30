package entity

import (
	"errors"
	"strconv"
	"time"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
)

var errInvalidShiftStatus = errors.New("entity: invalid shift status")

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
	Decided   bool        `gorm:""`
	Status    ShiftStatus `gorm:"-"`
	OpenAt    time.Time   `gorm:""`
	EndAt     time.Time   `gorm:""`
	CreatedAt time.Time   `gorm:"<-:create"`
	UpdatedAt time.Time   `gorm:""`
}

type ShiftSummaries []*ShiftSummary

func NewShiftSummary(yearMonth int32, openAt, endAt int64) *ShiftSummary {
	return &ShiftSummary{
		YearMonth: yearMonth,
		OpenAt:    jst.ParseFromUnix(openAt),
		EndAt:     jst.ParseFromUnix(endAt),
	}
}

func (s *ShiftSummary) Year() (int, error) {
	date, err := jst.Parse("200601", strconv.Itoa(int(s.YearMonth)))
	if err != nil {
		return 0, err
	}
	return date.Year(), nil
}

func (s *ShiftSummary) Month() (int, error) {
	date, err := jst.Parse("200601", strconv.Itoa(int(s.YearMonth)))
	if err != nil {
		return 0, err
	}
	return int(date.Month()), nil
}

func (s *ShiftSummary) EnabledSubmit() bool {
	return s.Status == ShiftStatusAccepting
}

func (s *ShiftSummary) Fill(now time.Time) {
	switch {
	case now.Before(s.OpenAt): // 開始前: now < OpenAt
		s.Status = ShiftStatusWaiting
	case now.After(s.EndAt): // 締切後: CloseAt < now
		s.Status = ShiftStatusFinished
	default: // 募集中: OpenAt <= now && now <= CloseAt
		s.Status = ShiftStatusAccepting
	}
}

func (s *ShiftSummary) Proto() *lesson.ShiftSummary {
	return &lesson.ShiftSummary{
		Id:        s.ID,
		YearMonth: s.YearMonth,
		Decided:   s.Decided,
		Status:    lesson.ShiftStatus(s.Status),
		OpenAt:    s.OpenAt.Unix(),
		EndAt:     s.EndAt.Unix(),
		CreatedAt: s.CreatedAt.Unix(),
		UpdatedAt: s.UpdatedAt.Unix(),
	}
}

func (ss ShiftSummaries) Map() map[int64]*ShiftSummary {
	res := make(map[int64]*ShiftSummary, len(ss))
	for _, s := range ss {
		res[s.ID] = s
	}
	return res
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

func NewShiftStatus(status lesson.ShiftStatus) (ShiftStatus, error) {
	switch status {
	case lesson.ShiftStatus_SHIFT_STATUS_WAITING:
		return ShiftStatusWaiting, nil
	case lesson.ShiftStatus_SHIFT_STATUS_ACCEPTING:
		return ShiftStatusAccepting, nil
	case lesson.ShiftStatus_SHIFT_STATUS_FINISHED:
		return ShiftStatusFinished, nil
	default:
		return ShiftStatusUnknown, errInvalidShiftStatus
	}
}
