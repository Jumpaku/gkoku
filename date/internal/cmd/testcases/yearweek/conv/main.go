package main

import (
	"fmt"
	"github.com/Jumpaku/gkoku/date/internal/cmd/testcases/yearweek"
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

	yws := yearweek.GetExampleYearWeeks()

	fmt.Fprintln(out, len(yws))
	for _, sut := range yws {
		// year month
		fmt.Fprintf(out, "%d %d\n", sut.Year, sut.Week)
	}
	log.Println("YearWeek conversion testcases successfully generated in " + outputPath)
}
