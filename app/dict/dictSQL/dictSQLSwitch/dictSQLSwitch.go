package dictSQLSwitch

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/dict/dictSQL/dictSQLData/dictSQLUser/dictSQLGetMobileNoByUserId"
	"github.com/kanelinweihao/lwhFrameGo/app/dict/dictSQL/dictSQLData/dictSQLUser/dictSQLGetOrgIdByUserId"
	"github.com/kanelinweihao/lwhFrameGo/app/dict/dictSQL/dictSQLData/dictSQLUser/dictSQLGetUserIdByMobileNoAndOrgId"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
)

func GetDictSQL(sqlName string) (arrArgName []string, entityDBData base.TEntityDBData, fileNamePrefix string, paramsExcelTitle base.AttrS2) {
	switch sqlName {
	case "GetMobileNoByUserId":
		arrArgName, entityDBData, fileNamePrefix, paramsExcelTitle = dictSQLGetMobileNoByUserId.GetDictSQL()
	case "GetOrgIdByUserId":
		arrArgName, entityDBData, fileNamePrefix, paramsExcelTitle = dictSQLGetOrgIdByUserId.GetDictSQL()
	case "GetUserIdByMobileNoAndOrgId":
		arrArgName, entityDBData, fileNamePrefix, paramsExcelTitle = dictSQLGetUserIdByMobileNoAndOrgId.GetDictSQL()
	default:
		msgError := fmt.Sprintf(
			"The dictSQL of sqlName |%s| not found",
			sqlName)
		err.ErrPanic(msgError)
	}
	return arrArgName, entityDBData, fileNamePrefix, paramsExcelTitle
}
