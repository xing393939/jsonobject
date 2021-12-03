package jsonobject

import (
	"encoding/json"
)

type JsonObject struct {
	p *interface{}
}

func NewJsonObject(str string) *JsonObject {
	newObj := interface{}(nil)
	_ = json.Unmarshal([]byte(str), &newObj)
	newJo := &JsonObject{
		&newObj,
	}
	return newJo
}

func (jo *JsonObject) GetString(params ...string) string {
	myObj := jo.getObject(params...)
	myStr, ok := myObj.(string)
	if !ok {
		return ""
	}
	return myStr
}

func (jo *JsonObject) GetInt(params ...string) int {
	myObj := jo.GetFloat64(params...)
	return int(myObj)
}

func (jo *JsonObject) GetInt64(params ...string) int64 {
	myObj := jo.GetFloat64(params...)
	return int64(myObj)
}

func (jo *JsonObject) GetFloat64(params ...string) float64 {
	myObj := jo.getObject(params...)
	myFloat, ok := myObj.(float64)
	if !ok {
		return 0
	}
	return myFloat
}

func (jo *JsonObject) GetBool(params ...string) bool {
	myObj := jo.getObject(params...)
	myBool, ok := myObj.(bool)
	if !ok {
		return false
	}
	return myBool
}

func (jo *JsonObject) GetStringMap(params ...string) map[string]interface{} {
	myObj := jo.getObject(params...)
	myMap, ok := myObj.(map[string]interface{})
	if !ok {
		return nil
	}
	return myMap
}

func (jo *JsonObject) GetJsonObjectSlice(params ...string) []*JsonObject {
	myObj := jo.getObject(params...)
	mySlice, ok := myObj.([]interface{})
	var newJoSlice []*JsonObject
	if !ok {
		return newJoSlice
	}
	for k := range mySlice {
		newJoSlice = append(newJoSlice, &JsonObject{&mySlice[k]})
	}
	return newJoSlice
}

func (jo *JsonObject) GetJsonObject(params ...string) *JsonObject {
	myObj := jo.getObject(params...)
	newJo := &JsonObject{
		&myObj,
	}
	return newJo
}

func (jo *JsonObject) getObject(params ...string) interface{} {
	myObj := *jo.p
	if len(params) == 0 {
		return myObj
	}
	myMap, ok := myObj.(map[string]interface{})
	if !ok {
		return nil
	}
	return myMap[params[0]]
}
