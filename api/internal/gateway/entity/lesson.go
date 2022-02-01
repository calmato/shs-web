package entity

import (
	"github.com/calmato/shs-web/api/pkg/set"
	"github.com/calmato/shs-web/api/proto/lesson"
)

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

func (ls Lessons) GroupByTeacherID() map[string]Lessons {
	res := make(map[string]Lessons)
	for _, l := range ls {
		if _, ok := res[l.TeacherId]; !ok {
			res[l.TeacherId] = make(Lessons, 0)
		}
		res[l.TeacherId] = append(res[l.TeacherId], l)
	}
	return res
}

func (ls Lessons) GroupByStudentID() map[string]Lessons {
	res := make(map[string]Lessons)
	for _, l := range ls {
		if _, ok := res[l.StudentId]; !ok {
			res[l.StudentId] = make(Lessons, 0)
		}
		res[l.StudentId] = append(res[l.StudentId], l)
	}
	return res
}
