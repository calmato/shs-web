package response

import "github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"

type TeacherSubmissionsResponse struct {
	Summaries entity.TeacherSubmissions `json:"summaries"` // シフト募集概要一覧
}

type TeacherShiftsResponse struct {
	Summary *entity.TeacherSubmission  `json:"summary"` // シフト募集概要
	Shifts  entity.TeacherShiftDetails `json:"shifts"`  // 募集シフト一覧
}

type StudentShiftsResponse struct {
	Summary          *entity.StudentSubmission      `json:"summary"`          // 授業希望募集概要
	Shifts           entity.StudentShiftDetails     `json:"shifts"`           // 募集授業一覧
	SuggestedLessons entity.StudentSuggestedLessons `json:"suggestedLessons"` // 授業毎の受講希望回数
}
