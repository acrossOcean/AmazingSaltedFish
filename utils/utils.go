package utils

import "encoding/json"

// GetJSONOutput 提供所给变量的 json 格式化输出字符串
func GetJSONOutput(v interface{}) string {
	b, _ := json.MarshalIndent(v, " ", "")
	return string(b)
}
