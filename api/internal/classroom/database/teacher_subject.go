package database

import (
	"context"
	"time"

	"github.com/calmato/shs-web/api/internal/classroom/entity"
	"github.com/calmato/shs-web/api/pkg/database"
	"github.com/calmato/shs-web/api/pkg/jst"
)

const teacherSubjectTable = "teacher_subjects"

var teacherSubjectFields = []string{
	"teacher_id", "subject_id", "created_at", "updated_at",
}

type teacherSubject struct {
	db  *database.Client
	now func() time.Time
}

func NewTeacherSubject(db *database.Client) TeacherSubject {
	return &teacherSubject{
		db:  db,
		now: jst.Now,
	}
}

func (s *teacherSubject) ListByTeacherIDs(
	ctx context.Context, teacherIDs []string, fields ...string,
) (entity.TeacherSubjects, error) {
	var subjects entity.TeacherSubjects
	if len(fields) == 0 {
		fields = teacherSubjectFields
	}

	stmt := s.db.DB.Table(teacherSubjectTable).Select(fields).
		Where("teacher_id IN (?)", teacherIDs)

	err := stmt.Find(&subjects).Error
	return subjects, dbError(err)
}

func (s *teacherSubject) Replace(
	ctx context.Context, schoolType entity.SchoolType, subjects entity.TeacherSubjects,
) error {
	now := s.now()
	for i := range subjects {
		subjects[i].CreatedAt = now
		subjects[i].UpdatedAt = now
	}

	tx, err := s.db.Begin()
	if err != nil {
		return dbError(err)
	}
	defer s.db.Close(tx)

	teacherID := subjects.TeacherID()

	var subjectIDs []int64
	err = tx.Table(teacherSubjectTable).Select("subjects.id").
		Joins("LEFT JOIN subjects ON teacher_subjects.subject_id = subjects.id").
		Where("teacher_subjects.teacher_id = ?", teacherID).
		Where("subjects.school_type = ?", schoolType).
		Find(&subjectIDs).Error
	if err != nil {
		return dbError(err)
	}

	err = tx.Table(teacherSubjectTable).
		Where("teacher_id = ?", teacherID).
		Where("subject_id IN (?)", subjectIDs).
		Delete(&entity.TeacherSubject{}).Error
	if err != nil {
		tx.Rollback()
		return dbError(err)
	}

	err = tx.Table(teacherSubjectTable).Create(&subjects).Error
	if err != nil {
		tx.Rollback()
		return dbError(err)
	}
	return dbError(tx.Commit().Error)
}
