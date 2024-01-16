package conf

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

const RouteTypeFile int = 1
const RouteTypeApi int = 2
const RouteTypeWeb int = 3

var RouteTypeDefault int = RouteTypeWeb
var ParamsInDefaultDefault typeMap.AttrT1 = typeMap.AttrT1{}
var ParamsOutDefaultDefault typeMap.AttrT1 = typeMap.AttrT1{}
var ArrEntitySectionIn []typeStruct.EntitySection = nil
var ArrEntitySectionOut []typeStruct.EntitySection = nil
var ParamsOutAppendWeb typeMap.AttrT1 = typeMap.AttrT1{
	"RouteNameCN":         "",
	"ArrEntitySectionIn":  ArrEntitySectionIn,
	"ArrEntitySectionOut": ArrEntitySectionOut,
}
var FuncServiceDefault typeStruct.FuncService = nil
var PathTmplDefault string = "./res/view/projectInfo.tmpl"
