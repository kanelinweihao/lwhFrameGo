package repoDB

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
)

type EntityRepoDB struct {
	typeInterface.EntityRepoDB
}

func (self *EntityRepoDB) GetPropertiesNeed() (paramsAppend typeMap.AttrT1) {
	attrBase := self.EntityRepoDB.ToAttr()
	attrT3DBData := attrBase["AttrT3DBData"].(typeMap.AttrT3)
	userId := self.getUserIdFromAttrT3DBData(attrT3DBData)
	paramsAppend = make(typeMap.AttrT1, 1)
	paramsAppend["UserId"] = userId
	return paramsAppend
}

func (self *EntityRepoDB) getUserIdFromAttrT3DBData(attrT3DBData typeMap.AttrT3) (userId int) {
	sqlName := "GetUserIdByMobileNoAndOrgId"
	row := len(attrT3DBData[sqlName])
	if row == 0 {
		return userId
	}
	userId, ok := attrT3DBData[sqlName]["0"]["UserId"].(int)
	if !ok {
		msgError := fmt.Sprintf(
			"The userId |%v| of |%v| is invalid",
			userId,
			attrT3DBData)
		err.ErrPanic(msgError)
	}
	return userId
}
