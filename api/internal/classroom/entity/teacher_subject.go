package entity

import (
	"time"

	"github.com/calmato/shs-web/api/pkg/set"
	"github.com/calmato/shs-web/api/proto/classroom"
)

type TeacherSubject struct {
	TeacherID string    `gorm:""`
	SubjectID int64     `gorm:""`
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time `gorm:""`
}

type TeacherSubjects []*TeacherSubject

func NewTeacherSubjects(teacherID string, subjectIDs []int64) TeacherSubjects {
	subjects := make(TeacherSubjects, len(subjectIDs))
	for i := range subjectIDs {
		subject := &TeacherSubject{
			TeacherID: teacherID,
			SubjectID: subjectIDs[i],
		}
		subjects[i] = subject
	}
	return subjects
}

func (ss TeacherSubjects) TeacherID() string {
	if len(ss) == 0 {
		return ""
	}
	return ss[0].TeacherID
}

func (ss TeacherSubjects) SubjectIDs() []int64 {
	set := set.New(len(ss))
	for _, s := range ss {
		set.AddInt64s(s.SubjectID)
	}
	return set.Int64s()
}

func (ss TeacherSubjects) GroupByTeacherID() map[string]TeacherSubjects {
	subjects := make(map[string]TeacherSubjects)
	for _, s := range ss {
		if _, ok := subjects[s.TeacherID]; !ok {
			subjects[s.TeacherID] = make(TeacherSubjects, 0)
		}
		subjects[s.TeacherID] = append(subjects[s.TeacherID], s)
	}
	return subjects
}

func (ss TeacherSubjects) TeacherProto() *classroom.TeacherSubject {
	if len(ss) == 0 {
		return nil
	}

	teacherID := ss[0].TeacherID
	subjects := make([]int64, len(ss))
	for i := range ss {
		subjects[i] = ss[i].SubjectID
	}

	return &classroom.TeacherSubject{
		TeacherId:  teacherID,
		SubjectIds: subjects,
	}
}

func (ss TeacherSubjects) TeachersProto() []*classroom.TeacherSubject {
	subjectsMap := ss.GroupByTeacherID()

	res := make([]*classroom.TeacherSubject, 0, len(subjectsMap))
	for _, subjects := range subjectsMap {
		subject := subjects.TeacherProto()
		res = append(res, subject)
	}
	return res
}
