package entity

import "time"

type Lesson struct {
	ID             int64     `gorm:"primaryKey;autoIncrement;<-:create"`
	ShiftSummaryID int64     `gorm:""`
	ShiftID        int64     `gorm:""`
	SubjectID      int64     `gorm:""`
	RoomID         int       `gorm:""`
	TeacherID      string    `gorm:""`
	StudentID      string    `gorm:""`
	Notes          string    `gorm:""`
	CreatedAt      time.Time `gorm:"<-:create"`
	UpdatedAt      time.Time `gorm:""`
}

type Lessons []*Lesson
