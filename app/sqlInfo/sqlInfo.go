package sqlInfo

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
)

func GetArrArgNameBySQLName(sqlName string) (arrArgName []string) {
	arrArgName, _, _, _ = GetSQLInfoBySQLName(sqlName)
	return arrArgName
}

func GetEntityDBDataBySQLName(sqlName string) (entityDBData base.EntityDBData) {
	_, entityDBData, _, _ = GetSQLInfoBySQLName(sqlName)
	return entityDBData
}

func GetFileNamePrefixBySQLName(sqlName string) (fileNamePrefix string) {
	_, _, fileNamePrefix, _ = GetSQLInfoBySQLName(sqlName)
	return fileNamePrefix
}

func GetParamsExcelTitleBySQLName(sqlName string) (paramsExcelTitle base.AttrS2) {
	_, _, _, paramsExcelTitle = GetSQLInfoBySQLName(sqlName)
	return paramsExcelTitle
}

func GetSQLInfoBySQLName(sqlName string) (arrArgName []string, entityDBData base.EntityDBData, fileNamePrefix string, paramsExcelTitle base.AttrS2) {
	switch sqlName {
	case "GetMobileNoByUserId":
		arrArgName = ArrNameArgsForQueryGetMobileNoByUserId
		entityDBData = new(EntityGetMobileNoByUserId)
		fileNamePrefix = FileNamePrefixGetMobileNoByUserId
		paramsExcelTitle = ParamsExcelTitleGetMobileNoByUserId
	case "GetOrgIdByUserId":
		arrArgName = ArrNameArgsForQueryGetOrgIdByUserId
		entityDBData = new(EntityGetOrgIdByUserId)
		fileNamePrefix = FileNamePrefixGetOrgIdByUserId
		paramsExcelTitle = ParamsExcelTitleGetOrgIdByUserId
	default:
		msgError := fmt.Sprintf(
			"The sqlInfo of sqlName |%s| not found",
			sqlName)
		err.ErrPanic(msgError)
	}
	return arrArgName, entityDBData, fileNamePrefix, paramsExcelTitle
}
