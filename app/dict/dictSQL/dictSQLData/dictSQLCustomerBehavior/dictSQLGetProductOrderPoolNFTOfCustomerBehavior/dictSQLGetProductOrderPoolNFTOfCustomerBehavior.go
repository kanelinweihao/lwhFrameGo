package dictSQLGetProductOrderPoolNFTOfCustomerBehavior

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

type EntityDBData struct {
	BuyAt          string `db:"购买时间(北京时间)"`
	UserId         int    `db:"UID"`
	MobileNo       string `db:"手机号"`
	SkuId          int    `db:"藏品集编号"`
	Title          string `db:"商品名称"`
	PriceUnit      string `db:"价格单位"`
	Price          string `db:"商品价格"`
	ProductOrderNo string `db:"订单号"`
	GoodsNo        string `db:"藏品唯一编号"`
	UserIdHold     int    `db:"藏品当前持有人UID"`
	MsgHold        string `db:"是否持有"`
}

var ArrArgName = []string{
	"UserId",
}

var FileNamePrefix string = "数据导出_客诉_首发留存"

var ParamsExcelTitle = typeMap.AttrS2{
	"BuyAt": {
		"field": "BuyAt",
		"title": "购买时间(北京时间)",
		"sort":  "1",
	},
	"UserId": {
		"field": "UserId",
		"title": "UID",
		"sort":  "2",
	},
	"MobileNo": {
		"field": "MobileNo",
		"title": "手机号",
		"sort":  "3",
	},
	"SkuId": {
		"field": "SkuId",
		"title": "藏品集编号",
		"sort":  "4",
	},
	"Title": {
		"field": "Title",
		"title": "商品名称",
		"sort":  "5",
	},
	"PriceUnit": {
		"field": "PriceUnit",
		"title": "价格单位",
		"sort":  "6",
	},
	"Price": {
		"field": "Price",
		"title": "商品价格",
		"sort":  "7",
	},
	"ProductOrderNo": {
		"field": "ProductOrderNo",
		"title": "订单号",
		"sort":  "8",
	},
	"GoodsNo": {
		"field": "GoodsNo",
		"title": "藏品唯一编号",
		"sort":  "9",
	},
	"UserIdHold": {
		"field": "UserIdHold",
		"title": "藏品当前持有人UID",
		"sort":  "10",
	},
	"MsgHold": {
		"field": "MsgHold",
		"title": "是否持有",
		"sort":  "11",
	},
}

func GetDictSQL() (arrArgName []string, entityDBData typeStruct.EntityDBData, fileNamePrefix string, paramsExcelTitle typeMap.AttrS2) {
	arrArgName = ArrArgName
	entityDBData = new(EntityDBData)
	fileNamePrefix = FileNamePrefix
	paramsExcelTitle = ParamsExcelTitle
	return arrArgName, entityDBData, fileNamePrefix, paramsExcelTitle
}
