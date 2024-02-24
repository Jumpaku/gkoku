package exact

import "math"

func DivTrunc(x, y int64) (result int64, mod int64, err error) {
	if y == 0 {
		return 0, 0, NewZeroDivisionError(x, y)
	}
	if x == math.MinInt64 && y == -1 {
		return x / y, x % y, NewOverflowError(OperatorDivFloor, x, y)
	}

	if x == 0 {
		return 0, 0, nil
	}
	result, mod = x/y, x%y
	return
}
