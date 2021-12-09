package jst

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNow(t *testing.T) {
	t.Parallel()
	assert.NotNil(t, Now())
}

func TestDate(t *testing.T) {
	t.Parallel()
	now := Now()
	assert.Equal(t, now, Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond()))
}

func TestToTime(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                 string
		now                  time.Time
		expectDay            time.Time
		expectWeek           time.Time
		expectBeggingOfMonth time.Time
		expectEndOfMonth     time.Time
	}{
		{
			name:                 "success",
			now:                  time.Date(2021, 8, 2, 18, 30, 0, 0, jst),
			expectDay:            time.Date(2021, 8, 2, 0, 0, 0, 0, jst),
			expectWeek:           time.Date(2021, 8, 1, 0, 0, 0, 0, jst),
			expectBeggingOfMonth: time.Date(2021, 8, 1, 0, 0, 0, 0, jst),
			expectEndOfMonth:     time.Date(2021, 8, 31, 0, 0, 0, 0, jst),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expectDay, BegginingOfDay(tt.now))
			assert.Equal(t, tt.expectWeek, BegginingOfWeek(tt.expectWeek, 7))
			assert.Equal(t, tt.expectBeggingOfMonth, BeginningOfMonth(tt.now.Year(), int(tt.now.Month())))
			assert.Equal(t, tt.expectEndOfMonth, EndOfMonth(tt.now.Year(), int(tt.now.Month())))
		})
	}
}

func TestToString(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		now            time.Time
		format         string
		expectFormat   string
		expectYYYYMMDD string
		expectYYYYMM   string
		expectHHMMSS   string
	}{
		{
			name:           "success",
			now:            time.Date(2021, 8, 2, 18, 30, 0, 0, jst),
			format:         "2006/01/02 15:04:05",
			expectFormat:   "2021/08/02 18:30:00",
			expectYYYYMMDD: "20210802",
			expectYYYYMM:   "202108",
			expectHHMMSS:   "183000",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expectFormat, Format(tt.now, tt.format))
			assert.Equal(t, tt.expectYYYYMMDD, FormatYYYYMMDD(tt.now))
			assert.Equal(t, tt.expectYYYYMM, FormatYYYYMM(tt.now))
			assert.Equal(t, tt.expectHHMMSS, FormatHHMMSS(tt.now))
		})
	}
}

func TestParse(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		format    string
		target    string
		expect    time.Time
		expectErr bool
	}{
		{
			name:      "success",
			format:    "20060102",
			target:    "20210802",
			expect:    time.Date(2021, 8, 2, 0, 0, 0, 0, jst),
			expectErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual, err := Parse(tt.format, tt.target)
			assert.Equal(t, tt.expectErr, err != nil, err)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestParseFromYYYYMMDD(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		target    string
		expect    time.Time
		expectErr bool
	}{
		{
			name:      "success",
			target:    "20210802",
			expect:    time.Date(2021, 8, 2, 0, 0, 0, 0, jst),
			expectErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual, err := ParseFromYYYYMMDD(tt.target)
			assert.Equal(t, tt.expectErr, err != nil, err)
			assert.Equal(t, tt.expect, actual)
		})
	}
}
