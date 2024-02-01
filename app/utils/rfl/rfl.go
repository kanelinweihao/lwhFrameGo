package rfl

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"reflect"
)

func IsPtr(value interface{}) (isPtr bool) {
	t := reflect.TypeOf(value)
	k := t.Kind()
	isPtr = k == reflect.Ptr
	return isPtr
}

func GetTypeInfo(value interface{}) (isPtr bool, typeName string, typeKindName string) {
	t := reflect.TypeOf(value)
	if t == nil {
		msgError := fmt.Sprintf(
			"The value |%v| has not been instantiated yet",
			value)
		err.ErrPanic(msgError)
	}
	isPtr = t.Kind() == reflect.Ptr
	typeName = t.Name()
	typeKindName = t.Kind().String()
	if !isPtr {
		return isPtr, typeName, typeKindName
	}
	tt := t.Elem()
	typeName = "*" + tt.Name()
	typeKindName = "*" + tt.Kind().String()
	return isPtr, typeName, typeKindName
}

func ShowTypeInfo(value interface{}) {
	_, typeName, typeKindName := GetTypeInfo(value)
	fmt.Print("type = ")
	fmt.Println(typeName)
	fmt.Print("kind = ")
	fmt.Println(typeKindName)
	fmt.Print("value = ")
	fmt.Printf("%#v", value)
	fmt.Println()
	fmt.Println()
}

func ErrPanicFormat(value interface{}, valueName string, typeNameValid string) {
	_, typeName, _ := GetTypeInfo(value)
	msgError := fmt.Sprintf(
		"The type of |%s| is |%s|, it should be |%s|",
		valueName,
		typeName,
		typeNameValid)
	err.ErrPanic(msgError)
}
