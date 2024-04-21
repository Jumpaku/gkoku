package iter

import "github.com/Jumpaku/gkoku/date"

type DateIterator interface {
	Get() date.Date
	Copy() DateIterator
	Move(days int)
	Diff(from DateIterator) int
	Year() YearIterator
	YearMonth() YearMonthIterator
	YearWeek() YearWeekIterator
}

func OfDate(date date.Date) DateIterator {
	return &dateIter{date: date}
}

type dateIter struct {
	date date.Date
}

func (i *dateIter) Get() date.Date {
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
