package conf

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/pack"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"strings"
)

func GetParamsDirSQL() (paramsDirSQL typeMap.AttrS1) {
	paramsDirSQL = make(typeMap.AttrS1)
	pathSQL := getPathSQL()
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
		pathDirSQLRel := arrPart[1]
		paramsDirSQL[sqlName] = pathDirSQLRel
	}
	return paramsDirSQL
}

func GetPathDirSQL(sqlName string) (pathDirSQL string) {
	paramsDirSQL := GetParamsDirSQL()
	pathDirSQL, ok := paramsDirSQL[sqlName]
	if !ok {
		msgError := fmt.Sprintf(
			"The sql |%s| not found in sql.env",
			sqlName)
		err.ErrPanic(msgError)
	}
	if len(pathDirSQL) == 0 {
		msgError := fmt.Sprintf(
			"The sql of |%s| is empty in sql.env",
			sqlName)
		err.ErrPanic(msgError)
	}
	return pathDirSQL
}
