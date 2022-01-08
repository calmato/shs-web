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

type ScheduleToUpdate struct {
	Weekday  time.Weekday    `json:"weekday,omitempty"`
	IsClosed bool            `json:"isClosed,omitempty"`
	Lessons  ScheduleLessons `json:"lessons,omitempty"`
}

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

func NewScheduleToUpdate(req *ScheduleToUpdate) *classroom.ScheduleToUpdate {
	return &classroom.ScheduleToUpdate{
		Weekday:  int32(req.Weekday),
		IsClosed: req.IsClosed,
		Lessons:  newScheduleLessonsToUpdate(req.Lessons, req.IsClosed),
	}
}

func NewSchedulesToUpdate(req []*ScheduleToUpdate) []*classroom.ScheduleToUpdate {
	ss := make([]*classroom.ScheduleToUpdate, len(req))
	for i := range req {
		ss[i] = NewScheduleToUpdate(req[i])
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

func newScheduleLessonToUpdate(lesson *ScheduleLesson) *classroom.ScheduleToUpdate_Lesson {
	return &classroom.ScheduleToUpdate_Lesson{
		StartTime: lesson.StartTime,
		EndTime:   lesson.EndTime,
	}
}

func newScheduleLessonsToUpdate(lessons ScheduleLessons, isClosed bool) []*classroom.ScheduleToUpdate_Lesson {
	if isClosed {
		return nil
	}
	ls := make([]*classroom.ScheduleToUpdate_Lesson, len(lessons))
	for i := range lessons {
		ls[i] = newScheduleLessonToUpdate(lessons[i])
	}
	return ls
}
