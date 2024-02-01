package params

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
)

type EntityParams struct {
	typeInterface.EntityParams `validate:"-"`
	StrArrAmount               string `validate:"required,min=3"`
	ErrorPrecision          int    `validate:"required,min=2,max=8"`
	IRR                     string `validate:"omitempty,min=0"`
}
