package database

import (
	"context"
	"time"

	"github.com/calmato/shs-web/api/internal/user/entity"
	"github.com/calmato/shs-web/api/pkg/database"
	"github.com/calmato/shs-web/api/pkg/firebase/authentication"
	"github.com/calmato/shs-web/api/pkg/jst"
)

const teacherTable = "teachers"

// var teacherFields = []string{
// 	"id", "last_name", "first_name", "last_name_kana", "first_name_kana",
// 	"mail", "role", "created_at", "updated_at", "deleted_at",
// }

type teacher struct {
	db   *database.Client
	auth authentication.Client
	now  func() time.Time
}

func NewTeacher(db *database.Client, auth authentication.Client) Teacher {
	return &teacher{
		db:   db,
		auth: auth,
		now:  jst.Now,
	}
}

func (t *teacher) Create(ctx context.Context, teacher *entity.Teacher) error {
	teacher.CreatedAt = t.now()
	teacher.UpdatedAt = t.now()

	_, err := t.auth.CreateUser(ctx, teacher.ID, teacher.Mail, teacher.Password)
	if err != nil {
		return dbError(err)
	}

	tx, err := t.db.Begin()
	if err != nil {
		return dbError(err)
	}
	defer t.db.Close(tx)

	err = tx.Table(teacherTable).Create(&teacher).Error
	if err != nil {
		tx.Rollback()
		return dbError(err)
	}
	return dbError(tx.Commit().Error)
}
