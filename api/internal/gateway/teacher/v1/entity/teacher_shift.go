package entity

import (
	"time"

	"github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
)

type TeacherShift struct {
	ID        int64  `json:"id"`
	Enabled   bool   `json:"enabled"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type TeacherShifts []*TeacherShift

type TeacherShiftDetail struct {
	Date     string        `json:"date"`
	IsClosed bool          `json:"isClosed"`
	Lessons  TeacherShifts `json:"lessons"`
}

type TeacherShiftDetails []*TeacherShiftDetail

func NewTeacherShift(shift *entity.Shift, enabled bool) *TeacherShift {
	return &TeacherShift{
		ID:        shift.Id,
		Enabled:   enabled,
		StartTime: shift.StartTime,
		EndTime:   shift.EndTime,
	}
}

func NewTeacherShifts(shifts entity.Shifts, teacherShiftMap map[int64]*entity.TeacherShift) TeacherShifts {
	ss := make(TeacherShifts, len(shifts))
	for i, shift := range shifts {
		_, enabled := teacherShiftMap[shift.Id]
		ss[i] = NewTeacherShift(shift, enabled)
	}
	return ss
}

func NewTeacherShiftDetail(
	shifts entity.Shifts, date time.Time, isClosed bool, teacherShiftMap map[int64]*entity.TeacherShift,
) *TeacherShiftDetail {
	return &TeacherShiftDetail{
		Date:     jst.FormatYYYYMMDD(date),
		IsClosed: isClosed,
		Lessons:  NewTeacherShifts(shifts, teacherShiftMap),
	}
}

func NewTeacherShiftDetailsForMonth(
	summary *ShiftSummary, shiftsMap map[time.Time]entity.Shifts, teacherShiftMap map[int64]*entity.TeacherShift,
) TeacherShiftDetails {
	const maxDays = 31
	firstDate, finalDate := newFirstAndFinalOfMonth(summary)

	details := make(TeacherShiftDetails, 0, maxDays)
	for date := firstDate; date.Before(finalDate); date = date.AddDate(0, 0, 1) {
		shifts, isOpened := shiftsMap[date] // 取得できない == 休校
		detail := NewTeacherShiftDetail(shifts, date, !isOpened, teacherShiftMap)
		details = append(details, detail)
	}
	return details
}
