package database

import (
	"context"
	"time"

	"github.com/calmato/shs-web/api/internal/user/entity"
	"github.com/calmato/shs-web/api/pkg/database"
	"github.com/calmato/shs-web/api/pkg/firebase/authentication"
	"github.com/calmato/shs-web/api/pkg/jst"
)

const studentTable = "students"

var studentFields = []string{
	"id", "last_name", "first_name", "last_name_kana", "first_name_kana",
	"mail", "birth_year", "created_at", "updated_at", "deleted_at",
}

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

func (s *student) Get(ctx context.Context, id string, fields ...string) (*entity.Student, error) {
	var student *entity.Student
	if len(fields) == 0 {
		fields = studentFields
	}

	stmt := s.db.DB.Table(studentTable).Select(fields).
		Where("id = ?", id)

	err := stmt.First(&student).Error
	if err != nil {
		return nil, dbError(err)
	}
	return student, nil
}
