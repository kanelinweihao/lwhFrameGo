package repoDB

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/db"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
)

type EntityRepoDB struct {
	AttrT3DBData     base.AttrT3
	ParamsAppend     base.AttrT1
	ParamsRepoDB     base.AttrT1
	ArrSQLName       []string
	AttrArgsForQuery base.AttrS1
}

func (self *EntityRepoDB) GetDBData() (attrT3DBData base.AttrT3, paramsAppend base.AttrT1) {
	self.setAttrT3DBData().setPropertiesNeed()
	attrT3DBData = self.AttrT3DBData
	paramsAppend = self.ParamsAppend
	return attrT3DBData, paramsAppend
}

func (self *EntityRepoDB) setAttrT3DBData() *EntityRepoDB {
	arrSqlName := self.ArrSQLName
	attrArgsForQuery := self.AttrArgsForQuery
	entityDB := db.InitEntityDB()
	defer entityDB.CloseDB()
	attrT3DBData := entityDB.BatchGetDataFromDB(arrSqlName, attrArgsForQuery)
	self.AttrT3DBData = attrT3DBData
	return self
}

func (self *EntityRepoDB) setPropertiesNeed() *EntityRepoDB {
	mobileNo := self.getMobileNoFromAttrT3DBData()
	orgId := self.getOrgIdFromAttrT3DBData()
	paramsAppend := make(base.AttrT1, 2)
	paramsAppend["MobileNo"] = mobileNo
	paramsAppend["OrgId"] = orgId
	self.ParamsAppend = paramsAppend
	return self
}

func (self *EntityRepoDB) getMobileNoFromAttrT3DBData() (mobileNo string) {
	attrT3DBData := self.AttrT3DBData
	countMobileNo := len(attrT3DBData["GetMobileNoByUserId"])
	if countMobileNo == 0 {
		return mobileNo
	}
	mobileNo, ok := attrT3DBData["GetMobileNoByUserId"]["0"]["MobileNo"].(string)
	if !ok {
		msgError := fmt.Sprintf(
			"The mobileNo |%v| of |%v| is invalid",
			mobileNo,
			attrT3DBData)
		err.ErrPanic(msgError)
	}
	return mobileNo
}

func (self *EntityRepoDB) getOrgIdFromAttrT3DBData() (orgId int) {
	attrT3DBData := self.AttrT3DBData
	countOrgId := len(attrT3DBData["GetOrgIdByUserId"])
	if countOrgId == 0 {
		return orgId
	}
	orgId, ok := attrT3DBData["GetOrgIdByUserId"]["0"]["OrgId"].(int)
	if !ok {
		msgError := fmt.Sprintf(
			"The orgId |%v| of |%v| is invalid",
			orgId,
			attrT3DBData)
		err.ErrPanic(msgError)
	}
	return orgId
}
