package entity

import (
	"time"

	"github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/proto/classroom"
)

var (
	// TODO: コマ設定の変更機能はまだ不要とのことなので固定で
	weekdayLessons = ScheduleLessons{
		{StartTime: "1700", EndTime: "1830"},
		{StartTime: "1830", EndTime: "2000"},
		{StartTime: "2000", EndTime: "2130"},
	}
	holidayLessons = ScheduleLessons{
		{StartTime: "1530", EndTime: "1700"},
		{StartTime: "1700", EndTime: "1830"},
		{StartTime: "1830", EndTime: "2000"},
		{StartTime: "2000", EndTime: "2130"},
	}
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
	Weekday  time.Weekday `json:"weekday,omitempty"`
	IsClosed bool         `json:"isClosed,omitempty"`
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
	s := &classroom.ScheduleToUpdate{
		Weekday:  int32(req.Weekday),
		IsClosed: req.IsClosed,
	}
	if s.IsClosed {
		return s
	}
	if isHoliday(req.Weekday) {
		s.Lessons = newScheduleLessonsToUpdate(holidayLessons)
	} else {
		s.Lessons = newScheduleLessonsToUpdate(weekdayLessons)
	}
	return s
}

func isHoliday(weekday time.Weekday) bool {
	return weekday == time.Sunday || weekday == time.Saturday
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

func newScheduleLessonsToUpdate(lessons ScheduleLessons) []*classroom.ScheduleToUpdate_Lesson {
	ls := make([]*classroom.ScheduleToUpdate_Lesson, len(lessons))
	for i := range lessons {
		ls[i] = newScheduleLessonToUpdate(lessons[i])
	}
	return ls
}
