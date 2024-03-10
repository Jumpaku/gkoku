package exact

import (
	"errors"
	"fmt"
)

type Operator string

const (
	OperatorUnspecified Operator = ""
	OperatorAdd         Operator = "Add"
	OperatorSub         Operator = "Sub"
	OperatorMul         Operator = "Mul"
	OperatorDivFloor    Operator = "DivFloor"
)

var (
	ErrOverflow     error = overflowError{}
	ErrZeroDivision error = zeroDivisionError{}
)

func NewOverflowError(operator Operator, left, right int64) error {
	return overflowError{operator: operator, left: left, right: right}
}
func NewZeroDivisionError(numerator, denominator int64) error {
	return zeroDivisionError{numerator: numerator, denominator: denominator}
}

type overflowError struct {
	operator Operator
	left     int64
	right    int64
}

func (err overflowError) Error() string {
	return fmt.Sprintf(`overflow on %s(%d, %d)`, err.operator, err.left, err.right)
}
func (err overflowError) Is(target error) bool {
	var overflowError overflowError
	ok := errors.As(target, &overflowError)
	return ok
}

type zeroDivisionError struct {
	numerator   int64
	denominator int64
}

func (err zeroDivisionError) Error() string {
	return fmt.Sprintf(`division by zero occurred at %d / %d`, err.numerator, err.denominator)
}
func (err zeroDivisionError) Is(target error) bool {
	_, ok := target.(zeroDivisionError)
	return ok
}
