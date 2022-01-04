package jsonobject

import (
	"encoding/json"
	"fmt"
	"reflect"
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
	case map[string]interface{}:
		newObj = obj
	default:
		break
	}
	newJo := &JsonObject{
		&newObj,
	}
	return newJo
}

func (jo *JsonObject) Set(key string, value interface{}) {
	myObj := jo.getObject()
	if reflect.ValueOf(myObj).Kind() == reflect.Invalid {
		myObj = map[string]interface{}{}
		jo.p = &myObj
	}
	if reflect.ValueOf(myObj).Kind() != reflect.Map {
		return
	}
	if reflect.TypeOf(myObj).Key() != reflect.TypeOf(key) {
		return
	}
	if reflect.TypeOf(myObj).Elem() != reflect.TypeOf(value) {
		if reflect.TypeOf(myObj).Elem().String() != "interface {}" {
			return
		}
	}
	myMap := reflect.ValueOf(myObj)
	myKey := reflect.ValueOf(key)
	myVal := reflect.ValueOf(value)
	myMap.SetMapIndex(myKey, myVal)
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
	if jo.p == nil {
		return nil
	}
	myObj := *jo.p
	if len(params) == 0 {
		return myObj
	}
	if reflect.ValueOf(myObj).Kind() != reflect.Map {
		return nil
	}
	if reflect.TypeOf(myObj).Key() != reflect.TypeOf(params[0]) {
		return nil
	}
	myMap := reflect.ValueOf(myObj)
	myVal := myMap.MapIndex(reflect.ValueOf(params[0]))
	if !myVal.IsValid() {
		return nil
	}
	return myVal.Interface()
}
