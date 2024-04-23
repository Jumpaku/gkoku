package tokiope

import (
	"errors"
	"github.com/Jumpaku/tokiope/internal/exact"
)

func divFloor(x, y int64) (d int64, m int64, state State) {
	d, m, err := exact.DivFloor(x, y)
	if errors.Is(err, exact.ErrOverflow) {
		state |= StateOverflow
	}
	if errors.Is(err, exact.ErrZeroDivision) {
		panic(err)
	}
	return
}

func divTrunc(x, y int64) (d int64, m int64, state State) {
	d, m, err := exact.DivTrunc(x, y)
	if errors.Is(err, exact.ErrOverflow) {
		state |= StateOverflow
	}
	if errors.Is(err, exact.ErrZeroDivision) {
		panic(err)
	}
	return
}

func add(x, y int64) (a int64, state State) {
	a, err := exact.Add(x, y)
	if errors.Is(err, exact.ErrOverflow) {
		state |= StateOverflow
	}
	return
}
func sub(x, y int64) (s int64, state State) {
	s, err := exact.Sub(x, y)
	if errors.Is(err, exact.ErrOverflow) {
		state |= StateOverflow
	}
	return
}
func mul(x, y int64) (m int64, state State) {
	m, err := exact.Mul(x, y)
	if errors.Is(err, exact.ErrOverflow) {
		state |= StateOverflow
	}
	return
}
