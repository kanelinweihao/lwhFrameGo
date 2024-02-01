package params

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
)

type EntityParams struct {
	typeInterface.EntityParams `validate:"-"`
	ShortUserId                int    `validate:"required,min=1"`
	UserId                  int    `validate:"required,min=1000001"`
	MobileNo                string `validate:"omitempty,min=11,max=11"`
	OrgId                   int    `validate:"omitempty,min=31"`
}
