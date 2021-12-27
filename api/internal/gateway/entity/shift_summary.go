package entity

import "github.com/calmato/shs-web/api/proto/lesson"

type ShiftSummary struct {
	*lesson.ShiftSummary
}

type ShiftSummaries []*ShiftSummary

func NewShiftSummary(summary *lesson.ShiftSummary) *ShiftSummary {
	return &ShiftSummary{
		ShiftSummary: summary,
	}
}

func NewShiftSummaries(summaries []*lesson.ShiftSummary) ShiftSummaries {
	ss := make(ShiftSummaries, len(summaries))
	for i := range summaries {
		ss[i] = NewShiftSummary(summaries[i])
	}
	return ss
}
