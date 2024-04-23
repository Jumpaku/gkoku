package main

import (
	"fmt"
	"github.com/Jumpaku/tokiope/date/internal/cmd/testcases/yearmonth"
	"github.com/Jumpaku/tokiope/internal/console"
	"log"
	"os"
)

func main() {
	console.PanicIf(len(os.Args) != 2, "positional argument <output_path> is required")
	outputPath := os.Args[1]
	out, err := os.Create(os.Args[1])
	console.PanicIfError(err, "failed to create output file: %s", outputPath)
	defer out.Close()

	ym := yearmonth.GetExampleYearMonths()
	in := []int{
		-1, -11, -12, -13 - 100_000,
		0,
		1, 11, 12, 13, 100_000,
	}

	fmt.Fprintln(out, len(ym)*len(in))
	for _, sut := range ym {
		for _, in := range in {
			fmt.Fprintf(out, "%d %d %d\n",
				sut.Year, sut.Month, in)
		}
	}
	log.Println("YearMonth add sub testcases successfully generated in " + outputPath)
}
