package calendar

import (
	"cmp"
	"github.com/Jumpaku/tokiope/internal/exact"
)

type YearMonth struct {
	// Months since 0000-01
	months int64
}

func YearMonthOf(year int, month Month) YearMonth {
	return YearMonth{months: int64(year*12) + (int64(month-1) % 12)}
}

var _ interface {
	YyyyMm() (year int, month Month)
	Days() int
	Year() Year
	Date(dayOfMonth int) Date
	FirstDate() Date
	LastDate() Date
	Add(months int) YearMonth
	Sub(months int) YearMonth
	MonthsUntil(endExclusive YearMonth) int64
	ContainsDay(dayOfMonth int) bool
	Cmp(other YearMonth) int
	Equal(other YearMonth) bool
	Before(other YearMonth) bool
	After(other YearMonth) bool
	String() string
} = YearMonth{}

func (ym YearMonth) YyyyMm() (year int, month Month) {
	y, m, _ := exact.DivFloor(ym.months, 12)
	return int(y), Month(m + 1)
}

func (ym YearMonth) Year() Year {
	y, _ := ym.YyyyMm()
	return Year(y)
}

func (ym YearMonth) Date(dayOfMonth int) Date {
	y, m := ym.YyyyMm()
	return DateOfYMD(y, m, dayOfMonth)
}

func (ym YearMonth) FirstDate() Date {
	y, m := ym.YyyyMm()
	return DateOfYMD(y, m, 1)
}

func (ym YearMonth) LastDate() Date {
	y, m := ym.YyyyMm()
	return DateOfYMD(y, m, ym.Days())
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

func (ym YearMonth) Add(months int) YearMonth {
	return YearMonth{months: ym.months + int64(months)}
}

func (ym YearMonth) Sub(months int) YearMonth {
	return YearMonth{months: ym.months - int64(months)}

}

func (ym YearMonth) MonthsUntil(endExclusive YearMonth) int64 {
	return endExclusive.months - ym.months
}

func (ym YearMonth) String() string {
	return FormatYearMonth(ym)
}
