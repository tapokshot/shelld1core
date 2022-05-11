package log

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

//Logger wrap zap logger
type Logger struct {
	zapLog *zap.Logger
}

type Config struct {
	// LogLevel zapLog level (debug, info .etc)
	LogLevel string `yaml:"log_level"`
}

var rootLog *Logger

func Log(ctx context.Context) *Logger {
	return &Logger{
		zapLog: rootLog.zapLog.With(traceIdFieldFromCtx(ctx)),
	}
}

func SetRootLog(log *Logger) {
	rootLog = log
}

func NewLogger(config *Config) (*Logger, error) {
	zapLog, err := createZapLog()
	if err != nil {
		return nil, fmt.Errorf("error create zap zapLog %v", err)
	}
	return &Logger{
		zapLog: zapLog,
	}, nil
}

func createZapLog() (*zap.Logger, error) {
	zapConf := zap.NewProductionConfig()
	zapConf.OutputPaths = []string{"stdout"}
	zapConf.DisableStacktrace = true
	return zapConf.Build(
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zap.FatalLevel),
	)
}

func (l *Logger) With(fields ...Field) *Logger {
	return &Logger{
		zapLog: l.zapLog.With(extractZapFields(fields...)...),
	}
}

func (l *Logger) Debug(msg string) {
	l.zapLog.Debug(msg)
}

func (l *Logger) Info(msg string, fields ...Field) {
	zapFields := extractZapFields(fields...)
	l.zapLog.Info(msg, zapFields...)
}

func (l *Logger) Error(msg string, err error, fields ...Field) {
	zapFields := extractZapFields(fields...)
	zapFields = append(zapFields, zap.Error(err))
	l.zapLog.Error(msg, zapFields...)
}
