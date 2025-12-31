package x

import (
	"fmt"
	"io"

	"github.com/winterant/gox/pkg/xcaller"
)

type fundamental struct {
	message string
	stack   *xcaller.CallerStack
}

func (e *fundamental) Error() string { return e.message }

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

// Recover catch panic and recover. call it at start of goroutine.
// For example:
//
//	go func(){
//		defer Recover(func(err error){  // must start with defer !!!!!!
//			log.Println(err)
//		})
//		...
//	}()
//
// Wrong example:
//
//	defer func(){
//		Recover(func(err error){  // must start with defer !!!!!! but not! it will not catch panic!!!!!!
//			log.Println(err)
//		})
//	}()
func Recover(f func(error)) {
	if r := recover(); r != nil {
		err := &fundamental{
			message: fmt.Sprintf("%+v", r),
			stack:   xcaller.GetCallerStack(2),
		}
		f(err)
	}
}
