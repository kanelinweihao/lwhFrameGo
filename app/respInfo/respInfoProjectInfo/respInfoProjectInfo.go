package respInfoProjectInfo

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
)

var ParamsInDefaultProjectInfo base.AttrT1 = base.AttrT1{}

var ParamsOutDefaultProjectInfo base.AttrT1 = base.AttrT1{
	"ProjectTitle":   "未知项目名称",
	"ProjectVersion": "v1.0.0",
}

var FuncServiceProjectInfo conf.FuncService = nil

var PathTmplProjectInfo string = "./res/view/projectInfo.tmpl"
