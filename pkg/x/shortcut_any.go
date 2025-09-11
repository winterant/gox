package x

import "encoding/json"

func AnyTo[T any](value any, target *T) {
	if target == nil {
		return
	}
	bytes := ToJsonBytes(value)
	_ = json.Unmarshal(bytes, target)
}
