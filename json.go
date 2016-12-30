package utils

import (
	"encoding/json"
)

// JSONEncode json encode
func JSONEncode(v interface{}) string {
	if v == nil {
		return ""
	}
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}

// JSONDecode json decode
func JSONDecode(v []byte) Map {
	if v == nil {
		return nil
	}
	var m  = make(Map)
	if err := json.Unmarshal(v, &m); err != nil {
		return nil
	}
	return m
}
