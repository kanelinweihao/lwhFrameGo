package repoDB

import (
	"github.com/kanelinweihao/lwhFrameGo/app/service/base/repoBase/repoDBBase"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

var ArrSQLName = []string{
	"GetUserIdByMobileNoAndOrgId",
}

func InitEntityRepoDB(paramsRepoDB typeMap.AttrT1) (entityRepoDB typeStruct.EntityRepoDB) {
	entityRepoDB = new(EntityRepoDB)
	entityRepoDBBase := new(repoDBBase.EntityRepoDBBase)
	entityRepoDB.Load(entityRepoDBBase).Init(paramsRepoDB)
	return entityRepoDB
}

func (self *EntityRepoDB) Load(entityRepoDBBase typeStruct.EntityRepoDB) typeStruct.EntityRepoDB {
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
	orgId := conv.ToStr(paramsRepoDB["OrgId"])
	mobileNo := conv.ToStr(paramsRepoDB["MobileNo"])
	attrArgsForQuery = typeMap.AttrS1{
		"OrgId":    orgId,
		"MobileNo": mobileNo,
	}
	return attrArgsForQuery
}
