package database

import (
	"context"
	"time"

	"github.com/calmato/shs-web/api/internal/classroom/entity"
	"github.com/calmato/shs-web/api/pkg/database"
	"github.com/calmato/shs-web/api/pkg/jst"
)

const studentSubjectTable = "student_subjects"

var studentSubjectFields = []string{
	"student_id", "subject_id", "created_at", "updated_at",
}

type studentSubject struct {
	db  *database.Client
	now func() time.Time
}

func NewStudentSubject(db *database.Client) StudentSubject {
	return &studentSubject{
		db:  db,
		now: jst.Now,
	}
}

func (s *studentSubject) ListByStudentIDs(
	ctx context.Context, studentIDs []string, fields ...string,
) (entity.StudentSubjects, error) {
	var subjects entity.StudentSubjects
	if len(fields) == 0 {
		fields = studentSubjectFields
	}

	stmt := s.db.DB.Table(studentSubjectTable).Select(fields).
		Where("student_id IN (?)", studentIDs)

	err := stmt.Find(&subjects).Error
	return subjects, dbError(err)
}

func (s *studentSubject) Replace(
	ctx context.Context, schoolType entity.SchoolType, subjects entity.StudentSubjects,
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

	studentID := subjects.StudentID()

	var subjectIDs []int64
	err = tx.Table(studentSubjectTable).Select("subjects.id").
		Joins("LEFT JOIN subjects ON student_subjects.subject_id = subjects.id").
		Where("student_subjects.student_id = ?", studentID).
		Where("subjects.school_type = ?", schoolType).
		Find(&subjectIDs).Error
	if err != nil {
		return dbError(err)
	}

	err = tx.Table(studentSubjectTable).
		Where("student_id = ?", studentID).
		Where("subject_id IN (?)", subjectIDs).
		Delete(&entity.StudentSubject{}).Error
	if err != nil {
		tx.Rollback()
		return dbError(err)
	}

	if len(subjects) > 0 {
		err = tx.Table(studentSubjectTable).Create(&subjects).Error
		if err != nil {
			tx.Rollback()
			return dbError(err)
		}
	}
	return dbError(tx.Commit().Error)
}
