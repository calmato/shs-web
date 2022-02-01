package database

import (
	"context"
	"time"

	"github.com/calmato/shs-web/api/internal/lesson/entity"
	"github.com/calmato/shs-web/api/pkg/database"
	"github.com/calmato/shs-web/api/pkg/jst"
	"gorm.io/gorm"
)

const lessonTable = "lessons"

var lessonFields = []string{
	"id", "shift_summary_id", "shift_id", "room_id", "subject_id",
	"teacher_id", "student_id", "notes", "created_at", "updated_at",
}

type lesson struct {
	db  *database.Client
	now func() time.Time
}

func NewLesson(db *database.Client) Lesson {
	return &lesson{
		db:  db,
		now: jst.Now,
	}
}

func (l *lesson) List(ctx context.Context, params *ListLessonsParams, fields ...string) (entity.Lessons, error) {
	var lessons entity.Lessons
	if len(fields) == 0 {
		fields = lessonFields
	}

	stmt := l.db.DB.Table(lessonTable).Select(fields)
	params.Statement(stmt)

	err := stmt.Find(&lessons).Error
	return lessons, dbError(err)
}

func (l *lesson) Create(ctx context.Context, lesson *entity.Lesson) error {
	now := l.now()
	lesson.CreatedAt = now
	lesson.UpdatedAt = now

	tx, err := l.db.Begin()
	if err != nil {
		return dbError(err)
	}
	defer l.db.Close(tx)

	err = tx.Table(lessonTable).Create(&lesson).Error
	if err != nil {
		tx.Rollback()
		return dbError(err)
	}
	return dbError(tx.Commit().Error)
}

func (l *lesson) Update(ctx context.Context, lessonID int64, lesson *entity.Lesson) error {
	tx, err := l.db.Begin()
	if err != nil {
		return dbError(err)
	}
	defer l.db.Close(tx)

	var current *entity.Lesson
	err = tx.Table(lessonTable).
		Select([]string{"id", "created_at"}).
		Where("id = ?", lessonID).
		First(&current).Error
	if err != nil {
		return dbError(err)
	}

	lesson.ID = current.ID
	lesson.CreatedAt = current.CreatedAt
	lesson.UpdatedAt = l.now()

	err = tx.Table(lessonTable).Save(&lesson).Error
	if err != nil {
		tx.Rollback()
		return dbError(err)
	}
	return dbError(tx.Commit().Error)
}

func (l *lesson) Delete(ctx context.Context, lessonID int64) error {
	tx, err := l.db.Begin()
	if err != nil {
		return dbError(err)
	}
	defer l.db.Close(tx)

	err = tx.Table(lessonTable).
		Where("id = ?", lessonID).
		Delete(&entity.Lesson{}).Error
	if err != nil {
		tx.Rollback()
		return dbError(err)
	}
	return dbError(tx.Commit().Error)
}

func (l *lesson) Count(ctx context.Context, params *ListLessonsParams) (int64, error) {
	var total int64
	stmt := l.db.DB.Table(lessonTable)
	params.Statement(stmt)

	err := stmt.Count(&total).Error
	return total, dbError(err)
}

func (p *ListLessonsParams) Statement(tx *gorm.DB) {
	if p.TeacherID != "" {
		tx.Where("teacher_id = ?", p.TeacherID)
	}
	if p.StudentID != "" {
		tx.Where("student_id = ?", p.StudentID)
	}
	if p.ShiftSummaryID > 0 {
		tx.Where("shift_summary_id = ?", p.ShiftSummaryID)
	}
	if p.ShiftID > 0 {
		tx.Where("shift_id = ?", p.ShiftID)
	}
	if len(p.ShiftIDs) > 0 {
		tx.Where("shift_id IN (?)", p.ShiftIDs)
	}
	if p.Limit > 0 {
		tx.Limit(p.Limit)
	}
	if p.Offset > 0 {
		tx.Limit(p.Offset)
	}
}
