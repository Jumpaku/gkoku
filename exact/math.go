package exact

import (
	"math"
)

func signAndAbs(a int64) (sign int, abs uint64) {
	if a == math.MinInt64 {
		return -1, 1 << 63
	}
	if a < 0 {
		return -1, uint64(-a)
	}
	if a > 0 {
		return 1, uint64(a)
	}
	return 0, 0
}
