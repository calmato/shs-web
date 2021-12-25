package response

import "github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"

type SchedulesResponse struct {
	Schedules entity.Schedules `json:"schedules"` // 授業スケジュール一覧
}
