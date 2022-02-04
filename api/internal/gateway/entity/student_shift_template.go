package entity

import (
	"fmt"
	"time"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/pkg/set"
	"github.com/calmato/shs-web/api/proto/lesson"
)

type StudentShiftTemplate struct {
	StudentID string
	Schedules StudentShiftSchedules
	Lessons   StudentSuggestedLessons
	CreatedAt time.Time
	UpdatedAt time.Time
}

type StudentShiftSchedule struct {
	*lesson.ShiftSchedule
}

type StudentShiftSchedules []*StudentShiftSchedule

type StudentSuggestedLesson struct {
	*lesson.SuggestedLesson
}

type StudentSuggestedLessons []*StudentSuggestedLesson

func LessonKey(weekday time.Weekday, startTime, endTime string) string {
	return fmt.Sprintf("%d:%s:%s", weekday, startTime, endTime)
}

func NewStudentShiftTemplate(template *lesson.StudentShiftTemplate) *StudentShiftTemplate {
	return &StudentShiftTemplate{
		StudentID: template.StudentId,
		Schedules: NewStudentShiftSchedules(template.Schedules),
		Lessons:   NewStudentSuggestedLessons(template.SuggestedLessons),
		CreatedAt: jst.ParseFromUnix(template.CreatedAt),
		UpdatedAt: jst.ParseFromUnix(template.UpdatedAt),
	}
}

func NewStudentShiftSchedule(schedule *lesson.ShiftSchedule) *StudentShiftSchedule {
	return &StudentShiftSchedule{
		ShiftSchedule: schedule,
	}
}

func NewStudentShiftSchedules(schedules []*lesson.ShiftSchedule) StudentShiftSchedules {
	res := make(StudentShiftSchedules, len(schedules))
	for i := range schedules {
		res[i] = NewStudentShiftSchedule(schedules[i])
	}
	return res
}

func NewStudentSuggestedLesson(sl *lesson.SuggestedLesson) *StudentSuggestedLesson {
	return &StudentSuggestedLesson{
		SuggestedLesson: sl,
	}
}

func NewStudentSuggestedLessons(sls []*lesson.SuggestedLesson) StudentSuggestedLessons {
	res := make(StudentSuggestedLessons, len(sls))
	for i := range sls {
		res[i] = NewStudentSuggestedLesson(sls[i])
	}
	return res
}

func (ss StudentShiftSchedules) LessonKeys() *set.Set {
	const maxLesson = 70 // 想定最大授業数: 10/day -> 10 * 7day = 70
	set := set.New(maxLesson)
	for i := range ss {
		for _, schedule := range ss[i].Lessons {
			key := LessonKey(time.Weekday(ss[i].Weekday), schedule.StartTime, schedule.EndTime)
			set.AddStrings(key)
		}
	}
	return set
}
