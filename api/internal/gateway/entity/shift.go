package entity

import (
	"time"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
)

type Shift struct {
	*lesson.Shift
}

type Shifts []*Shift

func NewShift(shift *lesson.Shift) *Shift {
	return &Shift{
		Shift: shift,
	}
}

func NewShifts(shifts []*lesson.Shift) Shifts {
	ss := make(Shifts, len(shifts))
	for i := range shifts {
		ss[i] = NewShift(shifts[i])
	}
	return ss
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
