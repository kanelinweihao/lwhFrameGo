package dd

import (
	"errors"
	"fmt"
	"os"
	"path"
	"reflect"
	"runtime"
	"strconv"
)

func IGNORE(args ...interface{}) {}

func DIE(args ...interface{}) {
	defer os.Exit(0)
}

func DDD(args ...interface{}) {
	defer os.Exit(0)
	defer throwErrorOfDD()
	ddMsgBefore()
	for _, arg := range args {
		ddMsg(arg)
	}
	ddMsgAfter()
}

func DD(arg interface{}) {
	defer os.Exit(0)
	defer throwErrorOfDD()
	ddMsgBefore()
	ddMsg(arg)
	ddMsgAfter()
}

func ddMsgBefore() {
	fmt.Println()
}

func ddMsgAfter() {
	ddMsgLine()
	fmt.Println("\n")
}

func ddMsgLine() {
	fmt.Println("----------")
}

func ddMsg(arg interface{}) {
	ddMsgLine()
	ddMsgSource()
	ddMsgArg(arg)
}

func ddMsgSource() {
	fileName, funcName, codeLine := getLocationOfDD()
	strCodeLine := strconv.Itoa(codeLine)
	msgSource := ""
	msgSource += "fileName = " + fileName + "\n"
	msgSource += "funcName = " + funcName + "\n"
	msgSource += "codeLine = " + strCodeLine + "\n"
	fmt.Print(msgSource)
}

func ddMsgArg(arg interface{}) {
	_, typeName, typeKindName := getTypeInfoOfDD(arg)
	strArg := fmt.Sprintf("%#v", arg)
	msgArg := ""
	msgArg += "type = " + typeName + "\n"
	msgArg += "kind = " + typeKindName + "\n"
	msgArg += "value = " + strArg + "\n"
	fmt.Print(msgArg)
}

// ->err.getLocationOfErr()
func getLocationOfDD() (fileName string, funcName string, codeLine int) {
	skip := 4
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return
	}
	fileName = path.Base(file)
	funcName = runtime.FuncForPC(pc).Name()
	codeLine = line
	return fileName, funcName, codeLine
}

// ->rfl.GetTypeInfo(value interface{})
func getTypeInfoOfDD(value interface{}) (isPtr bool, typeName string, typeKindName string) {
	t := reflect.TypeOf(value)
	if t == nil {
		msgError := "The value has not been instantiated yet"
		errDD := errors.New(msgError)
		panic(errDD)
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

// ->err.ThrowError()
func throwErrorOfDD() {
	err := recover()
	if err == nil {
		return
	}
	fileName, funcName, codeLine := getLocationOfErr()
	fmt.Println("\n")
	fmt.Println("----------")
	fmt.Print("fileName = ")
	fmt.Println(fileName)
	fmt.Print("funcName = ")
	fmt.Println(funcName)
	fmt.Print("codeLine = ")
	fmt.Println(codeLine)
	fmt.Print("msgError = ")
	fmt.Println(err)
	fmt.Println("----------")
	fmt.Println("\n")
	fmt.Println("\n")
}

// ->err.getLocationOfErr()
func getLocationOfErr() (fileName string, funcName string, codeLine int) {
	skip := 8
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return
	}
	fileName = path.Base(file)
	funcName = runtime.FuncForPC(pc).Name()
	codeLine = line
	return fileName, funcName, codeLine
}
