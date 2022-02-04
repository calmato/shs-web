package entity

import (
	"time"

	"github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/pkg/set"
)

type StudentShift struct {
	ID        int64  `json:"id"`
	Enabled   bool   `json:"enabled"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type StudentShifts []*StudentShift

type StudentShiftDetail struct {
	Date     string        `json:"date"`
	IsClosed bool          `json:"isClosed"`
	Lessons  StudentShifts `json:"lessons"`
}

type StudentShiftDetails []*StudentShiftDetail

func NewStudentShift(shift *entity.Shift, enabled bool) *StudentShift {
	return &StudentShift{
		ID:        shift.Id,
		Enabled:   enabled,
		StartTime: shift.StartTime,
		EndTime:   shift.EndTime,
	}
}

func NewStudentShifts(shifts entity.Shifts, studentShiftMap map[int64]*entity.StudentShift) StudentShifts {
	ss := make(StudentShifts, 0, len(shifts))
	for _, shift := range shifts {
		_, enabled := studentShiftMap[shift.Id]
		ss = append(ss, NewStudentShift(shift, enabled))
	}
	return ss
}

func NewStudentShiftsWithTemplate(shifts entity.Shifts, weekday time.Weekday, schedules *set.Set) StudentShifts {
	ss := make(StudentShifts, 0, len(shifts))
	for _, shift := range shifts {
		key := entity.LessonKey(weekday, shift.StartTime, shift.EndTime)
		enabled := schedules.Contains(key)
		ss = append(ss, NewStudentShift(shift, enabled))
	}
	return ss
}

func NewStudentShiftDetailForMonth(
	shifts entity.Shifts, date time.Time, isClosed bool, studentShiftMap map[int64]*entity.StudentShift,
) *StudentShiftDetail {
	return &StudentShiftDetail{
		Date:     jst.FormatYYYYMMDD(date),
		IsClosed: isClosed,
		Lessons:  NewStudentShifts(shifts, studentShiftMap),
	}
}

func NewStudentShiftDetailWithTemplate(
	shifts entity.Shifts, date time.Time, isClosed bool, schedules *set.Set,
) *StudentShiftDetail {
	return &StudentShiftDetail{
		Date:     jst.FormatYYYYMMDD(date),
		IsClosed: isClosed,
		Lessons:  NewStudentShiftsWithTemplate(shifts, date.Weekday(), schedules),
	}
}

func NewStudentShiftDetailsForMonth(
	summary *ShiftSummary, shiftsMap map[time.Time]entity.Shifts, studentShiftMap map[int64]*entity.StudentShift,
) StudentShiftDetails {
	const maxDays = 31
	firstDate, finalDate := newFirstAndFinalOfMonth(summary)

	details := make(StudentShiftDetails, 0, maxDays)
	for date := firstDate; date.Before(finalDate); date = date.AddDate(0, 0, 1) {
		shifts, isOpened := shiftsMap[date] // 取得できない == 休校
		detail := NewStudentShiftDetailForMonth(shifts, date, !isOpened, studentShiftMap)
		details = append(details, detail)
	}
	return details
}

func NewStudentShiftDetailsWithTemplate(
	summary *ShiftSummary, shiftsMap map[time.Time]entity.Shifts, schedules entity.StudentShiftSchedules,
) StudentShiftDetails {
	const maxDays = 31
	firstDate, finalDate := newFirstAndFinalOfMonth(summary)
	keys := schedules.LessonKeys()

	details := make(StudentShiftDetails, 0, maxDays)
	for date := firstDate; date.Before(finalDate); date = date.AddDate(0, 0, 1) {
		shifts, isOpened := shiftsMap[date] // 取得できない == 休校
		detail := NewStudentShiftDetailWithTemplate(shifts, date, !isOpened, keys)
		details = append(details, detail)
	}
	return details
}

func newFirstAndFinalOfMonth(summary *ShiftSummary) (time.Time, time.Time) {
	firstDate := jst.BeginningOfMonth(int(summary.Year), int(summary.Month))
	finalDate := firstDate.AddDate(0, 1, 0)
	return firstDate, finalDate
}
