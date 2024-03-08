package calender_test

import (
	"bytes"
	_ "embed"
	"fmt"
	. "github.com/Jumpaku/gkoku/calender"
	"github.com/Jumpaku/gkoku/internal/tests"
	"github.com/Jumpaku/gkoku/internal/tests/assert"
	"testing"
)

//go:embed testcases/testdata/yearmonth_until.txt
var testdataYearMonthUntil []byte

type testcaseYearMonthUntil struct {
	sutYear, sutMonth int
	inYear, inMonth   int
}

var testcasesYearMonthUntil = func() (testcases []testcaseYearMonthUntil) {
	s := tests.Scanner{Data: bytes.NewBuffer(testdataYearMonthUntil)}
	nTestcases := s.ScanInt()
	for i := 0; i < nTestcases; i++ {
		ints := s.ScanInts(4)
		testcases = append(testcases, testcaseYearMonthUntil{
			sutYear: ints[0], sutMonth: ints[1],
			inYear: ints[2], inMonth: ints[3],
		})
	}
	return testcases
}()

func TestYearMonth_MonthsUntil(t *testing.T) {
	for _, testcase := range testcasesYearMonthUntil {
		name := fmt.Sprintf("(%d-%d)-(%d-%d)", testcase.inYear, testcase.inMonth, testcase.sutYear, testcase.sutMonth)
		t.Run(name, func(t *testing.T) {
			sut := YearMonthOf(testcase.sutYear, Month(testcase.sutMonth))
			in := YearMonthOf(testcase.inYear, Month(testcase.inMonth))
			until := sut.MonthsUntil(in)
			want := ToTime(testcase.sutYear, Month(testcase.sutMonth), 1).AddDate(0, int(until), 0)
			{
				assert.Equal(t, want.Year(), testcase.inYear)
				assert.Equal(t, int(want.Month()), testcase.inMonth)
			}
		})
	}
}

func TestYearMonth_WholeYearsUntil(t *testing.T) {
	for _, testcase := range testcasesYearMonthUntil {
		name := fmt.Sprintf("(%d-%d)-(%d-%d)", testcase.inYear, testcase.inMonth, testcase.sutYear, testcase.sutMonth)
		t.Run(name, func(t *testing.T) {
			sut := YearMonthOf(testcase.sutYear, Month(testcase.sutMonth))
			in := YearMonthOf(testcase.inYear, Month(testcase.inMonth))
			until := sut.WholeYearsUntil(in)
			sutTime := ToTime(testcase.sutYear, Month(testcase.sutMonth), 1)
			inTime := ToTime(testcase.inYear, Month(testcase.inMonth), 1)
			switch {
			case until == 0:
				est2 := sutTime.AddDate(int(until+1), 0, 0)
				if !est2.After(inTime) {
					t.Errorf(`<sut> + <until> + 1y must be after <in>: got %d`, until)
				}

				est := sutTime.AddDate(int(until-1), 0, 0)
				if !est.Before(inTime) {
					t.Errorf(`<sut> + <until> - 1y must be before <in>: got %d`, until)
				}
			case until >= 0:
				est := sutTime.AddDate(int(until), 0, 0)
				if est.After(inTime) {
					t.Errorf(`<sut> + <until> must be not after <in>: got %d`, until)
				}

				est2 := sutTime.AddDate(int(until+1), 0, 0)
				if !est2.After(inTime) {
					t.Errorf(`<sut> + <until> + 1y must be after <in>: got %d`, until)
				}
			case until <= 0:
				est := sutTime.AddDate(int(until), 0, 0)
				if est.Before(inTime) {
					t.Errorf(`<sut> + <until> must be not before <in>: got %d`, until)
				}

				est2 := sutTime.AddDate(int(until-1), 0, 0)
				if !est2.Before(inTime) {
					t.Errorf(`<sut> + <until> - 1y must be before <in>: got %d`, until)
				}
			}
		})
	}
}

//go:embed testcases/testdata/yearmonth_addsub.txt
var testdataYearMonthAddSub []byte

type testcaseYearMonthAddSub struct {
	sutYear, sutMonth int
	in                int
}

var testcasesYearMonthAddSub = func() (testcases []testcaseYearMonthAddSub) {
	s := tests.Scanner{Data: bytes.NewBuffer(testdataYearMonthAddSub)}
	nTestcases := s.ScanInt()
	for i := 0; i < nTestcases; i++ {
		ints := s.ScanInts(3)
		testcases = append(testcases, testcaseYearMonthAddSub{
			sutYear: ints[0], sutMonth: ints[1],
			in: ints[2],
		})
	}
	return testcases
}()

func TestYearMonth_Add(t *testing.T) {
	for _, testcase := range testcasesYearMonthAddSub {
		name := fmt.Sprintf("(%d-%d)+%dd", testcase.sutYear, testcase.sutMonth, testcase.in)
		t.Run(name, func(t *testing.T) {
			sut := YearMonthOf(testcase.sutYear, Month(testcase.sutMonth))
			got := sut.Add(testcase.in)
			want := ToTime(testcase.sutYear, Month(testcase.sutMonth), 1).AddDate(0, testcase.in, 0)
			AssertEqualYearMonth(t, YearMonthOf(want.Year(), Month(want.Month())), got)
		})
	}
}

func TestYearMonth_Sub(t *testing.T) {
	for _, testcase := range testcasesYearMonthAddSub {
		name := fmt.Sprintf("(%d-%d)+%dd", testcase.sutYear, testcase.sutMonth, testcase.in)
		t.Run(name, func(t *testing.T) {
			sut := YearMonthOf(testcase.sutYear, Month(testcase.sutMonth))
			got := sut.Sub(testcase.in)
			want := ToTime(testcase.sutYear, Month(testcase.sutMonth), 1).AddDate(0, -testcase.in, 0)
			AssertEqualYearMonth(t, YearMonthOf(want.Year(), Month(want.Month())), got)
		})
	}
}
