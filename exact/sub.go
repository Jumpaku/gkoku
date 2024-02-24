package exact

func Sub(x, y int64) (result int64, err error) {
	result = x - y
	// HD 2-12 Overflow iff the arguments have different signs and
	// the sign of the result is different from the sign of x
	if ((x ^ y) & (x ^ result)) < 0 {
		err = NewOverflowError(OperatorSub, x, y)
	}
	return
}
