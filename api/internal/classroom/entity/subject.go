package entity

import (
	"time"

	"github.com/calmato/shs-web/api/proto/classroom"
)

type Subject struct {
	ID        int64     `gorm:"primaryKey;autoIncrement;<-:create"`
	Name      string    `gorm:""`
	Color     string    `gorm:""`
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time `gorm:""`
}

type Subjects []*Subject

func (s *Subject) Proto() *classroom.Subject {
	return &classroom.Subject{
		Id:        s.ID,
		Name:      s.Name,
		Color:     s.Color,
		CreatedAt: s.CreatedAt.Unix(),
		UpdatedAt: s.UpdatedAt.Unix(),
	}
}

func (ss Subjects) Proto() []*classroom.Subject {
	subjects := make([]*classroom.Subject, len(ss))
	for i := range ss {
		subjects[i] = ss[i].Proto()
	}
	return subjects
}
