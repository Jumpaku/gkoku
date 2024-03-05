package calender

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/Jumpaku/gkoku/internal/tests"
	"github.com/Jumpaku/gkoku/internal/tests/assert"
	"testing"
)

//go:embed testcases/testdata/date_yyyymmdd.txt
var testdataYyyyMmDd []byte

func Test_OfYyyyMmDd(t *testing.T) {
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
			sut := YyyyMmDd(testcase.sutYear, Month(testcase.sutMonth), testcase.sutDay)
			gotYear, gotMonth, gotDay := sut.YyyyMmDd()
			assert.Equal(t, testcase.wantYear, gotYear)
			assert.Equal(t, testcase.wantMonth, int(gotMonth))
			assert.Equal(t, testcase.wantDay, gotDay)
		})
	}
}

//go:embed testcases/testdata/date_yyyywwd.txt
var testdataYyyyWwD []byte

func Test_OfYyyyWwD(t *testing.T) {
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
			sut := YyyyWwD(testcase.sutYear, testcase.sutWeek, DayOfWeek(testcase.sutDay))
			{
				gotYear, gotMonth, gotDay := sut.YyyyMmDd()
				assert.Equal(t, testcase.wantYear, gotYear)
				assert.Equal(t, testcase.wantMonth, int(gotMonth))
				assert.Equal(t, testcase.wantDay, gotDay)
			}
			{
				gotYear, gotWeek, gotDay := sut.YyyyWwD()
				assert.Equal(t, testcase.sutYear, gotYear)
				assert.Equal(t, testcase.sutWeek, gotWeek)
				assert.Equal(t, testcase.sutDay, int(gotDay))
			}
		})
	}
}

//go:embed testcases/testdata/date_yyyyddd.txt
var testdataYyyyDdd []byte

func Test_OfYyyyDdd(t *testing.T) {
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
			sut := YyyyDdd(testcase.sutYear, testcase.sutDay)
			{
				gotYear, gotMonth, gotDay := sut.YyyyMmDd()
				assert.Equal(t, testcase.wantYear, gotYear)
				assert.Equal(t, testcase.wantMonth, int(gotMonth))
				assert.Equal(t, testcase.wantDay, gotDay)
			}
			{
				gotYear, gotDay := sut.YyyyDdd()
				assert.Equal(t, testcase.sutYear, gotYear)
				assert.Equal(t, testcase.sutDay, gotDay)
			}
		})
	}
}

//go:embed testcases/testdata/date_compare.txt
var testdataCompare []byte

func Test_Compare(t *testing.T) {
	type testcase struct {
		sutYear, sutMonth, sutDay int
		inYear, inMonth, inDay    int
		want                      int
	}

	s := tests.Scanner{Data: bytes.NewBuffer(testdataCompare)}
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
			sut := YyyyMmDd(testcase.sutYear, Month(testcase.sutMonth), testcase.sutDay)
			in := YyyyMmDd(testcase.inYear, Month(testcase.inMonth), testcase.inDay)
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
