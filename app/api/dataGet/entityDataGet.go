package dataGet

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/db"
)

var arrSQLName = []string{
	"GetMobileNoByUserId",
	"GetOrgIdByUserId",
}

type EntityDataGet struct {
	BoxFromDB        base.AttrS3
	ParamsOut        base.AttrT1
	ArrSQLName       []string
	AttrArgsForQuery base.AttrS1
	BoxForDB         base.BoxData
}

func (self *EntityDataGet) GetData() (boxFromDB base.AttrS3) {
	boxForDB := self.BoxForDB
	entityDB := db.MakeEntityDB()
	defer entityDB.CloseDB()
	boxFromDB = entityDB.BatchGetDataFromDB(boxForDB)
	self.BoxFromDB = boxFromDB
	return boxFromDB
}
