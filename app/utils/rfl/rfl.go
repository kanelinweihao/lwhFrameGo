package rfl

import (
	"fmt"
	_ "github.com/kanelinweihao/lwhFrameGo/app/utils/dd"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"reflect"
)

func GetTypeInfo(value interface{}) (isPtr bool, typeName string, typeKindName string) {
	t := reflect.TypeOf(value)
	if t == nil {
		msgError := "The value has not been instantiated yet"
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
