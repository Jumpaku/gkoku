package iter_test

import (
	"github.com/Jumpaku/tokiope/date"
	. "github.com/Jumpaku/tokiope/date/iter"
	"github.com/Jumpaku/tokiope/internal/tests/assert"
	calendar_test "github.com/Jumpaku/tokiope/internal/tests/date"
	"testing"
)

func TestYearMonthIterator_Move(t *testing.T) {
	tests := []struct {
		name string
		sut  YearMonthIterator
		in   int
		want YearMonthIterator
	}{
		{
			name: "stay",
			sut:  OfYearMonth(date.YearMonthOf(2024, date.MonthMarch)),
			in:   0,
			want: OfYearMonth(date.YearMonthOf(2024, date.MonthMarch)),
		},
		{
			name: "prev",
			sut:  OfYearMonth(date.YearMonthOf(2024, date.MonthMarch)),
			in:   -1,
			want: OfYearMonth(date.YearMonthOf(2024, date.MonthFebruary)),
		},
		{
			name: "next",
			sut:  OfYearMonth(date.YearMonthOf(2024, date.MonthMarch)),
			in:   1,
			want: OfYearMonth(date.YearMonthOf(2024, date.MonthApril)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := tt.sut.Copy()
			sut.Move(tt.in)
			calendar_test.AssertEqualYearMonth(t, tt.want.Get(), sut.Get())
		})
	}
}

func TestYearMonthIterator_Diff(t *testing.T) {
	tests := []struct {
		name string
		sut  YearMonthIterator
		in   YearMonthIterator
		want int
	}{
		{
			name: "zero",
			sut:  OfYearMonth(date.YearMonthOf(2024, date.MonthMarch)),
			in:   OfYearMonth(date.YearMonthOf(2024, date.MonthMarch)),
			want: 0,
		},
		{
			name: "positive",
			sut:  OfYearMonth(date.YearMonthOf(2024, date.MonthMarch)),
			in:   OfYearMonth(date.YearMonthOf(2024, date.MonthFebruary)),
			want: 1,
		},
		{
			name: "negative",
			sut:  OfYearMonth(date.YearMonthOf(2024, date.MonthMarch)),
			in:   OfYearMonth(date.YearMonthOf(2024, date.MonthApril)),
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.sut.Diff(tt.in))
		})
	}
}

func TestYearMonthIterator_Date(t *testing.T) {
	tests := []struct {
		name string
		sut  YearMonthIterator
		in   int
		want DateIterator
	}{
		{
			name: "9",
			sut:  OfYearMonth(date.YearMonthOf(2024, date.MonthMarch)),
			in:   9,
			want: OfDate(date.YyyyMmDd(2024, date.MonthMarch, 9)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calendar_test.AssertEqualDate(t, tt.want.Get(), tt.sut.Date(tt.in).Get())
		})
	}
}

func TestYearMonthIterator_Days(t *testing.T) {
	tests := []struct {
		name string
		sut  YearMonthIterator
		want int
	}{
		{
			name: "leap Feb",
			sut:  OfYearMonth(date.YearMonthOf(2024, date.MonthFebruary)),
			want: 29,
		},
		{
			name: "non leap Feb",
			sut:  OfYearMonth(date.YearMonthOf(2023, date.MonthFebruary)),
			want: 28,
		},
		{
			name: "Jan",
			sut:  OfYearMonth(date.YearMonthOf(2023, date.MonthJanuary)),
			want: 31,
		},
		{
			name: "June",
			sut:  OfYearMonth(date.YearMonthOf(2023, date.MonthJune)),
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.sut.Days())
		})
	}
}

func TestYearMonthIterator_FirstDate(t *testing.T) {
	tests := []struct {
		name string
		sut  YearMonthIterator
		want DateIterator
	}{
		{
			name: "2024-03",
			sut:  OfYearMonth(date.YearMonthOf(2024, date.MonthMarch)),
			want: OfDate(date.YyyyMmDd(2024, date.MonthMarch, 1)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calendar_test.AssertEqualDate(t, tt.want.Get(), tt.sut.FirstDate().Get())
		})
	}
}

func TestYearMonthIterator_LastDate(t *testing.T) {
	tests := []struct {
		name string
		sut  YearMonthIterator
		want DateIterator
	}{
		{
			name: "2024-03",
			sut:  OfYearMonth(date.YearMonthOf(2024, date.MonthMarch)),
			want: OfDate(date.YyyyMmDd(2024, date.MonthMarch, 31)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calendar_test.AssertEqualDate(t, tt.want.Get(), tt.sut.LastDate().Get())
		})
	}
}

func TestYearMonthIterator_Year(t *testing.T) {
	tests := []struct {
		name string
		sut  YearMonthIterator
		want YearIterator
	}{
		{
			name: "2024-03",
			sut:  OfYearMonth(date.YearMonthOf(2024, date.MonthMarch)),
			want: OfYear(date.Year(2024)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want.Get(), tt.sut.Year().Get())
		})
	}
}
