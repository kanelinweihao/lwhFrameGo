package dataPut

import (
	"fmt"
	"go.lwh.com/linweihao/lwhFrameGo/app/conf"
	"go.lwh.com/linweihao/lwhFrameGo/app/factory/factoryOfExcel"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/base"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/conv"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/dd"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/err"
	_ "go.lwh.com/linweihao/lwhFrameGo/app/utils/excel"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/file"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/time"
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

type EntityDataPut struct {
	UserId          string
	PathDirThisTime string
	ParamsPathFile  base.AttrS1
	BoxExcelData    base.AttrS3
	BoxExcelTitle   base.AttrS3
}

////
// SET
////

func (self *EntityDataPut) SetUserId(paramsOut base.AttrT1) {
	userId, ok := paramsOut["UID"].(string)
	if !ok {
		jsonParamsOut := conv.ToJsonFromAttr(paramsOut)
		msgError := fmt.Sprintf(
			"The |UID| of paramsOut |%s| not found",
			jsonParamsOut)
		err.ErrPanic(msgError)
	}
	self.UserId = userId
	// dd.DD(self.UserId)
	return
}

func (self *EntityDataPut) SetBoxExcelData(boxExcelData base.AttrS3) {
	self.BoxExcelData = boxExcelData
	return
}

func (self *EntityDataPut) SetBoxExcelTitle() {
	boxExcelTitle := base.AttrS3{}
	boxExcelTitle["UID"] = paramsExcelTitleUID
	boxExcelTitle["OID"] = paramsExcelTitleOID
	self.BoxExcelTitle = boxExcelTitle
	// dd.DD(self.BoxExcelTitle)
	return
}

func (self *EntityDataPut) SetPathDirThisTime() {
	pathDirPutExcel := conf.GetPathDirPutExcel()
	file.MakeDir(pathDirPutExcel)
	userId := self.getUserId()
	pathDirThisTime := getPathDirThisTime(userId)
	file.MakeDir(pathDirThisTime)
	self.PathDirThisTime = pathDirThisTime
	// dd.DD(self.PathDirThisTime)
	return
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
	PathDir = file.GetFilePathAbs(PathDir)
	// dd.DD(PathDir)
	return PathDir
}

func (self *EntityDataPut) SetParamsPathFile() {
	userId := self.getUserId()
	pathDirThisTime := self.PathDirThisTime
	paramsPathFile := base.AttrS1{}
	for shortName, fileNamePrefix := range paramsFileNamePrefix {
		fileName := getFileName(fileNamePrefix, userId)
		pathFileRel := pathDirThisTime + "/" + fileName
		pathFileAbs := file.GetFilePathAbs(pathFileRel)
		// fmt.Println(pathFileNum)
		paramsPathFile[shortName] = pathFileAbs
	}
	self.ParamsPathFile = paramsPathFile
	// dd.DD(self.ParamsPathFile)
	return
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
	// dd.DD(fileName)
	return fileName
}

////
// GET
////

func (self *EntityDataPut) getUserId() (userId string) {
	userId = self.UserId
	if userId == "0" {
		msgError := "The |userId| is required"
		err.ErrPanic(msgError)
	}
	return userId
}

func errPanicExcelParamsRequired(shortName string, part string) {
	msgError := fmt.Sprintf(
		"The |%s| of |%s| is required",
		part,
		shortName)
	err.ErrPanic(msgError)
}

////
// EXEC
////

func (self *EntityDataPut) BatchPutExcel() {
	paramsPathFile := self.ParamsPathFile
	boxExcelTitle := self.BoxExcelTitle
	boxExcelData := self.BoxExcelData
	// dd.DD(boxExcelData)
	for shortName, pathFile := range paramsPathFile {
		attrS2ExcelTitle, ok := boxExcelTitle[shortName]
		if !ok {
			errPanicExcelParamsRequired(shortName, "excelTitle")
		}
		attrS2ExcelData, ok := boxExcelData[shortName]
		if !ok {
			errPanicExcelParamsRequired(shortName, "excelData")
		}
		// dd.DDD(shortName, pathFile, attrS2ExcelTitle, attrS2ExcelData)
		go putOne(
			shortName,
			pathFile,
			attrS2ExcelTitle,
			attrS2ExcelData)
	}
	return
}

func putOne(shortName string, pathFile string, attrS2ExcelTitle base.AttrS2, attrS2ExcelData base.AttrS2) {
	entityExcel := factoryOfExcel.MakeEntityOfExcelNew(pathFile)
	defer entityExcel.CloseExcel()
	entityExcel.WriteExcel(attrS2ExcelTitle, attrS2ExcelData)
	readOne(pathFile)
	return
}

func readOne(pathFile string) {
	entityExcel := factoryOfExcel.MakeEntityOfExcelOpen(pathFile)
	defer entityExcel.CloseExcel()
	attrS2ExcelDataRead := entityExcel.ReadExcel()
	dd.IGNORE(attrS2ExcelDataRead)
}
