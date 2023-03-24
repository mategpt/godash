package godash

import "encoding/json"

func ToMap(src interface{}) map[string]interface{} {
	b, _ := json.Marshal(&src)
	var m map[string]interface{}
	_ = json.Unmarshal(b, &m)
	return m
}
