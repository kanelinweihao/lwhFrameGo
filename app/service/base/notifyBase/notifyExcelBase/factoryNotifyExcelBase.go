package notifyExcelBase

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/dict/dictSQL/dictSQL"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/excel"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/file"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/times"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

func (self *EntityNotifyExcelBase) Load(entityNotifyExcelDerived typeStruct.EntityNotifyExcel) typeStruct.EntityNotifyExcel {
	self.Derived = entityNotifyExcelDerived
	return self.Derived
}

func (self *EntityNotifyExcelBase) Init(paramsToNotify typeMap.AttrT1, attrT3DBData typeMap.AttrT3) typeStruct.EntityNotifyExcel {
	self.setPropertiesIn(paramsToNotify, attrT3DBData).setPropertiesMore()
	return self
}

func (self *EntityNotifyExcelBase) setPropertiesIn(paramsToNotify typeMap.AttrT1, attrT3DBData typeMap.AttrT3) *EntityNotifyExcelBase {
	self.ParamsToNotify = paramsToNotify
	self.AttrT3DBData = attrT3DBData
	return self
}

func (self *EntityNotifyExcelBase) setPropertiesMore() *EntityNotifyExcelBase {
	self.setUserId().setArrSQLName().setBoxExcelTitle()
	self.setPathDirThisTime().setParamsPathFile()
	self.setAttrEntityForExcel().setBoxToExcel()
	return self
}

func (self *EntityNotifyExcelBase) setUserId() *EntityNotifyExcelBase {
	paramsToNotify := self.ParamsToNotify
	userId, ok := paramsToNotify["UserId"].(int)
	if !ok {
		jsonParamsToNotify := conv.ToJsonFromAttr(paramsToNotify)
		msgError := fmt.Sprintf(
			"The |userId| of paramsToNotify |%s| not found",
			jsonParamsToNotify)
		err.ErrPanic(msgError)
	}
	strUserId := conv.ToStrFromInt(userId)
	self.UserId = userId
	self.StrUserId = strUserId
	return self
}

func (self *EntityNotifyExcelBase) setArrSQLName() *EntityNotifyExcelBase {
	attrT3DBData := self.AttrT3DBData
	arrSQLName := make([]string, 0, len(attrT3DBData))
	for sqlName, _ := range attrT3DBData {
		arrSQLName = append(arrSQLName, sqlName)
	}
	self.ArrSQLName = arrSQLName
	return self
}

func (self *EntityNotifyExcelBase) setBoxExcelTitle() *EntityNotifyExcelBase {
	arrSQLName := self.ArrSQLName
	boxExcelTitle := make(typeMap.AttrS3)
	for _, sqlName := range arrSQLName {
		excelTitle := getParamsExcelTitle(sqlName)
		boxExcelTitle[sqlName] = excelTitle
	}
	self.BoxExcelTitle = boxExcelTitle
	return self
}

func getParamsExcelTitle(sqlName string) (paramsExcelTitle typeMap.AttrS2) {
	paramsExcelTitle = dictSQL.GetParamsExcelTitleBySQLName(sqlName)
	return paramsExcelTitle
}

func (self *EntityNotifyExcelBase) setPathDirThisTime() *EntityNotifyExcelBase {
	pathDirPutExcel := conf.GetPathDirPutExcel()
	file.MakeDir(pathDirPutExcel)
	strUserId := self.StrUserId
	remarkUser := "用户" + strUserId
	dirNamePrefix := "数据导出"
	suffix := times.GetTimeSuffix()
	dirName := fmt.Sprintf(
		"%s_%s_%s",
		dirNamePrefix,
		remarkUser,
		suffix)
	PathDirPutExcel := conf.GetPathDirPutExcel()
	pathDirThisTimeRel := PathDirPutExcel + "/" + dirName
	pathDirThisTime := file.GetFilePathEmbed(pathDirThisTimeRel)
	file.MakeDir(pathDirThisTime)
	self.PathDirThisTime = pathDirThisTime
	return self
}

func (self *EntityNotifyExcelBase) setParamsPathFile() *EntityNotifyExcelBase {
	arrSQLName := self.ArrSQLName
	strUserId := self.StrUserId
	pathDirThisTime := self.PathDirThisTime
	paramsPathFile := make(typeMap.AttrS1)
	for _, sqlName := range arrSQLName {
		pathFile := getPathFile(
			sqlName,
			strUserId,
			pathDirThisTime)
		paramsPathFile[sqlName] = pathFile
	}
	self.ParamsPathFile = paramsPathFile
	return self
}

func getPathFile(sqlName string, strUserId string, pathDirThisTime string) (pathFile string) {
	fileName := getFileName(sqlName, strUserId)
	pathFile = pathDirThisTime + "/" + fileName
	pathFile = file.GetFilePathEmbed(pathFile)
	return pathFile
}

func getFileName(sqlName string, userId string) (fileName string) {
	fileNamePrefix := getFileNamePrefix(sqlName)
	remarkUser := "用户" + userId
	timeSuffix := times.GetTimeSuffix()
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
	fileNamePrefix = dictSQL.GetFileNamePrefixBySQLName(sqlName)
	return fileNamePrefix
}

func (self *EntityNotifyExcelBase) setAttrEntityForExcel() *EntityNotifyExcelBase {
	arrSQLName := self.ArrSQLName
	paramsPathFile := self.ParamsPathFile
	boxExcelTitle := self.BoxExcelTitle
	attrT3DBData := self.AttrT3DBData
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
		attrT2ExcelData, ok := attrT3DBData[sqlName]
		if !ok {
			errPanicExcelParamsRequired(sqlName, "excelData")
		}
		attrS2ExcelData := conv.ToAttrS2FromAttrT2(attrT2ExcelData)
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

func (self *EntityNotifyExcelBase) setBoxToExcel() *EntityNotifyExcelBase {
	attrEntityForExcel := self.AttrEntityForExcel
	boxToExcel := make(typeMap.BoxData)
	for sqlName, entityForExcel := range attrEntityForExcel {
		attrT1ForExcel := conv.ToAttrFromEntity(entityForExcel)
		boxToExcel[sqlName] = attrT1ForExcel
	}
	self.BoxToExcel = boxToExcel
	return self
}
