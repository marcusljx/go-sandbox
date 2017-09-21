package xjlm

import "encoding/json"

// JSON wraps json unmarshal with panic on error
func JSON(v interface{}) []byte {
	b, err := json.Marshal(v)
	CheckErrPanic(err)
	return b
}

