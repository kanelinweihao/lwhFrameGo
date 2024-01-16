package params

import "github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"

type EntityParams struct {
	typeStruct.EntityParams `validate:"-"`
	MobileNo                string `validate:"required,min=11,max=11"`
	OrgId                   int    `validate:"required,min=31"`
	UserId                  int    `validate:"omitempty,min=1000001"`
}
