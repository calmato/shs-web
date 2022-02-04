package request

import "github.com/calmato/shs-web/api/internal/gateway/student/v1/entity"

type SuggestedLesson struct {
	SubjectID int64 `json:"subjectId,omitempty"` // 授業科目ID
	Total     int64 `json:"total"`               // 授業希望回数
}

type UpsertSubmissionRequest struct {
	SuggestedLessons []*SuggestedLesson `json:"suggestedLessons,omitempty"` // 希望授業一覧
	ShiftIDs         []int64            `json:"shiftIds,omitempty"`         // 授業希望シフトID一覧
}

type UpsertSubmissionTemplateRequest struct {
	Schedules        entity.StudentSchedules        `json:"schedules"`        // 授業希望(曜日毎)一覧
	SuggestedLessons entity.StudentSuggestedLessons `json:"suggestedlessons"` // 授業希望(科目毎)一覧
}
