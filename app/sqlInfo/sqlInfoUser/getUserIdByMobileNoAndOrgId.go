package sqlInfoUser

import (
	"database/sql"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
)

type EntityGetUserIdByMobileNoAndOrgId struct {
	UserId   int            `db:"用户编号"`
	OrgId    sql.NullInt64  `db:"机构编号"`
	MobileNo sql.NullString `db:"手机号"`
}

var ArrNameArgsForQueryGetUserIdByMobileNoAndOrgId []string = []string{
	"MobileNo",
	"OrgId",
}

var FileNamePrefixGetUserIdByMobileNoAndOrgId string = "数据导出_用户编号"

var ParamsExcelTitleGetUserIdByMobileNoAndOrgId = base.AttrS2{
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
	"MobileNo": {
		"field": "MobileNo",
		"title": "手机号",
		"sort":  "3",
	},
}
