package repoDB

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/db"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
)

type EntityRepoDB struct {
	AttrT3DBData     base.AttrT3
	UserId           int
	ParamsRepoDB     base.AttrT1
	ArrSQLName       []string
	AttrArgsForQuery base.AttrS1
}

func (self *EntityRepoDB) GetDBData() (attrT3DBData base.AttrT3, userId int) {
	self.setAttrT3DBData().setPropertiesNeed()
	attrT3DBData, userId = self.getPropertiesNeed()
	return attrT3DBData, userId
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
	countUserId := len(attrT3DBData["GetUserIdByMobileNoAndOrgId"])
	if countUserId > 0 {
		userId, ok := attrT3DBData["GetUserIdByMobileNoAndOrgId"]["0"]["UserId"].(int)
		if !ok {
			msgError := fmt.Sprintf(
				"The userId |%v| of |%v| is invalid",
				userId,
				attrT3DBData)
			err.ErrPanic(msgError)
		}
		self.UserId = userId
	}
	return self
}

func (self *EntityRepoDB) getPropertiesNeed() (attrT3DBData base.AttrT3, userId int) {
	attrT3DBData = self.AttrT3DBData
	userId = self.UserId
	return attrT3DBData, userId
}
