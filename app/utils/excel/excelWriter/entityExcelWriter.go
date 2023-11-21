package excelWriter

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/goroutine"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/rfl"
)

type EntityExcelWriter struct {
	ArrPathFile              []string
	BoxForExcel              base.BoxData
	AttrEntityExcelDataWrite map[string]*EntityExcelDataWrite
	AttrEntityChannel        base.AttrT1
}

func (self *EntityExcelWriter) BatchWriteExcel() (arrPathFile []string) {
	self.writeAttrEntityChannel().readAttrEntityChannel()
	arrPathFile = self.ArrPathFile
	return arrPathFile
}

func (self *EntityExcelWriter) writeAttrEntityChannel() *EntityExcelWriter {
	attrEntityExcelDataWrite := self.AttrEntityExcelDataWrite
	attrEntityChannel := make(base.AttrT1)
	for pathFile, entityExcelDataWrite := range attrEntityExcelDataWrite {
		entityChannel := goroutine.MakeEntityChannel()
		attrEntityChannel[pathFile] = entityChannel
		go entityExcelDataWrite.WriteToChannelOfWriteExcelData(entityChannel)
	}
	self.AttrEntityChannel = attrEntityChannel
	return self
}

func (self *EntityExcelWriter) readAttrEntityChannel() *EntityExcelWriter {
	attrEntityChannel := self.AttrEntityChannel
	arrPathFile := make([]string, 0, len(attrEntityChannel))
	for _, entityChannelToAssign := range attrEntityChannel {
		entityChannel := entityChannelToAssign.(*goroutine.EntityChannel)
		pathFile := readFromChannelOfWriteExcelData(entityChannel)
		arrPathFile = append(arrPathFile, pathFile)
	}
	self.ArrPathFile = arrPathFile
	return self
}

func readFromChannelOfWriteExcelData(entityChannel *goroutine.EntityChannel) (pathFile string) {
	dataOnce := entityChannel.ReadOnce()
	pathFile, ok := dataOnce.(string)
	if !ok {
		rfl.ErrPanicFormat(pathFile, "pathFile", "string")
	}
	return pathFile
}
