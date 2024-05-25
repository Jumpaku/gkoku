package iter

import "github.com/Jumpaku/tokiope/calendar"

// DateIterator is an iterator iterating on dates of the calendar.
type DateIterator interface {
	// Get returns the current date.
	Get() calendar.Date
	// Copy returns a copy of this iterator.
	Copy() DateIterator
	// Move moves this iterator by the given days.
	Move(days int)
	// Diff returns the number of days from the given iterator to this iterator.
	Diff(from DateIterator) int
	// Year returns a YearIterator of the year including the date.
	Year() YearIterator
	// YearMonth returns a YearMonthIterator of the year and month including the date.
	YearMonth() YearMonthIterator
	// YearWeek returns a YearWeekIterator of the year and week including the date.
	YearWeek() YearWeekIterator
}

// OfDate creates a DateIterator on the date.
func OfDate(date calendar.Date) DateIterator {
	return &dateIter{date: date}
}

type dateIter struct {
	date calendar.Date
}

func (i *dateIter) Get() calendar.Date {
	return i.date
}

func (i *dateIter) Copy() DateIterator {
	return OfDate(i.Get())
}

func (i *dateIter) Move(days int) {
	i.date = i.Get().Add(days)
}

func (i *dateIter) Diff(from DateIterator) int {
	return int(from.Get().DaysUntil(i.Get()))
}

func (i *dateIter) Year() YearIterator {
	return OfYear(i.Get().Year())
}

func (i *dateIter) YearMonth() YearMonthIterator {
	return OfYearMonth(i.Get().YearMonth())
}

func (i *dateIter) YearWeek() YearWeekIterator {
	return OfYearWeek(i.Get().YearWeek())
}
