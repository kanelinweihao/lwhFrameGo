package conf

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	_ "github.com/kanelinweihao/lwhFrameGo/app/utils/dd"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/pack"
	"strings"
)

func getParamsEnv() (paramsEnv base.AttrS1) {
	paramsEnv = make(base.AttrS1)
	pathEnv := getPathEnv()
	strFromEnv := pack.ReadFileEmbedAsString(pathEnv)
	sliceEnvLine := strings.Split(strFromEnv, "\n")
	strAnnotation := "#"
	strSeparator := "="
	for _, strEnvLine := range sliceEnvLine {
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

func getEntityConfig[T interface{}](paramsKey base.AttrS1, entity *T) {
	paramsNeed := getParamsEnvNeed(paramsKey)
	conv.ToEntityFromAttr(paramsNeed, entity)
	return
}
