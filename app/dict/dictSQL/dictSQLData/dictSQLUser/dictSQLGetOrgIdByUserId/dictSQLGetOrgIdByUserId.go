package dictSQLGetOrgIdByUserId

import (
	"database/sql"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

type EntityDBData struct {
	UserId int           `db:"用户编号"`
	OrgId  sql.NullInt64 `db:"机构编号"`
}

var ArrArgName = []string{
	"UserId",
}

var FileNamePrefix string = "数据导出_用户机构编号"

var ParamsExcelTitle = typeMap.AttrS2{
	"UserId": {
		"field": "UserId",
		"title": "用户编号",
		"sort":  "1",
	},
	"OrgId": {
		"field": "OrgId",
		"title": "机构编号",
		"sort":  "2",
	},
}

func GetDictSQL() (arrArgName []string, entityDBData typeStruct.EntityDBData, fileNamePrefix string, paramsExcelTitle typeMap.AttrS2) {
	arrArgName = ArrArgName
	entityDBData = new(EntityDBData)
	fileNamePrefix = FileNamePrefix
	paramsExcelTitle = ParamsExcelTitle
	return arrArgName, entityDBData, fileNamePrefix, paramsExcelTitle
}
