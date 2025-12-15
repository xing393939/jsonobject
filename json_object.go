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
	myObj := jo.GetInterface()
	if myMap, ok := myObj.(map[string]interface{}); ok {
		myMap[key] = value
		return true
	}
	return false
}

func (jo *JsonObject) Delete(key string) bool {
	myObj := jo.GetInterface()
	if myMap, ok := myObj.(map[string]interface{}); ok {
		delete(myMap, key)
		return true
	}
	return false
}

func (jo *JsonObject) GetString(params ...string) string {
	myObj := jo.GetInterface(params...)
	myStr := ""
	switch myObj.(type) {
	case string:
		myStr = myObj.(string)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		myStr = fmt.Sprint(myObj)
	}
	return myStr
}

func (jo *JsonObject) GetInt(params ...string) int {
	myObj := jo.GetInterface(params...)
	myNumber := getNumber[int64](myObj)
	return int(myNumber)
}

func (jo *JsonObject) GetInt64(params ...string) int64 {
	myObj := jo.GetInterface(params...)
	myNumber := getNumber[int64](myObj)
	return myNumber
}

func (jo *JsonObject) GetFloat64(params ...string) float64 {
	myObj := jo.GetInterface(params...)
	myNumber := getNumber[float64](myObj)
	return myNumber
}

func (jo *JsonObject) GetBool(params ...string) bool {
	myObj := jo.GetInterface(params...)
	myBool, ok := myObj.(bool)
	if !ok {
		return false
	}
	return myBool
}

func (jo *JsonObject) GetStringMap(params ...string) map[string]interface{} {
	myObj := jo.GetInterface(params...)
	myMap, ok := myObj.(map[string]interface{})
	if !ok {
		myMap = make(map[string]interface{})
	}
	return myMap
}

func (jo *JsonObject) GetStringSlice(params ...string) []string {
	input := jo.GetInterfaceSlice(params...)
	output := make([]string, 0, len(input))
	for i := range input {
		stringValue, isString := input[i].(string)
		if isString {
			output = append(output, stringValue)
		}
	}
	return output
}

func (jo *JsonObject) GetInterfaceSlice(params ...string) []interface{} {
	myObj := jo.GetInterface(params...)
	mySli, _ := myObj.([]interface{})
	return mySli
}

func (jo *JsonObject) GetJsonObjectSlice(params ...string) []*JsonObject {
	myObj := jo.GetInterface(params...)
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
	myObj := jo.GetInterface(params...)
	newJo := &JsonObject{
		&myObj,
	}
	return newJo
}

func (jo *JsonObject) GetInterface(params ...string) interface{} {
	if jo == nil || jo.p == nil {
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
	return jo.GetInterface(params...) == nil
}

func (jo *JsonObject) Marshal(params ...string) string {
	myObj := jo.GetInterface(params...)
	bytes, err := json.Marshal(myObj)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (jo *JsonObject) MarshalJSON() ([]byte, error) {
	myObj := jo.GetInterface()
	bytes, err := json.Marshal(myObj)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func getNumber[T int64 | uint64 | float64](src any) (dist T) {
	switch src.(type) {
	case int:
		dist = (T)(src.(int))
	case int8:
		dist = (T)(src.(int8))
	case int16:
		dist = (T)(src.(int16))
	case int32:
		dist = (T)(src.(int32))
	case int64:
		dist = (T)(src.(int64))
	case uint:
		dist = (T)(src.(uint))
	case uint8:
		dist = (T)(src.(uint8))
	case uint16:
		dist = (T)(src.(uint16))
	case uint32:
		dist = (T)(src.(uint32))
	case uint64:
		dist = (T)(src.(uint64))
	case float32:
		dist = (T)(src.(float32))
	case float64:
		dist = (T)(src.(float64))
	case string:
		temp := src.(string)
		switch any(dist).(type) {
		case int64:
			v, _ := strconv.ParseInt(temp, 10, 64)
			dist = T(v)
		case uint64:
			v, _ := strconv.ParseUint(temp, 10, 64)
			dist = T(v)
		case float64:
			v, _ := strconv.ParseFloat(temp, 64)
			dist = T(v)
		}
	}
	return
}
