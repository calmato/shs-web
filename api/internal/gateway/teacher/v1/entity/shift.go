package entity

import (
	"time"

	"github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
)

type Shift struct {
	ID        int64  `json:"id"`
	Date      string `json:"date"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type Shifts []*Shift

type ShiftDetail struct {
	IsClosed bool                 `json:"isClosed"`
	Lessons  []*ShiftDetailLesson `json:"lessons"`
}

type ShiftDetailLesson struct {
	ID        int64  `json:"id"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

func NewShift(shift *entity.Shift) *Shift {
	return &Shift{
		ID:        shift.Id,
		Date:      shift.Date,
		StartTime: shift.StartTime,
		EndTime:   shift.EndTime,
	}
}

func NewShifts(shifts entity.Shifts) Shifts {
	ss := make(Shifts, len(shifts))
	for i := range shifts {
		ss[i] = NewShift(shifts[i])
	}
	return ss
}

func NewShiftDetailsForMonth(summary *ShiftSummary, shiftsMap map[time.Time]Shifts) map[string]*ShiftDetail {
	const maxDays = 31
	firstDate, finalDate := newFirstAndFinalOfMonth(summary)

	details := make(map[string]*ShiftDetail, maxDays)
	for date := firstDate; date.Before(finalDate); date = date.AddDate(0, 0, 1) {
		day := jst.FormatYYYYMMDD(date)
		shifts, ok := shiftsMap[date] // 取得できない == 休校
		details[day] = newShiftDetail(shifts, !ok)
	}
	return details
}

func newFirstAndFinalOfMonth(summary *ShiftSummary) (time.Time, time.Time) {
	firstDate := jst.BeginningOfMonth(int(summary.Year), int(summary.Month))
	finalDate := firstDate.AddDate(0, 1, 0)
	return firstDate, finalDate
}

func newShiftDetail(shifts Shifts, isClosed bool) *ShiftDetail {
	lessons := make([]*ShiftDetailLesson, len(shifts))
	for i := range shifts {
		lesson := &ShiftDetailLesson{
			ID:        shifts[i].ID,
			StartTime: shifts[i].StartTime,
			EndTime:   shifts[i].EndTime,
		}
		lessons[i] = lesson
	}
	return &ShiftDetail{
		IsClosed: isClosed,
		Lessons:  lessons,
	}
}

func (ss Shifts) GroupByDate() (map[time.Time]Shifts, error) {
	const maxDays = 31
	const maxLessons = 5
	res := make(map[time.Time]Shifts, maxDays)
	for _, s := range ss {
		date, err := jst.ParseFromYYYYMMDD(s.Date)
		if err != nil {
			return nil, err
		}
		if _, ok := res[date]; !ok {
			res[date] = make(Shifts, 0, maxLessons)
		}
		res[date] = append(res[date], s)
	}
	return res, nil
}
