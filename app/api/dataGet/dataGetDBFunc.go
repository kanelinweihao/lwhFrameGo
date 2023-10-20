package dataGet

import (
	"fmt"
	"go.lwh.com/linweihao/lwhFrameGo/app/factory/factoryOfDB"
	"go.lwh.com/linweihao/lwhFrameGo/app/factory/factoryOfGoroutine"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/base"
	_ "go.lwh.com/linweihao/lwhFrameGo/app/utils/conv"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/db"
	_ "go.lwh.com/linweihao/lwhFrameGo/app/utils/dd"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/err"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/goroutine"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/rfl"
)

func (self *EntityDataGet) GetDataFromDB() {
	// entityDB := db.InitDB()
	entityDB := factoryOfDB.MakeEntityOfDB()
	defer entityDB.CloseDB()
	// init
	boxData := base.AttrS3{}
	// userId := conv.ToIntFromStr(self.UID)
	userId := self.UID
	// write
	var arrEntityUID []EntityUID
	// dd.DD(arrEntityUID)
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
	// entityChannel = goroutine.InitEntityChannel()
	entityChannel = factoryOfGoroutine.MakeEntityOfGoroutine()
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
