package repoDB

import (
	"github.com/kanelinweihao/lwhFrameGo/app/service/base/repoBase/repoDBBase"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
)

var ArrSQLName = []string{
	"GetProductOrderPoolNFTOfCustomerBehavior",
	"GetProductOrderNFTBuyOfCustomerBehavior",
	"GetProductOrderNFTSellOfCustomerBehavior",
	"GetUserCrystalReceivingOfCustomerBehavior",
}

func InitEntityRepoDB(paramsRepoDB typeMap.AttrT1) (entityRepoDB typeInterface.EntityRepoDB) {
	entityRepoDB = new(EntityRepoDB)
	entityRepoDBBase := new(repoDBBase.EntityRepoDBBase)
	entityRepoDB.Load(entityRepoDBBase).Init(paramsRepoDB)
	return entityRepoDB
}

func (self *EntityRepoDB) Load(entityRepoDBBase typeInterface.EntityRepoDB) typeInterface.EntityRepoDB {
	self.EntityRepoDB = entityRepoDBBase
	entityRepoDBBase.Load(self)
	return self
}

func (self *EntityRepoDB) GetArrSQLName() (arrSQLName []string) {
	arrSQLName = ArrSQLName
	return arrSQLName
}

func (self *EntityRepoDB) GetAttrArgsForQuery() (attrArgsForQuery typeMap.AttrS1) {
	attrBase := self.EntityRepoDB.ToAttr()
	paramsRepoDB := attrBase["ParamsRepoDB"].(typeMap.AttrT1)
	userId := paramsRepoDB["UserId"].(int)
	strUserId := conv.ToStrFromInt(userId)
	attrArgsForQuery = typeMap.AttrS1{
		"UserId": strUserId,
	}
	return attrArgsForQuery
}
