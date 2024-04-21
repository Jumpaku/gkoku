package iter

import "github.com/Jumpaku/gkoku/date"

type YearWeekIterator interface {
	Get() date.YearWeek
	Copy() YearWeekIterator
	Move(weeks int)
	Diff(from YearWeekIterator) int
	Year() YearIterator
	Date(dayOfWeek date.DayOfWeek) DateIterator
	FirstDate() DateIterator
	LastDate() DateIterator
}

func OfYearWeek(yearWeek date.YearWeek) YearWeekIterator {
	return &yearWeekIter{yearWeek: yearWeek}
}

type yearWeekIter struct {
	yearWeek date.YearWeek
}

func (y *yearWeekIter) Get() date.YearWeek {
	return y.yearWeek
}

func (y *yearWeekIter) Copy() YearWeekIterator {
	return OfYearWeek(y.Get())
}

func (y *yearWeekIter) Move(weeks int) {
	y.yearWeek = y.Get().Date(date.DayOfWeekMonday).Add(weeks * 7).YearWeek()
}

func (y *yearWeekIter) Diff(from YearWeekIterator) int {
	b := from.Get().Date(date.DayOfWeekMonday)
	e := y.Get().Date(date.DayOfWeekMonday)
	days := b.DaysUntil(e)
	return int(days / 7)
}

func (y *yearWeekIter) Year() YearIterator {
	return OfYear(y.Get().Year())
}

func (y *yearWeekIter) Date(dayOfWeek date.DayOfWeek) DateIterator {
	return OfDate(y.Get().Date(dayOfWeek))
}

func (y *yearWeekIter) FirstDate() DateIterator {
	return OfDate(y.Get().Date(date.DayOfWeekMonday))

}

func (y *yearWeekIter) LastDate() DateIterator {
	return OfDate(y.Get().Date(date.DayOfWeekSunday))
}
