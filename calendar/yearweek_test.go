package calendar_test

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/Jumpaku/tokiope/calendar"
	"github.com/Jumpaku/tokiope/internal/tests"
	"github.com/Jumpaku/tokiope/internal/tests/assert"
	calendar_test "github.com/Jumpaku/tokiope/internal/tests/date"
	"testing"
)

//go:embed testdata/yearweek_yyyyww.txt
var testdataYearWeekYyyyWw []byte

func TestYearWeek_YW(t *testing.T) {
	type testcase struct {
		sutYear, sutWeek int
	}

	s := tests.Scanner{Data: bytes.NewBuffer(testdataYearWeekYyyyWw)}
	nTestcases := s.ScanInt()
	var testcases []testcase
	for i := 0; i < nTestcases; i++ {
		ints := s.ScanInts(2)
		testcases = append(testcases, testcase{
			sutYear: ints[0], sutWeek: ints[1],
		})
	}

	for _, testcase := range testcases {
		t.Run(fmt.Sprintf("%d-%d", testcase.sutYear, testcase.sutWeek), func(t *testing.T) {
			sut := calendar.YearWeekOf(testcase.sutYear, testcase.sutWeek)
			{
				gotYear, gotWeek := sut.YW()
				assert.Equal(t, testcase.sutYear, gotYear)
				assert.Equal(t, testcase.sutWeek, gotWeek)
			}
			{
				gotYear := sut.Year()
				assert.Equal(t, testcase.sutYear, int(gotYear))
			}
		})
	}
}

//go:embed testdata/yearweek_day.txt
var testdataYearWeekDay []byte

type testcaseYearWeek struct {
	sutYear, sutWeek int
}

var testcasesYearWeek = func() (testcases []testcaseYearWeek) {
	s := tests.Scanner{Data: bytes.NewBuffer(testdataYearWeekDay)}
	nTestcases := s.ScanInt()
	for i := 0; i < nTestcases; i++ {
		ints := s.ScanInts(2)
		testcases = append(testcases, testcaseYearWeek{
			sutYear: ints[0], sutWeek: ints[1],
		})
	}
	return testcases
}()

func TestYearWeek_Date(t *testing.T) {
	for _, testcase := range testcasesYearWeek {
		sut := calendar.YearWeekOf(testcase.sutYear, testcase.sutWeek)
		for day := calendar.DayOfWeekMonday; day <= calendar.DayOfWeekSunday; day++ {
			t.Run(fmt.Sprintf("%d-W%d-%d", testcase.sutYear, testcase.sutWeek, int(day)), func(t *testing.T) {
				got := sut.Date(day)
				want := calendar.DateOfYWD(testcase.sutYear, testcase.sutWeek, day)
				calendar_test.AssertEqualDate(t, want, got)
			})
		}
	}
}

func TestYearWeek_FirstDate(t *testing.T) {
	for _, testcase := range testcasesYearWeek {
		sut := calendar.YearWeekOf(testcase.sutYear, testcase.sutWeek)
		t.Run(fmt.Sprintf("%d-%d", testcase.sutYear, testcase.sutWeek), func(t *testing.T) {
			got := sut.FirstDate()
			want := calendar.DateOfYWD(testcase.sutYear, testcase.sutWeek, calendar.DayOfWeekMonday)
			calendar_test.AssertEqualDate(t, want, got)
		})
	}
}

func TestYearWeek_LastDate(t *testing.T) {
	for _, testcase := range testcasesYearWeek {
		sut := calendar.YearWeekOf(testcase.sutYear, testcase.sutWeek)
		t.Run(fmt.Sprintf("%d-%d", testcase.sutYear, testcase.sutWeek), func(t *testing.T) {
			got := sut.LastDate()
			want := calendar.DateOfYWD(testcase.sutYear, testcase.sutWeek, calendar.DayOfWeekSunday)
			calendar_test.AssertEqualDate(t, want, got)
		})
	}
}

//go:embed testdata/yearweek_compare.txt
var testdataYearWeekCompare []byte

func TestYearWeek_Compare(t *testing.T) {
	type testcase struct {
		sutYear, sutWeek int
		inYear, inWeek   int
		want             int
	}

	s := tests.Scanner{Data: bytes.NewBuffer(testdataYearWeekCompare)}
	nTestcases := s.ScanInt()
	var testcases []testcase
	for i := 0; i < nTestcases; i++ {
		ints := s.ScanInts(5)
		testcases = append(testcases, testcase{
			sutYear: ints[0], sutWeek: ints[1],
			inYear: ints[2], inWeek: ints[3],
			want: ints[4],
		})
	}

	for _, testcase := range testcases {
		sut := calendar.YearWeekOf(testcase.sutYear, testcase.sutWeek)
		in := calendar.YearWeekOf(testcase.inYear, testcase.inWeek)
		t.Run(fmt.Sprintf("%d-W%d <=> %d-W%d", testcase.sutYear, testcase.sutWeek, testcase.inYear, testcase.inWeek), func(t *testing.T) {
			got := sut.Cmp(in)
			assert.Equal(t, testcase.want, got)
		})
		t.Run(fmt.Sprintf("%d-W%d == %d-W%d", testcase.sutYear, testcase.sutWeek, testcase.inYear, testcase.inWeek), func(t *testing.T) {
			got := sut.Equal(in)
			assert.Equal(t, testcase.want == 0, got)
		})
		t.Run(fmt.Sprintf("%d-W%d < %d-W%d", testcase.sutYear, testcase.sutWeek, testcase.inYear, testcase.inWeek), func(t *testing.T) {
			got := sut.Before(in)
			assert.Equal(t, testcase.want < 0, got)
		})
		t.Run(fmt.Sprintf("%d-W%d > %d-W%d", testcase.sutYear, testcase.sutWeek, testcase.inYear, testcase.inWeek), func(t *testing.T) {
			got := sut.After(in)
			assert.Equal(t, testcase.want > 0, got)
		})
	}
}
