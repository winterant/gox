package x

// Recover catch panic and recover. call it at start of goroutine.
// For example:
//
//	go func(){
//		defer Recover()
//		...
//	}()
func Recover(f func(panicMessage interface{}, callstack *CallerStack)) {
	if r := recover(); r != nil {
		f(r, GetCallerStack(2))
	}
}
