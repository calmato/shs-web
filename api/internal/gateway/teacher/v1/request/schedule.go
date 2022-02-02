package request

import "time"

type ScheduleLesson struct {
	StartTime string `json:"startTime,omitempty"` // 授業開始時間
	EndTime   string `json:"endTime,omitempty"`   // 授業終了時間
}

type ScheduleToUpdate struct {
	Weekday  time.Weekday      `json:"weekday,omitempty"`  // 曜日
	IsClosed bool              `json:"isClosed,omitempty"` // 休校フラグ
	Lessons  []*ScheduleLesson `json:"lessons,omitempty"`
}

type UpdateSchedulesRequest struct {
	Schedules []*ScheduleToUpdate `json:"schedules,omitempty"` // 授業スケジュール一覧
}
