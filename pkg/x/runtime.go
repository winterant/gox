package x

import (
	"fmt"
	"runtime"
	"strings"
)

func RuntimeStack(skip int) string {
	callers := make([]uintptr, 32)
	numFrames := runtime.Callers(2+skip, callers)
	frames := runtime.CallersFrames(callers[:numFrames])

	var sb strings.Builder
	for {
		frame, more := frames.Next()
		sb.WriteString(fmt.Sprintf("%s\n\t%s:%d\n", frame.Function, frame.File, frame.Line))
		if !more {
			break
		}
	}
	return sb.String()
}
