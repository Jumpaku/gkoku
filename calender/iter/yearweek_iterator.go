package iter

import "github.com/Jumpaku/gkoku/calender"

type YearWeekIterator interface {
	Get() calender.YearWeek
	Move(weeks int)
	Diff(from YearWeekIterator) int
	Year() YearIterator
	Date(dayOfWeek calender.DayOfWeek) DateIterator
	FirstDate() DateIterator
	LastDate() DateIterator
}

func OfYearWeek(yearWeek calender.YearWeek) YearWeekIterator {
	return &yearWeekIter{yearWeek: yearWeek}
}

type yearWeekIter struct {
	yearWeek calender.YearWeek
}

func (y *yearWeekIter) Get() calender.YearWeek {
	return y.yearWeek
}

func (y *yearWeekIter) Move(weeks int) {
	y.yearWeek = y.Get().Date(calender.DayOfWeekMonday).Add(weeks * 7).YearWeek()
}

func (y *yearWeekIter) Diff(from YearWeekIterator) int {
	days := y.Get().Date(calender.DayOfWeekMonday).DaysUntil(from.Get().Date(calender.DayOfWeekMonday))
	return int(days / 7)
}

func (y *yearWeekIter) Year() YearIterator {
	return OfYear(y.Get().Year())
}

func (y *yearWeekIter) Date(dayOfWeek calender.DayOfWeek) DateIterator {
	return OfDate(y.Get().Date(dayOfWeek))
}

func (y *yearWeekIter) FirstDate() DateIterator {
	return OfDate(y.Get().Date(calender.DayOfWeekMonday))

}

func (y *yearWeekIter) LastDate() DateIterator {
	return OfDate(y.Get().Date(calender.DayOfWeekSunday))
}
