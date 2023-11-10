package dataGet

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/db"
)

type EntityDataGet struct {
	ParamsOut        base.AttrT1
	Sign             string
	ShortUserId      string
	UserId           string
	ArrSQLName       []string
	AttrArgsForQuery base.AttrS1
	BoxData          base.AttrS3
	MsgOut           string
}

func (self *EntityDataGet) GetData() (boxData base.AttrS3, paramsOut base.AttrT1) {
	self.getDataFromDB().setMsgOut().setParamsOut()
	boxData = self.BoxData
	paramsOut = self.ParamsOut
	return boxData, paramsOut
}

func (self *EntityDataGet) getDataFromDB() *EntityDataGet {
	entityDB := db.MakeEntityOfDB()
	defer entityDB.CloseDB()
	boxData := entityDB.GetDBData(
		self.ArrSQLName,
		self.AttrArgsForQuery)
	self.BoxData = boxData
	return self
}

func (self *EntityDataGet) setMsgOut() *EntityDataGet {
	msgOut := fmt.Sprintf(
		"\n%s\n%s\n",
		self.UserId,
		"Success")
	self.MsgOut = msgOut
	return self
}

func (self *EntityDataGet) setParamsOut() *EntityDataGet {
	paramsOut := make(base.AttrT1)
	paramsOut = base.AttrT1{
		"Sign":        self.Sign,
		"ShortUserId": self.ShortUserId,
		"UserId":      self.UserId,
		"MsgOut":      self.MsgOut,
	}
	self.ParamsOut = paramsOut
	return self
}
