package log

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerParams struct {
	Level string
	Path  string
}

// NewLogger - ログ出力用クライアントの生成
func NewLogger(params *LoggerParams) (*zap.Logger, error) {
	level := getLogLevel(params.Level)
	encoderConfig := zapcore.EncoderConfig{}

	// 標準出力設定
	consoleCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(os.Stdout),
		level,
	)

	// Path==""のとき、標準出力のみ
	if params.Path == "" {
		logger := zap.New(zapcore.NewTee(consoleCore))
		return logger, nil
	}

	// logPath!==""のとき、ファイル出力も追加
	outputPath := fmt.Sprintf("%s/outputs.log", params.Path)
	file, err := os.OpenFile(outputPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		return nil, err
	}

	// ファイル出力設定
	logCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(file),
		level,
	)

	logger := zap.New(zapcore.NewTee(consoleCore, logCore))
	return logger, nil
}

func getLogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}
