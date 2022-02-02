package response

import "github.com/calmato/shs-web/api/internal/gateway/student/v1/entity"

type SubmissionResponse struct {
	Summary          *entity.StudentSubmission      `json:"summary"`          // 授業希望概要
	Shifts           entity.StudentShiftDetails     `json:"shifts"`           // 授業募集(日毎)一覧
	SuggestedLessons entity.StudentSuggestedLessons `json:"suggestedLessons"` // 授業募集(科目毎)一覧
}

type SubmissionsResponse struct {
	Summaries entity.StudentSubmissions `json:"summaries"` // 授業希望募集概要一覧
}
