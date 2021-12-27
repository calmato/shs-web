package entity

import "github.com/calmato/shs-web/api/internal/gateway/entity"

type Shift struct {
	ID        int64  `json:"id"`
	Date      string `json:"date"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type Shifts []*Shift

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

func (ss Shifts) GroupByDate() map[string]Shifts {
	const maxDays = 31
	const maxLessons = 5
	res := make(map[string]Shifts, maxDays)
	for _, s := range ss {
		if _, ok := res[s.Date]; !ok {
			res[s.Date] = make(Shifts, 0, maxLessons)
		}
		res[s.Date] = append(res[s.Date], s)
	}
	return res
}
