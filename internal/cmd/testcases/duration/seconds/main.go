package main

import (
	"fmt"
	"github.com/Jumpaku/gkoku/internal/cmd/testcases"
	"log"
	"math"
	"math/big"
	"math/rand"
	"os"

	"github.com/Jumpaku/gkoku/internal/console"
)

func main() {
	console.PanicIf(len(os.Args) != 2, "positional argument <output_path> is required")
	outputPath := os.Args[1]
	out, err := os.Create(os.Args[1])
	console.PanicIfError(err, "failed to create output file: %s", outputPath)
	defer out.Close()

	nanos := []int64{0, math.MinInt64, math.MinInt64 + 1, math.MaxInt64 - 1, math.MaxInt64, 1, -1, 1_000_000_000, -1_000_000_000, 1_000_000_001, -1_000_000_001, 999_999_999, -999_999_999}
	r := rand.New(rand.NewSource(1))
	order := int64(10)
	for i := 0; i < 18; i++ {
		sign := r.Int63n(2)*2 - 1
		nanos = append(nanos, sign*r.Int63n(order))
		order *= 10
	}

	fmt.Fprintln(out, len(nanos))
	for _, nano := range nanos {
		sutSec, sutNano, ok := testcases.Decompose(big.NewInt(nano))
		if !ok {
			log.Panicf("%+v", nano)
		}
		wantSeconds, wantNano := (&big.Int{}).DivMod(big.NewInt(nano), big.NewInt(1_000_000_000), &big.Int{})
		fmt.Fprintf(out, "%d %d %d %d\n", sutSec, sutNano, wantSeconds, wantNano)
	}
	log.Println("func (Duration) Seconds testcases successfully generated in " + outputPath)
}
