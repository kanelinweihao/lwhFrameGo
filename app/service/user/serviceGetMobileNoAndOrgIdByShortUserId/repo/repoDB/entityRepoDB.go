package repoDB

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

type EntityRepoDB struct {
	typeStruct.EntityRepoDB
}

func (self *EntityRepoDB) GetPropertiesNeed() (paramsAppend typeMap.AttrT1) {
	attrBase := self.EntityRepoDB.ToAttr()
	attrT3DBData := attrBase["AttrT3DBData"].(typeMap.AttrT3)
	mobileNo := self.getMobileNoFromAttrT3DBData(attrT3DBData)
	orgId := self.getOrgIdFromAttrT3DBData(attrT3DBData)
	paramsAppend = make(typeMap.AttrT1, 2)
	paramsAppend["MobileNo"] = mobileNo
	paramsAppend["OrgId"] = orgId
	return paramsAppend
}

func (self *EntityRepoDB) getMobileNoFromAttrT3DBData(attrT3DBData typeMap.AttrT3) (mobileNo string) {
	sqlName := "GetMobileNoByUserId"
	row := len(attrT3DBData[sqlName])
	if row == 0 {
		return mobileNo
	}
	mobileNo, ok := attrT3DBData[sqlName]["0"]["MobileNo"].(string)
	if !ok {
		msgError := fmt.Sprintf(
			"The mobileNo |%v| of |%v| is invalid",
			mobileNo,
			attrT3DBData)
		err.ErrPanic(msgError)
	}
	return mobileNo
}

func (self *EntityRepoDB) getOrgIdFromAttrT3DBData(attrT3DBData typeMap.AttrT3) (orgId int) {
	sqlName := "GetOrgIdByUserId"
	row := len(attrT3DBData[sqlName])
	if row == 0 {
		return orgId
	}
	orgId, ok := attrT3DBData[sqlName]["0"]["OrgId"].(int)
	if !ok {
		msgError := fmt.Sprintf(
			"The orgId |%v| of |%v| is invalid",
			orgId,
			attrT3DBData)
		err.ErrPanic(msgError)
	}
	return orgId
}
