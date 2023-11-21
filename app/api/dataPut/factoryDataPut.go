package dataPut

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/sqlInfo"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/excel"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/file"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/time"
)

func MakeEntityDataPut(paramsOut base.AttrT1, boxExcelData base.AttrS3) (entityDataPut *EntityDataPut) {
	entityDataPut = new(EntityDataPut)
	entityDataPut.Init(paramsOut, boxExcelData)
	return entityDataPut
}

func (self *EntityDataPut) Init(paramsOut base.AttrT1, boxExcelData base.AttrS3) *EntityDataPut {
	self.setParamsIn(paramsOut, boxExcelData).setParamsMore()
	return self
}

func (self *EntityDataPut) setParamsIn(paramsOut base.AttrT1, boxExcelData base.AttrS3) *EntityDataPut {
	self.ParamsOut = paramsOut
	self.BoxExcelData = boxExcelData
	return self
}

func (self *EntityDataPut) setParamsMore() *EntityDataPut {
	self.setUserId().setArrSQLName().setBoxExcelTitle()
	self.setPathDirThisTime().setParamsPathFile()
	self.setAttrEntityForExcel().setBoxForExcel()
	return self
}

func (self *EntityDataPut) setUserId() *EntityDataPut {
	paramsOut := self.ParamsOut
	userId, ok := paramsOut["UserId"].(string)
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

func (self *EntityDataPut) setArrSQLName() *EntityDataPut {
	boxExcelData := self.BoxExcelData
	arrSQLName := make([]string, 0, len(boxExcelData))
	for sqlName, _ := range boxExcelData {
		arrSQLName = append(arrSQLName, sqlName)
	}
	self.ArrSQLName = arrSQLName
	return self
}

func (self *EntityDataPut) setBoxExcelTitle() *EntityDataPut {
	arrSQLName := self.ArrSQLName
	boxExcelTitle := make(base.AttrS3)
	for _, sqlName := range arrSQLName {
		excelTitle := getParamsExcelTitle(sqlName)
		boxExcelTitle[sqlName] = excelTitle
	}
	self.BoxExcelTitle = boxExcelTitle
	return self
}

func getParamsExcelTitle(sqlName string) (paramsExcelTitle base.AttrS2) {
	paramsExcelTitle = sqlInfo.GetParamsExcelTitleBySQLName(sqlName)
	return paramsExcelTitle
}

func (self *EntityDataPut) setPathDirThisTime() *EntityDataPut {
	pathDirPutExcel := conf.GetPathDirPutExcel()
	file.MakeDir(pathDirPutExcel)
	userId := self.UserId
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

func (self *EntityDataPut) setParamsPathFile() *EntityDataPut {
	arrSQLName := self.ArrSQLName
	userId := self.UserId
	pathDirThisTime := self.PathDirThisTime
	paramsPathFile := make(base.AttrS1)
	for _, sqlName := range arrSQLName {
		pathFile := getPathFile(
			sqlName,
			userId,
			pathDirThisTime)
		paramsPathFile[sqlName] = pathFile
	}
	self.ParamsPathFile = paramsPathFile
	return self
}

func getPathFile(sqlName string, userId string, pathDirThisTime string) (pathFile string) {
	fileName := getFileName(sqlName, userId)
	pathFile = pathDirThisTime + "/" + fileName
	pathFile = file.GetFilePathEmbed(pathFile)
	return pathFile
}

func getFileName(sqlName string, userId string) (fileName string) {
	fileNamePrefix := getFileNamePrefix(sqlName)
	remarkUser := "用户" + userId
	timeSuffix := time.GetTimeSuffix()
	ext := excel.ExtExcel
	fileName = fmt.Sprintf(
		"%s_%s_%s.%s",
		fileNamePrefix,
		remarkUser,
		timeSuffix,
		ext)
	return fileName
}

func getFileNamePrefix(sqlName string) (fileNamePrefix string) {
	fileNamePrefix = sqlInfo.GetFileNamePrefixBySQLName(sqlName)
	return fileNamePrefix
}

func (self *EntityDataPut) setAttrEntityForExcel() *EntityDataPut {
	arrSQLName := self.ArrSQLName
	paramsPathFile := self.ParamsPathFile
	boxExcelTitle := self.BoxExcelTitle
	boxExcelData := self.BoxExcelData
	attrEntityForExcel := make(map[string]*EntityForExcel)
	for _, sqlName := range arrSQLName {
		pathFile, ok := paramsPathFile[sqlName]
		if !ok {
			errPanicExcelParamsRequired(sqlName, "excelPath")
		}
		attrS2ExcelTitle, ok := boxExcelTitle[sqlName]
		if !ok {
			errPanicExcelParamsRequired(sqlName, "excelTitle")
		}
		attrS2ExcelData, ok := boxExcelData[sqlName]
		if !ok {
			errPanicExcelParamsRequired(sqlName, "excelData")
		}
		entityForExcel := new(EntityForExcel)
		entityForExcel.SQLName = sqlName
		entityForExcel.PathFile = pathFile
		entityForExcel.SheetName = excel.SheetNameDefault
		entityForExcel.AttrS2ExcelTitle = attrS2ExcelTitle
		entityForExcel.AttrS2ExcelData = attrS2ExcelData
		attrEntityForExcel[pathFile] = entityForExcel
	}
	self.AttrEntityForExcel = attrEntityForExcel
	return self
}

func errPanicExcelParamsRequired(sqlName string, part string) {
	msgError := fmt.Sprintf(
		"The |%s| of |%s| is required",
		part,
		sqlName)
	err.ErrPanic(msgError)
}

func (self *EntityDataPut) setBoxForExcel() *EntityDataPut {
	attrEntityForExcel := self.AttrEntityForExcel
	boxForExcel := make(base.BoxData)
	for sqlName, entityForExcel := range attrEntityForExcel {
		attrT1ForExcel := conv.ToAttrFromEntity(entityForExcel)
		boxForExcel[sqlName] = attrT1ForExcel
	}
	self.BoxForExcel = boxForExcel
	return self
}
