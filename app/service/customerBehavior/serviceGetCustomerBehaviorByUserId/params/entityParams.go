package params

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
)

type EntityParams struct {
	typeInterface.EntityParams `validate:"-"`
	UserId                     int    `validate:"required,min=1000001"`
	CountProductOrderPoolNFT int    `validate:"omitempty,min=0"`
	CountProductOrderNFTBuy  int    `validate:"omitempty,min=0"`
	CountProductOrderNFTSell int    `validate:"omitempty,min=0"`
	SumAmountGot             string `validate:"omitempty,min=0"`
}
