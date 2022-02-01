package entity

import (
	"fmt"
	"time"

	"github.com/calmato/shs-web/api/pkg/set"
	"github.com/calmato/shs-web/api/proto/lesson"
)

type Lesson struct {
	ID             int64     `gorm:"primaryKey;autoIncrement;<-:create"`
	ShiftSummaryID int64     `gorm:""`
	ShiftID        int64     `gorm:""`
	SubjectID      int64     `gorm:""`
	RoomID         int32     `gorm:""`
	TeacherID      string    `gorm:""`
	StudentID      string    `gorm:""`
	Notes          string    `gorm:""`
	CreatedAt      time.Time `gorm:"<-:create"`
	UpdatedAt      time.Time `gorm:""`
}

type Lessons []*Lesson

func (l *Lesson) Proto() *lesson.Lesson {
	return &lesson.Lesson{
		Id:             l.ID,
		ShiftSummaryId: l.ShiftSummaryID,
		ShiftId:        l.ShiftID,
		SubjectId:      l.SubjectID,
		RoomId:         l.RoomID,
		TeacherId:      l.TeacherID,
		StudentId:      l.StudentID,
		Notes:          l.Notes,
		CreatedAt:      l.CreatedAt.Unix(),
		UpdatedAt:      l.CreatedAt.Unix(),
	}
}

func (ls Lessons) ShiftSummaryIDs() []int64 {
	set := set.New(len(ls))
	for i := range ls {
		set.Add(ls[i].ShiftSummaryID)
	}
	return set.Int64s()
}

func (ls Lessons) ShiftIDs() []int64 {
	set := set.New(len(ls))
	for i := range ls {
		set.Add(ls[i].ShiftID)
	}
	return set.Int64s()
}

func (ls Lessons) Decided(summaries map[int64]*ShiftSummary) (Lessons, error) {
	res := make(Lessons, 0, len(ls))
	for i := range ls {
		summary, ok := summaries[ls[i].ShiftSummaryID]
		if !ok {
			return nil, fmt.Errorf("shift summary is not found")
		}
		if summary.Decided {
			res = append(res, ls[i])
		}
	}
	return res, nil
}

func (ls Lessons) Proto() []*lesson.Lesson {
	lessons := make([]*lesson.Lesson, len(ls))
	for i := range ls {
		lessons[i] = ls[i].Proto()
	}
	return lessons
}
