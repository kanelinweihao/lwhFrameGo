package sqlInfoUser

import (
	"database/sql"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
)

type EntityGetOrgIdByUserId struct {
	UserId int           `db:"用户编号"`
	OrgId  sql.NullInt64 `db:"机构编号"`
}

var ArrNameArgsForQueryGetOrgIdByUserId []string = []string{
	"UserId",
}

var FileNamePrefixGetOrgIdByUserId string = "数据导出_用户机构编号"

var ParamsExcelTitleGetOrgIdByUserId = base.AttrS2{
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
