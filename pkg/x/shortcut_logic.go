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

func In[T comparable](val T, list []T) bool {
	for _, v := range list {
		if v == val {
			return true
		}
	}
	return false
}
