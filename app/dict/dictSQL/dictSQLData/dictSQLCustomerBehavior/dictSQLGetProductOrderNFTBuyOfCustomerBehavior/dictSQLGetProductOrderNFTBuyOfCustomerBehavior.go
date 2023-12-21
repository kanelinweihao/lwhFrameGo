package dictSQLGetProductOrderNFTBuyOfCustomerBehavior

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

type EntityDBData struct {
	BuyAt          string `db:"成交时间(北京时间)"`
	UserIdSell     string `db:"卖方UID"`
	MobileNoSell   string `db:"卖方手机号"`
	UserIdBuy      string `db:"买方UID"`
	MobileNoBuy    string `db:"买方手机号"`
	SkuId          string `db:"藏品集编号"`
	Title          string `db:"藏品名称"`
	CodeId         string `db:"铸造编号"`
	Price          string `db:"成交价格"`
	ProductOrderNo string `db:"订单单号"`
}

var ArrArgName = []string{
	"UserId",
}

var FileNamePrefix string = "数据导出_客诉_寄售买入"

var ParamsExcelTitle = typeMap.AttrS2{
	"BuyAt": {
		"field": "BuyAt",
		"title": "成交时间(北京时间)",
		"sort":  "1",
	},
	"UserIdSell": {
		"field": "UserIdSell",
		"title": "卖方UID",
		"sort":  "2",
	},
	"MobileNoSell": {
		"field": "MobileNoSell",
		"title": "卖方手机号",
		"sort":  "3",
	},
	"UserIdBuy": {
		"field": "UserIdBuy",
		"title": "买方UID",
		"sort":  "4",
	},
	"MobileNoBuy": {
		"field": "MobileNoBuy",
		"title": "买方手机号",
		"sort":  "5",
	},
	"SkuId": {
		"field": "SkuId",
		"title": "藏品集编号",
		"sort":  "6",
	},
	"Title": {
		"field": "Title",
		"title": "藏品名称",
		"sort":  "7",
	},
	"CodeId": {
		"field": "CodeId",
		"title": "铸造编号",
		"sort":  "8",
	},
	"Price": {
		"field": "Price",
		"title": "成交价格",
		"sort":  "9",
	},
	"ProductOrderNo": {
		"field": "ProductOrderNo",
		"title": "订单单号",
		"sort":  "10",
	},
}

func GetDictSQL() (arrArgName []string, entityDBData typeStruct.EntityDBData, fileNamePrefix string, paramsExcelTitle typeMap.AttrS2) {
	arrArgName = ArrArgName
	entityDBData = new(EntityDBData)
	fileNamePrefix = FileNamePrefix
	paramsExcelTitle = ParamsExcelTitle
	return arrArgName, entityDBData, fileNamePrefix, paramsExcelTitle
}
