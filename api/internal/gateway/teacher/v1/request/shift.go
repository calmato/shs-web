package request

type UpdateShiftSummaryScheduleRequest struct {
	OpenDate string `json:"openDate,omitempty"` // 募集開始日
	EndDate  string `json:"endDate,omitempty"`  // 募集締切日
}

type UpdateShiftSummaryDecidedRequest struct {
	Decided bool `json:"decided,omitempty"` // 授業スケジュール確定フラグ
}

type CreateShiftsRequest struct {
	YearMonth   string   `json:"yearMonth,omitempty"`   // シフト募集年月
	OpenDate    string   `json:"openDate,omitempty"`    // 募集開始日
	EndDate     string   `json:"endDate,omitempty"`     // 募集締切日
	ClosedDates []string `json:"closedDates,omitempty"` // 休講日一覧
}
