package calc

import (
	"github.com/shopspring/decimal"
	_ "go.lwh.com/linweihao/lwhFrameGo/app/utils/dd"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/err"
	_ "math/big"
)

func Add(x string, y string) (z string) {
	X, errDeciaml := decimal.NewFromString(x)
	err.ErrCheck(errDeciaml)
	Y, errDeciaml := decimal.NewFromString(y)
	err.ErrCheck(errDeciaml)
	Z := X.Add(Y)
	z = Z.String()
	return z
}

func Sub(x string, y string) (z string) {
	X, errDeciaml := decimal.NewFromString(x)
	err.ErrCheck(errDeciaml)
	Y, errDeciaml := decimal.NewFromString(y)
	err.ErrCheck(errDeciaml)
	Z := X.Sub(Y)
	z = Z.String()
	return z
}

func Mul(x string, y string) (z string) {
	X, errDeciaml := decimal.NewFromString(x)
	err.ErrCheck(errDeciaml)
	Y, errDeciaml := decimal.NewFromString(y)
	err.ErrCheck(errDeciaml)
	Z := X.Mul(Y)
	z = Z.String()
	return z
}

func Div(x string, y string) (z string) {
	X, errDeciaml := decimal.NewFromString(x)
	err.ErrCheck(errDeciaml)
	Y, errDeciaml := decimal.NewFromString(y)
	err.ErrCheck(errDeciaml)
	Z := X.Div(Y)
	z = Z.String()
	return z
}
