package x

import (
	"fmt"
	"runtime"
	"strings"
)

type CallerStack struct {
	Stack []callerFrame
}

func (c *CallerStack) String() string {
	var buf strings.Builder
	for i := range c.Stack {
		buf.WriteString("\t")
		buf.WriteString(c.Stack[i].String())
		if i < len(c.Stack)-1 {
			buf.WriteString("\n")
		}
	}
	return buf.String()
}

type callerFrame struct {
	Function string
	File     string
	Line     int
}

func (r *callerFrame) String() string {
	return fmt.Sprintf("%s:%d %s", r.File, r.Line, r.Function)
}

func GetCallerStack(skip int) *CallerStack {
	callers := make([]uintptr, 32)
	numFrames := runtime.Callers(2+skip, callers)
	frames := runtime.CallersFrames(callers[:numFrames])

	stack := make([]callerFrame, 0, numFrames)
	for {
		frame, more := frames.Next()
		stack = append(stack, callerFrame{
			Function: frame.Function,
			File:     frame.File,
			Line:     frame.Line,
		})
		if !more {
			break
		}
	}
	return &CallerStack{
		Stack: stack,
	}
}
