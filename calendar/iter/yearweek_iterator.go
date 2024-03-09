package iter

import "github.com/Jumpaku/gkoku/calendar"

type YearWeekIterator interface {
	Get() calendar.YearWeek
	Copy() YearWeekIterator
	Move(weeks int)
	Diff(from YearWeekIterator) int
	Year() YearIterator
	Date(dayOfWeek calendar.DayOfWeek) DateIterator
	FirstDate() DateIterator
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
