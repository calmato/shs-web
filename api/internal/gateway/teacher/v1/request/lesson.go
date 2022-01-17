package request

type CreateLessonRequest struct {
	ShiftID   int64  `json:"shiftId,omitempty"`   // 授業スケジュールID
	SubjectID int64  `json:"subjectId,omitempty"` // 授業科目ID
	Room      int32  `json:"room,omitempty"`      // 教室番号
	TeacherID string `json:"teacherId,omitempty"` // 講師ID
	StudentID string `json:"studentId,omitempty"` // 生徒ID
}
