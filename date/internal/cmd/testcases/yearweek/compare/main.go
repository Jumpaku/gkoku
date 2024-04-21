package main

import (
	"fmt"
	"github.com/Jumpaku/gkoku/date/internal/cmd/testcases/date"
	"github.com/Jumpaku/gkoku/internal/console"
	"log"
	"os"
	"time"
)

func main() {
	console.PanicIf(len(os.Args) != 2, "positional argument <output_path> is required")
	outputPath := os.Args[1]
	out, err := os.Create(os.Args[1])
	console.PanicIfError(err, "failed to create output file: %s", outputPath)
	defer out.Close()

	dates := date.GetExampleDates()

	periods := []date.PeriodField{}
	// positive
	{
		pDay := []int{
			0,
			1, 6, 7, 8,
			360, 361, 362, 363, 364, 365, 366, 367, 368, 369, 370,
		}
		pYear := []int{
			0,
			1, 2, 3, 4, 5,
			99, 100, 101,
			399, 400, 401,
			9999,
		}
		for _, y := range pYear {
			for _, d := range pDay {
				periods = append(periods, date.PeriodField{Years: y, Months: 0, Days: d})
			}
		}
	}
	// negative
	{
		pDay := []int{

			-370, -369, -368, -367, -366, -365, -364, -363, -362, -361, -360,
			-8, -7, -6, -1,
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
			for _, d := range pDay {
				periods = append(periods, date.PeriodField{Years: y, Months: 0, Days: d})
			}
		}
	}

	type testcase struct {
		sut  time.Time
		in   time.Time
		want int
	}
	testcases := []testcase{}
	for _, sut := range dates {
		sut := date.GoDate(sut.Year, sut.Month, sut.Day)
		for _, p := range periods {
			in := sut.AddDate(p.Years, p.Months, p.Days)
			testcases = append(testcases, testcase{
				sut:  sut,
				in:   in,
				want: sut.Compare(in),
			})
		}
	}
	fmt.Fprintln(out, len(testcases))
	for _, t := range testcases {
		sutY, sutW := t.sut.ISOWeek()
		inY, inW := t.sut.ISOWeek()
		if sutY == inY && sutW == inW {
			t.want = 0
		}
		fmt.Fprintf(out, "%d %d %d %d %d\n", sutY, sutW, inY, inW, t.want)
	}
	log.Println("YearWeek comparison testcases successfully generated in " + outputPath)
}
