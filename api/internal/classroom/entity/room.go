package entity

import "time"

type Room struct {
	ID        int32     `gorm:"primaryKey;<-:create"`
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time `gorm:""`
}

type Rooms []*Room
