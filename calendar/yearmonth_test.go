package calendar_test

import (
	"bytes"
	_ "embed"
	"fmt"
	. "github.com/Jumpaku/gkoku/calendar"
	"github.com/Jumpaku/gkoku/internal/tests"
	calendar_test "github.com/Jumpaku/gkoku/internal/tests/calendar"
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

type testcaseYearMonthDay struct {
	sutYear, sutMonth int
	lastDay           int
}

var testcasesYearMonthDay = func() (testcases []testcaseYearMonthDay) {
	s := tests.Scanner{Data: bytes.NewBuffer(testdataYearMonthDay)}
	nTestcases := s.ScanInt()
	for i := 0; i < nTestcases; i++ {
		ints := s.ScanInts(3)
		testcases = append(testcases, testcaseYearMonthDay{
			sutYear: ints[0], sutMonth: ints[1],
			lastDay: ints[2],
		})
	}
	return testcases
}()

func TestYearMonth_Days(t *testing.T) {
	for _, testcase := range testcasesYearMonthDay {
		sut := YearMonthOf(testcase.sutYear, Month(testcase.sutMonth))
		t.Run(fmt.Sprintf("%d-%d", testcase.sutYear, testcase.sutMonth), func(t *testing.T) {
			{
				got := sut.Days()
				assert.Equal(t, testcase.lastDay, got)
			}
		})
	}
}

func TestYearMonth_ContainsDay(t *testing.T) {
	for _, testcase := range testcasesYearMonthDay {
		sut := YearMonthOf(testcase.sutYear, Month(testcase.sutMonth))
		t.Run(fmt.Sprintf("whether %d-%d contains -1", testcase.sutYear, testcase.sutMonth), func(t *testing.T) {
			got := sut.ContainsDay(-1)
			assert.Equal(t, false, got)
		})
		t.Run(fmt.Sprintf("whether %d-%d contains 0", testcase.sutYear, testcase.sutMonth), func(t *testing.T) {
			got := sut.ContainsDay(0)
			assert.Equal(t, false, got)
		})
		t.Run(fmt.Sprintf("whether %d-%d contains 1", testcase.sutYear, testcase.sutMonth), func(t *testing.T) {
			got := sut.ContainsDay(1)
			assert.Equal(t, true, got)
		})
		t.Run(fmt.Sprintf("whether %d-%d contains 28", testcase.sutYear, testcase.sutMonth), func(t *testing.T) {
			got := sut.ContainsDay(28)
			assert.Equal(t, true, got)
		})
		t.Run(fmt.Sprintf("whether %d-%d contains 29", testcase.sutYear, testcase.sutMonth), func(t *testing.T) {
			got := sut.ContainsDay(29)
			assert.Equal(t, 29 <= testcase.lastDay, got)
		})
		t.Run(fmt.Sprintf("whether %d-%d contains 30", testcase.sutYear, testcase.sutMonth), func(t *testing.T) {
			got := sut.ContainsDay(30)
			assert.Equal(t, 30 <= testcase.lastDay, got)
		})
		t.Run(fmt.Sprintf("whether %d-%d contains 31", testcase.sutYear, testcase.sutMonth), func(t *testing.T) {
			got := sut.ContainsDay(31)
			assert.Equal(t, 31 <= testcase.lastDay, got)
		})
		t.Run(fmt.Sprintf("whether %d-%d contains 32", testcase.sutYear, testcase.sutMonth), func(t *testing.T) {
			got := sut.ContainsDay(32)
			assert.Equal(t, false, got)
		})
	}
}

func TestYearMonth_Date(t *testing.T) {
	for _, testcase := range testcasesYearMonthDay {
		sut := YearMonthOf(testcase.sutYear, Month(testcase.sutMonth))
		for day := 1; day <= sut.Days(); day++ {
			t.Run(fmt.Sprintf("%d-%d-%d", testcase.sutYear, testcase.sutMonth, day), func(t *testing.T) {
				got := sut.Date(day)
				want := YyyyMmDd(testcase.sutYear, Month(testcase.sutMonth), day)
				calendar_test.AssertEqualDate(t, want, got)
			})
		}
	}
}

func TestYearMonth_FirstDate(t *testing.T) {
	for _, testcase := range testcasesYearMonthDay {
		sut := YearMonthOf(testcase.sutYear, Month(testcase.sutMonth))
		t.Run(fmt.Sprintf("%d-%d", testcase.sutYear, testcase.sutMonth), func(t *testing.T) {
			got := sut.FirstDate()
			want := YyyyMmDd(testcase.sutYear, Month(testcase.sutMonth), 1)
			calendar_test.AssertEqualDate(t, want, got)
		})
	}
}

func TestYearMonth_LastDate(t *testing.T) {
	for _, testcase := range testcasesYearMonthDay {
		sut := YearMonthOf(testcase.sutYear, Month(testcase.sutMonth))
		t.Run(fmt.Sprintf("%d-%d", testcase.sutYear, testcase.sutMonth), func(t *testing.T) {
			got := sut.LastDate()
			want := YyyyMmDd(testcase.sutYear, Month(testcase.sutMonth), sut.Days())
			calendar_test.AssertEqualDate(t, want, got)
		})
	}
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
		sut := YearMonthOf(testcase.sutYear, Month(testcase.sutMonth))
		in := YearMonthOf(testcase.inYear, Month(testcase.inMonth))
		t.Run(fmt.Sprintf("%d-%d <=> %d-%d", testcase.sutYear, testcase.sutMonth, testcase.inYear, testcase.inMonth), func(t *testing.T) {
			got := sut.Cmp(in)
			assert.Equal(t, testcase.want, got)
		})
		t.Run(fmt.Sprintf("%d-%d == %d-%d", testcase.sutYear, testcase.sutMonth, testcase.inYear, testcase.inMonth), func(t *testing.T) {
			got := sut.Equal(in)
			assert.Equal(t, testcase.want == 0, got)
		})
		t.Run(fmt.Sprintf("%d-%d < %d-%d", testcase.sutYear, testcase.sutMonth, testcase.inYear, testcase.inMonth), func(t *testing.T) {
			got := sut.Before(in)
			assert.Equal(t, testcase.want < 0, got)
		})
		t.Run(fmt.Sprintf("%d-%d > %d-%d", testcase.sutYear, testcase.sutMonth, testcase.inYear, testcase.inMonth), func(t *testing.T) {
			got := sut.After(in)
			assert.Equal(t, testcase.want > 0, got)
		})
	}
}

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
			calendar_test.AssertEqualYearMonth(t, YearMonthOf(want.Year(), Month(want.Month())), got)
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
			calendar_test.AssertEqualYearMonth(t, YearMonthOf(want.Year(), Month(want.Month())), got)
		})
	}
}
