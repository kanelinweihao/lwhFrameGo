package dictSQL

import (
	"github.com/kanelinweihao/lwhFrameGo/app/dict/dictSQL/dictSQLSwitch"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
)

func GetArrArgNameBySQLName(sqlName string) (arrArgName []string) {
	arrArgName, _, _, _ = getDictSQLBySQLName(sqlName)
	return arrArgName
}

func GetEntityDBDataBySQLName(sqlName string) (entityDBData typeInterface.EntityDBData) {
	_, entityDBData, _, _ = getDictSQLBySQLName(sqlName)
	return entityDBData
}

func GetFileNamePrefixBySQLName(sqlName string) (fileNamePrefix string) {
	_, _, fileNamePrefix, _ = getDictSQLBySQLName(sqlName)
	return fileNamePrefix
}

func GetParamsExcelTitleBySQLName(sqlName string) (paramsExcelTitle typeMap.AttrS2) {
	_, _, _, paramsExcelTitle = getDictSQLBySQLName(sqlName)
	return paramsExcelTitle
}

func getDictSQLBySQLName(sqlName string) (arrArgName []string, entityDBData typeInterface.EntityDBData, fileNamePrefix string, paramsExcelTitle typeMap.AttrS2) {
	arrArgName, entityDBData, fileNamePrefix, paramsExcelTitle = dictSQLSwitch.GetDictSQL(sqlName)
	return arrArgName, entityDBData, fileNamePrefix, paramsExcelTitle
}
