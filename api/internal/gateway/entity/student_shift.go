package entity

import (
	"github.com/calmato/shs-web/api/pkg/set"
	"github.com/calmato/shs-web/api/proto/lesson"
)

type StudentShift struct {
	*lesson.StudentShift
}

type StudentShifts []*StudentShift

func NewStudentShift(shift *lesson.StudentShift) *StudentShift {
	return &StudentShift{
		StudentShift: shift,
	}
}

func NewStudentShifts(shifts []*lesson.StudentShift) StudentShifts {
	ss := make(StudentShifts, len(shifts))
	for i := range shifts {
		ss[i] = NewStudentShift(shifts[i])
	}
	return ss
}

func (ss StudentShifts) StudentIDs() []string {
	set := set.New(len(ss))
	for i := range ss {
		set.AddStrings(ss[i].StudentId)
	}
	return set.Strings()
}

func (ss StudentShifts) MapByShiftID() map[int64]*StudentShift {
	res := make(map[int64]*StudentShift, len(ss))
	for _, s := range ss {
		res[s.ShiftId] = s
	}
	return res
}

func (ss StudentShifts) GroupByStudentID() map[string]StudentShifts {
	res := make(map[string]StudentShifts)
	for _, s := range ss {
		if _, ok := res[s.StudentId]; !ok {
			res[s.StudentId] = make(StudentShifts, 0)
		}
		res[s.StudentId] = append(res[s.StudentId], s)
	}
	return res
}
