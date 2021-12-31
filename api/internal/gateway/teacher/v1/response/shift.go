package response

import "github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"

type ShiftSummariesResponse struct {
	Summaries entity.ShiftSummaries `json:"summaries"` // シフト募集概要一覧
}

type ShiftsResponse struct {
	Summary *entity.ShiftSummary `json:"summary"` // シフト募集概要
	Shifts  entity.ShiftDetails  `json:"shifts"`  // 募集シフト一覧
}
