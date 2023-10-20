package conf

import (
	"fmt"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/base"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/conv"
	_ "go.lwh.com/linweihao/lwhFrameGo/app/utils/dd"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/err"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/file"
	_ "regexp"
	"strings"
)

var PathEnv string = "./res/env/env.env"

func getPathEnv() (pathEnv string) {
	pathEnv = file.GetFilePathAbs(PathEnv)
	// dd.DD(pathEnv)
	return pathEnv
}

func getParamsEnv() (paramsEnv base.AttrS1) {
	paramsEnv = base.AttrS1{}
	pathEnv := getPathEnv()
	strFromEnv := file.ReadFileAsString(pathEnv)
	// dd.DD(strFromEnv)
	sliceEnvLine := strings.Split(strFromEnv, "\n")
	// dd.DD(sliceEnvLine)
	strAnnotation := "#"
	strSeparator := "="
	for _, strEnvLine := range sliceEnvLine {
		// dd.DD(strEnvLine)
		if len(strEnvLine) == 0 {
			continue
		}
		firstLetter := string(strEnvLine[0])
		if firstLetter == strAnnotation {
			continue
		}
		sliceEnvKV := strings.Split(strEnvLine, strSeparator)
		if len(sliceEnvKV) != 2 {
			continue
		}
		paramsEnv[sliceEnvKV[0]] = sliceEnvKV[1]
	}
	return paramsEnv
}

func getEnvValue(envKey string) (envValue string) {
	paramsEnv := getParamsEnv()
	// dd.DD(paramsEnv)
	envValue, ok := paramsEnv[envKey]
	if !ok {
		// envValueDefault := ""
		// return envValueDefault
		msgError := fmt.Sprintf(
			"The |%s| not found in env.env",
			envKey)
		err.ErrPanic(msgError)
	}
	if len(envValue) == 0 {
		msgError := fmt.Sprintf(
			"The value of |%s| is empty in env.env",
			envKey)
		err.ErrPanic(msgError)
	}
	return envValue
}

func getParamsEnvNeed(paramsKey base.AttrS1) (paramsNeed base.AttrT1) {
	paramsNeed = base.AttrT1{}
	paramsEnv := getParamsEnv()
	for needKey, envKey := range paramsKey {
		envValue, ok := paramsEnv[envKey]
		if !ok {
			msgError := fmt.Sprintf(
				"The |%s| not found in env",
				envKey)
			err.ErrPanic(msgError)
		}
		paramsNeed[needKey] = envValue
	}
	return paramsNeed
}

func getEntityConfig[T interface{}](paramsKey base.AttrS1, entity *T) {
	paramsNeed := getParamsEnvNeed(paramsKey)
	// dd.DD(paramsNeed)
	conv.ToEntityFromAttr(paramsNeed, entity)
	// strJson := conv.ToJsonFromEntity(entity)
	// dd.DD(strJson)
	return
}
