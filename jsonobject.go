package jsonobject

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type JsonObject struct {
	p *interface{}
}

func NewJsonObject(obj interface{}) *JsonObject {
	newObj := interface{}(nil)
	switch obj.(type) {
	case string:
		_ = json.Unmarshal([]byte(obj.(string)), &newObj)
	case []byte:
		_ = json.Unmarshal(obj.([]byte), &newObj)
	case map[string]interface{}:
		newObj = obj
	}
	newJo := &JsonObject{
		&newObj,
	}
	return newJo
}

func (jo *JsonObject) Set(key string, value interface{}) bool {
	myObj := jo.getObject()
	if myMap, ok := myObj.(map[string]interface{}); ok {
		myMap[key] = value
	}
	return true
}

func (jo *JsonObject) GetString(params ...string) string {
	myObj := jo.getObject(params...)
	myStr, ok := myObj.(string)
	if !ok {
		myFloat, ok := myObj.(float64)
		if !ok {
			return ""
		}
		myStr = fmt.Sprint(myFloat)
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
	myFloat := float64(0)
	switch myObj.(type) {
	case float64:
		myFloat = myObj.(float64)
	case float32:
		myFloat = float64(myObj.(float32))
	case int64:
		myFloat = float64(myObj.(int64))
	case int:
		myFloat = float64(myObj.(int))
	case string:
		myFloat, _ = strconv.ParseFloat(myObj.(string), 64)
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
		myMap = make(map[string]interface{})
	}
	return myMap
}

func (jo *JsonObject) GetJsonObjectSlice(params ...string) []*JsonObject {
	myObj := jo.getObject(params...)
	var newJoSlice []*JsonObject
	if mySlice, ok := myObj.([]interface{}); ok {
		for k := range mySlice {
			newJoSlice = append(newJoSlice, &JsonObject{&mySlice[k]})
		}
	}
	if mySlice, ok := myObj.([]map[string]interface{}); ok {
		for k := range mySlice {
			tmpMap := interface{}(mySlice[k])
			newJoSlice = append(newJoSlice, &JsonObject{&tmpMap})
		}
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
	if jo.p == nil {
		return nil
	}
	myObj := *jo.p
	if len(params) == 0 {
		return myObj
	}
	if myMap, ok := myObj.(map[string]interface{}); ok {
		return myMap[params[0]]
	}
	return nil
}

func (jo *JsonObject) IsNil(params ...string) bool {
	return jo.getObject(params...) == nil
}

func (jo *JsonObject) Marshal(params ...string) string {
	myObj := jo.getObject(params...)
	bytes, err := json.Marshal(myObj)
	if err != nil {
		return ""
	}
	return string(bytes)
}
