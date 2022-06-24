package util

import "encoding/json"

func Response(data any, err error) []byte {
	response := map[string]any{
		"data":  data,
		"error": nil,
	}
	if err != nil {
		response = map[string]any{
			"data":  data,
			"error": err.Error(),
		}
	}
	marshal, _ := json.Marshal(response)
	return marshal
}
