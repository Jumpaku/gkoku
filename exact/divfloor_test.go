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

func getWantDivFloor(a, b int64) (wantDiv int64, wantMod int64, wantErr error) {
	if b == 0 {
		return 0, 0, ErrZeroDivision
	}

	wq, wr := (&big.Int{}).DivMod(big.NewInt(a), big.NewInt(b), &big.Int{})
	if wq.Cmp(big.NewInt(math.MinInt64)) < 0 || wq.Cmp(big.NewInt(math.MaxInt64)) > 0 {
		wantErr = ErrOverflow
	}
	wantDiv, wantMod = wq.Int64(), wr.Int64()
	return
}

type testInputDivFloor struct {
	a int64
	b int64
}

func TestDivFloor_Corner(t *testing.T) {
	var inputs []testInputDivFloor
	for _, a := range []int64{0, math.MinInt64, math.MinInt64 + 1, math.MaxInt64 - 1, math.MaxInt64, 1, -1} {
		for _, b := range []int64{0, math.MinInt64, math.MinInt64 + 1, math.MaxInt64 - 1, math.MaxInt64, 1, -1} {
			inputs = append(inputs, testInputDivFloor{a: a, b: b})
		}
	}

	for i, input := range inputs {
		a, b := input.a, input.b
		t.Run(fmt.Sprintf(`%3d: DivFloor(%d,%d)`, i, a, b), func(t *testing.T) {
			wantDiv, wantMod, wantErr := getWantDivFloor(a, b)

			gotDiv, gotMod, err := DivFloor(input.a, input.b)

			assert.True(t, errors.Is(err, wantErr))
			assert.Equal(t, wantDiv, gotDiv)
			assert.Equal(t, wantMod, gotMod)
		})
	}

}

func TestDivFloor_Small(t *testing.T) {
	var inputs []testInputDivFloor
	for a := -24; a <= 24; a++ {
		for b := -24; b <= 24; b++ {
			inputs = append(inputs, testInputDivFloor{a: int64(a), b: int64(b)})
		}
	}

	for i, input := range inputs {
		a, b := input.a, input.b
		t.Run(fmt.Sprintf(`%3d: DivFloor(%d,%d)`, i, a, b), func(t *testing.T) {
			wantDiv, wantMod, wantErr := getWantDivFloor(a, b)

			gotDiv, gotMod, err := DivFloor(input.a, input.b)

			assert.True(t, errors.Is(err, wantErr))
			assert.Equal(t, wantDiv, gotDiv)
			assert.Equal(t, wantMod, gotMod)
		})
	}

}

func TestDivFloor_Random(t *testing.T) {
	var inputs []testInputDivFloor
	r := rand.New(rand.NewSource(1))
	for i := 0; i < 1000; i++ {
		inputs = append(inputs, testInputDivFloor{a: r.Int63(), b: r.Int63()})
	}

	for i, input := range inputs {
		a, b := input.a, input.b
		t.Run(fmt.Sprintf(`%3d: DivFloor(%d,%d)`, i, a, b), func(t *testing.T) {
			wantDiv, wantMod, wantErr := getWantDivFloor(a, b)

			gotDiv, gotMod, err := DivFloor(input.a, input.b)

			assert.True(t, errors.Is(err, wantErr))
			assert.Equal(t, wantDiv, gotDiv)
			assert.Equal(t, wantMod, gotMod)
		})
	}
}

func TestDivFloor_Example(t *testing.T) {
	type testcaseDivFloor struct{ a, b, wantDiv, wantMod int64 }
	testcases := []testcaseDivFloor{
		{-12, 1, -12, 0},
		{-12, 5, -3, 3},
		{-12, 11, -2, 10},
		{-12, 12, -1, 0},
		{-12, 13, -1, 1},
		{12, 1, 12, 0},
		{12, 5, 2, 2},
		{12, 11, 1, 1},
		{12, 12, 1, 0},
		{12, 13, 0, 12},
		{-12, -1, 12, 0},
		{-12, -5, 3, 3},
		{-12, -11, 2, 10},
		{-12, -12, 1, 0},
		{-12, -13, 1, 1},
		{12, -1, -12, 0},
		{12, -5, -2, 2},
		{12, -11, -1, 1},
		{12, -12, -1, 0},
		{12, -13, 0, 12},
	}

	for i, testcase := range testcases {
		a, b, wantDiv, wantMod := testcase.a, testcase.b, testcase.wantDiv, testcase.wantMod
		t.Run(fmt.Sprintf(`%3d: DivFloor(%d,%d)`, i, a, b), func(t *testing.T) {
			gotDiv, gotMod, err := DivFloor(a, b)

			assert.Nil(t, err)
			assert.Equal(t, wantDiv, gotDiv)
			assert.Equal(t, wantMod, gotMod)
		})
	}

}
