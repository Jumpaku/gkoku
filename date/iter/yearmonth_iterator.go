package iter

import "github.com/Jumpaku/tokiope/date"

type YearMonthIterator interface {
	Get() date.YearMonth
	Copy() YearMonthIterator
	Move(months int)
	Diff(from YearMonthIterator) int
	Year() YearIterator
	Days() int
	Date(dayOfMonth int) DateIterator
	FirstDate() DateIterator
	LastDate() DateIterator
}

func OfYearMonth(yearMonth date.YearMonth) YearMonthIterator {
	return &yearMonthIter{yearMonth: yearMonth}
}

type yearMonthIter struct {
	yearMonth date.YearMonth
}

func (y *yearMonthIter) Get() date.YearMonth {
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
