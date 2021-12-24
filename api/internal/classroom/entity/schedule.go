package entity

import (
	"encoding/json"
	"time"

	"github.com/calmato/shs-web/api/proto/classroom"
	"gorm.io/datatypes"
)

type Schedule struct {
	Weekday     time.Weekday   `gorm:"primaryKey;<-:create"`
	IsClosed    bool           `gorm:""`
	Lessons     Lessons        `gorm:"-"`
	LessonsJSON datatypes.JSON `gorm:"column:lessons"`
	CreatedAt   time.Time      `gorm:"<-:create"`
	UpdatedAt   time.Time      `gorm:""`
}

type Schedules []*Schedule

type Lesson struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type Lessons []*Lesson

func (s *Schedule) Fill() error {
	var lessons Lessons
	if err := json.Unmarshal(s.LessonsJSON, &lessons); err != nil {
		return err
	}
	s.Lessons = lessons
	return nil
}

func (s *Schedule) FillJSON() error {
	v, err := json.Marshal(s.Lessons)
	if err != nil {
		return err
	}
	s.LessonsJSON = datatypes.JSON(v)
	return nil
}

func (s *Schedule) Proto() *classroom.Schedule {
	return &classroom.Schedule{
		Weekday:   int32(s.Weekday),
		IsClosed:  s.IsClosed,
		Lessons:   s.Lessons.proto(),
		CreatedAt: s.CreatedAt.Unix(),
		UpdatedAt: s.UpdatedAt.Unix(),
	}
}

func (ss Schedules) Fill() error {
	for i := range ss {
		if err := ss[i].Fill(); err != nil {
			return err
		}
	}
	return nil
}

func (ss Schedules) Proto() []*classroom.Schedule {
	schedules := make([]*classroom.Schedule, len(ss))
	for i := range ss {
		schedules[i] = ss[i].Proto()
	}
	return schedules
}

func (l *Lesson) proto() *classroom.Schedule_Lesson {
	return &classroom.Schedule_Lesson{
		StartTime: l.StartTime,
		EndTime:   l.EndTime,
	}
}

func (ls Lessons) proto() []*classroom.Schedule_Lesson {
	lessons := make([]*classroom.Schedule_Lesson, len(ls))
	for i := range ls {
		lessons[i] = ls[i].proto()
	}
	return lessons
}
