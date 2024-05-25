package iter

import "github.com/Jumpaku/tokiope/calendar"

// YearWeekIterator is an iterator iterating on year-weeks of the calendar.
type YearWeekIterator interface {
	// Get returns the current year-week.
	Get() calendar.YearWeek
	// Copy returns a copy of this iterator.
	Copy() YearWeekIterator
	// Move moves this iterator by the given weeks.
	Move(weeks int)
	// Diff returns the number of weeks from the given iterator to this iterator.
	Diff(from YearWeekIterator) int
	// Year returns a YearIterator of the year including the year-week.
	Year() YearIterator
	// Date returns a DateIterator of the day of week in the year-week.
	Date(dayOfWeek calendar.DayOfWeek) DateIterator
	// FirstDate returns the first date in the year-week.
	FirstDate() DateIterator
	// LastDate returns the last date in the year-week.
	LastDate() DateIterator
}

func OfYearWeek(yearWeek calendar.YearWeek) YearWeekIterator {
	return &yearWeekIter{yearWeek: yearWeek}
}

type yearWeekIter struct {
	yearWeek calendar.YearWeek
}

func (y *yearWeekIter) Get() calendar.YearWeek {
	return y.yearWeek
}

func (y *yearWeekIter) Copy() YearWeekIterator {
	return OfYearWeek(y.Get())
}

func (y *yearWeekIter) Move(weeks int) {
	y.yearWeek = y.Get().Date(calendar.DayOfWeekMonday).Add(weeks * 7).YearWeek()
}

func (y *yearWeekIter) Diff(from YearWeekIterator) int {
	b := from.Get().Date(calendar.DayOfWeekMonday)
	e := y.Get().Date(calendar.DayOfWeekMonday)
	days := b.DaysUntil(e)
	return int(days / 7)
}

func (y *yearWeekIter) Year() YearIterator {
	return OfYear(y.Get().Year())
}

func (y *yearWeekIter) Date(dayOfWeek calendar.DayOfWeek) DateIterator {
	return OfDate(y.Get().Date(dayOfWeek))
}

func (y *yearWeekIter) FirstDate() DateIterator {
	return OfDate(y.Get().Date(calendar.DayOfWeekMonday))

}

func (y *yearWeekIter) LastDate() DateIterator {
	return OfDate(y.Get().Date(calendar.DayOfWeekSunday))
}
