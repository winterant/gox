package x

// If returns trueVal if condition is true, otherwise falseVal.
//
// For example:
//
//	If(1<2, "yes", "no") // returns "yes"
//	If(1>2, "yes", "no") // returns "no"
func If[T any](condition bool, trueVal T, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}
