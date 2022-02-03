package entity

import (
	"encoding/json"
	"time"

	"github.com/calmato/shs-web/api/proto/lesson"
	"gorm.io/datatypes"
)

type StudentShiftTemplate struct {
	StudentID            string           `gorm:"primaryKey;<-:create"`
	Schedules            ShiftSchedules   `gorm:"-"`
	SchedulesJSON        datatypes.JSON   `gorm:"column:schedules"`
	SuggestedLessons     SuggestedLessons `gorm:"-"`
	SuggestedLessonsJSON datatypes.JSON   `gorm:"column:suggested_lessons"`
	CreatedAt            time.Time        `gorm:"<-:create"`
	UpdatedAt            time.Time        `gorm:""`
}

type ShiftSchedule struct {
	Weekday time.Weekday    `json:"weekday"`
	Lessons LessonSchedules `json:"lessons"`
}

type ShiftSchedules []*ShiftSchedule

type LessonSchedule struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type LessonSchedules []*LessonSchedule

func NewStudentShiftTemplate(studentID string, template *lesson.StudentShiftTemplateToUpsert) *StudentShiftTemplate {
	return &StudentShiftTemplate{
		StudentID:        studentID,
		Schedules:        newShiftSchedules(template.Schedules),
		SuggestedLessons: NewSuggestedLessons(template.SuggestedLessons),
	}
}

func newShiftSchedule(schedule *lesson.StudentShiftTemplateToUpsert_Schedule) *ShiftSchedule {
	return &ShiftSchedule{
		Weekday: time.Weekday(schedule.Weekday),
		Lessons: newLessonSchedules(schedule.Lessons),
	}
}

func newShiftSchedules(schedules []*lesson.StudentShiftTemplateToUpsert_Schedule) ShiftSchedules {
	res := make(ShiftSchedules, len(schedules))
	for i := range schedules {
		res[i] = newShiftSchedule(schedules[i])
	}
	return res
}

func newLessonSchedule(lesson *lesson.StudentShiftTemplateToUpsert_Lesson) *LessonSchedule {
	return &LessonSchedule{
		StartTime: lesson.StartTime,
		EndTime:   lesson.EndTime,
	}
}

func newLessonSchedules(lessons []*lesson.StudentShiftTemplateToUpsert_Lesson) LessonSchedules {
	res := make(LessonSchedules, len(lessons))
	for i := range lessons {
		res[i] = newLessonSchedule(lessons[i])
	}
	return res
}

func (t *StudentShiftTemplate) Fill() error {
	if err := t.fillSchedules(); err != nil {
		return err
	}
	if err := t.fillSuggestedLessons(); err != nil {
		return err
	}
	return nil
}

func (t *StudentShiftTemplate) fillSchedules() error {
	var schedules ShiftSchedules
	if err := json.Unmarshal(t.SchedulesJSON, &schedules); err != nil {
		return err
	}
	t.Schedules = schedules
	return nil
}

func (t *StudentShiftTemplate) fillSuggestedLessons() error {
	var lessons SuggestedLessons
	if err := json.Unmarshal(t.SuggestedLessonsJSON, &lessons); err != nil {
		return err
	}
	t.SuggestedLessons = lessons
	return nil
}

func (t *StudentShiftTemplate) FillJSON() error {
	if err := t.fillSchedulesJSON(); err != nil {
		return err
	}
	if err := t.fillSuggestedLessonsJSON(); err != nil {
		return err
	}
	return nil
}

func (t *StudentShiftTemplate) fillSchedulesJSON() error {
	v, err := json.Marshal(t.Schedules)
	if err != nil {
		return err
	}
	t.SchedulesJSON = datatypes.JSON(v)
	return nil
}

func (t *StudentShiftTemplate) fillSuggestedLessonsJSON() error {
	v, err := json.Marshal(t.SuggestedLessons)
	if err != nil {
		return err
	}
	t.SuggestedLessonsJSON = datatypes.JSON(v)
	return nil
}

func (t *StudentShiftTemplate) Proto() *lesson.StudentShiftTemplate {
	return &lesson.StudentShiftTemplate{
		StudentId:       t.StudentID,
		Schedules:       t.Schedules.Proto(),
		SuggesteLessons: t.SuggestedLessons.Proto(),
		CreatedAt:       t.CreatedAt.Unix(),
		UpdatedAt:       t.UpdatedAt.Unix(),
	}
}

func (s *ShiftSchedule) Proto() *lesson.ShiftSchedule {
	lessons := make([]*lesson.LessonSchedule, len(s.Lessons))
	for i := range s.Lessons {
		lessons[i] = &lesson.LessonSchedule{
			StartTime: s.Lessons[i].StartTime,
			EndTime:   s.Lessons[i].EndTime,
		}
	}

	return &lesson.ShiftSchedule{
		Weekday: int32(s.Weekday),
		Lessons: lessons,
	}
}

func (ss ShiftSchedules) Proto() []*lesson.ShiftSchedule {
	res := make([]*lesson.ShiftSchedule, len(ss))
	for i := range ss {
		res[i] = ss[i].Proto()
	}
	return res
}
