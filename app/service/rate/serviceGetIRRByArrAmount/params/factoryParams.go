package params

import (
	"github.com/kanelinweihao/lwhFrameGo/app/service/base/paramsBase"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/calc"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/funcArr"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

var IRRToCheckPositive = "0.1024"
var IRRToCheckNegative = "-0.1024"

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
	strArrAmount := self.StrArrAmount
	precision := self.ErrorPrecision + 8
	arrAmountWithZero := conv.ToArrStrFromStr(strArrAmount, ",")
	arrAmount := funcArr.TrimZero(arrAmountWithZero)
	sumAmount := calc.BatchAdd(arrAmount)
	if sumAmount == "0" {
		self.IRR = "0"
		return self
	}
	IRR := "0"
	if calc.IsNonNegativeNumber(sumAmount) {
		IRR = getIRRPositive(arrAmount, precision)
	} else {
		IRR = getIRRNegative(arrAmount, precision)
	}
	IRR = calc.UsePrecision(IRR, self.ErrorPrecision)
	self.IRR = IRR
	return self
}

func getIRRPositive(arrAmount []string, precision int) (IRR string) {
	min := "0"
	max := ""
	IRRToCheck := IRRToCheckPositive
	IRR = getIRR(arrAmount, precision, min, max, IRRToCheck)
	return IRR
}

func getIRRNegative(arrAmount []string, precision int) (IRR string) {
	min := ""
	max := "0"
	IRRToCheck := IRRToCheckNegative
	IRR = getIRR(arrAmount, precision, min, max, IRRToCheck)
	return IRR
}

func getIRR(arrAmount []string, precision int, min string, max string, IRRToCheck string) (IRR string) {
	for {
		isContinue, isNeedUp := checkIRR(arrAmount, precision, IRRToCheck)
		if !isContinue {
			break
		}
		if min == max {
			break
		}
		if isNeedUp {
			min = IRRToCheck
			if max == "" {
				IRRToCheck = calc.Mul(IRRToCheck, "2")
			}
		} else {
			max = IRRToCheck
			if min == "" {
				IRRToCheck = calc.Div(IRRToCheck, "2")
			}
		}
		if min != "" && max != "" {
			IRRToCheck = calc.Div(calc.Add(min, max), "2")
			IRRToCheck = calc.UsePrecision(IRRToCheck, precision)
			if IRRToCheck == min {
				break
			}
		}
	}
	IRR = IRRToCheck
	return IRR
}

func checkIRR(arrAmount []string, precision int, IRRToCheck string) (isContinue bool, isNeedUp bool) {
	z := "0"
	r := calc.Add("1", IRRToCheck)
	length := len(arrAmount)
	for i := 0; i < length; i++ {
		x := arrAmount[i]
		// z = (z + x) * r
		z = calc.Mul(calc.Add(z, x), r)
	}
	if z == "0" {
		return false, true
	}
	if calc.IsNonNegativeNumber(z) {
		return true, true
	}
	return true, false
}
