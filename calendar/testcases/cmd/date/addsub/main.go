package main

import (
	"fmt"
	"github.com/Jumpaku/gkoku/calendar/testcases/cmd/date"
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

	dates := date.GetExampleDates()

	in := []int{
		-1, -6, -7, -8, -27, -28, -29, -30, -31, -32, -364, -365, -366, -367, -4_000_000,
		0,
		1, 6, 7, 8, 27, 28, 29, 30, 31, 32, 364, 365, 366, 367, 4_000_000,
	}

	fmt.Fprintln(out, len(dates)*len(in))
	for _, sut := range dates {
		for _, in := range in {
			fmt.Fprintf(out, "%d %d %d %d\n",
				sut.Year, sut.Month, sut.Day, in)
		}
	}
	log.Println("Date add sub testcases successfully generated in " + outputPath)
}
