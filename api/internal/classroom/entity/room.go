package entity

import (
	"time"

	"github.com/calmato/shs-web/api/proto/classroom"
)

type Room struct {
	ID        int32     `gorm:"primaryKey;<-:create"`
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time `gorm:""`
}

type Rooms []*Room

func NewRooms(size int) Rooms {
	rooms := make(Rooms, size)
	for i := 0; i < size; i++ {
		room := &Room{ID: int32(i + 1)}
		rooms[i] = room
	}
	return rooms
}

func (r *Room) Proto() *classroom.Room {
	return &classroom.Room{
		Id:        r.ID,
		CreatedAt: r.CreatedAt.Unix(),
		UpdatedAt: r.UpdatedAt.Unix(),
	}
}
