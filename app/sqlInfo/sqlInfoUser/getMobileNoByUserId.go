package sqlInfoUser

import (
	"database/sql"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
)

type EntityGetMobileNoByUserId struct {
	UserId   int            `db:"用户编号"`
	MobileNo sql.NullString `db:"手机号"`
}

var ArrNameArgsForQueryGetMobileNoByUserId []string = []string{
	"UserId",
}

var FileNamePrefixGetMobileNoByUserId string = "数据导出_用户手机号"

var ParamsExcelTitleGetMobileNoByUserId = base.AttrS2{
	"UserId": {
		"field": "UserId",
		"title": "用户编号",
		"sort":  "1",
	},
	"MobileNo": {
		"field": "MobileNo",
		"title": "手机号",
		"sort":  "2",
	},
}
