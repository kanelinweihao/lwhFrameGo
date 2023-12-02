package excelReader

func InitEntityExcelReader(arrPathFile []string) (entityExcelReader *EntityExcelReader) {
	entityExcelReader = new(EntityExcelReader)
	entityExcelReader.Init(arrPathFile)
	return entityExcelReader
}

func (self *EntityExcelReader) Init(arrPathFile []string) *EntityExcelReader {
	self.setPropertiesIn(arrPathFile).setPropertiesMore()
	return self
}

func (self *EntityExcelReader) setPropertiesIn(arrPathFile []string) *EntityExcelReader {
	self.ArrPathFile = arrPathFile
	return self
}

func (self *EntityExcelReader) setPropertiesMore() *EntityExcelReader {
	self.setAttrEntityExcelDataRead()
	return self
}

func (self *EntityExcelReader) setAttrEntityExcelDataRead() *EntityExcelReader {
	arrPathFile := self.ArrPathFile
	attrEntityExcelDataRead := make(map[string]*EntityExcelDataRead)
	for _, pathFile := range arrPathFile {
		entityExcelDataRead := InitEntityExcelDataRead(pathFile)
		attrEntityExcelDataRead[pathFile] = entityExcelDataRead
	}
	self.AttrEntityExcelDataRead = attrEntityExcelDataRead
	return self
}
