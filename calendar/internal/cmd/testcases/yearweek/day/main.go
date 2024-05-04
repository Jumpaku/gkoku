package main

import (
	"fmt"
	"github.com/Jumpaku/tokiope/calendar/internal/cmd/testcases/yearweek"
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

	yms := yearweek.GetExampleYearWeeks()

	fmt.Fprintln(out, len(yms))
	for _, sut := range yms {
		// year week
		fmt.Fprintf(out, "%d %d\n", sut.Year, sut.Week)
	}
	log.Println("YearWeek day testcases successfully generated in " + outputPath)
}
