Easy JSON parser for Go. No custom structs, no code generation, no reflection

# Installing
go get github.com/xing393939/jsonobject

# Usage
```go
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
```

# Methods 

| Method  | Zero-value  |
| ------------ | ------------ |
| GetString  |  "" |
| GetInt  | 0  |
| GetBool  |  false |
| GetInt64  | 0  |
| GetFloat64  | 0  |
| GetInt64  | 0  |
| GetJsonObject  | *JsonObject  |
| GetJsonObjectSlice  | []*JsonObject  |