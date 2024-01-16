package params

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

type EntityParams struct {
	typeStruct.EntityParams `validate:"-"`
	X                       int    `validate:"required,min=2"`
	Rate                    string `validate:"omitempty,min=0"`
}
