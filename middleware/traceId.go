package middleware

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/tapokshot/shelld1core/traceutils"
)

var errorTraceIdMultiValue = errors.New("Header [TRACE_ID] must be single value")

// SetTraceId set trace id into context, if not exist then generate new traceId
func SetTraceId() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			traceId, err := getTraceIdFromRequest(ctx.Request())
			if err != nil {
				return err
			}
			r := ctx.Request().WithContext(
				context.WithValue(
					ctx.Request().Context(), traceutils.TraceIdKey, traceId,
				),
			)
			ctx.SetRequest(r)
			return next(ctx)
		}
	}
}

func getTraceIdFromRequest(request *http.Request) (string, error) {
	if vArr, ok := request.Header[traceIdHeader]; ok {
		if len(vArr) != 1 {
			return "", errorTraceIdMultiValue
		}
		return vArr[0], nil
	} else {
		return traceutils.GenerateTraceId(), nil
	}
}
