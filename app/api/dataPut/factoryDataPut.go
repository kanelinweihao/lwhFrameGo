package dataPut

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	_ "github.com/kanelinweihao/lwhFrameGo/app/utils/dd"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/file"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/time"
)

var ext = "xlsx"
var paramsFileNamePrefix = base.AttrS1{
	"UID": "示例_用户手机号",
	"OID": "示例_用户所属机构",
}
var paramsExcelTitleUID = base.AttrS2{
	"UserId": {
		"field": "UserId",
		"title": "UID",
		"sort":  "1",
	},
	"MobileNo": {
		"field": "MobileNo",
		"title": "手机号",
		"sort":  "2",
	},
}
var paramsExcelTitleOID = base.AttrS2{
	"UserId": {
		"field": "UserId",
		"title": "UID",
		"sort":  "1",
	},
	"OrgId": {
		"field": "OrgId",
		"title": "所属机构编号",
		"sort":  "2",
	},
}

func MakeEntityOfDataPut(paramsOut base.AttrT1, boxExcelData base.AttrS3) (entityDataPut *EntityDataPut) {
	entityDataPut = initDataPut(paramsOut, boxExcelData)
	return entityDataPut
}

func initDataPut(paramsOut base.AttrT1, boxExcelData base.AttrS3) (entityDataPut *EntityDataPut) {
	entityDataPut = new(EntityDataPut)
	entityDataPut.Init(paramsOut, boxExcelData)
	return entityDataPut
}

func (self *EntityDataPut) Init(paramsOut base.AttrT1, boxExcelData base.AttrS3) *EntityDataPut {
	self.SetUserId(paramsOut).SetBoxExcelData(boxExcelData).SetBoxExcelTitle().SetPathDirThisTime().SetParamsPathFile()
	return self
}

func (self *EntityDataPut) SetUserId(paramsOut base.AttrT1) *EntityDataPut {
	userId, ok := paramsOut["UID"].(string)
	if !ok {
		jsonParamsOut := conv.ToJsonFromAttr(paramsOut)
		msgError := fmt.Sprintf(
			"The |UID| of paramsOut |%s| not found",
			jsonParamsOut)
		err.ErrPanic(msgError)
	}
	self.UserId = userId
	return self
}

func (self *EntityDataPut) SetBoxExcelData(boxExcelData base.AttrS3) *EntityDataPut {
	self.BoxExcelData = boxExcelData
	return self
}

func (self *EntityDataPut) SetBoxExcelTitle() *EntityDataPut {
	boxExcelTitle := make(base.AttrS3)
	boxExcelTitle["UID"] = paramsExcelTitleUID
	boxExcelTitle["OID"] = paramsExcelTitleOID
	self.BoxExcelTitle = boxExcelTitle
	return self
}

func (self *EntityDataPut) SetPathDirThisTime() *EntityDataPut {
	pathDirPutExcel := conf.GetPathDirPutExcel()
	file.MakeDir(pathDirPutExcel)
	userId := self.getUserId()
	pathDirThisTime := getPathDirThisTime(userId)
	file.MakeDir(pathDirThisTime)
	self.PathDirThisTime = pathDirThisTime
	return self
}

func getPathDirThisTime(userId string) (PathDir string) {
	remarkUser := "用户" + userId
	dirNamePrefix := "数据导出"
	suffix := time.GetTimeSuffix()
	dirName := fmt.Sprintf(
		"%s_%s_%s",
		dirNamePrefix,
		remarkUser,
		suffix)
	PathDirPutExcel := conf.GetPathDirPutExcel()
	PathDir = PathDirPutExcel + "/" + dirName
	PathDir = file.GetFilePathEmbed(PathDir)
	return PathDir
}

func (self *EntityDataPut) SetParamsPathFile() *EntityDataPut {
	userId := self.getUserId()
	pathDirThisTime := self.PathDirThisTime
	paramsPathFile := make(base.AttrS1)
	for shortName, fileNamePrefix := range paramsFileNamePrefix {
		fileName := getFileName(fileNamePrefix, userId)
		pathFileOne := pathDirThisTime + "/" + fileName
		pathFileOne = file.GetFilePathEmbed(pathFileOne)
		paramsPathFile[shortName] = pathFileOne
	}
	self.ParamsPathFile = paramsPathFile
	return self
}

func getFileName(fileNamePrefix string, userId string) (fileName string) {
	timeSuffix := time.GetTimeSuffix()
	remarkUser := "用户" + userId
	fileName = fmt.Sprintf(
		"%s_%s_%s.%s",
		fileNamePrefix,
		remarkUser,
		timeSuffix,
		ext)
	return fileName
}
