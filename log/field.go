package log

import (
	"context"
	"time"

	"github.com/shelld1t/core/traceutils"
	"go.uber.org/zap"
)

// Field field of log struct
type Field struct {
	zapField zap.Field
}

func StringField(key, value string) Field {
	return Field{zapField: zap.String(key, value)}
}

func IntField(key string, val int) Field {
	return Field{zapField: zap.Int64(key, int64(val))}
}

func DurationField(key string, val time.Duration) Field {
	return Field{zapField: zap.Duration(key, val)}
}

func traceIdFieldFromCtx(ctx context.Context) zap.Field {
	if traceId := traceutils.FromCtx(ctx); traceId != nil {
		return zap.String("traceId", *traceId)
	}
	return zap.Skip()
}

func extractZapFields(fields ...Field) []zap.Field {
	if len(fields) == 0 {
		return []zap.Field{}
	}
	var zapFields = make([]zap.Field, 0, len(fields))
	for _, f := range fields {
		zapFields = append(zapFields, f.zapField)
	}
	return zapFields
}
