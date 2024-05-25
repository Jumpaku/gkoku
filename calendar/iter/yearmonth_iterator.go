package iter

import "github.com/Jumpaku/tokiope/calendar"

// YearMonthIterator is an iterator iterating on year-months of the calendar.
type YearMonthIterator interface {
	// Get returns the current year-month.
	Get() calendar.YearMonth
	// Copy returns a copy of this iterator.
	Copy() YearMonthIterator
	// Move moves this iterator by the given months.
	Move(months int)
	// Diff returns the number of months from the given iterator to this iterator.
	Diff(from YearMonthIterator) int
	// Year returns a YearIterator of the year including the year-month.
	Year() YearIterator
	// Days returns the number of days in the year-month.
	Days() int
	// Date returns a DateIterator of the day of month in the year-month.
	Date(dayOfMonth int) DateIterator
	// FirstDate returns the first date in the year-month.
	FirstDate() DateIterator
	// LastDate returns the last date in the year-month.
	LastDate() DateIterator
}

func OfYearMonth(yearMonth calendar.YearMonth) YearMonthIterator {
	return &yearMonthIter{yearMonth: yearMonth}
}

type yearMonthIter struct {
	yearMonth calendar.YearMonth
}

func (y *yearMonthIter) Get() calendar.YearMonth {
	return y.yearMonth
}

func (y *yearMonthIter) Copy() YearMonthIterator {
	return OfYearMonth(y.Get())
}

func (y *yearMonthIter) Move(months int) {
	y.yearMonth = y.Get().Add(months)
}

func (y *yearMonthIter) Diff(from YearMonthIterator) int {
	return int(from.Get().MonthsUntil(y.Get()))
}

func (y *yearMonthIter) Year() YearIterator {
	return OfYear(y.Get().Year())
}

func (y *yearMonthIter) Days() int {
	return y.Get().Days()
}

func (y *yearMonthIter) Date(dayOfMonth int) DateIterator {
	return OfDate(y.Get().Date(dayOfMonth))
}

func (y *yearMonthIter) FirstDate() DateIterator {
	return OfDate(y.Get().Date(1))
}

func (y *yearMonthIter) LastDate() DateIterator {
	return OfDate(y.Get().Date(y.Days()))
}
