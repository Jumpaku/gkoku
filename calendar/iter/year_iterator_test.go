package iter_test

import (
	"github.com/Jumpaku/gkoku/calendar"
	. "github.com/Jumpaku/gkoku/calendar/iter"
	assert2 "github.com/Jumpaku/gkoku/internal/tests/assert"
	calendar_test "github.com/Jumpaku/gkoku/internal/tests/calendar"
	"testing"
)

func TestYearIterator_Diff(t *testing.T) {
	tests := []struct {
		name string
		sut  YearIterator
		in   YearIterator
		want int
	}{
		{
			name: "zero",
			sut:  OfYear(calendar.Year(2024)),
			in:   OfYear(calendar.Year(2024)),
			want: 0,
		},
		{
			name: "positive",
			sut:  OfYear(calendar.Year(2024)),
			in:   OfYear(calendar.Year(2023)),
			want: 1,
		},
		{
			name: "negative",
			sut:  OfYear(calendar.Year(2024)),
			in:   OfYear(calendar.Year(2025)),
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert2.Equal(t, tt.want, tt.sut.Diff(tt.in))
		})
	}
}

func TestYearIterator_Move(t *testing.T) {
	tests := []struct {
		name string
		sut  YearIterator
		in   int
		want YearIterator
	}{
		{
			name: "stay",
			sut:  OfYear(calendar.Year(2024)),
			in:   0,
			want: OfYear(calendar.Year(2024)),
		},
		{
			name: "prev",
			sut:  OfYear(calendar.Year(2024)),
			in:   -1,
			want: OfYear(calendar.Year(2023)),
		},
		{
			name: "next",
			sut:  OfYear(calendar.Year(2024)),
			in:   1,
			want: OfYear(calendar.Year(2025)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := tt.sut.Copy()
			sut.Move(tt.in)
			assert2.Equal(t, tt.want.Get(), sut.Get())
		})
	}
}

func TestYearIterator_Days(t *testing.T) {
	tests := []struct {
		name string
		sut  YearIterator
		want int
	}{
		{
			name: "leap",
			sut:  OfYear(calendar.Year(2024)),
			want: 366,
		},
		{
			name: "non leap",
			sut:  OfYear(calendar.Year(2025)),
			want: 365,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert2.Equal(t, tt.want, tt.sut.Days())
		})
	}
}

func TestYearIterator_Date(t *testing.T) {
	tests := []struct {
		name string
		sut  YearIterator
		in   int
		want DateIterator
	}{
		{
			name: "69",
			sut:  OfYear(calendar.Year(2024)),
			in:   69,
			want: OfDate(calendar.YyyyMmDd(2024, calendar.MonthMarch, 9)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calendar_test.AssertEqualDate(t, tt.want.Get(), tt.sut.Date(tt.in).Get())
		})
	}
}

func TestYearIterator_FirstDate(t *testing.T) {
	tests := []struct {
		name string
		sut  YearIterator
		want DateIterator
	}{
		{
			name: "2024",
			sut:  OfYear(calendar.Year(2024)),
			want: OfDate(calendar.YyyyMmDd(2024, calendar.MonthJanuary, 1)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calendar_test.AssertEqualDate(t, tt.want.Get(), tt.sut.FirstDate().Get())
		})
	}
}

func TestYearIterator_LastDate(t *testing.T) {
	tests := []struct {
		name string
		sut  YearIterator
		want DateIterator
	}{
		{
			name: "2024",
			sut:  OfYear(calendar.Year(2024)),
			want: OfDate(calendar.YyyyMmDd(2024, calendar.MonthDecember, 31)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calendar_test.AssertEqualDate(t, tt.want.Get(), tt.sut.LastDate().Get())
		})
	}
}

func TestYearIterator_Week(t *testing.T) {
	tests := []struct {
		name string
		sut  YearIterator
		in   int
		want YearWeekIterator
	}{
		{
			name: "10",
			sut:  OfYear(calendar.Year(2024)),
			in:   10,
			want: OfYearWeek(calendar.YearWeekOf(2024, 10)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calendar_test.AssertEqualYearWeek(t, tt.want.Get(), tt.sut.Week(tt.in).Get())
		})
	}
}

func TestYearIterator_Weeks(t *testing.T) {
	tests := []struct {
		name string
		sut  YearIterator
		want int
	}{
		{
			name: "52",
			sut:  OfYear(calendar.Year(2024)),
			want: 52,
		},
		{
			name: "53",
			sut:  OfYear(calendar.Year(2032)),
			want: 53,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert2.Equal(t, tt.want, tt.sut.Weeks())
		})
	}
}

func TestYearIterator_FirstWeek(t *testing.T) {
	tests := []struct {
		name string
		sut  YearIterator
		want YearWeekIterator
	}{
		{
			name: "2024",
			sut:  OfYear(calendar.Year(2024)),
			want: OfYearWeek(calendar.YearWeekOf(2024, 1)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calendar_test.AssertEqualYearWeek(t, tt.want.Get(), tt.sut.FirstWeek().Get())
		})
	}
}

func TestYearIterator_LastWeek(t *testing.T) {
	tests := []struct {
		name string
		sut  YearIterator
		want YearWeekIterator
	}{
		{
			name: "2024",
			sut:  OfYear(calendar.Year(2024)),
			want: OfYearWeek(calendar.YearWeekOf(2024, 52)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calendar_test.AssertEqualYearWeek(t, tt.want.Get(), tt.sut.LastWeek().Get())
		})
	}
}

func TestYearIterator_FirstMonth(t *testing.T) {
	tests := []struct {
		name string
		sut  YearIterator
		want YearMonthIterator
	}{
		{
			name: "2024",
			sut:  OfYear(calendar.Year(2024)),
			want: OfYearMonth(calendar.YearMonthOf(2024, calendar.MonthJanuary)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calendar_test.AssertEqualYearMonth(t, tt.want.Get(), tt.sut.FirstMonth().Get())
		})
	}
}
func TestYearIterator_LastMonth(t *testing.T) {
	tests := []struct {
		name string
		sut  YearIterator
		want YearMonthIterator
	}{
		{
			name: "2024",
			sut:  OfYear(calendar.Year(2024)),
			want: OfYearMonth(calendar.YearMonthOf(2024, calendar.MonthDecember)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calendar_test.AssertEqualYearMonth(t, tt.want.Get(), tt.sut.LastMonth().Get())
		})
	}
}

func TestYearIterator_Month(t *testing.T) {
	tests := []struct {
		name string
		sut  YearIterator
		in   calendar.Month
		want YearMonthIterator
	}{
		{
			name: "2024",
			sut:  OfYear(calendar.Year(2024)),
			in:   calendar.MonthMarch,
			want: OfYearMonth(calendar.YearMonthOf(2024, calendar.MonthMarch)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calendar_test.AssertEqualYearMonth(t, tt.want.Get(), tt.sut.Month(tt.in).Get())
		})
	}
}
