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

func getWantMul(a, b int64) (want int64, shouldErr bool) {
	w := (&big.Int{}).Mul(big.NewInt(a), big.NewInt(b))
	shouldErr = w.Cmp(big.NewInt(math.MinInt64)) < 0 || w.Cmp(big.NewInt(math.MaxInt64)) > 0
	want = w.Int64()
	if shouldErr {
		want = a * b
	}
	return
}

type testInputMul struct {
	a int64
	b int64
}

func TestMul_Corner(t *testing.T) {
	var inputs []testInputMul
	for _, a := range []int64{0, math.MinInt64, math.MinInt64 + 1, math.MaxInt64 - 1, math.MaxInt64, 1, -1} {
		for _, b := range []int64{0, math.MinInt64, math.MinInt64 + 1, math.MaxInt64 - 1, math.MaxInt64, 1, -1} {
			inputs = append(inputs, testInputMul{a: a, b: b})
		}
	}

	for i, input := range inputs {
		a, b := input.a, input.b
		t.Run(fmt.Sprintf(`%3d: Mul(%d,%d)`, i, a, b), func(t *testing.T) {
			want, shouldErr := getWantMul(a, b)

			got, err := Mul(input.a, input.b)

			assert.Equal(t, shouldErr, errors.Is(err, ErrOverflow))
			assert.Equal(t, want, got)
		})
	}

}

func TestMul_Small(t *testing.T) {
	var inputs []testInputMul
	for a := -24; a <= 24; a++ {
		for b := -24; b <= 24; b++ {
			inputs = append(inputs, testInputMul{a: int64(a), b: int64(b)})
		}
	}

	for i, input := range inputs {
		a, b := input.a, input.b
		t.Run(fmt.Sprintf(`%3d: Mul(%d,%d)`, i, a, b), func(t *testing.T) {
			want, shouldErr := getWantMul(a, b)

			got, err := Mul(input.a, input.b)

			assert.Equal(t, shouldErr, errors.Is(err, ErrOverflow))
			assert.Equal(t, want, got)
		})
	}

}

func TestMul_Random(t *testing.T) {
	var inputs []testInputMul
	r := rand.New(rand.NewSource(1))
	for i := 0; i < 1000; i++ {
		inputs = append(inputs, testInputMul{a: r.Int63(), b: r.Int63()})
	}

	for i, input := range inputs {
		a, b := input.a, input.b
		t.Run(fmt.Sprintf(`%3d: Mul(%d,%d)`, i, a, b), func(t *testing.T) {
			want, shouldErr := getWantMul(a, b)

			got, err := Mul(a, b)

			assert.Equal(t, shouldErr, errors.Is(err, ErrOverflow))
			assert.Equal(t, want, got)
		})
	}
}
