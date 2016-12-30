package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Map type map
type Map map[string]interface{}

// Int ...
func (m Map) Int(key string) int64 {
	if v, ok := m[key]; ok {
		return V2Int64(v)
	}
	return int64(0)
}

// Float ...
func (m Map) Float(key string) float64 {
	if v, ok := m[key]; ok {
		return V2Float64(v)
	}
	return float64(0)
}

// Uint ...
func (m Map) Uint(key string) uint64 {
	if v, ok := m[key]; ok {
		return V2Uint64(v)
	}
	return 0
}

// String ...
func (m Map) String(key string) string {
	if v, ok := m[key]; ok {
		return V2Str(v)
	}
	return ""
}

// Array ...
func (m Map) Array(key string) []interface{} {
	if v, ok := m[key]; ok {
		return V2Array(v)
	}
	return nil
}

// Map ...
func (m Map) Map(key string) Map {
	if v, ok := m[key]; ok {
		return V2Map(v)
	}
	return nil
}

// StringP ...
func (m Map) StringP(path string) string {
	val, err := m.ValP(path)
	if err != nil {
		return ""
	}
	return V2Str(val)
}

// MapP ...
func (m Map) MapP(path string) Map {
	val, err := m.ValP(path)
	if err != nil {
		return nil
	}
	return V2Map(val)
}

// Val ...
func (m Map) Val(key string) interface{} {
	if val, ok := m[key]; ok {
		return val
	}
	return nil

}

// ValP ...
func (m Map) ValP(path string) (interface{}, error) {
	path = strings.TrimPrefix(path, "/")
	paths := strings.Split(path, "/")
	return m.valP(paths)
}

func (m Map) valP(paths []string) (interface{}, error) {
	lens := len(paths)
	var v interface{} = m
	for i := 0; i < lens; i++ {
		switch reflect.TypeOf(v).Kind() {
		case reflect.Map:
			tmp := V2Map(v)
			if tmp == nil {
				return nil, fmt.Errorf(fmt.Sprintf("invalid map in path(/%v)", strings.Join(paths[:i], "/")))
			}
			v = tmp.Val(paths[i])
		default:
			return nil, fmt.Errorf(fmt.Sprintf("invalid type(%v) in path(/%v)",
				reflect.TypeOf(v).Kind(), strings.Join(paths[:i], "/")))
		}
	}
	if v == nil {
		return nil, fmt.Errorf(fmt.Sprintf(fmt.Sprintf(
			"value not found in path(/%v)", strings.Join(paths, "/"),
		)))
	}
	return v, nil
}

// IsExit ...
func (m Map) IsExit(key string) bool {
	if _, ok := m[key]; ok {
		return true
	}
	return false
}

// SetIfNotExit ...
func (m Map) SetIfNotExit(key string, val interface{}) bool {
	if m.IsExit(key) {
		return false
	}
	m[key] = val
	return true
}

// Set ...
func (m Map) Set(key string, val interface{}) {
	m[key] = val
}

// Del ...
func (m Map) Del(key string) {
	delete(m, key)
}

// V2Int64 ...
func V2Int64(v interface{}) int64 {
	val, err := IntVal(v)
	if err == nil {
		return val
	}
	return int64(0)
}

// IntVal ...
func IntVal(v interface{}) (int64, error) {
	if v == nil {
		return 0, fmt.Errorf("arg value is null")
	}
	k := reflect.TypeOf(v)
	switch k.Kind() {
	case reflect.Int:
		return int64(v.(int)), nil
	case reflect.Int8:
		return int64(v.(int8)), nil
	case reflect.Int16:
		return int64(v.(int16)), nil
	case reflect.Int32:
		return int64(v.(int32)), nil
	case reflect.Int64:
		return v.(int64), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int64(V2Uint64(v)), nil
	case reflect.Float32, reflect.Float64:
		return int64(V2Float64(v)), nil
	case reflect.String:
		fv, err := strconv.ParseInt(v.(string), 10, 64)
		if err == nil {
			return fv, nil
		}
		return 0, err
	case reflect.Struct:
		return 0, fmt.Errorf("incompactable kind(%v)", k.Kind())
	default:
		return 0, fmt.Errorf("incompactable kind(%v)", k.Kind())
	}
}

// V2Float64 ...
func V2Float64(v interface{}) float64 {
	val, err := FloatVal(v)
	if err == nil {
		return val
	}
	return 0
}

// FloatVal ...
func FloatVal(v interface{}) (float64, error) {
	if v == nil {
		return 0, fmt.Errorf("arg value is null")
	}
	k := reflect.TypeOf(v)
	switch k.Kind() {
	case reflect.Float32:
		return float64(v.(float32)), nil
	case reflect.Float64:
		return float64(v.(float64)), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(V2Uint64(v)), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(V2Int64(v)), nil
	case reflect.String:
		fv, err := strconv.ParseFloat(v.(string), 64)
		if err == nil {
			return fv, nil
		}
		return 0, err
	default:
		return 0, fmt.Errorf("incompactable kind(%v)", k.Kind().String())
	}
}

// V2Uint64 ...
func V2Uint64(v interface{}) uint64 {
	val, err := Uint64Val(v)
	if err == nil {
		return val
	}
	return 0
}

// Uint64Val ...
func Uint64Val(v interface{}) (uint64, error) {
	if v == nil {
		return 0, fmt.Errorf("arg value is null")
	}
	k := reflect.TypeOf(v)
	switch k.Kind() {
	case reflect.Uint:
		return uint64(v.(uint)), nil
	case reflect.Uint8:
		return uint64(v.(uint8)), nil
	case reflect.Uint16:
		return uint64(v.(uint16)), nil
	case reflect.Uint32:
		return uint64(v.(uint32)), nil
	case reflect.Uint64:
		return v.(uint64), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return uint64(V2Int64(v)), nil
	case reflect.Float32, reflect.Float64:
		return uint64(V2Float64(v)), nil
	case reflect.String:
		fv, err := strconv.ParseUint(v.(string), 10, 64);
		if err == nil {
			return fv, nil
		}
		return 0, err

	default:
		return 0, fmt.Errorf("incompactable kind(%v)", k.Kind().String())
	}
}

// V2Str ...
func V2Str(v interface{}) string {
	if v == nil {
		return ""
	}
	str := ""
	switch reflect.TypeOf(v).Kind() {
	case reflect.String:
		str = v.(string)
	case reflect.Array:
		vals := reflect.ValueOf(v)
		vs := []string{}
		for i := 0; i < vals.Len(); i++ {
			vs = append(vs, fmt.Sprintf("%v", vals.Index(i).Interface()))
		}
		str = strings.Join(vs, ",")

	default:
		str = fmt.Sprintf("%v", v)
	}
	return str
}

// V2Map ...
func V2Map(v interface{}) Map {
	if m, ok := v.(Map); ok {
		return m
	} else if m, ok := v.(map[string]interface{}); ok {
		return Map(m)
	} else {
		return nil
	}
}

// V2Array ...
func V2Array(v interface{}) []interface{} {
	if vals, ok := v.([]interface{}); ok {
		return vals
	}
	vs := reflect.ValueOf(v)
	if vs.Kind() != reflect.Array {
		return nil
	}
	vals := []interface{}{}
	for i := 0; i < vs.Len(); i++ {
		vals = append(vals, fmt.Sprintf("%v", vs.Index(i).Interface()))
	}
	return vals
}
