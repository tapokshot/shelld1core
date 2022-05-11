package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/tapokshot/shelld1core/common"
	"github.com/tapokshot/shelld1core/log"
)

const (
	traceIdHeader = "Trace-Id"
)

// RequestLogger log request with meta data
func RequestLogger(skipPaths []string) echo.MiddlewareFunc {
	skipPathsMap := make(map[string]bool, len(skipPaths))
	for _, path := range skipPaths {
		skipPathsMap[path] = true
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			path := ctx.Request().URL.Path
			query := ctx.Request().URL.RawQuery

			requestLog := log.Log(ctx.Request().Context())
			start := time.Now()

			if err := next(ctx); err != nil {
				requestLog.Error("error request",
					err,
					log.IntField("status", ctx.Response().Status),
					log.StringField("st", common.GetStackTraceStr(err)),
					log.StringField("method", ctx.Request().Method),
					log.StringField("path", path),
					log.StringField("query", query),
					log.StringField("ip", ctx.RealIP()),
					log.StringField("user-agent", ctx.Request().UserAgent()))
				return err
			}

			if _, ok := skipPathsMap[path]; !ok {
				latency := time.Now().Sub(start)
				requestLog.Info("success request",
					log.IntField("status", ctx.Response().Status),
					log.StringField("method", ctx.Request().Method),
					log.StringField("path", path),
					log.StringField("query", query),
					log.StringField("ip", ctx.RealIP()),
					log.StringField("user-agent", ctx.Request().UserAgent()),
					log.DurationField("latency", latency),
				)
			}
			return nil
		}
	}
}

// Recovery returns a gin.HandlerFunc (middleware)
// that recovers from any panics and logs requests using uber-go/zap.
// All errors are logged using zap.Error().
// stack means whether output the stack info.
// The stack info is easy to find where the error occurs but the stack info is too large.
//func Recovery(log *log.Logger) gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		defer func() {
//			if errMsg := recover(); errMsg != nil {
//				// Check for a broken connection, as it is not really a
//				// condition that warrants a panic stack trace.
//				var brokenPipe bool
//				if ne, ok := errMsg.(*net.OpError); ok {
//					if se, ok := ne.Err.(*os.SyscallError); ok {
//						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
//							brokenPipe = true
//						}
//					}
//				}
//
//				_, _ = httputil.DumpRequest(ctx.Request, false)
//				if brokenPipe {
//					//log.Error(ctx.Request.URL.Path,
//					//	errMsg,
//					//	logger.StringField("request", string(httpRequest)),
//					//)
//					// If the connection is dead, we can't write a status to it.
//					ctx.Error(errMsg.(error)) // nolint: errcheck
//					ctx.Abort()
//					return
//				}
//
//				err := errors.New(fmt.Sprintf("%v", errMsg))
//				log.Error("[Recovery from panic]",
//					err,
//					//zap.Any("error", errMsg),
//					//logger.StringField("request", string(httpRequest)),
//					//logger.StringField("stack", string(debug.Stack())),
//				)
//				ctx.Error(err)
//				ctx.Abort()
//			}
//		}()
//		ctx.Next()
//	}
//}
