package database

import (
	"context"
	"time"

	"github.com/calmato/shs-web/api/internal/user/entity"
	"github.com/calmato/shs-web/api/pkg/database"
	"github.com/calmato/shs-web/api/pkg/firebase/authentication"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/pkg/uuid"
)

const teacherTable = "teachers"

var teacherFields = []string{
	"id", "last_name", "first_name", "last_name_kana", "first_name_kana",
	"mail", "role", "created_at", "updated_at", "deleted_at",
}

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

func (t *teacher) List(ctx context.Context, params *ListTeachersParams, fields ...string) (entity.Teachers, error) {
	var teachers entity.Teachers
	if len(fields) == 0 {
		fields = teacherFields
	}

	stmt := t.db.DB.Table(teacherTable).Select(fields)
	if params.Limit > 0 {
		stmt.Limit(params.Limit)
	}
	if params.Offset > 0 {
		stmt.Offset(params.Offset)
	}

	err := stmt.Find(&teachers).Error
	return teachers, dbError(err)
}

func (t *teacher) MultiGet(ctx context.Context, ids []string, fields ...string) (entity.Teachers, error) {
	var teachers entity.Teachers
	if len(fields) == 0 {
		fields = teacherFields
	}

	stmt := t.db.DB.Table(teacherTable).Select(fields).
		Where("id IN (?)", ids)

	err := stmt.Find(&teachers).Error
	return teachers, dbError(err)
}

func (t *teacher) Get(ctx context.Context, id string, fields ...string) (*entity.Teacher, error) {
	var teacher *entity.Teacher
	if len(fields) == 0 {
		fields = teacherFields
	}

	stmt := t.db.DB.Table(teacherTable).Select(fields).
		Where("id = ?", id)

	err := stmt.First(&teacher).Error
	if err != nil {
		return nil, dbError(err)
	}
	return teacher, nil
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

func (t *teacher) UpdateMail(ctx context.Context, teacherID string, mail string) error {
	_, err := t.auth.UpdateEmail(ctx, teacherID, mail)
	if err != nil {
		return dbError(err)
	}

	tx, err := t.db.Begin()
	if err != nil {
		return dbError(err)
	}
	defer t.db.Close(tx)

	params := map[string]interface{}{
		"mail":       mail,
		"updated_at": t.now(),
	}

	err = tx.Table(teacherTable).Where("id = ?", teacherID).Updates(params).Error
	if err != nil {
		tx.Rollback()
		return dbError(err)
	}
	return dbError(tx.Commit().Error)
}

func (t *teacher) UpdatePassword(ctx context.Context, teacherID string, password string) error {
	_, err := t.auth.UpdatePassword(ctx, teacherID, password)
	return dbError(err)
}

func (t *teacher) UpdateRole(ctx context.Context, teacherID string, role entity.Role) error {
	tx, err := t.db.Begin()
	if err != nil {
		return dbError(err)
	}
	defer t.db.Close(tx)

	params := map[string]interface{}{
		"role":       int32(role),
		"updated_at": t.now(),
	}

	err = tx.Table(teacherTable).Where("id = ?", teacherID).Updates(params).Error
	if err != nil {
		tx.Rollback()
		return dbError(err)
	}
	return dbError(tx.Commit().Error)
}

func (t *teacher) Delete(ctx context.Context, teacherID string) error {
	now := t.now()
	uid := uuid.Base58Encode(uuid.New())

	err := t.auth.DeleteUser(ctx, teacherID)
	if err != nil {
		return dbError(err)
	}

	tx, err := t.db.Begin()
	if err != nil {
		return dbError(err)
	}
	defer t.db.Close(tx)

	params := map[string]interface{}{
		"mail":       uid,
		"updated_at": now,
		"deleted_at": now,
	}

	err = tx.Table(teacherTable).Where("id = ?", teacherID).Updates(params).Error
	if err != nil {
		tx.Rollback()
		return dbError(err)
	}
	return dbError(tx.Commit().Error)
}

func (t *teacher) Count(ctx context.Context) (int64, error) {
	var total int64
	err := t.db.DB.Table(teacherTable).Where("deleted_at IS NULL").Count(&total).Error
	return total, dbError(err)
}
