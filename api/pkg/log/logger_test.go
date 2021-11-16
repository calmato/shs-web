package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
)

func TestLogger(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		params *Params
		isErr  bool
	}{
		{
			name: "success only stdout",
			params: &Params{
				Level: "debug",
				Path:  "",
			},
			isErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			logger, err := NewLogger(tt.params)
			if tt.isErr {
				assert.Error(t, err)
				assert.Nil(t, logger)
				return
			}
			assert.NoError(t, err)
			assert.NotNil(t, logger)
		})
	}
}

func TestLogger_GetLogLevel(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		level  string
		expect zapcore.Level
	}{
		{
			name:   "debug",
			level:  "debug",
			expect: zapcore.DebugLevel,
		},
		{
			name:   "info",
			level:  "info",
			expect: zapcore.InfoLevel,
		},
		{
			name:   "warn",
			level:  "warn",
			expect: zapcore.WarnLevel,
		},
		{
			name:   "error",
			level:  "error",
			expect: zapcore.ErrorLevel,
		},
		{
			name:   "default",
			level:  "",
			expect: zapcore.InfoLevel,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, getLogLevel(tt.level))
		})
	}
}
