package exact

func Add(x, y int64) (result int64, err error) {
	result = x + y
	// HD 2-12 Overflow iff both arguments have the opposite sign of the result
	if ((x ^ result) & (y ^ result)) < 0 {
		err = NewOverflowError(OperatorAdd, x, y)
	}
	return
}
