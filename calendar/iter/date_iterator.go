package iter

import "github.com/Jumpaku/tokiope/calendar"

type DateIterator interface {
	Get() calendar.Date
	Copy() DateIterator
	Move(days int)
	Diff(from DateIterator) int
	Year() YearIterator
	YearMonth() YearMonthIterator
	YearWeek() YearWeekIterator
}

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
