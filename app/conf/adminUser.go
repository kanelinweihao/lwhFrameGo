package conf

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/pack"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"strings"
)

type EntityAdminUser struct {
	AdminUserName     string `validate:"required"`
	AdminUserPassword string `validate:"required"`
	AdminUserId       int    `validate:"required"`
	AdminUserRole     int    `validate:"required"`
}

func (self *EntityAdminUser) ToJson() (json string) {
	attrT1 := conv.ToAttrFromEntity(self)
	delete(attrT1, "AdminUserPassword")
	json = conv.ToJsonFromAttr(attrT1)
	return json
}

func GetParamsAdminUser() (paramsAdminUser typeMap.AttrS1) {
	paramsAdminUser = make(typeMap.AttrS1)
	pathAdminUser := getPathAdminUser()
	strFromEnv := pack.ReadFileEmbedAsString(pathAdminUser)
	arrLine := strings.Split(strFromEnv, "\n")
	strAnnotation := "#"
	strSeparator := "|"
	for _, strLine := range arrLine {
		if len(strLine) == 0 {
			continue
		}
		firstLetter := string(strLine[0])
		if firstLetter == strAnnotation {
			continue
		}
		arrPart := strings.Split(strLine, strSeparator)
		if len(arrPart) != 2 {
			continue
		}
		adminUserName := arrPart[0]
		jsonAdminUser := arrPart[1]
		paramsAdminUser[adminUserName] = jsonAdminUser
	}
	return paramsAdminUser
}

func GetJsonAdminUser(adminUserName string) (jsonAdminUser string) {
	paramsAdminUser := GetParamsAdminUser()
	jsonAdminUser, ok := paramsAdminUser[adminUserName]
	if !ok {
		msgError := fmt.Sprintf(
			"The adminUserName |%s| not found in adminUser.env",
			adminUserName)
		err.ErrPanic(msgError)
	}
	if len(jsonAdminUser) == 0 {
		msgError := fmt.Sprintf(
			"The value of |%s| is empty in adminUser.env",
			adminUserName)
		err.ErrPanic(msgError)
	}
	return jsonAdminUser
}

func GetEntityAdminUser(adminUserName string) (entityAdminUser *EntityAdminUser) {
	jsonAdminUser := GetJsonAdminUser(adminUserName)
	entityAdminUser = new(EntityAdminUser)
	conv.ToEntityFromJson(jsonAdminUser, entityAdminUser)
	return entityAdminUser
}
