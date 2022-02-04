package entity

import (
	"time"

	"github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/pkg/set"
	"github.com/calmato/shs-web/api/proto/classroom"
)

type StudentSchedule struct {
	Weekday time.Weekday           `json:"weekday"`
	Lessons StudentScheduleLessons `json:"lessons"`
}

type StudentSchedules []*StudentSchedule

type StudentScheduleLesson struct {
	Enabled   bool   `json:"enabled"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type StudentScheduleLessons []*StudentScheduleLesson

type StudentSuggestedLesson struct {
	SubjectID int64 `json:"subjectId"` // 授業ID
	Total     int64 `json:"total"`     // 希望授業数
}

type StudentSuggestedLessons []*StudentSuggestedLesson

func NewStudentSchedule(schedule *entity.Schedule, keys *set.Set) *StudentSchedule {
	weekday := time.Weekday(schedule.Weekday)
	return &StudentSchedule{
		Weekday: weekday,
		Lessons: NewStudentScheduleLessons(weekday, schedule.Lessons, keys),
	}
}

func NewStudentSchedules(schedules entity.Schedules, studentSchedules entity.StudentShiftSchedules) StudentSchedules {
	keys := studentSchedules.LessonKeys()
	res := make(StudentSchedules, len(schedules))
	for i := range schedules {
		res[i] = NewStudentSchedule(schedules[i], keys)
	}
	return res
}

func NewStudentScheduleLesson(lesson *classroom.Schedule_Lesson, enabled bool) *StudentScheduleLesson {
	return &StudentScheduleLesson{
		Enabled:   enabled,
		StartTime: lesson.StartTime,
		EndTime:   lesson.EndTime,
	}
}

func NewStudentScheduleLessons(weekday time.Weekday, lessons []*classroom.Schedule_Lesson, keys *set.Set) StudentScheduleLessons {
	res := make(StudentScheduleLessons, len(lessons))
	for i := range lessons {
		key := entity.LessonKey(weekday, lessons[i].StartTime, lessons[i].EndTime)
		contain := keys.Contains(key)
		res[i] = NewStudentScheduleLesson(lessons[i], contain)
	}
	return res
}

func NewStudentSuggestedLesson(lesson *entity.StudentSuggestedLesson) *StudentSuggestedLesson {
	return &StudentSuggestedLesson{
		SubjectID: lesson.SubjectId,
		Total:     lesson.Total,
	}
}

func NewStudentSuggestedLessons(lessons entity.StudentSuggestedLessons) StudentSuggestedLessons {
	res := make(StudentSuggestedLessons, len(lessons))
	for i := range lessons {
		res[i] = NewStudentSuggestedLesson(lessons[i])
	}
	return res
}
