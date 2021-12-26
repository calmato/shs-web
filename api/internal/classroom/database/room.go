package database

import (
	"context"
	"time"

	"github.com/calmato/shs-web/api/internal/classroom/entity"
	"github.com/calmato/shs-web/api/pkg/database"
	"github.com/calmato/shs-web/api/pkg/jst"
	"gorm.io/gorm"
)

const roomTable = "rooms"

// var roomFields = []string{
// 	"id", "created_at", "updated_at",
// }

type room struct {
	db  *database.Client
	now func() time.Time
}

func NewRoom(db *database.Client) Room {
	return &room{
		db:  db,
		now: jst.Now,
	}
}

func (r *room) Replace(ctx context.Context, rooms entity.Rooms) error {
	now := r.now()
	for i := range rooms {
		rooms[i].CreatedAt = now
		rooms[i].UpdatedAt = now
	}

	tx, err := r.db.Begin()
	if err != nil {
		return dbError(err)
	}
	defer r.db.Close(tx)

	err = tx.Table(roomTable).
		Session(&gorm.Session{AllowGlobalUpdate: true}).
		Delete(&entity.Room{}).Error
	if err != nil {
		tx.Rollback()
		return dbError(err)
	}

	err = tx.Table(roomTable).Create(&rooms).Error
	if err != nil {
		tx.Rollback()
		return dbError(err)
	}
	return dbError(tx.Commit().Error)
}

func (r *room) Count(ctx context.Context) (int64, error) {
	var total int64
	err := r.db.DB.Table(roomTable).Count(&total).Error
	return total, dbError(err)
}
