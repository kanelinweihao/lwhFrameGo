package dictSQL

import (
	"github.com/kanelinweihao/lwhFrameGo/app/dict/dictSQL/dictSQLSwitch"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
)

func GetArrArgNameBySQLName(sqlName string) (arrArgName []string) {
	arrArgName, _, _, _ = getDictSQLBySQLName(sqlName)
	return arrArgName
}

func GetEntityDBDataBySQLName(sqlName string) (entityDBData base.TEntityDBData) {
	_, entityDBData, _, _ = getDictSQLBySQLName(sqlName)
	return entityDBData
}

func GetFileNamePrefixBySQLName(sqlName string) (fileNamePrefix string) {
	_, _, fileNamePrefix, _ = getDictSQLBySQLName(sqlName)
	return fileNamePrefix
}

func GetParamsExcelTitleBySQLName(sqlName string) (paramsExcelTitle base.AttrS2) {
	_, _, _, paramsExcelTitle = getDictSQLBySQLName(sqlName)
	return paramsExcelTitle
}

func getDictSQLBySQLName(sqlName string) (arrArgName []string, entityDBData base.TEntityDBData, fileNamePrefix string, paramsExcelTitle base.AttrS2) {
	arrArgName, entityDBData, fileNamePrefix, paramsExcelTitle = dictSQLSwitch.GetDictSQL(sqlName)
	return arrArgName, entityDBData, fileNamePrefix, paramsExcelTitle
}
