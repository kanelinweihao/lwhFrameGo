package dataGet

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/db"
	_ "github.com/kanelinweihao/lwhFrameGo/app/utils/dd"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/goroutine"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/rfl"
)

type EntityDataGet struct {
	Field1  string
	Field2  string
	Field3  string
	UID     string
	MsgOut  string
	BoxData base.AttrS3
}

func (self *EntityDataGet) GetData() (boxData base.AttrS3, paramsOut base.AttrT1) {
	self.setMsgOut().getDataFromDB()
	boxData = self.BoxData
	paramsOut = conv.ToAttrFromEntity(self)
	return boxData, paramsOut
}

func (self *EntityDataGet) setMsgOut() *EntityDataGet {
	MsgOut := fmt.Sprintf(
		"\n%s%s\n",
		self.Field1,
		self.UID)
	self.MsgOut = MsgOut
	return self
}

func (self *EntityDataGet) getDataFromDB() {
	entityDB := db.MakeEntityOfDB()
	defer entityDB.CloseDB()
	// init
	boxData := make(base.AttrS3)
	userId := self.UID
	// write
	var arrEntityUID []EntityUID
	channelReadUID := getAttrS2ExcelDataOfWrite(
		entityDB,
		arrEntityUID,
		queryUID,
		userId)
	var arrEntityOID []EntityOID
	channelReadOID := getAttrS2ExcelDataOfWrite(
		entityDB,
		arrEntityOID,
		queryOID,
		userId)
	// read
	attrS2ExcelDataUID := getAttrS2ExcelDataOfRead(channelReadUID)
	attrS2ExcelDataOID := getAttrS2ExcelDataOfRead(channelReadOID)
	boxData["UID"] = attrS2ExcelDataUID
	boxData["OID"] = attrS2ExcelDataOID
	self.BoxData = boxData
	return
}

func getAttrS2ExcelDataOfWrite[T TypeEntityData](entityDB *db.EntityDB, arrEntity []T, query string, userId string) (entityChannel *goroutine.EntityChannel) {
	entityChannel = goroutine.MakeEntityOfGoroutine()
	go db.GetArrAttrForExcelUseGoroutine(
		entityChannel,
		entityDB,
		arrEntity,
		query,
		userId)
	return entityChannel
}

func getAttrS2ExcelDataOfRead(entityChannel *goroutine.EntityChannel) (attrS2ExcelData base.AttrS2) {
	dataOnce := entityChannel.ReadOnce()
	attrS2ExcelData, ok := dataOnce.(base.AttrS2)
	if !ok {
		typeInvalid := rfl.GetTypeName(attrS2ExcelData)
		msgError := fmt.Sprintf(
			"The type of |%s| is |%s|, it should be |%s|",
			"attrS2ExcelData",
			typeInvalid,
			"base.AttrS2")
		err.ErrPanic(msgError)
	}
	return attrS2ExcelData
}
