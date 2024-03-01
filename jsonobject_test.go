package jsonobject

import (
	"encoding/json"
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("unexpected value obtained; got %q; want %q", a, b)
	}
}

func TestGetLeafNode(t *testing.T) {
	jsonContent := `{
		"bool": true, 
        "int": 1, 
        "int64": 64, 
        "float64": 1.64, 
        "string": "abc",
        "list": [true, 1, 64, 1.64, "abc"]
	}`
	jo := NewJsonObject(jsonContent)
	// item exists
	assertEqual(t, jo.GetBool("bool"), true)
	assertEqual(t, jo.GetInt("int"), 1)
	assertEqual(t, jo.GetInt64("int64"), int64(64))
	assertEqual(t, jo.GetFloat64("float64"), 1.64)
	assertEqual(t, jo.GetString("string"), "abc")
	assertEqual(t, jo.GetString("float64"), "1.64")
	assertEqual(t, len(jo.GetInterfaceSlice("list")), 5)
	// item not exists
	assertEqual(t, jo.GetBool("bool_extra"), false)
	assertEqual(t, jo.GetInt("int_extra"), 0)
	assertEqual(t, jo.GetInt64("int64_extra"), int64(0))
	assertEqual(t, jo.GetFloat64("float64_extra"), float64(0))
	assertEqual(t, jo.GetString("string_extra"), "")
	// get itself
	list := jo.GetJsonObjectSlice("list")
	assertEqual(t, list[0].GetBool(), true)
	assertEqual(t, list[1].GetInt(), 1)
	assertEqual(t, list[2].GetInt64(), int64(64))
	assertEqual(t, list[3].GetFloat64(), float64(1.64))
	assertEqual(t, list[4].GetString(), "abc")
}

func TestGetNonLeafNode(t *testing.T) {
	jsonContent := `{
		"obj": {
            "name": "John"
        },
        "list": [
            {"name": 1}, {"name": 2}, {"name": 3}
        ]
	}`
	jo := NewJsonObject(jsonContent)
	obj1, _ := json.Marshal(jo.GetJsonObject("obj").GetStringMap())
	obj2, _ := json.Marshal(map[string]interface{}{
		"name": "John",
	})
	assertEqual(t, string(obj1), string(obj2))
	list := jo.GetJsonObjectSlice("list")
	for i, item := range list {
		assertEqual(t, item.GetInt("name"), i+1)
	}
}

func TestIsNil(t *testing.T) {
	jo := NewJsonObject(map[string]interface{}{
		"float32": float32(1.32),
		"int64":   int64(1),
		"int":     1,
		"string":  "1.32",
	})
	assertEqual(t, jo.GetJsonObject("int").GetInt("none"), 0)
	assertEqual(t, jo.GetFloat64("float32"), float64(float32(1.32)))
	assertEqual(t, jo.GetInt64("int64"), int64(1))
	assertEqual(t, jo.GetInt("int"), 1)
	assertEqual(t, jo.GetFloat64("string"), 1.32)
	assertEqual(t, jo.IsNil(), false)

	jo = NewJsonObject(nil)
	assertEqual(t, jo.GetStringMap() == nil, false)
	assertEqual(t, jo.IsNil(), true)

	jo = NewJsonObject(`{"a":1}`)
	assertEqual(t, jo.IsNil(), false)
	assertEqual(t, jo.IsNil("a"), false)
	assertEqual(t, jo.IsNil("b"), true)
}

func TestMarshal(t *testing.T) {
	jsonContent := `{
		"obj": {
            "name": "John"
        },
        "list": [
            {"name": 1}, {"name": 2}, {"name": 3}
        ]
	}`
	jo := NewJsonObject(jsonContent)
	assertEqual(t, len(jo.Marshal()), 65)

	list := jo.GetJsonObjectSlice("list")
	list[0].Set("name", 11)
	assertEqual(t, len(jo.Marshal()), 66)

	obj := jo.GetJsonObject("obj")
	obj.Set("name", "John2")
	assertEqual(t, len(jo.Marshal()), 67)

	mapJo := NewJsonObject(map[string]interface{}{
		"a": "a",
		"b": []map[string]interface{}{
			{"childA": 1},
		},
	})
	mapJo.Set("a", "aa")
	assertEqual(t, mapJo.Marshal(), `{"a":"aa","b":[{"childA":1}]}`)
	mapJo.GetJsonObjectSlice("b")[0].Set("childA", 11)
	assertEqual(t, mapJo.Marshal(), `{"a":"aa","b":[{"childA":11}]}`)
}
