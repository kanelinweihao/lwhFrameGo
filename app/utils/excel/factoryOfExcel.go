package excel

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/xuri/excelize/v2"
)

/*
Init
*/

func MakeEntityOfExcelNew(pathFile string) (entityExcel *EntityExcel) {
	checkExt(pathFile)
	isNew := true
	entityExcel = makeEntityOfExcel(pathFile, isNew)
	return entityExcel
}

func MakeEntityOfExcelOpen(pathFile string) (entityExcel *EntityExcel) {
	checkExt(pathFile)
	isNew := false
	entityExcel = makeEntityOfExcel(pathFile, isNew)
	return entityExcel
}

func makeEntityOfExcel(pathFile string, isNew bool) (entityExcel *EntityExcel) {
	checkExt(pathFile)
	entityExcel = initEntityExcel(pathFile, isNew)
	return entityExcel
}

func initEntityExcel(pathFile string, isNew bool) (entityExcel *EntityExcel) {
	entityExcel = new(EntityExcel)
	entityExcel.Init(pathFile, isNew)
	return entityExcel
}

func (self *EntityExcel) Init(pathFile string, isNew bool) *EntityExcel {
	self.PathFile = pathFile
	self.IsNew = isNew
	self.setExcelFile()
	return self
}

func (self *EntityExcel) setExcelFile() *EntityExcel {
	isNew := self.IsNew
	if !isNew {
		pathFile := self.PathFile
		f, errExcelOpen := excelize.OpenFile(pathFile)
		err.ErrCheck(errExcelOpen)
		self.ExcelFile = f
		return self
	}
	f := excelize.NewFile()
	self.ExcelFile = f
	return self
}

/*
Close
*/

func (self *EntityExcel) CloseExcel() {
	errExcelClose := self.ExcelFile.Close()
	err.ErrCheck(errExcelClose)
	self.ExcelFile = nil
	return
}
