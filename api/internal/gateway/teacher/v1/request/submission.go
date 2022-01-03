package request

type UpsertTeacherShiftsRequest struct {
	ShiftIDs []int64 `json:"shiftIds,omitempty"` // 出勤可能シフトID一覧
}
