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
}