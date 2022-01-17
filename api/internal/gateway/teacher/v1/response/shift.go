package response

import "github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"

type ShiftSummariesResponse struct {
	Summaries entity.ShiftSummaries `json:"summaries"` // シフト募集概要一覧
}

type ShiftsResponse struct {
	Summary  *entity.ShiftSummary            `json:"summary"`  // シフト募集概要
	Shifts   entity.ShiftDetails             `json:"shifts"`   // 募集シフト一覧
	Rooms    int64                           `json:"rooms"`    // 教室数
	Teachers entity.TeacherSubmissionDetails `json:"teachers"` // 講師情報一覧
	Students entity.StudentSubmissionDetails `json:"students"` // 生徒情報一覧
	Lessons  entity.Lessons                  `json:"lessons"`  // 授業一覧
}
