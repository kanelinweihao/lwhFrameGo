package dataGet

import (
	"fmt"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/base"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/calc"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/conv"
	_ "go.lwh.com/linweihao/lwhFrameGo/app/utils/dd"
)

type EntityDataGet struct {
	Field1  string
	Field2  string
	Field3  string
	UID     string
	MsgOut  string
	BoxData base.AttrS3
}

////
// SET
////

func (self *EntityDataGet) Exec() {
	self.Init()
	self.GetDataFromDB()
	self.SetMsgOut()
	return
}

func (self *EntityDataGet) Init() {
	self.UID = calc.Add(self.Field2, "1000000")
	self.Field3 = "empty"
	return
}

func (self *EntityDataGet) SetMsgOut() {
	MsgOut := fmt.Sprintf(
		"\n%s%s\n",
		self.Field1,
		self.UID)
	self.MsgOut = MsgOut
	return
}

////
// SET
////

func (self *EntityDataGet) GetBoxData() (boxData base.AttrS3) {
	boxData = self.BoxData
	// dd.DD(boxData)
	return boxData
}

func (self *EntityDataGet) GetParamsOut() (paramsOut base.AttrT1) {
	paramsOut = conv.ToAttrFromEntity(self)
	// dd.DD(paramsOut)
	return paramsOut
}
