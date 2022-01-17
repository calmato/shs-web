package entity

import "github.com/calmato/shs-web/api/proto/lesson"

type Lesson struct {
	*lesson.Lesson
}

type Lessons []*Lesson

func NewLesson(lesson *lesson.Lesson) *Lesson {
	return &Lesson{
		Lesson: lesson,
	}
}

func NewLessons(lessons []*lesson.Lesson) Lessons {
	ls := make(Lessons, len(lessons))
	for i := range lessons {
		ls[i] = NewLesson(lessons[i])
	}
	return ls
}
