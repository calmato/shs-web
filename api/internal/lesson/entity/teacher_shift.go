package entity

import (
	"time"

	"github.com/calmato/shs-web/api/proto/lesson"
)

type TeacherShift struct {
	TeacherID      string    `gorm:"primaryKey;<-:create"`
	ShiftID        int64     `gorm:"primaryKey;<-:create"`
	ShiftSummaryID int64     `gorm:""`
	CreatedAt      time.Time `gorm:"<-:create"`
	UpdatedAt      time.Time `gorm:""`
}

type TeacherShifts []*TeacherShift

func NewTeacherShift(teacherID string, summaryID int64, shiftID int64) *TeacherShift {
	return &TeacherShift{
		TeacherID:      teacherID,
		ShiftID:        shiftID,
		ShiftSummaryID: summaryID,
	}
}

func NewTeacherShifts(teacherID string, summaryID int64, shiftIDs []int64) TeacherShifts {
	ss := make(TeacherShifts, len(shiftIDs))
	for i := range shiftIDs {
		ss[i] = NewTeacherShift(teacherID, summaryID, shiftIDs[i])
	}
	return ss
}

func (s *TeacherShift) Proto() *lesson.TeacherShift {
	return &lesson.TeacherShift{
		TeacherId:      s.TeacherID,
		ShiftId:        s.ShiftID,
		ShiftSummaryId: s.ShiftSummaryID,
		CreatedAt:      s.CreatedAt.Unix(),
		UpdatedAt:      s.CreatedAt.Unix(),
	}
}

func (ss TeacherShifts) Proto() []*lesson.TeacherShift {
	shifts := make([]*lesson.TeacherShift, len(ss))
	for i := range ss {
		shifts[i] = ss[i].Proto()
	}
	return shifts
}
