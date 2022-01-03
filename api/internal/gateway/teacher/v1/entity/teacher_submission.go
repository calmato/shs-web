package entity

import (
	"time"

	"github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
)

type TeacherSubmission struct {
	ID               int64                   `json:"id"`               // シフト募集ID
	Year             int32                   `json:"year"`             // 年
	Month            int32                   `json:"month"`            // 月
	ShiftStatus      ShiftStatus             `json:"shiftStatus"`      // シフト募集ステータス
	SubmissionStatus TeacherSubmissionStatus `json:"submissionStatus"` // シフト提出ステータス
	OpenAt           time.Time               `json:"openAt"`           // 募集開始日時
	EndAt            time.Time               `json:"endAt"`            // 募集締切日時
	CreatedAt        time.Time               `json:"createdAt"`        // 登録日時
	UpdatedAt        time.Time               `json:"updatedAt"`        // 更新日時
}

type TeacherSubmissions []*TeacherSubmission

type TeacherSubmissionStatus int32

const (
	TeacherSubmissionStatusUnknown   TeacherSubmissionStatus = 0
	TeacherSubmissionStatusWaiting   TeacherSubmissionStatus = 1
	TeacherSubmissionStatusSubmitted TeacherSubmissionStatus = 2
)

func NewTeacherSubmission(summary *entity.ShiftSummary, submission *entity.TeacherSubmission) *TeacherSubmission {
	return &TeacherSubmission{
		ID:               summary.Id,
		Year:             summary.YearMonth / 100,
		Month:            summary.YearMonth % 100,
		ShiftStatus:      NewShiftStatus(summary.Status),
		SubmissionStatus: NewTeacherSubmissionStatus(summary, submission),
		OpenAt:           jst.ParseFromUnix(summary.OpenAt),
		EndAt:            jst.ParseFromUnix(summary.EndAt),
		CreatedAt:        jst.ParseFromUnix(summary.CreatedAt),
		UpdatedAt:        jst.ParseFromUnix(summary.UpdatedAt),
	}
}

func NewTeacherSubmissions(
	summaries entity.ShiftSummaries, submissions map[int64]*entity.TeacherSubmission,
) TeacherSubmissions {
	ss := make(TeacherSubmissions, len(summaries))
	for i, s := range summaries {
		submission := submissions[s.Id] // null: 出勤不可
		ss[i] = NewTeacherSubmission(s, submission)
	}
	return ss
}

func NewTeacherSubmissionStatus(
	summary *entity.ShiftSummary, submission *entity.TeacherSubmission,
) TeacherSubmissionStatus {
	if summary == nil {
		return TeacherSubmissionStatusUnknown
	}
	if submission == nil || !submission.Decided {
		return TeacherSubmissionStatusWaiting
	}
	return TeacherSubmissionStatusSubmitted
}
