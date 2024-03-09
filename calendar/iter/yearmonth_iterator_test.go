package iter_test

import (
	"github.com/Jumpaku/gkoku/calendar"
	. "github.com/Jumpaku/gkoku/calendar/iter"
	"github.com/Jumpaku/gkoku/internal/tests/assert"
	calendar_test "github.com/Jumpaku/gkoku/internal/tests/calendar"
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
			sut:  OfYearMonth(calendar.YearMonthOf(2024, calendar.MonthMarch)),
			in:   0,
			want: OfYearMonth(calendar.YearMonthOf(2024, calendar.MonthMarch)),
		},
		{
			name: "prev",
			sut:  OfYearMonth(calendar.YearMonthOf(2024, calendar.MonthMarch)),
			in:   -1,
			want: OfYearMonth(calendar.YearMonthOf(2024, calendar.MonthFebruary)),
		},
		{
			name: "next",
			sut:  OfYearMonth(calendar.YearMonthOf(2024, calendar.MonthMarch)),
			in:   1,
			want: OfYearMonth(calendar.YearMonthOf(2024, calendar.MonthApril)),
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
			sut:  OfYearMonth(calendar.YearMonthOf(2024, calendar.MonthMarch)),
			in:   OfYearMonth(calendar.YearMonthOf(2024, calendar.MonthMarch)),
			want: 0,
		},
		{
			name: "positive",
			sut:  OfYearMonth(calendar.YearMonthOf(2024, calendar.MonthMarch)),
			in:   OfYearMonth(calendar.YearMonthOf(2024, calendar.MonthFebruary)),
			want: 1,
		},
		{
			name: "negative",
			sut:  OfYearMonth(calendar.YearMonthOf(2024, calendar.MonthMarch)),
			in:   OfYearMonth(calendar.YearMonthOf(2024, calendar.MonthApril)),
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
			sut:  OfYearMonth(calendar.YearMonthOf(2024, calendar.MonthMarch)),
			in:   9,
			want: OfDate(calendar.YyyyMmDd(2024, calendar.MonthMarch, 9)),
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
			sut:  OfYearMonth(calendar.YearMonthOf(2024, calendar.MonthFebruary)),
			want: 29,
		},
		{
			name: "non leap Feb",
			sut:  OfYearMonth(calendar.YearMonthOf(2023, calendar.MonthFebruary)),
			want: 28,
		},
		{
			name: "Jan",
			sut:  OfYearMonth(calendar.YearMonthOf(2023, calendar.MonthJanuary)),
			want: 31,
		},
		{
			name: "June",
			sut:  OfYearMonth(calendar.YearMonthOf(2023, calendar.MonthJune)),
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
			sut:  OfYearMonth(calendar.YearMonthOf(2024, calendar.MonthMarch)),
			want: OfDate(calendar.YyyyMmDd(2024, calendar.MonthMarch, 1)),
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
			sut:  OfYearMonth(calendar.YearMonthOf(2024, calendar.MonthMarch)),
			want: OfDate(calendar.YyyyMmDd(2024, calendar.MonthMarch, 31)),
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
			sut:  OfYearMonth(calendar.YearMonthOf(2024, calendar.MonthMarch)),
			want: OfYear(calendar.Year(2024)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want.Get(), tt.sut.Year().Get())
		})
	}
}
