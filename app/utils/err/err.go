package err

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"path"
	"runtime"
)

func ErrCheck(err error) {
	if err == nil {
		return
	}
	panic(err)
}

func ErrNew(msgError string) (err error) {
	err = errors.New(msgError)
	return err
}

func ErrPanic(msgError string) {
	err := ErrNew(msgError)
	panic(err)
}

func ThrowError() {
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

func getLocationOfErr() (fileName string, funcName string, codeLine int) {
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

func ErrValidate(err error) {
	if err == nil {
		return
	}
	for _, err := range err.(validator.ValidationErrors) {
		msgError := fmt.Sprintf(
			"The propertie |%s| of type |%s| should |%s=%s|, but it is |%v|",
			err.Namespace(),
			err.Kind(),
			err.Tag(),
			err.Param(),
			err.Value())
		ErrPanic(msgError)
	}
}
