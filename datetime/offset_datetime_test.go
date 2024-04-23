package datetime_test

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/Jumpaku/tokiope"
	"github.com/Jumpaku/tokiope/date"
	. "github.com/Jumpaku/tokiope/datetime"
	"github.com/Jumpaku/tokiope/internal/tests"
	"github.com/Jumpaku/tokiope/internal/tests/assert"
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestFormatOffsetDateTime(t *testing.T) {
	tests := []struct {
		in   OffsetDateTime
		want string
	}{
		{
			in: NewOffsetDateTime(
				date.YyyyMmDd(1, 1, 1),
				TimeOf(12, 34, 56, 0),
				OffsetMinutes(0),
			),
			want: `0001-01-01T12:34:56.000000000+00:00`,
		},
		{
			in: NewOffsetDateTime(
				date.YyyyMmDd(1, 1, 1),
				TimeOf(12, 34, 56, 0),
				OffsetMinutes(-(12*60 + 34)),
			),
			want: `0001-01-01T12:34:56.000000000-12:34`,
		},
		{
			in: NewOffsetDateTime(
				date.YyyyMmDd(2024, 3, 9),
				TimeOf(12, 34, 56, 0),
				OffsetMinutes(62),
			),
			want: `2024-03-09T12:34:56.000000000+01:02`,
		},
		{
			in: NewOffsetDateTime(
				date.YyyyMmDd(2024, 3, 9),
				TimeOf(12, 34, 56, 123456789),
				OffsetMinutes(62),
			),
			want: `2024-03-09T12:34:56.123456789+01:02`,
		},
		{
			in: NewOffsetDateTime(
				date.YyyyMmDd(2024, 3, 9),
				TimeOf(12, 34, 56, 0),
				OffsetMinutes(62),
			),
			want: `2024-03-09T12:34:56.000000000+01:02`,
		},
		{
			in: NewOffsetDateTime(
				date.YyyyMmDd(2024, 3, 9),
				TimeOf(12, 34, 56, 500_000_000),
				OffsetMinutes(62),
			),
			want: `2024-03-09T12:34:56.500000000+01:02`,
		},
		{
			in: NewOffsetDateTime(
				date.YyyyMmDd(2024, 3, 9),
				TimeOf(12, 34, 56, 50_000_000),
				OffsetMinutes(62),
			),
			want: `2024-03-09T12:34:56.050000000+01:02`,
		},
		{
			in: NewOffsetDateTime(
				date.YyyyMmDd(-2024, 3, 9),
				TimeOf(12, 34, 56, 50_000_000),
				OffsetMinutes(62),
			),
			want: `-2024-03-09T12:34:56.050000000+01:02`,
		},
		{
			in: NewOffsetDateTime(
				date.YyyyMmDd(12024, 3, 9),
				TimeOf(12, 34, 56, 50_000_000),
				OffsetMinutes(62),
			),
			want: `12024-03-09T12:34:56.050000000+01:02`,
		},
		{
			in: NewOffsetDateTime(
				date.YyyyMmDd(-12024, 3, 9),
				TimeOf(12, 34, 56, 50_000_000),
				OffsetMinutes(62),
			),
			want: `-12024-03-09T12:34:56.050000000+01:02`,
		},
		{
			in: NewOffsetDateTime(
				date.YyyyMmDd(1, 1, 1),
				TimeOf(0, 0, 0, 0),
				OffsetMinutes(0),
			),
			want: `0001-01-01T00:00:00.000000000+00:00`,
		},
		{
			in: NewOffsetDateTime(
				date.YyyyMmDd(-1, 1, 1),
				TimeOf(0, 0, 0, 0),
				OffsetMinutes(0),
			),
			want: `-0001-01-01T00:00:00.000000000+00:00`,
		},
		{
			in: NewOffsetDateTime(
				date.YyyyMmDd(-1, 1, 1),
				TimeOf(24, 0, 0, 0),
				OffsetMinutes(0),
			),
			want: `-0001-01-01T24:00:00.000000000+00:00`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			got := FormatOffsetDateTime(tt.in)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestParseOffsetDateTime(t *testing.T) {
	tests := []struct {
		in      string
		want    OffsetDateTime
		wantErr bool
	}{
		{
			in: `0001-01-01T12:34:56Z`,
			want: NewOffsetDateTime(
				date.YyyyMmDd(1, 1, 1),
				TimeOf(12, 34, 56, 0),
				OffsetMinutes(0),
			),
		},
		{
			in: `0001-01-01T12:34:56+00:00`,
			want: NewOffsetDateTime(
				date.YyyyMmDd(1, 1, 1),
				TimeOf(12, 34, 56, 0),
				OffsetMinutes(0),
			),
		},
		{
			in: `0001-01-01T12:34:56+0000`,
			want: NewOffsetDateTime(
				date.YyyyMmDd(1, 1, 1),
				TimeOf(12, 34, 56, 0),
				OffsetMinutes(0),
			),
		},
		{
			in: `0001-01-01T12:34:56-1234`,
			want: NewOffsetDateTime(
				date.YyyyMmDd(1, 1, 1),
				TimeOf(12, 34, 56, 0),
				OffsetMinutes(-(12*60 + 34)),
			),
		},
		{
			in: `0001-01-01T12:34:56+1234`,
			want: NewOffsetDateTime(
				date.YyyyMmDd(1, 1, 1),
				TimeOf(12, 34, 56, 0),
				OffsetMinutes(12*60+34),
			),
		},
		{
			in: `0001-01-01T12:34:56-12:34`,
			want: NewOffsetDateTime(
				date.YyyyMmDd(1, 1, 1),
				TimeOf(12, 34, 56, 0),
				OffsetMinutes(-(12*60 + 34)),
			),
		},
		{
			in: `2024-03-09T12:34:56+01:02`,
			want: NewOffsetDateTime(
				date.YyyyMmDd(2024, 3, 9),
				TimeOf(12, 34, 56, 0),
				OffsetMinutes(62),
			),
		},
		{
			in: `2024-W10-6T12:34:56+01:02`,
			want: NewOffsetDateTime(
				date.YyyyMmDd(2024, 3, 9),
				TimeOf(12, 34, 56, 0),
				OffsetMinutes(62),
			),
		},
		{
			in: `2024-069T12:34:56+01:02`,
			want: NewOffsetDateTime(
				date.YyyyMmDd(2024, 3, 9),
				TimeOf(12, 34, 56, 0),
				OffsetMinutes(62),
			),
		},
		{
			in: `2024-03-09T12:34:56.1+01:02`,
			want: NewOffsetDateTime(
				date.YyyyMmDd(2024, 3, 9),
				TimeOf(12, 34, 56, 100_000_000),
				OffsetMinutes(62),
			),
		},
		{
			in: `2024-03-09T12:34:56.123456789+01:02`,
			want: NewOffsetDateTime(
				date.YyyyMmDd(2024, 3, 9),
				TimeOf(12, 34, 56, 123456789),
				OffsetMinutes(62),
			),
		},
		{
			in: `2024-03-09T12:34:56.000000000+01:02`,
			want: NewOffsetDateTime(
				date.YyyyMmDd(2024, 3, 9),
				TimeOf(12, 34, 56, 0),
				OffsetMinutes(62),
			),
		},
		{
			in: `2024-03-09T12:34:56.50+01:02`,
			want: NewOffsetDateTime(
				date.YyyyMmDd(2024, 3, 9),
				TimeOf(12, 34, 56, 500_000_000),
				OffsetMinutes(62),
			),
		},
		{
			in: `2024-03-09T12:34:56.050+01:02`,
			want: NewOffsetDateTime(
				date.YyyyMmDd(2024, 3, 9),
				TimeOf(12, 34, 56, 50_000_000),
				OffsetMinutes(62),
			),
		},
		{
			in: `-2024-03-09T12:34:56.050+01:02`,
			want: NewOffsetDateTime(
				date.YyyyMmDd(-2024, 3, 9),
				TimeOf(12, 34, 56, 50_000_000),
				OffsetMinutes(62),
			),
		},
		{
			in: `12345-03-09T12:34:56.050+01:02`,
			want: NewOffsetDateTime(
				date.YyyyMmDd(12345, 3, 9),
				TimeOf(12, 34, 56, 50_000_000),
				OffsetMinutes(62),
			),
		},
		{
			in: `-12345-03-09T12:34:56.050+01:02`,
			want: NewOffsetDateTime(
				date.YyyyMmDd(-12345, 3, 9),
				TimeOf(12, 34, 56, 50_000_000),
				OffsetMinutes(62),
			),
		},
		{
			in: `-12345-03-09T24:00:00.00+01:02`,
			want: NewOffsetDateTime(
				date.YyyyMmDd(-12345, 3, 9),
				TimeOf(24, 0, 0, 0),
				OffsetMinutes(62),
			),
		},
		{
			in:      `1-03-09T12:34:56.050+01:02`,
			wantErr: true,
		},
		{
			in:      `2024-3-09T12:34:56.050+01:02`,
			wantErr: true,
		},
		{
			in:      `2024-03-9T12:34:56.050+01:02`,
			wantErr: true,
		},
		{
			in:      `2024-03-9T12:34:56.050+01:02`,
			wantErr: true,
		},
		{
			in:      `2024-03-09T1:34:56.050+01:02`,
			wantErr: true,
		},
		{
			in:      `2024-03-09T12:3:56.050+01:02`,
			wantErr: true,
		},
		{
			in:      `2024-03-09T12:34:5.050+01:02`,
			wantErr: true,
		},
		{
			in:      `2024-03-09T12:34:56.050+1:02`,
			wantErr: true,
		},
		{
			in:      `2024-03-09T12:34:56.050+01:2`,
			wantErr: true,
		},
		{
			in:      `2024-03-09T12:34:56.050+123`,
			wantErr: true,
		},
		{
			in:      `2024-00-09T12:34:56.050+123`,
			wantErr: true,
		},
		{
			in:      `2024-13-09T12:34:56.050+123`,
			wantErr: true,
		},
		{
			in:      `2024-03-09T12:34:56.050+123`,
			wantErr: true,
		},
		{
			in:      `2024-03-09T24:01:00.00+12:34`,
			wantErr: true,
		},
		{
			in:      `2024-03-09T24:00:01.00+12:34`,
			wantErr: true,
		},
		{
			in:      `2024-03-09T12:34:56.00`,
			wantErr: true,
		},
		{
			in:      `2024-03-09t12:34:56.00+00:00`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got, err := ParseOffsetDateTime(tt.in)
			if tt.wantErr {
				assert2.NotNil(t, err)
			} else {
				assert.Equal(t, tt.want.Offset(), got.Offset())
				assert.Equal(t, tt.want.Time(), got.Time())
				assert.Equal(t, tt.want.Date(), got.Date())
			}
		})
	}
}

//go:embed testdata/offsetdatetime.txt
var testdataOffsetDateTime []byte

func TestOffsetDateTime_Instant(t *testing.T) {
	type testcase struct {
		sutYear, sutMonth, sutDay                int
		sutHour, sutMinute, sutSecond, sutOffset int
		want                                     int64
	}

	s := tests.Scanner{Data: bytes.NewBuffer(testdataOffsetDateTime)}
	nTestcases := s.ScanInt()
	var testcases []testcase
	for i := 0; i < nTestcases; i++ {
		ints := s.ScanInt64s(8)
		testcases = append(testcases, testcase{
			sutYear: int(ints[0]), sutMonth: int(ints[1]), sutDay: int(ints[2]),
			sutHour: int(ints[3]), sutMinute: int(ints[4]), sutSecond: int(ints[5]),
			sutOffset: int(ints[6]),
			want:      ints[7],
		})
	}

	for number, testcase := range testcases {
		t.Run(fmt.Sprintf(`%d`, number), func(t *testing.T) {
			sut := NewOffsetDateTime(
				date.YyyyMmDd(testcase.sutYear, date.Month(testcase.sutMonth), testcase.sutDay),
				TimeOf(testcase.sutHour, testcase.sutMinute, testcase.sutSecond, 0),
				OffsetMinutes(testcase.sutOffset),
			)
			got, _ := sut.Instant().Unix()
			assert.Equal(t, testcase.want, got)
		})
	}
}

func TestFromInstant(t *testing.T) {
	type testcase struct {
		inUnix                           int64
		inOffset                         OffsetMinutes
		wantYear, wantMonth, wantDay     int
		wantHour, wantMinute, wantSecond int
	}

	s := tests.Scanner{Data: bytes.NewBuffer(testdataOffsetDateTime)}
	nTestcases := s.ScanInt()
	var testcases []testcase
	for i := 0; i < nTestcases; i++ {
		ints := s.ScanInt64s(8)
		testcases = append(testcases, testcase{
			inOffset: OffsetMinutes(ints[6]),
			inUnix:   ints[7],
			wantYear: int(ints[0]), wantMonth: int(ints[1]), wantDay: int(ints[2]),
			wantHour: int(ints[3]), wantMinute: int(ints[4]), wantSecond: int(ints[5]),
		})
	}

	for number, testcase := range testcases {
		t.Run(fmt.Sprintf(`%d`, number), func(t *testing.T) {
			got := FromInstant(tokiope.Unix(testcase.inUnix, 0), testcase.inOffset)
			gotY, gotM, gotD := got.Date().YyyyMmDd()
			assert.Equal(t, testcase.wantYear, gotY)
			assert.Equal(t, testcase.wantMonth, int(gotM))
			assert.Equal(t, testcase.wantDay, gotD)
			assert.Equal(t, testcase.wantHour, got.Time().Hour())
			assert.Equal(t, testcase.wantMinute, got.Time().Minute())
			assert.Equal(t, testcase.wantSecond, got.Time().Second())
			assert.Equal(t, 0, got.Time().Nano())
		})
	}
}
