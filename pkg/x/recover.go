package x

import "github.com/winterant/gox/pkg/xcaller"

// Recover Deprecated catch panic and recover. call it at start of goroutine.
// For example:
//
//	go func(){
//		defer Recover()
//		...
//	}()
//
// Deprecated: No more maintenance, directly use r:=recover
func Recover(f func(panicMessage any, callstack *xcaller.CallerStack)) {
	if r := recover(); r != nil {
		f(r, xcaller.GetCallerStack(2))
	}
}
