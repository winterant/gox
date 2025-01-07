package x

func MapKeys[K comparable, V any](value map[K]V) []K {
	keys := make([]K, 0, len(value))
	for k := range value {
		keys = append(keys, k)
	}
	return keys
}

func MapValues[K comparable, V any](value map[K]V) []V {
	values := make([]V, 0, len(value))
	for k := range value {
		values = append(values, value[k])
	}
	return values
}
