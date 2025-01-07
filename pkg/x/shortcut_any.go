package x

import "encoding/json"

func AnyTo[T any](value any) *T {
	bytes := ToJsonBytes(value)
	t := new(T)
	err := json.Unmarshal(bytes, t)
	if err != nil {
		return nil
	}
	return t
}
