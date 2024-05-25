package calendar_test

import (
	"bytes"
	_ "embed"
	"fmt"
	. "github.com/Jumpaku/tokiope/calendar"
	"github.com/Jumpaku/tokiope/internal/tests"
	"github.com/stretchr/testify/assert"
	"testing"
)

//go:embed testdata/date_conv.txt
var testdataDateConv []byte

func TestParseDate(t *testing.T) {
	type testcase struct {
		sutYear, sutMonth, sutDay int
	}

	s := tests.Scanner{Data: bytes.NewBuffer(testdataDateConv)}
	nTestcases := s.ScanInt()
	var testcases []testcase
	for i := 0; i < nTestcases; i++ {
		ints := s.ScanInts(3)
		testcases = append(testcases, testcase{
			sutYear: ints[0], sutMonth: ints[1], sutDay: ints[2],
		})
	}

	for _, testcase := range testcases {
		sut := DateOfYMD(testcase.sutYear, Month(testcase.sutMonth), testcase.sutDay)
		t.Run(fmt.Sprintf("%d-%d-%d yyyy-mm-dd", testcase.sutYear, testcase.sutMonth, testcase.sutDay), func(t *testing.T) {
			got := FormatDate(sut, DateFormatYMD)
			est, err := ParseDate(got, DateFormatYMD)
			assert.Nil(t, err)
			wantY, wantM, wantD := sut.YMD()
			gotY, gotM, gotD := est.YMD()
			assert.Equal(t, wantY, gotY)
			assert.Equal(t, wantM, gotM)
			assert.Equal(t, wantD, gotD)
		})
		t.Run(fmt.Sprintf("%d-%d-%d yyyy-ddd", testcase.sutYear, testcase.sutMonth, testcase.sutDay), func(t *testing.T) {
			got := FormatDate(sut, DateFormatYD)
			est, err := ParseDate(got, DateFormatYD)
			assert.Nil(t, err)
			wantY, wantM, wantD := sut.YMD()
			gotY, gotM, gotD := est.YMD()
			assert.Equal(t, wantY, gotY)
			assert.Equal(t, wantM, gotM)
			assert.Equal(t, wantD, gotD)
		})
		t.Run(fmt.Sprintf("%d-%d-%d yyyy-Www-d", testcase.sutYear, testcase.sutMonth, testcase.sutDay), func(t *testing.T) {
			got := FormatDate(sut, DateFormatYWD)
			est, err := ParseDate(got, DateFormatYWD)
			assert.Nil(t, err)
			wantY, wantM, wantD := sut.YMD()
			gotY, gotM, gotD := est.YMD()
			assert.Equal(t, wantY, gotY)
			assert.Equal(t, wantM, gotM)
			assert.Equal(t, wantD, gotD)
		})
	}
}

//go:embed testdata/yearmonth_conv.txt
var testdataConvYearMonth []byte

func TestParseYearMonth(t *testing.T) {
	type testcase struct {
		sutYear, sutMonth int
	}

	s := tests.Scanner{Data: bytes.NewBuffer(testdataConvYearMonth)}
	nTestcases := s.ScanInt()
	var testcases []testcase
	for i := 0; i < nTestcases; i++ {
		ints := s.ScanInts(2)
		testcases = append(testcases, testcase{
			sutYear: ints[0], sutMonth: ints[1],
		})
	}

	for _, testcase := range testcases {
		t.Run(fmt.Sprintf("%d-%d", testcase.sutYear, testcase.sutMonth), func(t *testing.T) {
			sut := YearMonthOf(testcase.sutYear, Month(testcase.sutMonth))
			got := FormatYearMonth(sut)
			est, err := ParseYearMonth(got)
			assert.Nil(t, err)
			wantY, wantM := sut.YM()
			gotY, gotM := est.YM()
			assert.Equal(t, wantY, gotY)
			assert.Equal(t, wantM, gotM)
		})
	}
}

//go:embed testdata/yearweek_conv.txt
var testdataConvYearWeek []byte

func TestParseYearWeek(t *testing.T) {
	type testcase struct {
		sutYear, sutWeek int
	}

	s := tests.Scanner{Data: bytes.NewBuffer(testdataConvYearWeek)}
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
			sut := YearWeekOf(testcase.sutYear, testcase.sutWeek)
			got := FormatYearWeek(sut)
			est, err := ParseYearWeek(got)
			assert.Nil(t, err)
			wantY, wantW := sut.YW()
			gotY, gotW := est.YW()
			assert.Equal(t, wantY, gotY)
			assert.Equal(t, wantW, gotW)
		})
	}
}

func TestParseYear(t *testing.T) {
	for sut := Year(-9999); sut <= 9999; sut++ {
		t.Run(fmt.Sprintf("%d", sut), func(t *testing.T) {
			got := FormatYear(sut)
			est, err := ParseYear(got)
			assert.Nil(t, err)
			wantY := sut
			gotY := est
			assert.Equal(t, wantY, gotY)
		})
	}
}
