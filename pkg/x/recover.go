package x

import (
	"fmt"

	"github.com/winterant/gox/pkg/xcaller"
)

// Recover Deprecated catch panic and recover. call it at start of goroutine.
// For example:
//
//	go func(){
//		defer Recover(func(err error){
//			log.Println(err)
//		})
//		...
//	}()
func Recover(f func(error)) {
	if r := recover(); r != nil {
		err := fmt.Errorf("%+v\n%s", r, xcaller.GetCallerStack(2).String())
		f(err)
	}
}
