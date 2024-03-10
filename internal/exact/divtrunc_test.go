package exact

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"math/big"
	"math/rand"
	"testing"
)

func getWantDivTrunc(a, b int64) (wantDiv int64, wantMod int64, wantErr error) {
	if b == 0 {
		return 0, 0, ErrZeroDivision
	}

	wq, wr := (&big.Int{}).QuoRem(big.NewInt(a), big.NewInt(b), &big.Int{})
	if wq.Cmp(big.NewInt(math.MinInt64)) < 0 || wq.Cmp(big.NewInt(math.MaxInt64)) > 0 {
		wantErr = ErrOverflow
	}
	wantDiv, wantMod = wq.Int64(), wr.Int64()
	return
}

type testInputDivTrunc struct {
	a int64
	b int64
}

func TestDivTrunc_Corner(t *testing.T) {
	var inputs []testInputDivTrunc
	for _, a := range []int64{0, math.MinInt64, math.MinInt64 + 1, math.MaxInt64 - 1, math.MaxInt64, 1, -1} {
		for _, b := range []int64{0, math.MinInt64, math.MinInt64 + 1, math.MaxInt64 - 1, math.MaxInt64, 1, -1} {
			inputs = append(inputs, testInputDivTrunc{a: a, b: b})
		}
	}

	for i, input := range inputs {
		a, b := input.a, input.b
		t.Run(fmt.Sprintf(`%3d: DivTrunc(%d,%d)`, i, a, b), func(t *testing.T) {
			wantDiv, wantMod, wantErr := getWantDivTrunc(a, b)

			gotDiv, gotMod, err := DivTrunc(input.a, input.b)

			assert.True(t, errors.Is(err, wantErr))
			assert.Equal(t, wantDiv, gotDiv)
			assert.Equal(t, wantMod, gotMod)
		})
	}

}

func TestDivTrunc_Small(t *testing.T) {
	var inputs []testInputDivTrunc
	for a := -24; a <= 24; a++ {
		for b := -24; b <= 24; b++ {
			inputs = append(inputs, testInputDivTrunc{a: int64(a), b: int64(b)})
		}
	}

	for i, input := range inputs {
		a, b := input.a, input.b
		t.Run(fmt.Sprintf(`%3d: DivTrunc(%d,%d)`, i, a, b), func(t *testing.T) {
			wantDiv, wantMod, wantErr := getWantDivTrunc(a, b)

			gotDiv, gotMod, err := DivTrunc(input.a, input.b)

			assert.True(t, errors.Is(err, wantErr))
			assert.Equal(t, wantDiv, gotDiv)
			assert.Equal(t, wantMod, gotMod)
		})
	}

}

func TestDivTrunc_Random(t *testing.T) {
	var inputs []testInputDivTrunc
	r := rand.New(rand.NewSource(1))
	for i := 0; i < 1000; i++ {
		inputs = append(inputs, testInputDivTrunc{a: r.Int63(), b: r.Int63()})
	}

	for i, input := range inputs {
		a, b := input.a, input.b
		t.Run(fmt.Sprintf(`%3d: DivTrunc(%d,%d)`, i, a, b), func(t *testing.T) {
			wantDiv, wantMod, wantErr := getWantDivTrunc(a, b)

			gotDiv, gotMod, err := DivTrunc(input.a, input.b)

			assert.True(t, errors.Is(err, wantErr))
			assert.Equal(t, wantDiv, gotDiv)
			assert.Equal(t, wantMod, gotMod)
		})
	}
}

func TestDivTrunc_Example(t *testing.T) {
	type testcaseDivTrunc struct{ a, b, wantDiv, wantMod int64 }
	testcases := []testcaseDivTrunc{
		{-12, 1, -12, 0},
		{-12, 5, -2, -2},
		{-12, 11, -1, -1},
		{-12, 12, -1, 0},
		{-12, 13, 0, -12},
		{12, 1, 12, 0},
		{12, 5, 2, 2},
		{12, 11, 1, 1},
		{12, 12, 1, 0},
		{12, 13, 0, 12},
		{-12, -1, 12, 0},
		{-12, -5, 2, -2},
		{-12, -11, 1, -1},
		{-12, -12, 1, 0},
		{-12, -13, 0, -12},
		{12, -1, -12, 0},
		{12, -5, -2, 2},
		{12, -11, -1, 1},
		{12, -12, -1, 0},
		{12, -13, 0, 12},
	}

	for i, testcase := range testcases {
		a, b, wantDiv, wantMod := testcase.a, testcase.b, testcase.wantDiv, testcase.wantMod
		t.Run(fmt.Sprintf(`%3d: DivTrunc(%d,%d)`, i, a, b), func(t *testing.T) {
			gotDiv, gotMod, err := DivTrunc(a, b)

			assert.Nil(t, err)
			assert.Equal(t, wantDiv, gotDiv)
			assert.Equal(t, wantMod, gotMod)
		})
	}

}
