package utils

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	str := `{"int":123,"string":"string","map":{"mapint":321,"map":{"string":"321"}},"array":["123","321"]}`
	m := JsonDecode([]byte(str))
	if m.IsExit("int") != true {
		t.Error("isexit err")
		return
	}
	if m.IsExit("tni") != false {
		t.Error("isexit false")
		return
	}
	if m.Map("map").IsExit("mapint") != true {
		t.Error("isexit error2")
		return
	}
	if m.Int("int") != 123 {
		t.Error("int err")
		return
	}
	if m.Int("string") != 0 {
		t.Error("int err,not 0")
		return
	}
	if m.Uint("int") != 123 {
		t.Error("uint err")
		return
	}
	if m.Uint("xxx") != 0 {
		t.Error("uint err2")
		return
	}
	if m.Uint("string") != 0 {
		t.Error("uint err3")
		return
	}
	if m.String("int") != "123" {
		t.Error("string err,", m.String("int"))
		return
	}
	if m.Map("string") != nil {
		t.Error("map err,", m.Map("string"))
		return
	}
	if m.Map("map").String("mapint") != "321" {
		t.Error("map err2")
		return
	}
	if m.Map("map").Map("map").String("string") != "321" {
		t.Error("map err3")
		return
	}
	if m.Map("map").Map("map").Map("map") != nil {
		t.Error("map err4")
		return
	}
	if m.Array("int") != nil {
		t.Error("array err,", m.Array("int"))
		return
	}
	if m.Array("array") == nil {
		t.Error("array err2")
		return
	}
	fmt.Println(m.Array("array"))

	if m.SetIfNotExit("array", "array") != false {
		t.Error("setifnotexit err")
		return
	}
	if m.SetIfNotExit("set", "not exit") != true {
		t.Error("setifnotexit err2")
		return
	}
	fmt.Println(JsonEncode(m))
	//test mapP
	if m.MapP("int/123") != nil {
		t.Error("mapP err")
		return
	}
	if m.MapP("map/mapint") != nil {
		t.Error("mapP err2")
		return
	}
	if m.MapP("map/map") == nil {
		t.Error("mapP err2")
		return
	}
	fmt.Println(m.MapP("map/map"))
	if m.StringP("/string/123") != "" {
		t.Error("stringP err")
		return
	}
	if m.StringP("/map/mapint") != "321" {
		t.Error("stringP err2")
		return
	}
	if m.StringP("/map/map/string") != "321" {
		t.Error("stringP err3")
		return
	}
	fmt.Println(m.StringP("/map/map/string"))
	m.Set("set", "set")
	if "set" != m.String("set") {
		t.Error("set err")
		return
	}
	fmt.Println(m.String("set"))
	m.Del("set")
	if "" != m.String("set") {
		t.Error("del err")
		return
	}
	m.Del("")
}

func TestJson(t *testing.T) {
	str := JsonEncode(nil)
	if str != "" {
		t.Error("str err", str)
		return
	}
	v := ""
	str = JsonEncode(v)
	if str != `""` {
		t.Error("str err2", str, len(str))
		return
	}
	str = JsonEncode(Map{
		"abc": "cba",
		"map": Map{"123": 321},
	})
	fmt.Println(str)
}

func TestM(t *testing.T)  {
	m := Map{}
	fmt.Println(m)
	f(&m)
	fmt.Println(m)
}

func f(m *Map) {
	m.Set("1",1)
}