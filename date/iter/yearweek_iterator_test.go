package iter

import (
	"github.com/Jumpaku/gkoku/date"
	calendar_test "github.com/Jumpaku/gkoku/internal/tests/date"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestYearWeekIterator_Move(t *testing.T) {
	tests := []struct {
		name string
		sut  YearWeekIterator
		in   int
		want YearWeekIterator
	}{
		{
			name: "stay",
			sut:  OfYearWeek(date.YearWeekOf(2024, 10)),
			in:   0,
			want: OfYearWeek(date.YearWeekOf(2024, 10)),
		},
		{
			name: "prev",
			sut:  OfYearWeek(date.YearWeekOf(2024, 10)),
			in:   -1,
			want: OfYearWeek(date.YearWeekOf(2024, 9)),
		},
		{
			name: "next",
			sut:  OfYearWeek(date.YearWeekOf(2024, 10)),
			in:   1,
			want: OfYearWeek(date.YearWeekOf(2024, 11)),
		},
		{
			name: "next year",
			sut:  OfYearWeek(date.YearWeekOf(2024, 52)),
			in:   1,
			want: OfYearWeek(date.YearWeekOf(2025, 1)),
		},
		{
			name: "prev year",
			sut:  OfYearWeek(date.YearWeekOf(2024, 1)),
			in:   -1,
			want: OfYearWeek(date.YearWeekOf(2023, 52)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := tt.sut.Copy()
			sut.Move(tt.in)
			calendar_test.AssertEqualYearWeek(t, tt.want.Get(), sut.Get())
		})
	}
}

func TestYearWeekIterator_Diff(t *testing.T) {
	tests := []struct {
		name string
		sut  YearWeekIterator
		in   YearWeekIterator
		want int
	}{
		{
			name: "zero",
			sut:  OfYearWeek(date.YearWeekOf(2024, 10)),
			in:   OfYearWeek(date.YearWeekOf(2024, 10)),
			want: 0,
		},
		{
			name: "positive",
			sut:  OfYearWeek(date.YearWeekOf(2024, 10)),
			in:   OfYearWeek(date.YearWeekOf(2024, 9)),
			want: 1,
		},
		{
			name: "negative",
			sut:  OfYearWeek(date.YearWeekOf(2024, 10)),
			in:   OfYearWeek(date.YearWeekOf(2024, 11)),
			want: -1,
		},
		{
			name: "prev year",
			sut:  OfYearWeek(date.YearWeekOf(2024, 1)),
			in:   OfYearWeek(date.YearWeekOf(2023, 52)),
			want: 1,
		},
		{
			name: "next year",
			sut:  OfYearWeek(date.YearWeekOf(2024, 52)),
			in:   OfYearWeek(date.YearWeekOf(2025, 1)),
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.sut.Diff(tt.in))
		})
	}
}

func TestYearWeekIterator_Date(t *testing.T) {
	tests := []struct {
		name string
		sut  YearWeekIterator
		in   date.DayOfWeek
		want DateIterator
	}{
		{
			name: "Wednesday",
			sut:  OfYearWeek(date.YearWeekOf(2024, 10)),
			in:   date.DayOfWeekWednesday,
			want: OfDate(date.YyyyMmDd(2024, date.MonthMarch, 6)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calendar_test.AssertEqualDate(t, tt.want.Get(), tt.sut.Date(tt.in).Get())
		})
	}
}

func TestYearWeekIterator_FirstDate(t *testing.T) {
	tests := []struct {
		name string
		sut  YearWeekIterator
		want DateIterator
	}{
		{
			name: "2024-W10",
			sut:  OfYearWeek(date.YearWeekOf(2024, 10)),
			want: OfDate(date.YyyyWwD(2024, 10, date.DayOfWeekMonday)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calendar_test.AssertEqualDate(t, tt.want.Get(), tt.sut.FirstDate().Get())
		})
	}
}

func TestYearWeekIterator_LastDate(t *testing.T) {
	tests := []struct {
		name string
		sut  YearWeekIterator
		want DateIterator
	}{
		{
			name: "2024-W10",
			sut:  OfYearWeek(date.YearWeekOf(2024, 10)),
			want: OfDate(date.YyyyWwD(2024, 10, date.DayOfWeekSunday)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calendar_test.AssertEqualDate(t, tt.want.Get(), tt.sut.LastDate().Get())
		})
	}
}

func TestYearWeekIterator_Year(t *testing.T) {
	tests := []struct {
		name string
		sut  YearWeekIterator
		want YearIterator
	}{
		{
			name: "2024-W10",
			sut:  OfYearWeek(date.YearWeekOf(2024, 10)),
			want: OfYear(date.Year(2024)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want.Get(), tt.sut.Year().Get())
		})
	}
}
