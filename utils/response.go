package util

import "encoding/json"

func Response(data any, err error) []byte {
	jsonData := map[string]any{
		"data":  data,
		"error": err,
	}
	marshal, _ := json.Marshal(jsonData)
	return marshal
}
