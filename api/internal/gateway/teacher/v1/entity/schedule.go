package entity

import (
	"time"

	"github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/proto/classroom"
)

type Schedule struct {
	Weekday  time.Weekday    `json:"weekday"`
	IsClosed bool            `json:"isClosed"`
	Lessons  ScheduleLessons `json:"lessons"`
}

type Schedules []*Schedule

type ScheduleLesson struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type ScheduleLessons []*ScheduleLesson

func NewSchedule(schedule *entity.Schedule) *Schedule {
	return &Schedule{
		Weekday:  time.Weekday(schedule.Weekday),
		IsClosed: schedule.IsClosed,
		Lessons:  newScheduleLessons(schedule.Lessons),
	}
}

func NewSchedules(schedules entity.Schedules) Schedules {
	ss := make(Schedules, len(schedules))
	for i := range schedules {
		ss[i] = NewSchedule(schedules[i])
	}
	return ss
}

func newScheduleLesson(lesson *classroom.Schedule_Lesson) *ScheduleLesson {
	return &ScheduleLesson{
		StartTime: lesson.StartTime,
		EndTime:   lesson.EndTime,
	}
}

func newScheduleLessons(lessons []*classroom.Schedule_Lesson) ScheduleLessons {
	ls := make(ScheduleLessons, len(lessons))
	for i := range lessons {
		ls[i] = newScheduleLesson(lessons[i])
	}
	return ls
}
