package conv

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/rfl"
	"reflect"
	"strconv"
)

func ToValidType(valueOld interface{}, typeNameNew string) (valueNew interface{}) {
	isPtr, typeNameOld, typeKindNameOld := rfl.GetTypeInfo(valueOld)
	if isPtr {
		errPanicIsPtr(
			valueOld,
			typeKindNameOld,
			typeNameOld,
			typeNameNew)
	}
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

func errPanicIsPtr(value interface{}, typeKindName string, typeNameOld string, typeNameNew string) {
	msgError := fmt.Sprintf(
		"Fail to format value |%v|, kind |%s| from typeName |%s| to typeName |%s|",
		value,
		typeKindName,
		typeNameOld,
		typeNameNew)
	err.ErrPanic(msgError)
}

func errPanicFormat(value interface{}, typeNameOld string, typeNameNew string) {
	msgError := fmt.Sprintf(
		"Fail to format value |%v| from typeName |%s| to typeName |%s|",
		value,
		typeNameOld,
		typeNameNew)
	err.ErrPanic(msgError)
}

// ToBool interface{}->bool
func ToBool(valueOld interface{}) (valueNew bool) {
	typeNameNew := "bool"
	isPtr, typeNameOld, typeKindNameOld := rfl.GetTypeInfo(valueOld)
	if isPtr {
		errPanicIsPtr(
			valueOld,
			typeKindNameOld,
			typeNameOld,
			typeNameNew)
	}
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

// ToBoolFromInt int->bool
func ToBoolFromInt(num int) (ok bool) {
	if num == 0 {
		ok = false
	} else {
		ok = true
	}
	return ok
}

// ToBoolFromStr string->bool
func ToBoolFromStr(str string) (ok bool) {
	ok, errParseBool := strconv.ParseBool(str)
	err.ErrCheck(errParseBool)
	return ok
}

// ToInt interface{}->int
func ToInt(valueOld interface{}) (valueNew int) {
	typeNameNew := "int"
	isPtr, typeNameOld, typeKindNameOld := rfl.GetTypeInfo(valueOld)
	if isPtr {
		errPanicIsPtr(
			valueOld,
			typeKindNameOld,
			typeNameOld,
			typeNameNew)
	}
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

// ToIntFromBool bool->int
func ToIntFromBool(ok bool) (num int) {
	if ok {
		num = 1
	} else {
		num = 0
	}
	return num
}

// ToIntFromStr string->int
func ToIntFromStr(str string) (num int) {
	num, errAtoi := strconv.Atoi(str)
	err.ErrCheck(errAtoi)
	return num
}

// ToStr interface{}->string
func ToStr(valueOld interface{}) (valueNew string) {
	typeNameNew := "string"
	isPtr, typeNameOld, typeKindNameOld := rfl.GetTypeInfo(valueOld)
	if isPtr {
		errPanicIsPtr(
			valueOld,
			typeKindNameOld,
			typeNameOld,
			typeNameNew)
	}
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
	case "NullInt64":
		valueInt, _ := valueOld.(sql.NullInt64).Value()
		valueNew = fmt.Sprintf(
			"%d",
			valueInt)
	case "NullString":
		valueStr, _ := valueOld.(sql.NullString).Value()
		valueNew = fmt.Sprintf(
			"%s",
			valueStr)
	default:
		errPanicFormat(
			valueNew,
			typeNameOld,
			typeNameNew)
	}
	return valueNew
}

// ToStrFromBool bool->string
func ToStrFromBool(ok bool) (str string) {
	str = strconv.FormatBool(ok)
	return str
}

// ToStrFromInt int->string
func ToStrFromInt(num int) (str string) {
	str = strconv.Itoa(num)
	return str
}

// ToJsonFromSlice slice->json
func ToJsonFromSlice(arrT []string) (strJson string) {
	arrByte, errTo := json.Marshal(arrT)
	err.ErrCheck(errTo)
	strJson = string(arrByte)
	return strJson
}

// ToJsonFromAttr map->json
func ToJsonFromAttr(attrT1 base.AttrT1) (strJson string) {
	arrByte, errTo := json.Marshal(attrT1)
	err.ErrCheck(errTo)
	strJson = string(arrByte)
	return strJson
}

// ToJsonFromEntity struct->json
func ToJsonFromEntity(entity base.TEntityBase) (strJson string) {
	attrT1 := ToAttrFromEntity(entity)
	strJson = ToJsonFromAttr(attrT1)
	return strJson
}

// ToArrFromAttr map->slice
func ToArrFromAttr(attrT1 base.AttrT1) (arrT []interface{}) {
	len := len(attrT1)
	arrT = make([]interface{}, len)
	for _, v := range attrT1 {
		arrT = append(arrT, v)
	}
	return arrT
}

// ToArrStrFromAttrStr map->slice
func ToArrStrFromAttrStr(attrS1 base.AttrS1) (arrS []string) {
	arrS = make([]string, 0, len(attrS1))
	for _, v := range attrS1 {
		arrS = append(arrS, v)
	}
	return arrS
}

// ToAttrFromJson json->map
func ToAttrFromJson(strJson string) (attrT1 base.AttrT1) {
	arrByte := []byte(strJson)
	errTo := json.Unmarshal(arrByte, &attrT1)
	err.ErrCheck(errTo)
	return attrT1
}

// ToAttrFromEntity struct->map
func ToAttrFromEntity(entity base.TEntityBase) (attrT1 base.AttrT1) {
	attrT1 = make(base.AttrT1)
	t := reflect.TypeOf(entity).Elem()
	v := reflect.ValueOf(entity).Elem()
	for i := 0; i < t.NumField(); i++ {
		key := t.Field(i).Name
		value := v.Field(i).Interface()
		var valueNeed interface{}
		switch value.(type) {
		case sql.NullInt64:
			valueInt64, _ := value.(sql.NullInt64).Value()
			valueNeed = int(valueInt64.(int64))
		case sql.NullString:
			valueStr, _ := value.(sql.NullString).Value()
			valueNeed = valueStr
		default:
			valueNeed = value
		}
		attrT1[key] = valueNeed
	}
	return attrT1
}

// ToEntityFromAttr map->struct
func ToEntityFromAttr(attrT1 base.AttrT1, entity base.TEntityBase) {
	structValue := reflect.ValueOf(entity).Elem()
	for k, v := range attrT1 {
		structFieldValue := structValue.FieldByName(k)
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
		typeStructFieldValue := structFieldValue.Type()
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

// ToAttrFromAttrStr AttrS1->AttrT1
func ToAttrFromAttrStr(attrS1 base.AttrS1) (attrT1 base.AttrT1) {
	attrT1 = make(base.AttrT1, len(attrS1))
	for key, value := range attrS1 {
		attrT1[key] = value
	}
	return attrT1
}

// ToAttrStrFromAttr AttrT1->AttrS1
func ToAttrStrFromAttr(attrT1 base.AttrT1) (attrS1 base.AttrS1) {
	attrS1 = make(base.AttrS1, len(attrT1))
	for key, value := range attrT1 {
		valueStr := ToStr(value)
		attrS1[key] = valueStr
	}
	return attrS1
}

// ToAttrS2FromAttrT2 AttrT2->AttrS2
func ToAttrS2FromAttrT2(attrT2 base.AttrT2) (attrS2 base.AttrS2) {
	attrS2 = make(base.AttrS2, len(attrT2))
	for k, attrT1 := range attrT2 {
		attrS := ToAttrStrFromAttr(attrT1)
		attrS2[k] = attrS
	}
	return attrS2
}
