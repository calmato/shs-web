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
