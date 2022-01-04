package main

import "github.com/xing393939/jsonobject"

func main() {
	jsonContent := `{
		"isMaster": false, 
		"metadata": {
			"name": "oracle"
		},
		"tags": ["db", "sql"]
	}`
	jo := jsonobject.NewJsonObject(jsonContent)
	println(jo.GetBool("isMaster"))
	println(jo.GetJsonObject("metadata").GetString("name"))
	joArr := jo.GetJsonObjectSlice("tags")
	for _, joRow := range joArr {
		println(joRow.GetString())
	}

	jsonContent = `[{"name":"John", "score":100}, {"name":"Tom", "score":200}]`
	jo = jsonobject.NewJsonObject(jsonContent)
	joArr = jo.GetJsonObjectSlice()
	for _, joRow := range joArr {
		println(joRow.GetString("name"), joRow.GetInt("score"))
	}

	joEntry := jsonobject.JsonObject{}
	jo = &joEntry
	// jo = jsonobject.NewJsonObject("")
	jo.Set("a", "a")
	println(jo.GetString("a"))

	jo = jsonobject.NewJsonObject(map[string]interface{}{
		"key": "",
		"stringMapInt": map[string]int{
			"string1": 1,
		},
	})
	jo.Set("key", "value")
	jo2 := jo.GetJsonObject("stringMapInt")
	jo2.Set("string1", 100)
	jo2.Set("string2", 200)
	jo2.Set("string3", 300)
	println(jo.GetString("key"))
	println(jo.GetJsonObject("stringMapInt").GetInt("string1"))
	println(jo.GetJsonObject("stringMapInt").GetInt("string2"))
	println(jo.GetJsonObject("stringMapInt").GetInt("string3"))
}
