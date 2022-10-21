[![Build](https://github.com/xing393939/jsonobject/actions/workflows/go.yml/badge.svg)](https://github.com/xing393939/jsonobject/actions/?query=branch%3Amain+event%3Apush)
[![codecov](https://codecov.io/gh/xing393939/jsonobject/branch/main/graph/badge.svg)](https://codecov.io/gh/xing393939/jsonobject)
[![Go Reference](https://pkg.go.dev/badge/github.com/xing393939/jsonobject.svg)](https://pkg.go.dev/github.com/xing393939/jsonobject)

Easy JSON parser for Go. No custom structs, no code generation.

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
