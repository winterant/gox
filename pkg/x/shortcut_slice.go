package x

func SliceMap[T any, R any](slice []T, f func(T) R) []R {
	result := make([]R, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}

func SliceMapAny[T any](slice []T) []any {
	return SliceMap(slice, func(v T) any { return v })
}
