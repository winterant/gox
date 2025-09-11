package x

import (
	"fmt"

	"github.com/winterant/gox/pkg/xcaller"
)

// TryCatch catches panic. Must be called in defer. For example: `defer TryCatch()`
// Deprecated: No more maintenance, directly use r:=recover
func TryCatch(f func(error)) {
	if r := recover(); r != nil {
		err := fmt.Errorf("%+v\n%+v", r, xcaller.GetCallerStack(2))
		f(err)
	}
}
