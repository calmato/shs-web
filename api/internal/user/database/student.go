package database

import (
	"time"

	"github.com/calmato/shs-web/api/pkg/database"
	"github.com/calmato/shs-web/api/pkg/firebase/authentication"
	"github.com/calmato/shs-web/api/pkg/jst"
)

type student struct {
	db   *database.Client
	auth authentication.Client
	now  func() time.Time
}

func NewStudent(db *database.Client, auth authentication.Client) Student {
	return &student{
		db:   db,
		auth: auth,
		now:  jst.Now,
	}
}
