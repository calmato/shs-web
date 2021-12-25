package entity

import "github.com/calmato/shs-web/api/proto/classroom"

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
