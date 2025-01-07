package x

import (
	"bytes"
	"encoding/json"
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
