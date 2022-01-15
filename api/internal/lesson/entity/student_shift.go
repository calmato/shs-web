package entity

import (
	"time"

	"github.com/calmato/shs-web/api/proto/lesson"
)

type StudentShift struct {
	StudentID      string    `gorm:"primaryKey;<-:create"`
	ShiftID        int64     `gorm:"primaryKey;<-:create"`
	ShiftSummaryID int64     `gorm:""`
	CreatedAt      time.Time `gorm:"<-:create"`
	UpdatedAt      time.Time `gorm:""`
}

type StudentShifts []*StudentShift

func NewStudentShift(studentID string, summaryID int64, shiftID int64) *StudentShift {
	return &StudentShift{
		StudentID:      studentID,
		ShiftID:        shiftID,
		ShiftSummaryID: summaryID,
	}
}

func NewStudentShifts(studentID string, summaryID int64, shiftIDs []int64) StudentShifts {
	ss := make(StudentShifts, len(shiftIDs))
	for i := range shiftIDs {
		ss[i] = NewStudentShift(studentID, summaryID, shiftIDs[i])
	}
	return ss
}

func (s *StudentShift) Proto() *lesson.StudentShift {
	return &lesson.StudentShift{
		StudentId:      s.StudentID,
		ShiftId:        s.ShiftID,
		ShiftSummaryId: s.ShiftSummaryID,
		CreatedAt:      s.CreatedAt.Unix(),
		UpdatedAt:      s.CreatedAt.Unix(),
	}
}

func (ss StudentShifts) Proto() []*lesson.StudentShift {
	shifts := make([]*lesson.StudentShift, len(ss))
	for i := range ss {
		shifts[i] = ss[i].Proto()
	}
	return shifts
}
