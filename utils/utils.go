package utils

import "encoding/json"

func GetJsonOutput(v interface{}) string {
	b, _ := json.MarshalIndent(v, " ", "")
	return string(b)
}
