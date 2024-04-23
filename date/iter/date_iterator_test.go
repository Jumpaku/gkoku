package iter_test

import (
	"github.com/Jumpaku/tokiope/date"
	. "github.com/Jumpaku/tokiope/date/iter"
	"github.com/Jumpaku/tokiope/internal/tests/assert"
	calendar_test "github.com/Jumpaku/tokiope/internal/tests/date"
	"testing"
)

func TestDateIterator_Diff(t *testing.T) {
	tests := []struct {
		name string
		sut  DateIterator
		in   DateIterator
		want int
	}{
		{
			name: "zero",
			sut:  OfDate(date.YyyyMmDd(2024, date.MonthMarch, 9)),
			in:   OfDate(date.YyyyMmDd(2024, date.MonthMarch, 9)),
			want: 0,
		},
		{
			name: "positive",
			sut:  OfDate(date.YyyyMmDd(2024, date.MonthMarch, 9)),
			in:   OfDate(date.YyyyMmDd(2024, date.MonthMarch, 8)),
			want: 1,
		},
		{
			name: "negative",
			sut:  OfDate(date.YyyyMmDd(2024, date.MonthMarch, 9)),
			in:   OfDate(date.YyyyMmDd(2024, date.MonthMarch, 10)),
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.sut.Diff(tt.in))
		})
	}
}

func TestDateIterator_Move(t *testing.T) {
	tests := []struct {
		name string
		sut  DateIterator
		in   int
		want DateIterator
	}{
		{
			name: "stay",
			sut:  OfDate(date.YyyyMmDd(2024, date.MonthMarch, 9)),
			in:   0,
			want: OfDate(date.YyyyMmDd(2024, date.MonthMarch, 9)),
		},
		{
			name: "next",
			sut:  OfDate(date.YyyyMmDd(2024, date.MonthMarch, 9)),
			in:   1,
			want: OfDate(date.YyyyMmDd(2024, date.MonthMarch, 10)),
		},
		{
			name: "prev",
			sut:  OfDate(date.YyyyMmDd(2024, date.MonthMarch, 9)),
			in:   -1,
			want: OfDate(date.YyyyMmDd(2024, date.MonthMarch, 8)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := tt.sut.Copy()
			sut.Move(tt.in)
			calendar_test.AssertEqualDate(t, tt.want.Get(), sut.Get())
		})
	}
}

func TestDateIterator_Year(t *testing.T) {
	tests := []struct {
		name string
		sut  DateIterator
		want YearIterator
	}{
		{
			name: "stay",
			sut:  OfDate(date.YyyyMmDd(2024, date.MonthMarch, 9)),
			want: OfYear(date.Year(2024)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want.Get(), tt.sut.Year().Get())
		})
	}
}

func TestDateIterator_YearMonth(t *testing.T) {
	tests := []struct {
		name string
		sut  DateIterator
		want YearMonthIterator
	}{
		{
			name: "stay",
			sut:  OfDate(date.YyyyMmDd(2024, date.MonthMarch, 9)),
			want: OfYearMonth(date.YearMonthOf(2024, date.MonthMarch)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calendar_test.AssertEqualYearMonth(t, tt.want.Get(), tt.sut.YearMonth().Get())
		})
	}
}

func TestDateIterator_YearWeek(t *testing.T) {
	tests := []struct {
		name string
		sut  DateIterator
		want YearWeekIterator
	}{
		{
			name: "stay",
			sut:  OfDate(date.YyyyMmDd(2024, date.MonthMarch, 9)),
			want: OfYearWeek(date.YearWeekOf(2024, 10)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calendar_test.AssertEqualYearWeek(t, tt.want.Get(), tt.sut.YearWeek().Get())
		})
	}
}
