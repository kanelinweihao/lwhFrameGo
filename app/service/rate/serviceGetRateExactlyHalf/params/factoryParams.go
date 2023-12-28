package params

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/service/base/paramsBase"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/calc"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

func InitEntityParams(paramsIn typeMap.AttrT1) (entityParams typeStruct.EntityParams) {
	entityParams = new(EntityParams)
	entityParamsBase := new(paramsBase.EntityParamsBase)
	entityParams.Load(entityParamsBase).Init(paramsIn)
	return entityParams
}

func (self *EntityParams) Load(entityParamsBase typeStruct.EntityParams) typeStruct.EntityParams {
	self.EntityParams = entityParamsBase
	entityParamsBase.Load(self)
	return self
}

func (self *EntityParams) SetParamsExec() typeStruct.EntityParams {
	x := self.X
	validateX(x)
	xHalf := x / 2
	decMin := 0
	decMax := 1 << x
	rate := getRate(xHalf, decMin, decMax)
	self.Rate = rate
	return self
}

func validateX(x int) {
	remainder := x % 2
	if remainder == 0 {
		return
	}
	msgError := fmt.Sprintf(
		"The x |%d| is invalid, it must be a multiple of 2",
		x)
	err.ErrPanic(msgError)
}

func getRate(xHalf int, decMin int, decMax int) (rate string) {
	timesWin := 0
	strNeed := conv.ToStrFromInt(xHalf)
	strDecMax := conv.ToStrFromInt(decMax)
	for dec := decMin; dec < decMax; dec++ {
		strBin := conv.ToStrBinFromDec(dec)
		// fmt.Println(strBin)
		arrStrNum := conv.ToArrStrFromStr(strBin, "")
		strSum := calc.BatchAdd(arrStrNum)
		// fmt.Println(strSum)
		if strSum == strNeed {
			timesWin++
		}
	}
	strWin := conv.ToStrFromInt(timesWin)
	rate = calc.Div(strWin, strDecMax)
	return rate
}
