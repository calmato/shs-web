package entity

import (
	"time"

	"github.com/calmato/shs-web/api/pkg/set"
	"github.com/calmato/shs-web/api/proto/classroom"
)

type StudentSubject struct {
	StudentID string    `gorm:"primaryKey;<-:create"`
	SubjectID int64     `gorm:"primaryKey;<-:create"`
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time `gorm:""`
}

type StudentSubjects []*StudentSubject

func NewStudentSubjects(studentID string, subjectIDs []int64) StudentSubjects {
	subjects := make(StudentSubjects, len(subjectIDs))
	for i := range subjectIDs {
		subject := &StudentSubject{
			StudentID: studentID,
			SubjectID: subjectIDs[i],
		}
		subjects[i] = subject
	}
	return subjects
}

func (ss StudentSubjects) StudentID() string {
	if len(ss) == 0 {
		return ""
	}
	return ss[0].StudentID
}

func (ss StudentSubjects) SubjectIDs() []int64 {
	set := set.New(len(ss))
	for _, s := range ss {
		set.AddInt64s(s.SubjectID)
	}
	return set.Int64s()
}

func (ss StudentSubjects) GroupByStudentID() map[string]StudentSubjects {
	subjects := make(map[string]StudentSubjects)
	for _, s := range ss {
		if _, ok := subjects[s.StudentID]; !ok {
			subjects[s.StudentID] = make(StudentSubjects, 0)
		}
		subjects[s.StudentID] = append(subjects[s.StudentID], s)
	}
	return subjects
}

func (ss StudentSubjects) StudentProto() *classroom.StudentSubject {
	if len(ss) == 0 {
		return nil
	}

	studentID := ss[0].StudentID
	subjects := make([]int64, len(ss))
	for i := range ss {
		subjects[i] = ss[i].SubjectID
	}

	return &classroom.StudentSubject{
		StudentId:  studentID,
		SubjectIds: subjects,
	}
}

func (ss StudentSubjects) StudentsProto() []*classroom.StudentSubject {
	subjectsMap := ss.GroupByStudentID()

	res := make([]*classroom.StudentSubject, 0, len(subjectsMap))
	for _, subjects := range subjectsMap {
		subject := subjects.StudentProto()
		res = append(res, subject)
	}
	return res
}
