package xerrors

import (
	"fmt"
	"github.com/winterant/gox/pkg/x"
	"io"
)

type fundamental struct {
	message string
	stack   *x.CallerStack
}

type withMessage struct {
	message string
	cause   error
}

type withStack struct {
	error
	*x.CallerStack
}

func (e *fundamental) Error() string { return e.message }

func (e *withMessage) Error() string { return e.message }

func (e *fundamental) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%s\n%s", e.message, e.stack.String())
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, e.message)
	case 'q':
		fmt.Fprintf(s, "%q", e.message)
	}
}

func (e *withMessage) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%s: %+v", e.message, e.cause)
			return
		}
		fallthrough
	case 's':
		fmt.Fprintf(s, "%s: %s", e.message, e.cause)
	case 'q':
		fmt.Fprintf(s, `"%s: %s"`, e.message, e.cause)
	}
}

func (e *withStack) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v\n%s", e.error, e.CallerStack.String())
			return
		}
		fallthrough
	case 's':
		fmt.Fprintf(s, "%s", e.error)
	case 'q':
		fmt.Fprintf(s, "%q", e.error)
	}
}

func New(message string) error {
	return &fundamental{
		message: message,
		stack:   x.GetCallerStack(1),
	}
}

func Errorf(format string, args ...any) error {
	return &fundamental{
		message: fmt.Sprintf(format, args...),
		stack:   x.GetCallerStack(1),
	}
}

func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	switch err.(type) {
	case *fundamental, *withMessage, *withStack: // already wrapped with stack
		return &withMessage{
			cause:   err,
			message: message,
		}
	default: // due to no stack, wrap with stack
		return &withStack{
			&withMessage{
				cause:   err,
				message: message,
			},
			x.GetCallerStack(1),
		}
	}
}

func Cause(err error) error {
	for {
		switch err.(type) {
		case *withMessage:
			err = err.(*withMessage).cause
		case *withStack:
			err = err.(*withStack).error
		default:
			return err
		}
	}
}
