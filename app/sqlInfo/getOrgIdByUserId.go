package sqlInfo

import (
	"database/sql"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
)

type EntityGetOrgIdByUserId struct {
	UserId int           `db:"UID"`
	OrgId  sql.NullInt64 `db:"机构编号"`
}

var ArrNameArgsForQueryGetOrgIdByUserId []string = []string{
	"UserId",
}

var ParamsExcelTitleGetOrgIdByUserId = base.AttrS2{
	"UserId": {
		"field": "UserId",
		"title": "用户编号",
		"sort":  "1",
	},
	"OrgId": {
		"field": "OrgId",
		"title": "所属机构编号",
		"sort":  "2",
	},
}
