package entity

import (
	"github.com/calmato/shs-web/api/pkg/set"
	"github.com/calmato/shs-web/api/proto/lesson"
)

type Lesson struct {
	*lesson.Lesson
}

type Lessons []*Lesson

func NewLesson(l *lesson.Lesson) *Lesson {
	return &Lesson{
		Lesson: l,
	}
}

func NewLessons(ls []*lesson.Lesson) Lessons {
	lessons := make(Lessons, len(ls))
	for i := range ls {
		lessons[i] = NewLesson(ls[i])
	}
	return lessons
}

func (ls Lessons) TeacherIDs() []string {
	set := set.New(len(ls))
	for i := range ls {
		set.AddStrings(ls[i].TeacherId)
	}
	return set.Strings()
}

func (ls Lessons) StudentIDs() []string {
	set := set.New(len(ls))
	for i := range ls {
		set.AddStrings(ls[i].StudentId)
	}
	return set.Strings()
}
