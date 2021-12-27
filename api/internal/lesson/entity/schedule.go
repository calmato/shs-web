package entity

import (
	"time"

	"github.com/calmato/shs-web/api/proto/classroom"
)

type Schedule struct {
	*classroom.Schedule
}

type Schedules []*Schedule

func NewSchedule(schedule *classroom.Schedule) *Schedule {
	return &Schedule{
		Schedule: schedule,
	}
}

func NewSchedules(schedules []*classroom.Schedule) Schedules {
	ss := make(Schedules, len(schedules))
	for i := range schedules {
		ss[i] = NewSchedule(schedules[i])
	}
	return ss
}

func (ss Schedules) Map() map[time.Weekday]*Schedule {
	res := make(map[time.Weekday]*Schedule, len(ss))
	for _, s := range ss {
		res[time.Weekday(s.Weekday)] = s
	}
	return res
}
