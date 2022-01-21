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

func (ss Subjects) Map() map[int64]*Subject {
	res := make(map[int64]*Subject, len(ss))
	for _, s := range ss {
		res[s.Id] = s
	}
	return res
}

func (ss Subjects) GroupByTeacher(teachers TeacherSubjects) map[string]Subjects {
	sm := ss.Map()
	res := make(map[string]Subjects, len(teachers))
	for _, t := range teachers {
		tss := make(Subjects, 0, len(t.SubjectIds))
		for _, sid := range t.SubjectIds {
			tss = append(tss, sm[sid])
		}
		res[t.TeacherId] = tss
	}
	return res
}

func (ss Subjects) GroupByStudent(students StudentSubjects) map[string]Subjects {
	sm := ss.Map()
	res := make(map[string]Subjects, len(students))
	for _, t := range students {
		sss := make(Subjects, 0, len(t.SubjectIds))
		for _, sid := range t.SubjectIds {
			sss = append(sss, sm[sid])
		}
		res[t.StudentId] = sss
	}
	return res
}
