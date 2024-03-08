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

//go:embed testcases/testdata/date_conv.txt
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
		sut := YyyyMmDd(testcase.sutYear, Month(testcase.sutMonth), testcase.sutDay)
		t.Run(fmt.Sprintf("%d-%d-%d yyyy-mm-dd", testcase.sutYear, testcase.sutMonth, testcase.sutDay), func(t *testing.T) {
			got := FormatDate(sut, DateFormatYyyyMmDd)
			est, err := ParseDate(got, DateFormatYyyyMmDd)
			assert.Nil(t, err)
			wantY, wantM, wantD := sut.YyyyMmDd()
			gotY, gotM, gotD := est.YyyyMmDd()
			assert.Equal(t, wantY, gotY)
			assert.Equal(t, wantM, gotM)
			assert.Equal(t, wantD, gotD)
		})
		t.Run(fmt.Sprintf("%d-%d-%d yyyy-ddd", testcase.sutYear, testcase.sutMonth, testcase.sutDay), func(t *testing.T) {
			got := FormatDate(sut, DateFormatYyyyDdd)
			est, err := ParseDate(got, DateFormatYyyyDdd)
			assert.Nil(t, err)
			wantY, wantM, wantD := sut.YyyyMmDd()
			gotY, gotM, gotD := est.YyyyMmDd()
			assert.Equal(t, wantY, gotY)
			assert.Equal(t, wantM, gotM)
			assert.Equal(t, wantD, gotD)
		})
		t.Run(fmt.Sprintf("%d-%d-%d yyyy-Www-d", testcase.sutYear, testcase.sutMonth, testcase.sutDay), func(t *testing.T) {
			got := FormatDate(sut, DateFormatYyyyWwD)
			est, err := ParseDate(got, DateFormatYyyyWwD)
			assert.Nil(t, err)
			wantY, wantM, wantD := sut.YyyyMmDd()
			gotY, gotM, gotD := est.YyyyMmDd()
			assert.Equal(t, wantY, gotY)
			assert.Equal(t, wantM, gotM)
			assert.Equal(t, wantD, gotD)
		})
	}
}

//go:embed testcases/testdata/yearmonth_conv.txt
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
			wantY, wantM := sut.YyyyMm()
			gotY, gotM := est.YyyyMm()
			assert.Equal(t, wantY, gotY)
			assert.Equal(t, wantM, gotM)
		})
	}
}

//go:embed testcases/testdata/yearweek_conv.txt
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
			wantY, wantW := sut.YyyyWw()
			gotY, gotW := est.YyyyWw()
			assert.Equal(t, wantY, gotY)
			assert.Equal(t, wantW, gotW)
		})
	}
}
