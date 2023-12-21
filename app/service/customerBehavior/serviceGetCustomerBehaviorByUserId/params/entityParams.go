package params

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

type EntityParams struct {
	typeStruct.EntityParams  `validate:"-"`
	UserId                   int    `validate:"required,min=1000001"`
	Sign                     string `validate:"-"`
	CountProductOrderPoolNFT int    `validate:"omitempty,min=0"`
	CountProductOrderNFTBuy  int    `validate:"omitempty,min=0"`
	CountProductOrderNFTSell int    `validate:"omitempty,min=0"`
	SumAmountGot             string `validate:"omitempty,min=0"`
}
