package main

import (
	"fmt"
	"github.com/Jumpaku/tokiope/date/internal/cmd/testcases/date"
	"github.com/Jumpaku/tokiope/internal/console"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	console.PanicIf(len(os.Args) != 2, "positional argument <output_path> is required")
	outputPath := os.Args[1]
	out, err := os.Create(os.Args[1])
	console.PanicIfError(err, "failed to create output file: %s", outputPath)
	defer out.Close()

	years := []int{
		-9999,
		-1,
		1,
		9999,
	}
	for y := 1900; y <= 2100; y++ {
		years = append(years, y)
	}
	for y := -400; y <= 2400; y += 100 {
		years = append(years, y)
	}
	r := rand.New(rand.NewSource(1))
	for i := 0; i < 71; i++ {
		sign := r.Intn(2)*2 - 1
		years = append(years, sign*r.Intn(10000))
	}

	var n int
	for _, year := range years {
		n += date.GoDate(year, 12, 31).YearDay()
	}
	fmt.Fprintln(out, n)
	for _, year := range years {
		y0101 := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
		for yd := 0; yd < date.GoDate(year, 12, 31).YearDay(); yd++ {
			in := y0101.AddDate(0, 0, yd)
			y, m, d := in.Date()

			fmt.Fprintf(out, "%d %d %d\n", y, m, d)
		}
	}
	log.Println("Date conversion testcases successfully generated in " + outputPath)
}
