package main

import (
	"fmt"
	"github.com/Jumpaku/tokiope/date/internal/cmd/testcases/date"
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

	yms := yearmonth.GetExampleYearMonths()

	fmt.Fprintln(out, len(yms))
	for _, sut := range yms {
		sut := date.GoDate(sut.Year, sut.Month, 32)
		sut = sut.AddDate(0, 0, -sut.Day())
		// year month lastDayOfMonth
		fmt.Fprintf(out, "%d %d %d\n", sut.Year(), sut.Month(), sut.Day())
	}
	log.Println("YearMonth day testcases successfully generated in " + outputPath)
}
