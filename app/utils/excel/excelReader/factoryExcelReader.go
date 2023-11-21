package excelReader

func MakeEntityExcelReader(arrPathFile []string) (entityExcelReader *EntityExcelReader) {
	entityExcelReader = new(EntityExcelReader)
	entityExcelReader.Init(arrPathFile)
	return entityExcelReader
}

func (self *EntityExcelReader) Init(arrPathFile []string) *EntityExcelReader {
	self.setParamsIn(arrPathFile).setParamsMore()
	return self
}

func (self *EntityExcelReader) setParamsIn(arrPathFile []string) *EntityExcelReader {
	self.ArrPathFile = arrPathFile
	return self
}

func (self *EntityExcelReader) setParamsMore() *EntityExcelReader {
	self.setAttrEntityExcelDataRead()
	return self
}

func (self *EntityExcelReader) setAttrEntityExcelDataRead() *EntityExcelReader {
	arrPathFile := self.ArrPathFile
	attrEntityExcelDataRead := make(map[string]*EntityExcelDataRead)
	for _, pathFile := range arrPathFile {
		entityExcelDataRead := MakeEntityExcelDataRead(pathFile)
		attrEntityExcelDataRead[pathFile] = entityExcelDataRead
	}
	self.AttrEntityExcelDataRead = attrEntityExcelDataRead
	return self
}
