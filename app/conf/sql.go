package conf

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/pack"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"strings"
)

func GetParamsPathQuery() (paramsPathQuery typeMap.AttrS1) {
	paramsPathQuery = make(typeMap.AttrS1)
	pathSQL := GetPathSQL()
	strFromEnv := pack.ReadFileEmbedAsString(pathSQL)
	arrLine := strings.Split(strFromEnv, "\n")
	strAnnotation := "#"
	strSeparator := "|"
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
		sqlName := arrPart[0]
		pathQueryRel := arrPart[1]
		paramsPathQuery[sqlName] = pathQueryRel
	}
	return paramsPathQuery
}

func GetPathQuery(sqlName string) (pathQuery string) {
	paramsPathQuery := GetParamsPathQuery()
	pathQuery, ok := paramsPathQuery[sqlName]
	if !ok {
		msgError := fmt.Sprintf(
			"The sql |%s| not found in sql.env",
			sqlName)
		err.ErrPanic(msgError)
	}
	if len(pathQuery) == 0 {
		msgError := fmt.Sprintf(
			"The sql of |%s| is empty in sql.env",
			sqlName)
		err.ErrPanic(msgError)
	}
	return pathQuery
}
