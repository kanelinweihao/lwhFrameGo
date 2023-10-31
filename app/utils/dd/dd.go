package dd

import (
	"fmt"
	"os"
	"path"
	"reflect"
	"runtime"
)

func IGNORE(args ...interface{}) {}

func DIE(args ...interface{}) {
	defer os.Exit(0)
}

func DDD(args ...interface{}) {
	defer DIE(args...)
	ddMsgBefore()
	for _, arg := range args {
		ddMsg(arg)
	}
	ddMsgAfter()
}

func DD(arg interface{}) {
	defer os.Exit(0)
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
	typeName := reflect.TypeOf(arg).Name()
	value := reflect.ValueOf(arg)
	fileName, funcName, codeLine := getLocationOfDD()
	ddMsgLine()
	fmt.Print("fileName = ")
	fmt.Println(fileName)
	fmt.Print("funcName = ")
	fmt.Println(funcName)
	fmt.Print("codeLine = ")
	fmt.Println(codeLine)
	if typeName != "" {
		fmt.Print("type = ")
		fmt.Println(typeName)
		fmt.Print("value = ")
		fmt.Println(value)
	} else {
		fmt.Println("value ->")
		fmt.Printf("%#v\n", value)
	}
}

func getLocationOfDD() (fileName string, funcName string, codeLine int) {
	skip := 3
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return
	}
	fileName = path.Base(file)
	funcName = runtime.FuncForPC(pc).Name()
	codeLine = line
	return fileName, funcName, codeLine
}
