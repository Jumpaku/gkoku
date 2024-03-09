package main

import (
	"fmt"
	"github.com/Jumpaku/gkoku/calendar/testcases/cmd/date"
	"github.com/Jumpaku/gkoku/internal/console"
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

	dates := date.GetExampleDates()

	periods := []date.PeriodField{}
	// positive
	{
		pDay := []int{
			0,
			1, 2, 3, 4,
			27, 28, 29, 30, 31,
		}
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
				for _, d := range pDay {
					periods = append(periods, date.PeriodField{Years: y, Months: m, Days: d})
				}
			}
		}
	}
	// negative
	{
		pDay := []int{
			-31, -30, -29, -28, -27,
			-4, -3, -2, -1,
			0,
		}
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
				for _, d := range pDay {
					periods = append(periods, date.PeriodField{Years: y, Months: m, Days: d})
				}
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
	nTestcases := 1_000_000
	fmt.Fprintln(out, nTestcases)
	r := rand.New(rand.NewSource(3))
	r.Shuffle(len(testcases), func(i, j int) { testcases[i], testcases[j] = testcases[j], testcases[i] })
	for _, t := range testcases[:nTestcases] {
		fmt.Fprintf(out, "%d %d %d %d %d %d %d\n",
			t.sut.Year(), t.sut.Month(), t.sut.Day(), t.in.Year(), t.in.Month(), t.in.Day(), t.want)
	}
	log.Println("Date comparison testcases successfully generated in " + outputPath)
}
