package dictSQLSwitch

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/dict/dictSQL/dictSQLData/dictSQLCustomerBehavior/dictSQLGetProductOrderNFTBuyOfCustomerBehavior"
	"github.com/kanelinweihao/lwhFrameGo/app/dict/dictSQL/dictSQLData/dictSQLCustomerBehavior/dictSQLGetProductOrderNFTSellOfCustomerBehavior"
	"github.com/kanelinweihao/lwhFrameGo/app/dict/dictSQL/dictSQLData/dictSQLCustomerBehavior/dictSQLGetProductOrderPoolNFTOfCustomerBehavior"
	"github.com/kanelinweihao/lwhFrameGo/app/dict/dictSQL/dictSQLData/dictSQLCustomerBehavior/dictSQLGetUserCrystalReceivingOfCustomerBehavior"
	"github.com/kanelinweihao/lwhFrameGo/app/dict/dictSQL/dictSQLData/dictSQLUser/dictSQLGetMobileNoByUserId"
	"github.com/kanelinweihao/lwhFrameGo/app/dict/dictSQL/dictSQLData/dictSQLUser/dictSQLGetOrgIdByUserId"
	"github.com/kanelinweihao/lwhFrameGo/app/dict/dictSQL/dictSQLData/dictSQLUser/dictSQLGetUserIdByMobileNoAndOrgId"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
)

func GetDictSQL(sqlName string) (arrArgName []string, entityDBData typeInterface.EntityDBData, fileNamePrefix string, paramsExcelTitle typeMap.AttrS2) {
	switch sqlName {
	case "GetMobileNoByUserId":
		arrArgName, entityDBData, fileNamePrefix, paramsExcelTitle = dictSQLGetMobileNoByUserId.GetDictSQL()
	case "GetOrgIdByUserId":
		arrArgName, entityDBData, fileNamePrefix, paramsExcelTitle = dictSQLGetOrgIdByUserId.GetDictSQL()
	case "GetUserIdByMobileNoAndOrgId":
		arrArgName, entityDBData, fileNamePrefix, paramsExcelTitle = dictSQLGetUserIdByMobileNoAndOrgId.GetDictSQL()
	case "GetProductOrderPoolNFTOfCustomerBehavior":
		arrArgName, entityDBData, fileNamePrefix, paramsExcelTitle = dictSQLGetProductOrderPoolNFTOfCustomerBehavior.GetDictSQL()
	case "GetProductOrderNFTBuyOfCustomerBehavior":
		arrArgName, entityDBData, fileNamePrefix, paramsExcelTitle = dictSQLGetProductOrderNFTBuyOfCustomerBehavior.GetDictSQL()
	case "GetProductOrderNFTSellOfCustomerBehavior":
		arrArgName, entityDBData, fileNamePrefix, paramsExcelTitle = dictSQLGetProductOrderNFTSellOfCustomerBehavior.GetDictSQL()
	case "GetUserCrystalReceivingOfCustomerBehavior":
		arrArgName, entityDBData, fileNamePrefix, paramsExcelTitle = dictSQLGetUserCrystalReceivingOfCustomerBehavior.GetDictSQL()
	default:
		msgError := fmt.Sprintf(
			"The dictSQL of sqlName |%s| not found",
			sqlName)
		err.ErrPanic(msgError)
	}
	return arrArgName, entityDBData, fileNamePrefix, paramsExcelTitle
}
