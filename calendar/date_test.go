package calendar_test

import (
	"bytes"
	_ "embed"
	"fmt"
	. "github.com/Jumpaku/tokiope/calendar"
	"github.com/Jumpaku/tokiope/internal/tests"
	"github.com/Jumpaku/tokiope/internal/tests/assert"
	"github.com/Jumpaku/tokiope/internal/tests/date"
	"testing"
)

func TestUnixDay(t *testing.T) {
	testcases := []struct {
		in   int64
		want Date
	}{
		{
			in:   0,
			want: DateOfYMD(1970, 1, 1),
		},
		{
			in:   1,
			want: DateOfYMD(1970, 1, 2),
		},
		{
			in:   -1,
			want: DateOfYMD(1969, 12, 31),
		},
	}
	for _, testcase := range testcases {
		t.Run(fmt.Sprintf(`%d`, testcase.in), func(t *testing.T) {
			got := UnixDay(testcase.in)
			date.AssertEqualDate(t, testcase.want, got)
		})
	}
}

//go:embed testdata/date_yyyymmdd.txt
var testdataYyyyMmDd []byte

func Test_DateOfYMD(t *testing.T) {
	type testcase struct {
		sutYear, sutMonth, sutDay    int
		wantYear, wantMonth, wantDay int
	}

	s := tests.Scanner{Data: bytes.NewBuffer(testdataYyyyMmDd)}
	nTestcases := s.ScanInt()
	var testcases []testcase
	for i := 0; i < nTestcases; i++ {
		ints := s.ScanInts(6)
		testcases = append(testcases, testcase{
			sutYear: ints[0], sutMonth: ints[1], sutDay: ints[2],
			wantYear: ints[3], wantMonth: ints[4], wantDay: ints[5],
		})
	}

	for _, testcase := range testcases {
		t.Run(fmt.Sprintf("%d-%d-%d", testcase.wantYear, testcase.wantMonth, testcase.wantDay), func(t *testing.T) {
			sut := DateOfYMD(testcase.sutYear, Month(testcase.sutMonth), testcase.sutDay)
			{
				gotYear, gotMonth, gotDay := sut.YMD()
				assert.Equal(t, testcase.wantYear, gotYear)
				assert.Equal(t, testcase.wantMonth, int(gotMonth))
				assert.Equal(t, testcase.wantDay, gotDay)
			}
			{
				gotYear, gotMonth := sut.YearMonth().YyyyMm()
				assert.Equal(t, testcase.wantYear, gotYear)
				assert.Equal(t, testcase.wantMonth, int(gotMonth))
			}
			{
				gotYear := sut.Year()
				assert.Equal(t, Year(testcase.wantYear), gotYear)
			}
		})
	}
}

//go:embed testdata/date_yyyywwd.txt
var testdataYyyyWwD []byte

func Test_DateOfYWD(t *testing.T) {
	type testcase struct {
		sutYear, sutWeek, sutDay     int
		wantYear, wantMonth, wantDay int
	}

	s := tests.Scanner{Data: bytes.NewBuffer(testdataYyyyWwD)}
	nTestcases := s.ScanInt()
	var testcases []testcase
	for i := 0; i < nTestcases; i++ {
		ints := s.ScanInts(6)
		testcases = append(testcases, testcase{
			sutYear: ints[0], sutWeek: ints[1], sutDay: ints[2],
			wantYear: ints[3], wantMonth: ints[4], wantDay: ints[5],
		})
	}

	for _, testcase := range testcases {
		t.Run(fmt.Sprintf("%d-%d-%d", testcase.wantYear, testcase.wantMonth, testcase.wantDay), func(t *testing.T) {
			sut := DateOfYWD(testcase.sutYear, testcase.sutWeek, DayOfWeek(testcase.sutDay))
			{
				gotYear, gotMonth, gotDay := sut.YMD()
				assert.Equal(t, testcase.wantYear, gotYear)
				assert.Equal(t, testcase.wantMonth, int(gotMonth))
				assert.Equal(t, testcase.wantDay, gotDay)
			}
			{
				gotYear, gotWeek, gotDay := sut.YWD()
				assert.Equal(t, testcase.sutYear, gotYear)
				assert.Equal(t, testcase.sutWeek, gotWeek)
				assert.Equal(t, testcase.sutDay, int(gotDay))
			}
			{
				wantYear, wantWeek, _ := sut.YWD()
				gotYear, gotWeek := sut.YearWeek().YyyyWw()
				assert.Equal(t, wantYear, gotYear)
				assert.Equal(t, wantWeek, gotWeek)
			}
			{
				gotYear := sut.Year()
				assert.Equal(t, Year(testcase.wantYear), gotYear)
			}
		})
	}
}

//go:embed testdata/date_yyyyddd.txt
var testdataYyyyDdd []byte

func Test_DateOfYD(t *testing.T) {
	type testcase struct {
		sutYear, sutDay              int
		wantYear, wantMonth, wantDay int
	}

	s := tests.Scanner{Data: bytes.NewBuffer(testdataYyyyDdd)}
	nTestcases := s.ScanInt()
	var testcases []testcase
	for i := 0; i < nTestcases; i++ {
		ints := s.ScanInts(5)
		testcases = append(testcases, testcase{
			sutYear: ints[0], sutDay: ints[1],
			wantYear: ints[2], wantMonth: ints[3], wantDay: ints[4],
		})
	}

	for _, testcase := range testcases {
		t.Run(fmt.Sprintf("%d-%d-%d", testcase.wantYear, testcase.wantMonth, testcase.wantDay), func(t *testing.T) {
			sut := DateOfYD(testcase.sutYear, testcase.sutDay)
			{
				gotYear, gotMonth, gotDay := sut.YMD()
				assert.Equal(t, testcase.wantYear, gotYear)
				assert.Equal(t, testcase.wantMonth, int(gotMonth))
				assert.Equal(t, testcase.wantDay, gotDay)
			}
			{
				gotYear, gotDay := sut.YD()
				assert.Equal(t, testcase.sutYear, gotYear)
				assert.Equal(t, testcase.sutDay, gotDay)
			}
			{
				gotYear := sut.Year()
				assert.Equal(t, Year(testcase.wantYear), gotYear)
			}
		})
	}
}

//go:embed testdata/date_compare.txt
var testdataDateCompare []byte

func TestDate_Compare(t *testing.T) {
	type testcase struct {
		sutYear, sutMonth, sutDay int
		inYear, inMonth, inDay    int
		want                      int
	}

	s := tests.Scanner{Data: bytes.NewBuffer(testdataDateCompare)}
	nTestcases := s.ScanInt()
	var testcases []testcase
	for i := 0; i < nTestcases; i++ {
		ints := s.ScanInts(7)
		testcases = append(testcases, testcase{
			sutYear: ints[0], sutMonth: ints[1], sutDay: ints[2],
			inYear: ints[3], inMonth: ints[4], inDay: ints[5],
			want: ints[6],
		})
	}

	for _, testcase := range testcases {
		t.Run(fmt.Sprintf("%d-%d-%d <=> %d-%d-%d", testcase.sutYear, testcase.sutMonth, testcase.sutDay, testcase.inYear, testcase.inMonth, testcase.inDay), func(t *testing.T) {
			sut := DateOfYMD(testcase.sutYear, Month(testcase.sutMonth), testcase.sutDay)
			in := DateOfYMD(testcase.inYear, Month(testcase.inMonth), testcase.inDay)
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

//go:embed testdata/date_until.txt
var testdataDateUntil []byte

type testcaseDateUntil struct {
	sutYear, sutMonth, sutDay int
	inYear, inMonth, inDay    int
}

var testcasesDateUntil = func() (testcases []testcaseDateUntil) {
	s := tests.Scanner{Data: bytes.NewBuffer(testdataDateUntil)}
	nTestcases := s.ScanInt()
	for i := 0; i < nTestcases; i++ {
		ints := s.ScanInts(6)
		testcases = append(testcases, testcaseDateUntil{
			sutYear: ints[0], sutMonth: ints[1], sutDay: ints[2],
			inYear: ints[3], inMonth: ints[4], inDay: ints[5],
		})
	}
	return testcases
}()

func TestDate_DaysUntil(t *testing.T) {
	for _, testcase := range testcasesDateUntil {
		name := fmt.Sprintf("(%d-%d-%d)-(%d-%d-%d)", testcase.inYear, testcase.inMonth, testcase.inDay, testcase.sutYear, testcase.sutMonth, testcase.sutDay)
		t.Run(name, func(t *testing.T) {
			sut := DateOfYMD(testcase.sutYear, Month(testcase.sutMonth), testcase.sutDay)
			in := DateOfYMD(testcase.inYear, Month(testcase.inMonth), testcase.inDay)
			until := sut.DaysUntil(in)
			want := ToTime(sut.YMD()).AddDate(0, 0, int(until))
			{
				assert.Equal(t, want.Year(), testcase.inYear)
				assert.Equal(t, int(want.Month()), testcase.inMonth)
				assert.Equal(t, want.Day(), testcase.inDay)
			}
		})
	}
}

//go:embed testdata/date_addsub.txt
var testdataDateAddSub []byte

type testcaseDateAddSub struct {
	sutYear, sutMonth, sutDay int
	in                        int
}

var testcasesDateAddSub = func() (testcases []testcaseDateAddSub) {
	s := tests.Scanner{Data: bytes.NewBuffer(testdataDateAddSub)}
	nTestcases := s.ScanInt()
	for i := 0; i < nTestcases; i++ {
		ints := s.ScanInts(4)
		testcases = append(testcases, testcaseDateAddSub{
			sutYear: ints[0], sutMonth: ints[1], sutDay: ints[2],
			in: ints[3],
		})
	}
	return testcases
}()

func TestDate_Add(t *testing.T) {
	for _, testcase := range testcasesDateAddSub {
		name := fmt.Sprintf("(%d-%d-%d)+%dd", testcase.sutYear, testcase.sutMonth, testcase.sutDay, testcase.in)
		t.Run(name, func(t *testing.T) {
			sut := DateOfYMD(testcase.sutYear, Month(testcase.sutMonth), testcase.sutDay)
			got := sut.Add(testcase.in)
			want := ToTime(sut.YMD()).AddDate(0, 0, testcase.in)
			date.AssertEqualDate(t, DateFromTime(want), got)
		})
	}
}

func TestDate_Sub(t *testing.T) {
	for _, testcase := range testcasesDateAddSub {
		name := fmt.Sprintf("(%d-%d-%d)-%dd", testcase.sutYear, testcase.sutMonth, testcase.sutDay, testcase.in)
		t.Run(name, func(t *testing.T) {
			sut := DateOfYMD(testcase.sutYear, Month(testcase.sutMonth), testcase.sutDay)
			got := sut.Sub(testcase.in)
			want := ToTime(sut.YMD()).AddDate(0, 0, -testcase.in)
			date.AssertEqualDate(t, DateFromTime(want), got)
		})
	}
}
