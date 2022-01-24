package entity

import (
	"encoding/json"
	"time"

	"github.com/calmato/shs-web/api/proto/lesson"
	"gorm.io/datatypes"
)

type StudentSubmission struct {
	StudentID            string           `gorm:"primaryKey;<-:create"`
	ShiftSummaryID       int64            `gorm:"primaryKey;<-:create"`
	Decided              bool             `gorm:""`
	SuggestedLessons     SuggestedLessons `gorm:"-"`
	SuggestedLessonsJSON datatypes.JSON   `gorm:"column:suggested_lessons"`
	CreatedAt            time.Time        `gorm:"<-:create"`
	UpdatedAt            time.Time        `gorm:""`
}

type StudentSubmissions []*StudentSubmission

type SuggestedLesson struct {
	SubjectID int64 `json:"subjectId"`
	Total     int64 `json:"total"`
}

type SuggestedLessons []*SuggestedLesson

func NewStudentSubmission(
	studentID string, summaryID int64, decided bool, lessons SuggestedLessons,
) *StudentSubmission {
	return &StudentSubmission{
		StudentID:        studentID,
		ShiftSummaryID:   summaryID,
		Decided:          decided,
		SuggestedLessons: lessons,
	}
}

func NewSuggestedLesson(l *lesson.StudentSuggestedLesson) *SuggestedLesson {
	return &SuggestedLesson{
		SubjectID: l.SubjectId,
		Total:     l.Total,
	}
}

func NewSuggestedLessons(ls []*lesson.StudentSuggestedLesson) SuggestedLessons {
	res := make(SuggestedLessons, len(ls))
	for i := range ls {
		res[i] = NewSuggestedLesson(ls[i])
	}
	return res
}

func (s *StudentSubmission) Fill() error {
	var lessons SuggestedLessons
	if err := json.Unmarshal(s.SuggestedLessonsJSON, &lessons); err != nil {
		return err
	}
	s.SuggestedLessons = lessons
	return nil
}

func (s *StudentSubmission) FillJSON() error {
	v, err := json.Marshal(s.SuggestedLessons)
	if err != nil {
		return err
	}
	s.SuggestedLessonsJSON = datatypes.JSON(v)
	return nil
}

func (s *StudentSubmission) Proto() *lesson.StudentSubmission {
	return &lesson.StudentSubmission{
		StudentId:        s.StudentID,
		ShiftSummaryId:   s.ShiftSummaryID,
		Decided:          s.Decided,
		SuggestedLessons: s.SuggestedLessons.Proto(),
		CreatedAt:        s.CreatedAt.Unix(),
		UpdatedAt:        s.UpdatedAt.Unix(),
	}
}

func (ss StudentSubmissions) Fill() error {
	for i := range ss {
		if err := ss[i].Fill(); err != nil {
			return err
		}
	}
	return nil
}

func (ss StudentSubmissions) FillJSON() error {
	for i := range ss {
		if err := ss[i].FillJSON(); err != nil {
			return err
		}
	}
	return nil
}

func (ss StudentSubmissions) Proto() []*lesson.StudentSubmission {
	submissions := make([]*lesson.StudentSubmission, len(ss))
	for i := range ss {
		submissions[i] = ss[i].Proto()
	}
	return submissions
}

func (l *SuggestedLesson) Proto() *lesson.SuggestedLesson {
	return &lesson.SuggestedLesson{
		SubjectId: l.SubjectID,
		Total:     l.Total,
	}
}

func (ls SuggestedLessons) SubjectIDs() []int64 {
	res := make([]int64, len(ls))
	for i := range ls {
		res[i] = ls[i].SubjectID
	}
	return res
}

func (ls SuggestedLessons) Proto() []*lesson.SuggestedLesson {
	res := make([]*lesson.SuggestedLesson, len(ls))
	for i := range ls {
		res[i] = ls[i].Proto()
	}
	return res
}
