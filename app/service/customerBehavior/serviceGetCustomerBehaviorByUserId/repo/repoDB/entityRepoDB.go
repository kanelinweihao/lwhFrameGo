package repoDB

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

type EntityRepoDB struct {
	typeStruct.EntityRepoDB
}

func (self *EntityRepoDB) GetPropertiesNeed() (paramsAppend typeMap.AttrT1) {
	attrBase := self.EntityRepoDB.ToAttr()
	attrT3DBData := attrBase["AttrT3DBData"].(typeMap.AttrT3)
	countProductOrderPoolNFT := self.getCountProductOrderPoolNFTFromAttrT3DBData(attrT3DBData)
	countProductOrderNFTBuy := self.getCountProductOrderNFTBuyFromAttrT3DBData(attrT3DBData)
	countProductOrderNFTSell := self.getCountProductOrderNFTSellFromAttrT3DBData(attrT3DBData)
	sumAmountGot := self.getSumAmountGotFromAttrT3DBData(attrT3DBData)
	paramsAppend = make(typeMap.AttrT1, 4)
	paramsAppend["CountProductOrderPoolNFT"] = countProductOrderPoolNFT
	paramsAppend["CountProductOrderNFTBuy"] = countProductOrderNFTBuy
	paramsAppend["CountProductOrderNFTSell"] = countProductOrderNFTSell
	paramsAppend["SumAmountGot"] = sumAmountGot
	return paramsAppend
}

func (self *EntityRepoDB) getCountProductOrderPoolNFTFromAttrT3DBData(attrT3DBData typeMap.AttrT3) (countProductOrderPoolNFT int) {
	sqlName := "GetProductOrderPoolNFTOfCustomerBehavior"
	row := len(attrT3DBData[sqlName])
	countProductOrderPoolNFT = row
	return countProductOrderPoolNFT
}

func (self *EntityRepoDB) getCountProductOrderNFTBuyFromAttrT3DBData(attrT3DBData typeMap.AttrT3) (countProductOrderNFTBuy int) {
	sqlName := "GetProductOrderNFTBuyOfCustomerBehavior"
	row := len(attrT3DBData[sqlName])
	countProductOrderNFTBuy = row
	return countProductOrderNFTBuy
}

func (self *EntityRepoDB) getCountProductOrderNFTSellFromAttrT3DBData(attrT3DBData typeMap.AttrT3) (countProductOrderNFTSell int) {
	sqlName := "GetProductOrderNFTSellOfCustomerBehavior"
	row := len(attrT3DBData[sqlName])
	countProductOrderNFTSell = row
	return countProductOrderNFTSell
}

func (self *EntityRepoDB) getSumAmountGotFromAttrT3DBData(attrT3DBData typeMap.AttrT3) (sumAmountGot string) {
	sqlName := "GetUserCrystalReceivingOfCustomerBehavior"
	row := len(attrT3DBData[sqlName])
	if row == 0 {
		return "0"
	}
	sumAmountGot, ok := attrT3DBData[sqlName]["0"]["SumAmountGot"].(string)
	if !ok {
		msgError := fmt.Sprintf(
			"The sumAmountGot |%v| of |%v| is invalid",
			sumAmountGot,
			attrT3DBData)
		err.ErrPanic(msgError)
	}
	return sumAmountGot
}
