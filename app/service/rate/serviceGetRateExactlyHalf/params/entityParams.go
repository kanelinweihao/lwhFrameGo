package params

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
)

type EntityParams struct {
	typeInterface.EntityParams `validate:"-"`
	X                          int    `validate:"required,min=2"`
	Rate                    string `validate:"omitempty,min=0"`
}
