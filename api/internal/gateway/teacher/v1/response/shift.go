package response

import "github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"

type ShiftsResponse struct {
	Summary *entity.ShiftSummary     `json:"summary"` // シフト募集概要
	Shifts  map[string]entity.Shifts `json:"shifts"`  // 募集シフト一覧
}
