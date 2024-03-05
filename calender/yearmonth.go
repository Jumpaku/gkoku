package calender

import (
	"cmp"
	"github.com/Jumpaku/gkoku/exact"
)

type YearMonth struct {
	// Months since 0000-01
	months int64
}

func YearMonthOf(year Year, month Month) YearMonth {
	return YearMonth{months: int64(year*12) + (int64(month-1) % 12)}
}

var _ interface {
	YyyyMm() (year int, month Month)
	Days() int
	Iter() YearIterator
	MonthsUntil(endExclusive YearMonth) int64
	WholeYearsUntil(endExclusive YearMonth) int64
	ContainsDay(dayOfMonth int) bool
	Cmp(other YearMonth) int
	Equal(other YearMonth) bool
	Before(other YearMonth) bool
	After(other YearMonth) bool
} = YearMonth{}

func (ym YearMonth) MonthsUntil(endExclusive YearMonth) int64 {
	return endExclusive.months - ym.months
}

func (ym YearMonth) WholeYearsUntil(endExclusive YearMonth) int64 {
	wm, _, _ := exact.DivTrunc(ym.MonthsUntil(endExclusive), 12)
	return wm
}

func (ym YearMonth) Iter() YearIterator {
	return nil
}

func (ym YearMonth) ContainsDay(dayOfMonth int) bool {
	return 1 <= dayOfMonth && dayOfMonth <= ym.Days()
}

func (ym YearMonth) Days() int {
	y, m := ym.YyyyMm()
	days := monthDays[m]
	if m == 2 && Year(y).IsLeap() {
		days++
	}
	return days
}

func (ym YearMonth) YyyyMm() (year int, month Month) {
	y, m, _ := exact.DivFloor(ym.months, 12)
	return int(y), Month(m + 1)
}

func (ym YearMonth) Cmp(other YearMonth) int {
	return cmp.Compare(ym.months, other.months)
}

func (ym YearMonth) Equal(other YearMonth) bool {
	return ym.months == other.months
}

func (ym YearMonth) Before(other YearMonth) bool {
	return ym.months < other.months
}

func (ym YearMonth) After(other YearMonth) bool {
	return ym.months > other.months
}
