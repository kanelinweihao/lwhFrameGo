package dictSQLGetUserCrystalReceivingOfCustomerBehavior

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

type EntityDBData struct {
	UserId          string `db:"UID"`
	SumAmountCanGet string `db:"可领取水晶总数"`
	SumAmountGot    string `db:"已领取水晶总数"`
}

var ArrArgName = []string{
	"UserId",
}

var FileNamePrefix string = "数据导出_客诉_水晶领取"

var ParamsExcelTitle = typeMap.AttrS2{
	"UserId": {
		"field": "UserId",
		"title": "UID",
		"sort":  "1",
	},
	"SumAmountCanGet": {
		"field": "SumAmountCanGet",
		"title": "可领取水晶总数",
		"sort":  "2",
	},
	"SumAmountGot": {
		"field": "SumAmountGot",
		"title": "已领取水晶总数",
		"sort":  "3",
	},
}

func GetDictSQL() (arrArgName []string, entityDBData typeStruct.EntityDBData, fileNamePrefix string, paramsExcelTitle typeMap.AttrS2) {
	arrArgName = ArrArgName
	entityDBData = new(EntityDBData)
	fileNamePrefix = FileNamePrefix
	paramsExcelTitle = ParamsExcelTitle
	return arrArgName, entityDBData, fileNamePrefix, paramsExcelTitle
}
