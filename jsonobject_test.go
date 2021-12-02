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
	assertEqual(t, jo.GetFloat64("float64"), float64(1.64))
	assertEqual(t, jo.GetString("string"), "abc")
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
	obj1, _ := json.Marshal(jo.GetJsonObject("obj").GetMapWithStringKey())
	obj2, _ := json.Marshal(map[string]interface{}{
		"name": "John",
	})
	assertEqual(t, string(obj1), string(obj2))
	list := jo.GetJsonObjectSlice("list")
	for i, item := range list {
		assertEqual(t, item.GetInt("name"), i+1)
	}
}
