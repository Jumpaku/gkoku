package iter

import "github.com/Jumpaku/gkoku/calender"

type YearIterator interface {
	Get() calender.Year
	Move(years int)
	Diff(from YearIterator) int
	Date(dayOfYear int) DateIterator
	Days() int
	FirstDate() DateIterator
	LastDate() DateIterator
	Weeks() int
	Week(week int) YearWeekIterator
	FirstWeek() YearWeekIterator
	LastWeek() YearWeekIterator
	Month(month calender.Month) YearMonthIterator
	FirstMonth() YearMonthIterator
	LastMonth() YearMonthIterator
}

func OfYear(year calender.Year) YearIterator {
	return &yearIter{year: year}
}

type yearIter struct {
	year calender.Year
}

func (y *yearIter) Get() calender.Year {
	return y.year
}

func (y *yearIter) Move(years int) {
	y.year = calender.Year(int(y.Get()) + years)
}

func (y *yearIter) Diff(from YearIterator) int {
	return int(y.Get() - from.Get())
}

func (y *yearIter) Date(dayOfYear int) DateIterator {
	return OfDate(y.Get().Date(dayOfYear))
}

func (y *yearIter) Days() int {
	return y.Get().Days()
}

func (y *yearIter) FirstDate() DateIterator {
	return y.Date(1)
}

func (y *yearIter) LastDate() DateIterator {
	return y.Date(y.Days())
}

func (y *yearIter) Weeks() int {
	return y.Get().Weeks()
}

func (y *yearIter) Week(week int) YearWeekIterator {
	return OfYearWeek(y.Get().Week(week))

}

func (y *yearIter) FirstWeek() YearWeekIterator {
	return y.Week(1)
}

func (y *yearIter) LastWeek() YearWeekIterator {
	return y.Week(y.Weeks())
}

func (y *yearIter) Month(month calender.Month) YearMonthIterator {
	return OfYearMonth(y.Get().Month(month))
}

func (y *yearIter) FirstMonth() YearMonthIterator {
	return y.Month(calender.MonthJanuary)
}

func (y *yearIter) LastMonth() YearMonthIterator {
	return y.Month(calender.MonthDecember)
}
