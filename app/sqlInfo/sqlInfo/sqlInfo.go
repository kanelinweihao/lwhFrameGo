package sqlInfo

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/sqlInfo/sqlInfoUser"
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
		arrArgName = sqlInfoUser.ArrNameArgsForQueryGetMobileNoByUserId
		entityDBData = new(sqlInfoUser.EntityGetMobileNoByUserId)
		fileNamePrefix = sqlInfoUser.FileNamePrefixGetMobileNoByUserId
		paramsExcelTitle = sqlInfoUser.ParamsExcelTitleGetMobileNoByUserId
	case "GetOrgIdByUserId":
		arrArgName = sqlInfoUser.ArrNameArgsForQueryGetOrgIdByUserId
		entityDBData = new(sqlInfoUser.EntityGetOrgIdByUserId)
		fileNamePrefix = sqlInfoUser.FileNamePrefixGetOrgIdByUserId
		paramsExcelTitle = sqlInfoUser.ParamsExcelTitleGetOrgIdByUserId
	case "GetUserIdByMobileNoAndOrgId":
		arrArgName = sqlInfoUser.ArrNameArgsForQueryGetUserIdByMobileNoAndOrgId
		entityDBData = new(sqlInfoUser.EntityGetUserIdByMobileNoAndOrgId)
		fileNamePrefix = sqlInfoUser.FileNamePrefixGetUserIdByMobileNoAndOrgId
		paramsExcelTitle = sqlInfoUser.ParamsExcelTitleGetUserIdByMobileNoAndOrgId
	default:
		msgError := fmt.Sprintf(
			"The sqlInfo of sqlName |%s| not found",
			sqlName)
		err.ErrPanic(msgError)
	}
	return arrArgName, entityDBData, fileNamePrefix, paramsExcelTitle
}
