package calender_test

import (
	"bytes"
	_ "embed"
	"fmt"
	. "github.com/Jumpaku/gkoku/calender"
	"github.com/Jumpaku/gkoku/internal/tests"
	"github.com/stretchr/testify/assert"
	"testing"
)

//go:embed testcases/testdata/yearmonth_yyyymm.txt
var testdataYearMonthYyyyMm []byte

func TestYearMonth_YyyyMm(t *testing.T) {
	type testcase struct {
		sutYear, sutMonth   int
		wantYear, wantMonth int
	}

	s := tests.Scanner{Data: bytes.NewBuffer(testdataYearMonthYyyyMm)}
	nTestcases := s.ScanInt()
	var testcases []testcase
	for i := 0; i < nTestcases; i++ {
		ints := s.ScanInts(4)
		testcases = append(testcases, testcase{
			sutYear: ints[0], sutMonth: ints[1],
			wantYear: ints[2], wantMonth: ints[3],
		})
	}

	for _, testcase := range testcases {
		t.Run(fmt.Sprintf("%d-%d", testcase.wantYear, testcase.wantMonth), func(t *testing.T) {
			sut := YearMonthOf(testcase.sutYear, Month(testcase.sutMonth))
			{
				gotYear, gotMonth := sut.YyyyMm()
				assert.Equal(t, testcase.wantYear, gotYear)
				assert.Equal(t, testcase.wantMonth, int(gotMonth))
			}
		})
	}
}

//go:embed testcases/testdata/yearmonth_day.txt
var testdataYearMonthDay []byte

func TestYearMonth_Day(t *testing.T) {
	type testcase struct {
		sutYear, sutMonth int
		lastDay           int
	}

	s := tests.Scanner{Data: bytes.NewBuffer(testdataYearMonthDay)}
	nTestcases := s.ScanInt()
	var testcases []testcase
	for i := 0; i < nTestcases; i++ {
		ints := s.ScanInts(3)
		testcases = append(testcases, testcase{
			sutYear: ints[0], sutMonth: ints[1],
			lastDay: ints[2],
		})
	}

	for _, testcase := range testcases {
		sut := YearMonthOf(testcase.sutYear, Month(testcase.sutMonth))
		t.Run(fmt.Sprintf("%d-%d", testcase.sutYear, testcase.sutMonth), func(t *testing.T) {
			{
				got := sut.Days()
				assert.Equal(t, testcase.lastDay, got)
			}
			{
				got := sut.ContainsDay(-1)
				assert.Equal(t, false, got)
			}
			{
				got := sut.ContainsDay(0)
				assert.Equal(t, false, got)
			}
			{
				got := sut.ContainsDay(1)
				assert.Equal(t, true, got)
			}
			{
				got := sut.ContainsDay(28)
				assert.Equal(t, true, got)
			}
			{
				got := sut.ContainsDay(29)
				assert.Equal(t, 29 <= testcase.lastDay, got)
			}
			{
				got := sut.ContainsDay(30)
				assert.Equal(t, 30 <= testcase.lastDay, got)
			}
			{
				got := sut.ContainsDay(31)
				assert.Equal(t, 31 <= testcase.lastDay, got)
			}
			{
				got := sut.ContainsDay(32)
				assert.Equal(t, false, got)
			}
		})
	}
}

func TestYearMonth_MonthsUntil(t *testing.T) {

}

func TestYearMonth_WholeYearsUntil(t *testing.T) {

}

//go:embed testcases/testdata/yearmonth_compare.txt
var testdataYearMonthCompare []byte

func TestYearMonth_Compare(t *testing.T) {
	type testcase struct {
		sutYear, sutMonth int
		inYear, inMonth   int
		want              int
	}

	s := tests.Scanner{Data: bytes.NewBuffer(testdataYearMonthCompare)}
	nTestcases := s.ScanInt()
	var testcases []testcase
	for i := 0; i < nTestcases; i++ {
		ints := s.ScanInts(5)
		testcases = append(testcases, testcase{
			sutYear: ints[0], sutMonth: ints[1],
			inYear: ints[2], inMonth: ints[3],
			want: ints[4],
		})
	}

	for _, testcase := range testcases {
		t.Run(fmt.Sprintf("%d-%d <=> %d-%d", testcase.sutYear, testcase.sutMonth, testcase.inYear, testcase.inMonth), func(t *testing.T) {
			sut := YearMonthOf(testcase.sutYear, Month(testcase.sutMonth))
			in := YearMonthOf(testcase.inYear, Month(testcase.inMonth))
			{
				got := sut.Cmp(in)
				assert.Equal(t, testcase.want, got)
			}
			{
				got := sut.Equal(in)
				assert.Equal(t, testcase.want == 0, got)
			}
			{
				got := sut.Before(in)
				assert.Equal(t, testcase.want < 0, got)
			}
			{
				got := sut.After(in)
				assert.Equal(t, testcase.want > 0, got)
			}
		})
	}
}
