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
		t.Run(fmt.Sprintf("%d-%d-%d", testcase.sutYear, testcase.sutMonth, testcase.sutDay), func(t *testing.T) {
			sut := YyyyMmDd(testcase.sutYear, Month(testcase.sutMonth), testcase.sutDay)
			{
				got := FormatDate(sut, DateFormatYyyyMmDd)
				est, err := ParseDate(got, DateFormatYyyyMmDd)
				assert.Nil(t, err)
				wantY, wantM, wantD := sut.YyyyMmDd()
				gotY, gotM, gotD := est.YyyyMmDd()
				assert.Equal(t, wantY, gotY)
				assert.Equal(t, wantM, gotM)
				assert.Equal(t, wantD, gotD)
			}
			{
				got := FormatDate(sut, DateFormatYyyyDdd)
				est, err := ParseDate(got, DateFormatYyyyDdd)
				assert.Nil(t, err)
				wantY, wantM, wantD := sut.YyyyMmDd()
				gotY, gotM, gotD := est.YyyyMmDd()
				assert.Equal(t, wantY, gotY)
				assert.Equal(t, wantM, gotM)
				assert.Equal(t, wantD, gotD)
			}
			{
				got := FormatDate(sut, DateFormatYyyyWwD)
				est, err := ParseDate(got, DateFormatYyyyWwD)
				assert.Nil(t, err)
				wantY, wantM, wantD := sut.YyyyMmDd()
				gotY, gotM, gotD := est.YyyyMmDd()
				assert.Equal(t, wantY, gotY)
				assert.Equal(t, wantM, gotM)
				assert.Equal(t, wantD, gotD)
			}
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
