package exact

import (
	"math"
	"math/bits"
)

func Mul(x, y int64) (result int64, err error) {
	result = x * y
	signX, absX := signAndAbs(x)
	signY, absY := signAndAbs(y)
	h, l := bits.Mul64(absX, absY)
	if h != 0 {
		err = NewOverflowError(OperatorMul, x, y)
	}
	sign := signX * signY
	if sign < 0 && l > (-math.MinInt64) {
		err = NewOverflowError(OperatorMul, x, y)
	}
	if sign > 0 && l > math.MaxInt64 {
		err = NewOverflowError(OperatorMul, x, y)
	}
	return
}
