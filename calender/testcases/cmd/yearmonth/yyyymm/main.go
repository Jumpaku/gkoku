package main

import (
	"fmt"
	"github.com/Jumpaku/gkoku/calender/testcases/cmd/yearmonth"
	"github.com/Jumpaku/gkoku/internal/console"
	"log"
	"os"
)

func main() {
	console.PanicIf(len(os.Args) != 2, "positional argument <output_path> is required")
	outputPath := os.Args[1]
	out, err := os.Create(os.Args[1])
	console.PanicIfError(err, "failed to create output file: %s", outputPath)
	defer out.Close()

	yms := yearmonth.GetExampleYearMonths()

	fmt.Fprintln(out, len(yms))
	for _, sut := range yms {
		// year month
		fmt.Fprintf(out, "%d %d %d %d\n", sut.Year, sut.Month, sut.Year, sut.Month)
	}
	log.Println("YearMonth testcases successfully generated in " + outputPath)
}
