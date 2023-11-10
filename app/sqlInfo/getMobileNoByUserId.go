package sqlInfo

import (
	"database/sql"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
)

type EntityGetMobileNoByUserId struct {
	UserId   int            `db:"UID"`
	MobileNo sql.NullString `db:"手机号"`
}

var ArrNameArgsForQueryGetMobileNoByUserId []string = []string{
	"UserId",
}

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
