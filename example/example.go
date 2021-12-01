package main

import "github.com/xing393939/jsonobject"

func main() {
	jsonContent := `{"name":"John", "score":100}`
	jo := jsonobject.NewJsonObject(jsonContent)
	println(jo.GetString("name"), jo.GetInt("score"))

	jsonContent = `[{"name":"John", "score":100}, {"name":"Tom", "score":200}]`
	jo = jsonobject.NewJsonObject(jsonContent)
	joArr := jo.GetJsonObjectSlice()
	for _, joRow := range joArr {
		println(joRow.GetString("name"), joRow.GetInt("score"))
	}
}
