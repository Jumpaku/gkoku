package main

import (
	"fmt"
	"github.com/Jumpaku/gkoku/calendar/testcases/cmd/date"
	"github.com/Jumpaku/gkoku/calendar/testcases/cmd/yearmonth"
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

	periods := []date.PeriodField{}
	// positive
	{
		pMonth := []int{
			0,
			1, 2, 10, 11,
		}
		pYear := []int{
			0,
			1, 2, 3, 4, 5,
			99, 100, 101,
			399, 400, 401,
			9999,
		}
		for _, y := range pYear {
			for _, m := range pMonth {
				periods = append(periods, date.PeriodField{Years: y, Months: m})
			}
		}
	}
	// negative
	{
		pMonth := []int{
			-11, -10, -2, -1,
			0,
		}
		pYear := []int{
			-9999,
			-401, -400, -399,
			-101, -100, -99,
			-5, -4, -3, -2, -1,
			0,
		}
		for _, y := range pYear {
			for _, m := range pMonth {
				periods = append(periods, date.PeriodField{Years: y, Months: m})
			}
		}
	}

	fmt.Fprintln(out, len(yms)*len(periods))
	for _, sut := range yms {
		sut := date.GoDate(sut.Year, sut.Month, 1)
		for _, p := range periods {
			in := sut.AddDate(p.Years, p.Months, 0)
			fmt.Fprintf(out, "%d %d %d %d %d\n", sut.Year(), sut.Month(), in.Year(), in.Month(), sut.Compare(in))
		}
	}
	log.Println("YearMonth comparison testcases successfully generated in " + outputPath)
}
