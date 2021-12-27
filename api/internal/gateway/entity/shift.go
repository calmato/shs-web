package entity

import "github.com/calmato/shs-web/api/proto/lesson"

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
