package conf

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/pack"
	"strings"
)

func getParamsEnv() (paramsEnv base.AttrS1) {
	paramsEnv = make(base.AttrS1)
	pathEnv := getPathEnv()
	strFromEnv := pack.ReadFileEmbedAsString(pathEnv)
	arrLine := strings.Split(strFromEnv, "\n")
	strAnnotation := "#"
	strSeparator := "="
	for _, strLine := range arrLine {
		if len(strLine) == 0 {
			continue
		}
		firstLetter := string(strLine[0])
		if firstLetter == strAnnotation {
			continue
		}
		arrPart := strings.Split(strLine, strSeparator)
		if len(arrPart) != 2 {
			continue
		}
		envKey := arrPart[0]
		envValue := arrPart[1]
		paramsEnv[envKey] = envValue
	}
	return paramsEnv
}

func getEnvValue(envKey string) (envValue string) {
	paramsEnv := getParamsEnv()
	envValue, ok := paramsEnv[envKey]
	if !ok {
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
	paramsNeed = make(base.AttrT1)
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

func getEntityConfig(paramsKey base.AttrS1, entity base.TEntityBase) {
	paramsNeed := getParamsEnvNeed(paramsKey)
	conv.ToEntityFromAttr(paramsNeed, entity)
	return
}
