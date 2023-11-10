package sqlInfo

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
)

func GetArrArgNameBySQLName(sqlName string) (arrArgName []string) {
	arrArgName, _, _ = GetSQLInfoBySQLName(sqlName)
	return arrArgName
}

func GetEntityDBDataBySQLName(sqlName string) (entityDBData base.EntityDBData) {
	_, entityDBData, _ = GetSQLInfoBySQLName(sqlName)
	return entityDBData
}

func GetParamsExcelTitleBySQLName(sqlName string) (paramsExcelTitle base.AttrS2) {
	_, _, paramsExcelTitle = GetSQLInfoBySQLName(sqlName)
	return paramsExcelTitle
}

func GetSQLInfoBySQLName(sqlName string) (arrArgName []string, entityDBData base.EntityDBData, paramsExcelTitle base.AttrS2) {
	switch sqlName {
	case "GetMobileNoByUserId":
		arrArgName = ArrNameArgsForQueryGetMobileNoByUserId
		entityDBData = new(EntityGetMobileNoByUserId)
		paramsExcelTitle = ParamsExcelTitleGetMobileNoByUserId
	case "GetOrgIdByUserId":
		arrArgName = ArrNameArgsForQueryGetOrgIdByUserId
		entityDBData = new(EntityGetOrgIdByUserId)
		paramsExcelTitle = ParamsExcelTitleGetOrgIdByUserId
	default:
		msgError := fmt.Sprintf(
			"The sqlInfo of sqlName |%s| not found",
			sqlName)
		err.ErrPanic(msgError)
	}
	return arrArgName, entityDBData, paramsExcelTitle
}
