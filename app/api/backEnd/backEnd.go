package backEnd

import (
	"github.com/kanelinweihao/lwhFrameGo/app/api/cacheSet"
	"github.com/kanelinweihao/lwhFrameGo/app/api/dataGet"
	"github.com/kanelinweihao/lwhFrameGo/app/api/dataPut"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	_ "github.com/kanelinweihao/lwhFrameGo/app/utils/dd"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/time"
)

var arrSQLName = []string{
	"GetMobileNoByUserId",
	"GetOrgIdByUserId",
}

func ExecBackEnd(paramsIn base.AttrT1) (paramsOut base.AttrT1) {
	paramsOut, boxData := dataGet.GetData(paramsIn, arrSQLName)
	time.ShowTimeAndMsg("Data get success")
	dataPut.PutDataToExcel(paramsOut, boxData)
	time.ShowTimeAndMsg("Data put success")
	cacheSet.SetCache(paramsOut)
	time.ShowTimeAndMsg("Cache set success")
	time.ShowTimeAndMsg("OK")
	return paramsOut
}
