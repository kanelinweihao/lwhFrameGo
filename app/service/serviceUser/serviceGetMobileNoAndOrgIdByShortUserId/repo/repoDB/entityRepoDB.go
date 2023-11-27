package repoDB

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/db"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
)

type EntityRepoDB struct {
	AttrT3DBData     base.AttrT3
	MobileNo         string
	OrgId            int
	ParamsRepoDB     base.AttrT1
	ArrSQLName       []string
	AttrArgsForQuery base.AttrS1
}

func (self *EntityRepoDB) GetDBData() (attrT3DBData base.AttrT3, mobileNo string, orgId int) {
	self.setAttrT3DBData().setPropertiesNeed()
	attrT3DBData, mobileNo, orgId = self.getPropertiesNeed()
	return attrT3DBData, mobileNo, orgId
}

func (self *EntityRepoDB) setAttrT3DBData() *EntityRepoDB {
	arrSqlName := self.ArrSQLName
	attrArgsForQuery := self.AttrArgsForQuery
	entityDB := db.MakeEntityDB()
	defer entityDB.CloseDB()
	attrT3DBData := entityDB.BatchGetDataFromDB(arrSqlName, attrArgsForQuery)
	self.AttrT3DBData = attrT3DBData
	return self
}

func (self *EntityRepoDB) setPropertiesNeed() *EntityRepoDB {
	attrT3DBData := self.AttrT3DBData
	countMobileNo := len(attrT3DBData["GetMobileNoByUserId"])
	if countMobileNo > 0 {
		mobileNo, ok := attrT3DBData["GetMobileNoByUserId"]["0"]["MobileNo"].(string)
		if !ok {
			msgError := fmt.Sprintf(
				"The mobileNo |%v| of |%v| is invalid",
				mobileNo,
				attrT3DBData)
			err.ErrPanic(msgError)
		}
		self.MobileNo = mobileNo
	}
	countOrgId := len(attrT3DBData["GetOrgIdByUserId"])
	if countOrgId > 0 {
		orgId, ok := attrT3DBData["GetOrgIdByUserId"]["0"]["OrgId"].(int)
		if !ok {

			msgError := fmt.Sprintf(
				"The orgId |%v| of |%v| is invalid",
				orgId,
				attrT3DBData)
			err.ErrPanic(msgError)
		}
		self.OrgId = orgId
	}
	return self
}

func (self *EntityRepoDB) getPropertiesNeed() (attrT3DBData base.AttrT3, mobileNo string, orgId int) {
	attrT3DBData = self.AttrT3DBData
	mobileNo = self.MobileNo
	orgId = self.OrgId
	return attrT3DBData, mobileNo, orgId
}
