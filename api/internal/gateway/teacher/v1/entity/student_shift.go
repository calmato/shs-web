package entity

import (
	"time"

	"github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
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

func NewStudentShifts(
	shifts entity.Shifts, studentShiftMap map[int64]*entity.StudentShift, onlyEnabled bool,
) StudentShifts {
	ss := make(StudentShifts, 0, len(shifts))
	for _, shift := range shifts {
		_, enabled := studentShiftMap[shift.Id]
		if onlyEnabled && !enabled {
			continue
		}
		ss = append(ss, NewStudentShift(shift, enabled))
	}
	return ss
}

func NewEnabledStudentShiftDetail(
	shifts entity.Shifts, date time.Time, studentShiftMap map[int64]*entity.StudentShift,
) *StudentShiftDetail {
	lessons := NewStudentShifts(shifts, studentShiftMap, true)
	if len(lessons) == 0 {
		return nil
	}
	return &StudentShiftDetail{
		Date:     jst.FormatYYYYMMDD(date),
		IsClosed: false,
		Lessons:  lessons,
	}
}

func NewEnabledStudentShiftDetails(
	summary *ShiftSummary, shiftsMap map[time.Time]entity.Shifts, studentShiftMap map[int64]*entity.StudentShift,
) StudentShiftDetails {
	const maxDays = 31
	firstDate, finalDate := newFirstAndFinalOfMonth(summary)

	details := make(StudentShiftDetails, 0, maxDays)
	for date := firstDate; date.Before(finalDate); date = date.AddDate(0, 0, 1) {
		shifts, isOpened := shiftsMap[date] // 取得できない == 休校
		if !isOpened {
			continue
		}
		detail := NewEnabledStudentShiftDetail(shifts, date, studentShiftMap)
		if detail == nil {
			continue
		}
		details = append(details, detail)
	}
	return details
}
