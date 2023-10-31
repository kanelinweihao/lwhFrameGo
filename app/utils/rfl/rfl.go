package rfl

import (
	"fmt"
	_ "github.com/kanelinweihao/lwhFrameGo/app/utils/dd"
	"reflect"
)

/*
Type
*/

func getType(value interface{}) (typeOfValue reflect.Type) {
	typeOfValue = reflect.TypeOf(value)
	return typeOfValue
}

func ShowType(value interface{}) {
	typeOfValue := getType(value)
	fmt.Println(typeOfValue)
	return
}

func GetTypeName(value interface{}) (typeName string) {
	typeOfValue := getType(value)
	typeName = typeOfValue.Name()
	return typeName
}

/*
Value
*/

func getValue(value interface{}) (valueOfValue reflect.Value) {
	valueOfValue = reflect.ValueOf(value)
	return valueOfValue
}
