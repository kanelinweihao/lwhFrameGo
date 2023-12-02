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
	userId := self.getUserIdFromAttrT3DBData()
	paramsAppend := make(base.AttrT1, 1)
	paramsAppend["UserId"] = userId
	self.ParamsAppend = paramsAppend
	return self
}

func (self *EntityRepoDB) getUserIdFromAttrT3DBData() (userId int) {
	attrT3DBData := self.AttrT3DBData
	countUserId := len(attrT3DBData["GetUserIdByMobileNoAndOrgId"])
	if countUserId == 0 {
		return userId
	}
	userId, ok := attrT3DBData["GetUserIdByMobileNoAndOrgId"]["0"]["UserId"].(int)
	if !ok {
		msgError := fmt.Sprintf(
			"The userId |%v| of |%v| is invalid",
			userId,
			attrT3DBData)
		err.ErrPanic(msgError)
	}
	return userId
}
