package x

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

func ToJsonBytes(data any) []byte {
	byteBuf := bytes.NewBuffer([]byte{})
	encoder := json.NewEncoder(byteBuf)
	encoder.SetEscapeHTML(false) // 不转义 特殊字符
	err := encoder.Encode(data)
	if err != nil {
		return nil
	}
	return byteBuf.Bytes()
}

func ToJson(data any) string {
	byteBuf := ToJsonBytes(data)
	return strings.TrimSpace(string(byteBuf))
}

func FromJson(jsonStr string, v any) (string, error) {
	if v == nil {
		return jsonStr, errors.New("the arg `v` must be a non-nil pointer")
	}
	var errs []error
	for i := 0; i < 3; i++ {
		err := json.Unmarshal([]byte(jsonStr), v)
		if err != nil {
			errs = append(errs, err)
			unquoted, err := strconv.Unquote(jsonStr)
			if err != nil {
				return jsonStr, errors.Join(errs...)
			}
			jsonStr = unquoted
			continue
		}
		return jsonStr, nil
	}
	return jsonStr, errors.Join(errs...)
}
