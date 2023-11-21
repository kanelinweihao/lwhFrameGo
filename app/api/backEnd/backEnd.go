package backEnd

import (
	"github.com/kanelinweihao/lwhFrameGo/app/api/cacheSet"
	"github.com/kanelinweihao/lwhFrameGo/app/api/dataGet"
	"github.com/kanelinweihao/lwhFrameGo/app/api/dataPut"
	"github.com/kanelinweihao/lwhFrameGo/app/api/emailSend"
	"github.com/kanelinweihao/lwhFrameGo/app/api/paramsOutGet"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/time"
)

func ExecBackEnd(paramsIn base.AttrT1) (paramsOut base.AttrT1) {
	paramsOut = paramsOutGet.GetParamsOut(paramsIn)
	time.ShowTimeAndMsg("ParamsOut get success")
	boxData := dataGet.GetData(paramsOut)
	time.ShowTimeAndMsg("Data get success")
	pathDirExcel, arrPathFileExcel := dataPut.PutData(paramsOut, boxData)
	time.ShowTimeAndMsg("Data put success")
	emailSend.SendEmail(pathDirExcel, arrPathFileExcel)
	time.ShowTimeAndMsg("Email send success")
	cacheSet.SetCache(arrPathFileExcel)
	time.ShowTimeAndMsg("Cache set success")
	time.ShowTimeAndMsg("OK")
	return paramsOut
}
