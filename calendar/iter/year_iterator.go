package iter

import "github.com/Jumpaku/tokiope/calendar"

// YearIterator is an iterator iterating on years of the calendar.
type YearIterator interface {
	// Get returns the current year.
	Get() calendar.Year
	// Copy returns a copy of this iterator.
	Copy() YearIterator
	// Move moves this iterator by the given years.
	Move(years int)
	// Diff returns the number of years from the given iterator to this iterator.
	Diff(from YearIterator) int
	// Date returns a DateIterator of the day of year in the year.
	Date(dayOfYear int) DateIterator
	// Days returns the number of days in the year.
	Days() int
	// FirstDate returns the first date in the year.
	FirstDate() DateIterator
	// LastDate returns the last date in the year.
	LastDate() DateIterator
	// Weeks returns the number of weeks in the year.
	Weeks() int
	// Week returns a YearWeekIterator of the week in the year.
	Week(week int) YearWeekIterator
	// FirstWeek returns the first week in the year.
	FirstWeek() YearWeekIterator
	// LastWeek returns the last week in the year.
	LastWeek() YearWeekIterator
	// Month returns a YearMonthIterator of the month in the year.
	Month(month calendar.Month) YearMonthIterator
	// FirstMonth returns the first month in the year.
	FirstMonth() YearMonthIterator
	// LastMonth returns the last month in the year.
	LastMonth() YearMonthIterator
}

func OfYear(year calendar.Year) YearIterator {
	return &yearIter{year: year}
}

type yearIter struct {
	year calendar.Year
}

func (y *yearIter) Get() calendar.Year {
	return y.year
}

func (y *yearIter) Copy() YearIterator {
	return OfYear(y.Get())
}

func (y *yearIter) Move(years int) {
	y.year = calendar.Year(int(y.Get()) + years)
}

func (y *yearIter) Diff(from YearIterator) int {
	return int(y.Get() - from.Get())
}

func (y *yearIter) Date(dayOfYear int) DateIterator {
	return OfDate(y.Get().Date(dayOfYear))
}

func (y *yearIter) Days() int {
	return y.Get().Days()
}

func (y *yearIter) FirstDate() DateIterator {
	return y.Date(1)
}

func (y *yearIter) LastDate() DateIterator {
	return y.Date(y.Days())
}

func (y *yearIter) Weeks() int {
	return y.Get().Weeks()
}

func (y *yearIter) Week(week int) YearWeekIterator {
	return OfYearWeek(y.Get().Week(week))

}

func (y *yearIter) FirstWeek() YearWeekIterator {
	return y.Week(1)
}

func (y *yearIter) LastWeek() YearWeekIterator {
	return y.Week(y.Weeks())
}

func (y *yearIter) Month(month calendar.Month) YearMonthIterator {
	return OfYearMonth(y.Get().Month(month))
}

func (y *yearIter) FirstMonth() YearMonthIterator {
	return y.Month(calendar.MonthJanuary)
}

func (y *yearIter) LastMonth() YearMonthIterator {
	return y.Month(calendar.MonthDecember)
}
