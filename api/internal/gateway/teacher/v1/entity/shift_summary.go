package entity

import (
	"errors"
	"time"

	"github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
)

var errInvalidShiftStatus = errors.New("entity: invalid shift status")

type ShiftSummary struct {
	ID        int64       `json:"id"`        // シフト募集ID
	Year      int32       `json:"year"`      // 年
	Month     int32       `json:"month"`     // 月
	Status    ShiftStatus `json:"status"`    // シフト募集ステータス
	OpenAt    time.Time   `json:"openAt"`    // 募集開始日時
	EndAt     time.Time   `json:"endAt"`     // 募集締切日時
	CreatedAt time.Time   `json:"createdAt"` // 登録日時
	UpdatedAt time.Time   `json:"updatedAt"` // 更新日時
}

type ShiftSummaries []*ShiftSummary

type ShiftStatus int32

const (
	ShiftStatusUnknown   ShiftStatus = 0
	ShiftStatusWaiting   ShiftStatus = 1
	ShiftStatusAccepting ShiftStatus = 2
	ShiftStatusFinished  ShiftStatus = 3
)

func NewShiftSummary(summary *entity.ShiftSummary) *ShiftSummary {
	return &ShiftSummary{
		ID:        summary.Id,
		Year:      summary.YearMonth / 100,
		Month:     summary.YearMonth % 100,
		Status:    NewShiftStatus(summary.Status),
		OpenAt:    jst.ParseFromUnix(summary.OpenAt),
		EndAt:     jst.ParseFromUnix(summary.EndAt),
		CreatedAt: jst.ParseFromUnix(summary.CreatedAt),
		UpdatedAt: jst.ParseFromUnix(summary.UpdatedAt),
	}
}

func NewShiftSummaries(summaries entity.ShiftSummaries) ShiftSummaries {
	ss := make(ShiftSummaries, len(summaries))
	for i := range summaries {
		ss[i] = NewShiftSummary(summaries[i])
	}
	return ss
}

func NewShiftStatus(status lesson.ShiftStatus) ShiftStatus {
	switch status {
	case lesson.ShiftStatus_SHIFT_STATUS_WAITING:
		return ShiftStatusWaiting
	case lesson.ShiftStatus_SHIFT_STATUS_ACCEPTING:
		return ShiftStatusAccepting
	case lesson.ShiftStatus_SHIFT_STATUS_FINISHED:
		return ShiftStatusFinished
	default:
		return ShiftStatusUnknown
	}
}

func (s ShiftStatus) LessonShiftStatus() (lesson.ShiftStatus, error) {
	switch s {
	case ShiftStatusWaiting:
		return lesson.ShiftStatus_SHIFT_STATUS_WAITING, nil
	case ShiftStatusAccepting:
		return lesson.ShiftStatus_SHIFT_STATUS_ACCEPTING, nil
	case ShiftStatusFinished:
		return lesson.ShiftStatus_SHIFT_STATUS_FINISHED, nil
	default:
		return lesson.ShiftStatus_SHIFT_STATUS_UNKNOWN, errInvalidSchoolType
	}
}
