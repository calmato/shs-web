package entity

import (
	"time"

	"github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
)

type StudentSubmission struct {
	ShiftSummaryID   int64                   `json:"id"`               // 授業希望募集ID
	Year             int32                   `json:"year"`             // 年
	Month            int32                   `json:"month"`            // 月
	ShiftStatus      ShiftStatus             `json:"shiftStatus"`      // 授業希望募集ステータス
	SubmissionStatus StudentSubmissionStatus `json:"submissionStatus"` // 授業希望提出ステータス
	OpenAt           time.Time               `json:"openAt"`           // 募集開始日時
	EndAt            time.Time               `json:"endAt"`            // 募集締切日時
	CreatedAt        time.Time               `json:"createdAt"`        // 登録日時
	UpdatedAt        time.Time               `json:"updatedAt"`        // 更新日時
}

type StudentSubmissions []*StudentSubmission

type StudentSubmissionDetail struct {
	Student               *Student                `json:"student"`               // 生徒情報
	SuggestedLessons      StudentSuggestedLessons `json:"suggestedLessons"`      // 希望受講授業一覧
	SuggestedLessonsTotal int64                   `json:"suggestedLessonsTotal"` // 希望受講授業数
	LessonTotal           int64                   `json:"lessonTotal"`           // 登録受講授業数
}

type StudentSubmissionDetails []*StudentSubmissionDetail

type StudentSubmissionStatus int32

const (
	StudentSubmissionStatusUnknown   StudentSubmissionStatus = 0
	StudentSubmissionStatusWaiting   StudentSubmissionStatus = 1
	StudentSubmissionStatusSubmitted StudentSubmissionStatus = 2
)

func NewStudentSubmission(summary *entity.ShiftSummary, submission *entity.StudentSubmission) *StudentSubmission {
	return &StudentSubmission{
		ShiftSummaryID:   summary.Id,
		Year:             summary.YearMonth / 100,
		Month:            summary.YearMonth % 100,
		ShiftStatus:      NewShiftStatus(summary.Status),
		SubmissionStatus: NewStudentSubmissionStatus(summary, submission),
		OpenAt:           jst.ParseFromUnix(summary.OpenAt),
		EndAt:            jst.ParseFromUnix(summary.EndAt),
		CreatedAt:        jst.ParseFromUnix(summary.CreatedAt),
		UpdatedAt:        jst.ParseFromUnix(summary.UpdatedAt),
	}
}

func NewStudentSubmissions(
	summaries entity.ShiftSummaries, submissions map[int64]*entity.StudentSubmission,
) StudentSubmissions {
	ss := make(StudentSubmissions, len(summaries))
	for i, s := range summaries {
		submission := submissions[s.Id] // null: 出勤不可
		ss[i] = NewStudentSubmission(s, submission)
	}
	return ss
}

func NewStudentSubmissionStatus(
	summary *entity.ShiftSummary, submission *entity.StudentSubmission,
) StudentSubmissionStatus {
	if summary == nil {
		return StudentSubmissionStatusUnknown
	}
	if submission != nil && submission.Decided {
		return StudentSubmissionStatusSubmitted
	}
	return StudentSubmissionStatusWaiting
}
