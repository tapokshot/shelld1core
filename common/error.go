package common

import (
	"fmt"

	"github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

// GetStackTraceStr get stack trace from pkg/errors
// see https://pkg.go.dev/github.com/pkg/errors#hdr-Retrieving_the_stack_trace_of_an_error_or_wrapper
func GetStackTraceStr(err error) string {
	if err, ok := errors.Cause(err).(stackTracer); ok {
		return fmt.Sprintf("%+v", err.StackTrace())
	}
	return ""
}
