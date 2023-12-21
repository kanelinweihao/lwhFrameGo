package params

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

type EntityParams struct {
	typeStruct.EntityParams `validate:"-"`
	StrArrAmount            string `validate:"required,min=3"`
	ErrorPrecision          int    `validate:"required,min=2,max=8"`
	Sign                    string `validate:"-"`
	IRR                     string `validate:"omitempty,min=0"`
}
