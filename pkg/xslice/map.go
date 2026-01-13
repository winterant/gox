package xslice

func Map[T any, R any](slice []T, f func(T) R) []R {
	result := make([]R, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}

func MapAsAny[T any](slice []T) []any {
	result := make([]R, len(slice))
	for i, v := range slice {
		result[i] = v
	}
	return result
}
