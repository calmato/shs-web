package validation

import (
	"testing"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/stretchr/testify/assert"
)

func TestCreateShifts(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *lesson.CreateShiftsRequest
		isErr bool
	}{
		{
			name: "success",
			req: &lesson.CreateShiftsRequest{
				YearMonth:   202202,
				OpenAt:      jst.Date(2022, 2, 1, 0, 0, 0, 0).Unix(),
				EndAt:       jst.Date(2022, 2, 14, 23, 59, 59, 0).Unix(),
				ClosedDates: []string{"20220212"},
			},
			isErr: false,
		},
		{
			name: "YearMonth is gte",
			req: &lesson.CreateShiftsRequest{
				YearMonth:   99999,
				OpenAt:      jst.Date(2022, 2, 1, 0, 0, 0, 0).Unix(),
				EndAt:       jst.Date(2022, 2, 14, 23, 59, 59, 0).Unix(),
				ClosedDates: []string{"20220212"},
			},
			isErr: true,
		},
		{
			name: "YearMonth is lt",
			req: &lesson.CreateShiftsRequest{
				YearMonth:   1000000,
				OpenAt:      jst.Date(2022, 2, 1, 0, 0, 0, 0).Unix(),
				EndAt:       jst.Date(2022, 2, 14, 23, 59, 59, 0).Unix(),
				ClosedDates: []string{"20220212"},
			},
			isErr: true,
		},
		{
			name: "OpenAt is gt",
			req: &lesson.CreateShiftsRequest{
				YearMonth:   202202,
				OpenAt:      0,
				EndAt:       jst.Date(2022, 2, 14, 23, 59, 59, 0).Unix(),
				ClosedDates: []string{"20220212"},
			},
			isErr: true,
		},
		{
			name: "EndAt is gt",
			req: &lesson.CreateShiftsRequest{
				YearMonth:   202202,
				OpenAt:      jst.Date(2022, 2, 1, 0, 0, 0, 0).Unix(),
				EndAt:       0,
				ClosedDates: []string{"20220212"},
			},
			isErr: true,
		},
		{
			name: "ClosedDates is unique",
			req: &lesson.CreateShiftsRequest{
				YearMonth:   202202,
				OpenAt:      jst.Date(2022, 2, 1, 0, 0, 0, 0).Unix(),
				EndAt:       jst.Date(2022, 2, 14, 23, 59, 59, 0).Unix(),
				ClosedDates: []string{"20220212", "20220212"},
			},
			isErr: true,
		},
		{
			name: "ClosedDates is items.len",
			req: &lesson.CreateShiftsRequest{
				YearMonth:   202202,
				OpenAt:      jst.Date(2022, 2, 1, 0, 0, 0, 0).Unix(),
				EndAt:       jst.Date(2022, 2, 14, 23, 59, 59, 0).Unix(),
				ClosedDates: []string{"0212"},
			},
			isErr: true,
		},
		{
			name: "ClosedDates is items.pattern",
			req: &lesson.CreateShiftsRequest{
				YearMonth:   202202,
				OpenAt:      jst.Date(2022, 2, 1, 0, 0, 0, 0).Unix(),
				EndAt:       jst.Date(2022, 2, 14, 23, 59, 59, 0).Unix(),
				ClosedDates: []string{"abcdefgh"},
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.CreateShifts(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}
