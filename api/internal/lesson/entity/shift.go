package entity

import (
	"fmt"
	"sync"
	"time"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/pkg/set"
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

func NewShift(date time.Time, startTime, endTime string) *Shift {
	return &Shift{
		Date:      date,
		StartTime: startTime,
		EndTime:   endTime,
	}
}

func NewShifts(summary *ShiftSummary, schedules map[time.Weekday]*Schedule, closedDates []string) (Shifts, error) {
	const maxShifts = 155 // 1ヶ月31日 * 1日コマ数5
	firstDate, finalDate, err := newFirstAndFinalOfMonth(summary)
	if err != nil {
		return nil, err
	}
	set := set.New(len(closedDates))
	set.AddStrings(closedDates...)

	var m sync.Mutex
	var wg sync.WaitGroup
	shifts := make(Shifts, 0, maxShifts)
	for date := firstDate; date.Before(finalDate); date = date.AddDate(0, 0, 1) {
		date := date
		wg.Add(1)
		go func() {
			defer wg.Done()
			target := jst.FormatYYYYMMDD(date)
			if set.Contains(target) {
				return
			}
			schedule, ok := schedules[date.Weekday()]
			fmt.Println(">>>>debug: ", target, date.Weekday(), schedule)
			if !ok || schedule.IsClosed {
				return
			}
			for _, lesson := range schedule.Lessons {
				shift := NewShift(date, lesson.StartTime, lesson.EndTime)
				m.Lock()
				shifts = append(shifts, shift)
				m.Unlock()
			}
		}()
	}
	wg.Wait()

	return shifts, nil
}

func newFirstAndFinalOfMonth(summary *ShiftSummary) (time.Time, time.Time, error) {
	year, err := summary.Year()
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	month, err := summary.Month()
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	firstDate := jst.BeginningOfMonth(year, month)
	finalDate := firstDate.AddDate(0, 1, 0)
	return firstDate, finalDate, nil
}

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

func (ss Shifts) IDs() []int64 {
	res := make([]int64, len(ss))
	for i := range ss {
		res[i] = ss[i].ID
	}
	return res
}

func (ss Shifts) GroupByShiftSummaryID() map[int64]Shifts {
	shifts := make(map[int64]Shifts)
	for _, s := range ss {
		if _, ok := shifts[s.ShiftSummaryID]; !ok {
			shifts[s.ShiftSummaryID] = make(Shifts, 0, len(ss))
		}
		shifts[s.ShiftSummaryID] = append(shifts[s.ShiftSummaryID], s)
	}
	return shifts
}

func (ss Shifts) Proto() []*lesson.Shift {
	shifts := make([]*lesson.Shift, len(ss))
	for i := range ss {
		shifts[i] = ss[i].Proto()
	}
	return shifts
}
