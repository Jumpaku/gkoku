package iter

import "github.com/Jumpaku/gkoku/calender"

type YearMonthIterator interface {
	Get() calender.YearMonth
	Move(months int)
	Diff(from YearMonthIterator) int
	Year() YearIterator
	Days() int
	Date(dayOfMonth int) DateIterator
	FirstDate() DateIterator
	LastDate() DateIterator
}

func OfYearMonth(yearMonth calender.YearMonth) YearMonthIterator {
	return &yearMonthIter{yearMonth: yearMonth}
}

type yearMonthIter struct {
	yearMonth calender.YearMonth
}

func (y *yearMonthIter) Get() calender.YearMonth {
	return y.yearMonth
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
