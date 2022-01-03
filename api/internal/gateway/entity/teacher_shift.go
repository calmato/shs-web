package entity

import "github.com/calmato/shs-web/api/proto/lesson"

type TeacherShift struct {
	*lesson.TeacherShift
}

type TeacherShifts []*TeacherShift

func NewTeacherShift(shift *lesson.TeacherShift) *TeacherShift {
	return &TeacherShift{
		TeacherShift: shift,
	}
}

func NewTeacherShifts(shifts []*lesson.TeacherShift) TeacherShifts {
	ss := make(TeacherShifts, len(shifts))
	for i := range shifts {
		ss[i] = NewTeacherShift(shifts[i])
	}
	return ss
}

func (ss TeacherShifts) MapByShiftID() map[int64]*TeacherShift {
	res := make(map[int64]*TeacherShift, len(ss))
	for _, s := range ss {
		res[s.ShiftId] = s
	}
	return res
}
