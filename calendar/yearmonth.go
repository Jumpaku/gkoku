package calendar

import (
	"cmp"
	"github.com/Jumpaku/tokiope/internal/exact"
)

// YearMonth represents a year and a month in the proleptic Gregorian calendar.
type YearMonth struct {
	// Months since 0000-01
	months int64
}

// YearMonthOf returns a YearMonth of the given year and month.
func YearMonthOf(year int, month Month) YearMonth {
	return YearMonth{months: int64(year*12) + (int64(month-1) % 12)}
}

var _ interface {
	YM() (year int, month Month)
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

// YM returns the year and the month.
func (ym YearMonth) YM() (year int, month Month) {
	y, m, _ := exact.DivFloor(ym.months, 12)
	return int(y), Month(m + 1)
}

// Year returns the year.
func (ym YearMonth) Year() Year {
	y, _ := ym.YM()
	return Year(y)
}

// Date returns a Date of the dayOfMonth in this year and month.
func (ym YearMonth) Date(dayOfMonth int) Date {
	y, m := ym.YM()
	return DateOfYMD(y, m, dayOfMonth)
}

// FirstDate returns the first date in this year and month.
func (ym YearMonth) FirstDate() Date {
	y, m := ym.YM()
	return DateOfYMD(y, m, 1)
}

// LastDate returns the last date in this year and month.
func (ym YearMonth) LastDate() Date {
	y, m := ym.YM()
	return DateOfYMD(y, m, ym.Days())
}

// ContainsDay returns true if the dayOfMonth is in this year and month.
func (ym YearMonth) ContainsDay(dayOfMonth int) bool {
	return 1 <= dayOfMonth && dayOfMonth <= ym.Days()
}

// Days returns the number of days in this year and month.
func (ym YearMonth) Days() int {
	y, m := ym.YM()
	days := monthDays[m]
	if m == 2 && Year(y).IsLeap() {
		days++
	}
	return days
}

// Cmp compares the year and month with the other year and month.
func (ym YearMonth) Cmp(other YearMonth) int {
	return cmp.Compare(ym.months, other.months)
}

// Equal returns whether the year and month is equal to the other year and month.
func (ym YearMonth) Equal(other YearMonth) bool {
	return ym.months == other.months
}

// Before returns whether the year and month is before the other year and month.
func (ym YearMonth) Before(other YearMonth) bool {
	return ym.months < other.months
}

// After returns whether the year and month is after the other year and month.
func (ym YearMonth) After(other YearMonth) bool {
	return ym.months > other.months
}

// Add returns the year and month going forward by the amount of months.
func (ym YearMonth) Add(months int) YearMonth {
	return YearMonth{months: ym.months + int64(months)}
}

// Sub returns the year and month going backward by the amount of months.
func (ym YearMonth) Sub(months int) YearMonth {
	return YearMonth{months: ym.months - int64(months)}

}

// MonthsUntil returns the number of months until the year and month of endExclusive.
func (ym YearMonth) MonthsUntil(endExclusive YearMonth) int64 {
	return endExclusive.months - ym.months
}

// String returns a textual representation of the YearMonth value formatted.
func (ym YearMonth) String() string {
	return FormatYearMonth(ym)
}
