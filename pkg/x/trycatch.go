package x

import (
	"fmt"
)

// TryCatch catches panic. Must be called in defer. For example: `defer TryCatch()`
func TryCatch(f func(error)) {
	if r := recover(); r != nil {
		err := fmt.Errorf("%+v\n%+v", r, RuntimeStack(2))
		f(err)
	}
}
