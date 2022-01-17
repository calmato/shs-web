package entity

import (
	"fmt"
	"time"

	"github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
)

type Lesson struct {
	ID        int64     `json:"id"`        // 授業ID
	ShiftID   int64     `json:"shiftId"`   // 授業スケジュールID
	SubjectID int64     `json:"subjectId"` // 授業科目ID
	Room      int32     `json:"room"`      // 教室番号
	TeacherID string    `json:"teacherId"` // 講師ID
	StudentID string    `json:"studentId"` // 生徒ID
	StartAt   time.Time `json:"startAt"`   // 受賞開始日時
	EndAt     time.Time `json:"endAt"`     // 授業終了日時
	Notes     string    `json:"notes"`     // 備考
	CreatedAt time.Time `json:"createdAt"` // 登録日時
	UpdatedAt time.Time `json:"updatedAt"` // 更新日時
}

type Lessons []*Lesson

func NewLesson(lesson *entity.Lesson, shift *entity.Shift) (*Lesson, error) {
	const format = "200601021504"
	startAt, err := jst.Parse(format, fmt.Sprintf("%s%s", shift.Date, shift.StartTime))
	if err != nil {
		return nil, err
	}
	endAt, err := jst.Parse(format, fmt.Sprintf("%s%s", shift.Date, shift.EndTime))
	if err != nil {
		return nil, err
	}
	return &Lesson{
		ID:        lesson.Id,
		ShiftID:   lesson.ShiftId,
		SubjectID: lesson.SubjectId,
		Room:      lesson.RoomId,
		TeacherID: lesson.TeacherId,
		StudentID: lesson.StudentId,
		StartAt:   startAt,
		EndAt:     endAt,
		Notes:     lesson.Notes,
		CreatedAt: jst.ParseFromUnix(lesson.CreatedAt),
		UpdatedAt: jst.ParseFromUnix(lesson.UpdatedAt),
	}, nil
}

func NewLessons(lessons entity.Lessons, shifts map[int64]*entity.Shift) (Lessons, error) {
	ls := make(Lessons, len(lessons))
	for i, lesson := range lessons {
		shift, ok := shifts[lesson.ShiftId]
		if !ok {
			return nil, fmt.Errorf("entity: shift is not found")
		}
		l, err := NewLesson(lesson, shift)
		if err != nil {
			return nil, err
		}
		ls[i] = l
	}
	return ls, nil
}
