package database

import (
	"time"

	"github.com/calmato/shs-web/api/pkg/database"
	"github.com/calmato/shs-web/api/pkg/jst"
)

type subject struct {
	db  *database.Client
	now func() time.Time
}

func NewSubject(db *database.Client) Subject {
	return &subject{
		db:  db,
		now: jst.Now,
	}
}
