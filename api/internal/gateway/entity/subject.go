package entity

import "github.com/calmato/shs-web/api/proto/classroom"

type Subject struct {
	*classroom.Subject
}

type Subjects []*Subject

func NewSubject(subject *classroom.Subject) *Subject {
	return &Subject{
		Subject: subject,
	}
}

func NewSubjects(subjects []*classroom.Subject) Subjects {
	ss := make(Subjects, len(subjects))
	for i := range subjects {
		ss[i] = NewSubject(subjects[i])
	}
	return ss
}
