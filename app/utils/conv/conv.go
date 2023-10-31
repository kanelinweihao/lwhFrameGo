package conv

import (
	"encoding/json"
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	_ "github.com/kanelinweihao/lwhFrameGo/app/utils/dd"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/rfl"
	"reflect"
	"strconv"
)

func ToValidType(valueOld interface{}, typeNameNew string) (valueNew interface{}) {
	typeNameOld := rfl.GetTypeName(valueOld)
	if typeNameOld == typeNameNew {
		valueNew = valueOld
		return valueNew
	}
	switch typeNameNew {
	case "bool":
		valueNew = ToBool(valueOld)
	case "int":
		valueNew = ToInt(valueOld)
	case "string":
		valueNew = ToStr(valueOld)
	default:
		errPanicFormat(
			valueOld,
			typeNameOld,
			typeNameNew)
	}
	return valueNew
}

func errPanicFormat(value interface{}, typeNameOld string, typeNameNew string) {
	msgError := fmt.Sprintf(
		"Fail to format value |%v| from typeName |%s| to typeName |%s|",
		value,
		typeNameOld,
		typeNameNew)
	err.ErrPanic(msgError)
}

/*
ToBool
*/

// interface{}->bool
func ToBool(valueOld interface{}) (valueNew bool) {
	typeNameNew := "bool"
	typeNameOld := rfl.GetTypeName(valueOld)
	if typeNameOld == typeNameNew {
		valueNew = valueOld.(bool)
		return valueNew
	}
	switch typeNameOld {
	case "int":
		valueNew = ToBoolFromInt(valueOld.(int))
	case "string":
		valueNew = ToBoolFromStr(valueOld.(string))
	case "bool":
		valueNew = valueOld.(bool)
	default:
		errPanicFormat(
			valueOld,
			typeNameOld,
			typeNameNew)
	}
	return valueNew
}

// int->bool
func ToBoolFromInt(num int) (ok bool) {
	if num == 0 {
		ok = false
	} else {
		ok = true
	}
	return ok
}

// string->bool
func ToBoolFromStr(str string) (ok bool) {
	ok, errParseBool := strconv.ParseBool(str)
	err.ErrCheck(errParseBool)
	return ok
}

/*
ToInt
*/

// interface{}->int
func ToInt(valueOld interface{}) (valueNew int) {
	typeNameNew := "int"
	typeNameOld := rfl.GetTypeName(valueOld)
	if typeNameOld == typeNameNew {
		valueNew = valueOld.(int)
		return valueNew
	}
	switch typeNameOld {
	case "bool":
		valueNew = ToIntFromBool(valueOld.(bool))
	case "string":
		valueNew = ToIntFromStr(valueOld.(string))
	case "int":
		valueNew = valueOld.(int)
	default:
		errPanicFormat(
			valueOld,
			typeNameOld,
			typeNameNew)
	}
	return valueNew
}

// bool->int
func ToIntFromBool(ok bool) (num int) {
	if ok {
		num = 1
	} else {
		num = 0
	}
	return num
}

// string->int
func ToIntFromStr(str string) (num int) {
	num, errAtoi := strconv.Atoi(str)
	err.ErrCheck(errAtoi)
	return num
}

/*
ToStr
*/

// interface{}->string
func ToStr(valueOld interface{}) (valueNew string) {
	typeNameNew := "string"
	typeNameOld := rfl.GetTypeName(valueOld)
	if typeNameOld == typeNameNew {
		valueNew = valueOld.(string)
		return valueNew
	}
	switch typeNameOld {
	case "int":
		valueNew = fmt.Sprintf(
			"%d",
			valueOld)
	case "bool":
		valueNew = fmt.Sprintf(
			"%t",
			valueOld)
	case "string":
		valueNew = fmt.Sprintf(
			"%s",
			valueOld)
	default:
		errPanicFormat(
			valueNew,
			typeNameOld,
			typeNameNew)
	}
	return valueNew
}

// bool->string
func ToStrFromBool(ok bool) (str string) {
	str = strconv.FormatBool(ok)
	return str
}

// int->string
func ToStrFromInt(num int) (str string) {
	str = strconv.Itoa(num)
	return str
}

/*
ToJson
*/

// map->json
func ToJsonFromAttr(attrT1 base.AttrT1) (strJson string) {
	arrByte, errTo := json.Marshal(attrT1)
	err.ErrCheck(errTo)
	strJson = string(arrByte)
	return strJson
}

// map->json
func ToJsonFromEntity[T interface{}](entity *T) (strJson string) {
	attrT1 := ToAttrFromEntity(entity)
	strJson = ToJsonFromAttr(attrT1)
	return strJson
}

/*
ToMap
*/

// map->json
func ToAttrFromJson(strJson string) (attrT1 base.AttrT1) {
	arrByte := []byte(strJson)
	errTo := json.Unmarshal(arrByte, &attrT1)
	err.ErrCheck(errTo)
	return attrT1
}

// struct->map
func ToAttrFromEntity[T interface{}](entity *T) (attrT1 base.AttrT1) {
	attrT1 = make(base.AttrT1)
	entityStruct := *entity
	t := reflect.TypeOf(entityStruct)
	v := reflect.ValueOf(entityStruct)
	for i := 0; i < t.NumField(); i++ {
		key := t.Field(i).Name
		value := v.Field(i).Interface()
		attrT1[key] = value
	}
	return attrT1
}

/*
ToStruct
*/

// map->struct
func ToEntityFromAttr[T interface{}](attrT1 base.AttrT1, entity *T) {
	structValue := reflect.ValueOf(entity).Elem()
	for k, v := range attrT1 {
		structFieldValue := structValue.FieldByName(k)
		typeStructFieldValue := structFieldValue.Type()
		if !structFieldValue.IsValid() {
			continue
		}
		if !structFieldValue.CanSet() {
			msgError := fmt.Sprintf(
				"The field |%s| can not set",
				k)
			err.ErrPanic(msgError)
		}
		val := reflect.ValueOf(v)
		typeVal := val.Type()
		if typeStructFieldValue != typeVal {
			typeNameStruct := typeStructFieldValue.Name()
			vNew := ToValidType(
				v,
				typeNameStruct)
			val = reflect.ValueOf(vNew)
		}
		structFieldValue.Set(val)
	}
	return
}

// []Struct->Map2
func ToAttrT2FromArrEntity[T interface{}](arrEntity []T) (attrT2 base.AttrT2) {
	attrT2 = make(base.AttrT2, len(arrEntity))
	for key, entityStruct := range arrEntity {
		k := ToStrFromInt(key)
		entity := &entityStruct
		attrT1 := ToAttrFromEntity(entity)
		attrT2[k] = attrT1
	}
	return attrT2
}

/*
ToMapNeed
*/

// AttrS1->AttrT1
func ToAttrFromAttrStr(attrS1 base.AttrS1) (attrT1 base.AttrT1) {
	attrT1 = make(base.AttrT1, len(attrS1))
	for key, value := range attrS1 {
		attrT1[key] = value
	}
	return attrT1
}

// AttrT1->AttrS1
func ToAttrStrFromAttr(attrT1 base.AttrT1) (attrS1 base.AttrS1) {
	attrS1 = make(base.AttrS1, len(attrT1))
	for key, value := range attrT1 {
		valueStr := ToStr(value)
		attrS1[key] = valueStr
	}
	return attrS1
}

// ArrAttr->AttrS2
func ToAttrS2FromAttrT2(attrT2 base.AttrT2) (attrS2 base.AttrS2) {
	attrS2 = make(base.AttrS2, len(attrT2))
	for k, attrT1 := range attrT2 {
		attrS := ToAttrStrFromAttr(attrT1)
		attrS2[k] = attrS
	}
	return attrS2
}
