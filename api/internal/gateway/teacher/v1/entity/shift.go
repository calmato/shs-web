package entity

import (
	"time"

	"github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
)

type Shift struct {
	ID        int64  `json:"id"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type Shifts []*Shift

type ShiftDetail struct {
	Date     string `json:"date"`
	IsClosed bool   `json:"isClosed"`
	Lessons  Shifts `json:"lessons"`
}

type ShiftDetails []*ShiftDetail

func NewShift(shift *entity.Shift) *Shift {
	return &Shift{
		ID:        shift.Id,
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

func NewShiftDetail(shifts entity.Shifts, date time.Time, isClosed bool) *ShiftDetail {
	return &ShiftDetail{
		Date:     jst.FormatYYYYMMDD(date),
		IsClosed: isClosed,
		Lessons:  NewShifts(shifts),
	}
}

func NewShiftDetailsForMonth(summary *ShiftSummary, shiftsMap map[time.Time]entity.Shifts) ShiftDetails {
	const maxDays = 31
	firstDate, finalDate := newFirstAndFinalOfMonth(summary)

	details := make(ShiftDetails, 0, maxDays)
	for date := firstDate; date.Before(finalDate); date = date.AddDate(0, 0, 1) {
		shifts, isOpened := shiftsMap[date] // 取得できない == 休校
		detail := NewShiftDetail(shifts, date, !isOpened)
		details = append(details, detail)
	}
	return details
}

func newFirstAndFinalOfMonth(summary *ShiftSummary) (time.Time, time.Time) {
	firstDate := jst.BeginningOfMonth(int(summary.Year), int(summary.Month))
	finalDate := firstDate.AddDate(0, 1, 0)
	return firstDate, finalDate
}
