package duration

import (
	"math"
	"math/big"
)

func Decompose(x *big.Int) (sec int64, nano int64, ok bool) {
	ok = true
	s, n := (&big.Int{}).DivMod(x, big.NewInt(1_000_000_000), &big.Int{})
	if s.Cmp(big.NewInt(math.MinInt64)) < 0 || s.Cmp(big.NewInt(math.MaxInt64)) > 0 {
		ok = false
	}
	if n.Cmp(big.NewInt(math.MinInt64)) < 0 || n.Cmp(big.NewInt(math.MaxInt64)) > 0 {
		ok = false
	}
	return s.Int64(), n.Int64(), ok
}
