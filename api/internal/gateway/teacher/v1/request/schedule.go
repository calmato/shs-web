package request

import "github.com/calmato/shs-web/api/internal/gateway/teacher/v1/entity"

type UpdateSchedulesRequest struct {
	Schedules []*entity.ScheduleToUpdate `json:"schedules,omitempty"` // 授業スケジュール一覧
}
