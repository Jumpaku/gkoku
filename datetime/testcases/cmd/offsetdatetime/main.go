package main

import (
	"fmt"
	"github.com/Jumpaku/gkoku/internal/console"
	"log"
	"math/rand"
	"os"
	"time"
)

func RandRange(r *rand.Rand, a, b int) int {
	return r.Intn(b+1-a) + a
}

func main() {
	console.PanicIf(len(os.Args) != 2, "positional argument <output_path> is required")
	outputPath := os.Args[1]
	out, err := os.Create(os.Args[1])
	console.PanicIfError(err, "failed to create output file: %s", outputPath)
	defer out.Close()

	n := 1_000_000
	fmt.Fprintln(out, n)
	r := rand.New(rand.NewSource(4))
	for i := 0; i < n; i++ {
		year := RandRange(r, -10000, 10000)
		month := RandRange(r, 1, 12)
		day := RandRange(r, 0, 31)
		hour := RandRange(r, 0, 24)
		minute := RandRange(r, 0, 60)
		second := RandRange(r, 0, 60)
		offset := RandRange(r, -14, 14) * 60

		loc := time.FixedZone(fmt.Sprintf(`zone_%d`, i), offset*60)
		sut := time.Date(year, time.Month(month), day, hour, minute, second, 0, loc)
		want := sut.Unix()
		fmt.Fprintf(out, "%d %d %d %d %d %d %d %d\n",
			sut.Year(), int(sut.Month()), sut.Day(), sut.Hour(), sut.Minute(), sut.Second(), offset, want)
	}

	log.Println("OffsetDateTime Instant() testcases successfully generated in " + outputPath)
}
