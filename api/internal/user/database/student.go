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

func (s *student) List(ctx context.Context, params *ListStudentsParams, fields ...string) (entity.Students, error) {
	var students entity.Students
	if len(fields) == 0 {
		fields = studentFields
	}

	stmt := s.db.DB.Table(studentTable).Select(fields)
	if params.Limit > 0 {
		stmt.Limit(params.Limit)
	}
	if params.Offset > 0 {
		stmt.Offset((params.Offset))
	}

	err := stmt.Find(&students).Error
	if err != nil {
		return nil, dbError(err)
	}
	students.Fill(s.now())
	return students, nil
}

func (s *student) MultiGet(ctx context.Context, ids []string, fields ...string) (entity.Students, error) {
	var students entity.Students
	if len(fields) == 0 {
		fields = studentFields
	}

	stmt := s.db.DB.Table(studentTable).Select(fields).
		Where("id IN (?)", ids)

	err := stmt.Find(&students).Error
	if err != nil {
		return nil, dbError(err)
	}
	students.Fill(s.now())
	return students, nil
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
	student.Fill(s.now())
	return student, nil
}

func (s *student) Create(ctx context.Context, student *entity.Student) error {
	student.CreatedAt = s.now()
	student.UpdatedAt = s.now()

	_, err := s.auth.CreateUser(ctx, student.ID, student.Mail, student.Password)
	if err != nil {
		return dbError(err)
	}

	tx, err := s.db.Begin()
	if err != nil {
		return dbError(err)
	}
	defer s.db.Close(tx)

	err = tx.Table(studentTable).Create(&student).Error
	if err != nil {
		tx.Rollback()
		return dbError(err)
	}
	return dbError(tx.Commit().Error)
}

func (s *student) Delete(ctx context.Context, studentID string) error {
	now := s.now()
	uid := uuid.Base58Encode(uuid.New())

	err := s.auth.DeleteUser(ctx, studentID)
	if err != nil {
		return dbError(err)
	}

	tx, err := s.db.Begin()
	if err != nil {
		return dbError(err)
	}
	defer s.db.Close(tx)

	params := map[string]interface{}{
		"mail":       uid,
		"updated_at": now,
		"deleted_at": now,
	}

	err = tx.Table(studentTable).Where("id = ?", studentID).Updates(params).Error
	if err != nil {
		tx.Rollback()
		return dbError(err)
	}
	return dbError(tx.Commit().Error)
}

func (s *student) Count(ctx context.Context) (int64, error) {
	var total int64
	err := s.db.DB.Table(studentTable).Where("deleted_at IS NULL").Count(&total).Error
	return total, dbError(err)
}
