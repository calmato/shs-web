package entity

import (
	"testing"
	"time"

	"github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/stretchr/testify/assert"
)

func TestShift(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 12, 1, 12, 30, 0, 0)
	tests := []struct {
		name   string
		shift  *entity.Shift
		expect *Shift
	}{
		{
			name: "success",
			shift: &entity.Shift{
				Shift: &lesson.Shift{
					Id:             1,
					ShiftSummaryId: 1,
					Date:           "20211226",
					StartTime:      "1700",
					EndTime:        "1830",
					CreatedAt:      now.Unix(),
					UpdatedAt:      now.Unix(),
				},
			},
			expect: &Shift{
				ID:        1,
				Date:      "20211226",
				StartTime: "1700",
				EndTime:   "1830",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewShift(tt.shift))
		})
	}
}

func TestShifts(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 12, 1, 12, 30, 0, 0)
	tests := []struct {
		name   string
		shifts entity.Shifts
		expect Shifts
	}{
		{
			name: "success",
			shifts: entity.Shifts{
				{
					Shift: &lesson.Shift{
						Id:             1,
						ShiftSummaryId: 1,
						Date:           "20211226",
						StartTime:      "1700",
						EndTime:        "1830",
						CreatedAt:      now.Unix(),
						UpdatedAt:      now.Unix(),
					},
				},
			},
			expect: Shifts{
				{
					ID:        1,
					Date:      "20211226",
					StartTime: "1700",
					EndTime:   "1830",
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewShifts(tt.shifts))
		})
	}
}

func TestShifts_GroupByDate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		shifts Shifts
		expect map[time.Time]Shifts
		isErr  bool
	}{
		{
			name: "success",
			shifts: Shifts{
				{
					ID:        1,
					Date:      "20211226",
					StartTime: "1700",
					EndTime:   "1830",
				},
				{
					ID:        2,
					Date:      "20211226",
					StartTime: "1830",
					EndTime:   "2000",
				},
				{
					ID:        3,
					Date:      "20211227",
					StartTime: "1700",
					EndTime:   "1830",
				},
			},
			expect: map[time.Time]Shifts{
				jst.Date(2021, 12, 26, 0, 0, 0, 0): {
					{
						ID:        1,
						Date:      "20211226",
						StartTime: "1700",
						EndTime:   "1830",
					},
					{
						ID:        2,
						Date:      "20211226",
						StartTime: "1830",
						EndTime:   "2000",
					},
				},
				jst.Date(2021, 12, 27, 0, 0, 0, 0): {
					{
						ID:        3,
						Date:      "20211227",
						StartTime: "1700",
						EndTime:   "1830",
					},
				},
			},
			isErr: false,
		},
		{
			name: "failed to parse date",
			shifts: Shifts{
				{
					ID:        1,
					Date:      "20211200",
					StartTime: "1700",
					EndTime:   "1830",
				},
			},
			expect: nil,
			isErr:  true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual, err := tt.shifts.GroupByDate()
			assert.Equal(t, tt.isErr, err != nil, err)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestShifts_NewShiftDetailsForMonth(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name    string
		summary *ShiftSummary
		shifts  map[time.Time]Shifts
		expect  map[string]*ShiftDetail
	}{
		{
			name: "success",
			summary: &ShiftSummary{
				ID:        1,
				Year:      2022,
				Month:     2,
				Status:    ShiftStatusAccepting,
				OpenAt:    jst.Date(2022, 1, 1, 0, 0, 0, 0),
				EndAt:     jst.Date(2022, 1, 15, 0, 0, 0, 0),
				CreatedAt: now,
				UpdatedAt: now,
			},
			shifts: map[time.Time]Shifts{
				jst.Date(2022, 2, 1, 0, 0, 0, 0): {
					{
						ID:        1,
						Date:      "20220201",
						StartTime: "1700",
						EndTime:   "1830",
					},
					{
						ID:        2,
						Date:      "20220201",
						StartTime: "1830",
						EndTime:   "2000",
					},
				},
				jst.Date(2022, 2, 3, 0, 0, 0, 0): {
					{
						ID:        3,
						Date:      "20220203",
						StartTime: "1700",
						EndTime:   "1830",
					},
				},
			},
			expect: map[string]*ShiftDetail{
				"20220201": {
					IsClosed: false,
					Lessons: []*ShiftDetailLesson{
						{ID: 1, StartTime: "1700", EndTime: "1830"},
						{ID: 2, StartTime: "1830", EndTime: "2000"},
					},
				},
				"20220202": {IsClosed: true, Lessons: []*ShiftDetailLesson{}},
				"20220203": {
					IsClosed: false,
					Lessons: []*ShiftDetailLesson{
						{ID: 3, StartTime: "1700", EndTime: "1830"},
					},
				},
				"20220204": {IsClosed: true, Lessons: []*ShiftDetailLesson{}},
				"20220205": {IsClosed: true, Lessons: []*ShiftDetailLesson{}},
				"20220206": {IsClosed: true, Lessons: []*ShiftDetailLesson{}},
				"20220207": {IsClosed: true, Lessons: []*ShiftDetailLesson{}},
				"20220208": {IsClosed: true, Lessons: []*ShiftDetailLesson{}},
				"20220209": {IsClosed: true, Lessons: []*ShiftDetailLesson{}},
				"20220210": {IsClosed: true, Lessons: []*ShiftDetailLesson{}},
				"20220211": {IsClosed: true, Lessons: []*ShiftDetailLesson{}},
				"20220212": {IsClosed: true, Lessons: []*ShiftDetailLesson{}},
				"20220213": {IsClosed: true, Lessons: []*ShiftDetailLesson{}},
				"20220214": {IsClosed: true, Lessons: []*ShiftDetailLesson{}},
				"20220215": {IsClosed: true, Lessons: []*ShiftDetailLesson{}},
				"20220216": {IsClosed: true, Lessons: []*ShiftDetailLesson{}},
				"20220217": {IsClosed: true, Lessons: []*ShiftDetailLesson{}},
				"20220218": {IsClosed: true, Lessons: []*ShiftDetailLesson{}},
				"20220219": {IsClosed: true, Lessons: []*ShiftDetailLesson{}},
				"20220220": {IsClosed: true, Lessons: []*ShiftDetailLesson{}},
				"20220221": {IsClosed: true, Lessons: []*ShiftDetailLesson{}},
				"20220222": {IsClosed: true, Lessons: []*ShiftDetailLesson{}},
				"20220223": {IsClosed: true, Lessons: []*ShiftDetailLesson{}},
				"20220224": {IsClosed: true, Lessons: []*ShiftDetailLesson{}},
				"20220225": {IsClosed: true, Lessons: []*ShiftDetailLesson{}},
				"20220226": {IsClosed: true, Lessons: []*ShiftDetailLesson{}},
				"20220227": {IsClosed: true, Lessons: []*ShiftDetailLesson{}},
				"20220228": {IsClosed: true, Lessons: []*ShiftDetailLesson{}},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewShiftDetailsForMonth(tt.summary, tt.shifts)
			assert.Equal(t, tt.expect, actual)
		})
	}
}
