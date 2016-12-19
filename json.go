package utils

import (
	"encoding/json"
)

func JsonEncode(v interface{}) string {
	if v == nil {
		return ""
	}
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}

func JsonDecode(v []byte) Map {
	if v == nil {
		return nil
	}
	var m Map = make(Map)
	if err := json.Unmarshal(v, &m); err != nil {
		return nil
	}
	return m
}
