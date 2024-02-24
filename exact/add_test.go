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

func getWantAdd(a, b int64) (want int64, shouldErr bool) {
	w := (&big.Int{}).Add(big.NewInt(a), big.NewInt(b))
	shouldErr = w.Cmp(big.NewInt(math.MinInt64)) < 0 || w.Cmp(big.NewInt(math.MaxInt64)) > 0
	want = w.Int64()
	if shouldErr {
		want = a + b
	}
	return
}

type testInputAdd struct {
	a int64
	b int64
}

func TestAdd_Corner(t *testing.T) {
	var inputs []testInputAdd
	for _, a := range []int64{0, math.MinInt64, math.MinInt64 + 1, math.MaxInt64 - 1, math.MaxInt64, 1, -1} {
		for _, b := range []int64{0, math.MinInt64, math.MinInt64 + 1, math.MaxInt64 - 1, math.MaxInt64, 1, -1} {
			inputs = append(inputs, testInputAdd{a: a, b: b})
		}
	}

	for i, input := range inputs {
		a, b := input.a, input.b
		t.Run(fmt.Sprintf(`%3d: Add(%d,%d)`, i, a, b), func(t *testing.T) {
			want, shouldErr := getWantAdd(a, b)

			got, err := Add(input.a, input.b)

			assert.Equal(t, shouldErr, errors.Is(err, ErrOverflow))
			assert.Equal(t, want, got)
		})
	}

}

func TestAdd_Small(t *testing.T) {
	var inputs []testInputAdd
	for a := -24; a <= 24; a++ {
		for b := -24; b <= 24; b++ {
			inputs = append(inputs, testInputAdd{a: int64(a), b: int64(b)})
		}
	}

	for i, input := range inputs {
		a, b := input.a, input.b
		t.Run(fmt.Sprintf(`%3d: Add(%d,%d)`, i, a, b), func(t *testing.T) {
			want, shouldErr := getWantAdd(a, b)

			got, err := Add(input.a, input.b)

			assert.Equal(t, shouldErr, errors.Is(err, ErrOverflow))
			assert.Equal(t, want, got)
		})
	}

}

func TestAdd_Random(t *testing.T) {
	var inputs []testInputAdd
	r := rand.New(rand.NewSource(1))
	for i := 0; i < 1000; i++ {
		inputs = append(inputs, testInputAdd{a: r.Int63(), b: r.Int63()})
	}

	for i, input := range inputs {
		a, b := input.a, input.b
		t.Run(fmt.Sprintf(`%3d: Add(%d,%d)`, i, a, b), func(t *testing.T) {
			want, shouldErr := getWantAdd(a, b)

			got, err := Add(a, b)

			assert.Equal(t, shouldErr, errors.Is(err, ErrOverflow))
			assert.Equal(t, want, got)
		})
	}
}
