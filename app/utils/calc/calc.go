package calc

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/shopspring/decimal"
)

func Add(x string, y string) (z string) {
	X, errDecimal := decimal.NewFromString(x)
	err.ErrCheck(errDecimal)
	Y, errDecimal := decimal.NewFromString(y)
	err.ErrCheck(errDecimal)
	Z := X.Add(Y)
	z = Z.String()
	return z
}

func BatchAdd(arrX []string) (z string) {
	z = "0"
	for _, x := range arrX {
		z = Add(z, x)
	}
	return z
}

func Sub(x string, y string) (z string) {
	X, errDecimal := decimal.NewFromString(x)
	err.ErrCheck(errDecimal)
	Y, errDecimal := decimal.NewFromString(y)
	err.ErrCheck(errDecimal)
	Z := X.Sub(Y)
	z = Z.String()
	return z
}

func Mul(x string, y string) (z string) {
	X, errDecimal := decimal.NewFromString(x)
	err.ErrCheck(errDecimal)
	Y, errDecimal := decimal.NewFromString(y)
	err.ErrCheck(errDecimal)
	Z := X.Mul(Y)
	z = Z.String()
	return z
}

func Div(x string, y string) (z string) {
	X, errDecimal := decimal.NewFromString(x)
	err.ErrCheck(errDecimal)
	Y, errDecimal := decimal.NewFromString(y)
	err.ErrCheck(errDecimal)
	Z := X.Div(Y)
	z = Z.String()
	return z
}

func IsNonNegativeNumber(x string) bool {
	X, errDecimal := decimal.NewFromString(x)
	err.ErrCheck(errDecimal)
	Zero, _ := decimal.NewFromString("0")
	cmp := X.Cmp(Zero)
	if cmp < 0 {
		return false
	}
	return true
}

func UsePrecision(x string, precision int) (z string) {
	X, errDecimal := decimal.NewFromString(x)
	err.ErrCheck(errDecimal)
	precision32 := int32(precision)
	Z := X.Truncate(precision32)
	z = Z.String()

	return z
}
