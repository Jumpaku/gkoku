package iter

import "github.com/Jumpaku/tokiope/calendar"

type YearIterator interface {
	Get() calendar.Year
	Copy() YearIterator
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
	Month(month calendar.Month) YearMonthIterator
	FirstMonth() YearMonthIterator
	LastMonth() YearMonthIterator
}

func OfYear(year calendar.Year) YearIterator {
	return &yearIter{year: year}
}

type yearIter struct {
	year calendar.Year
}

func (y *yearIter) Get() calendar.Year {
	return y.year
}

func (y *yearIter) Copy() YearIterator {
	return OfYear(y.Get())
}

func (y *yearIter) Move(years int) {
	y.year = calendar.Year(int(y.Get()) + years)
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

func (y *yearIter) Month(month calendar.Month) YearMonthIterator {
	return OfYearMonth(y.Get().Month(month))
}

func (y *yearIter) FirstMonth() YearMonthIterator {
	return y.Month(calendar.MonthJanuary)
}

func (y *yearIter) LastMonth() YearMonthIterator {
	return y.Month(calendar.MonthDecember)
}
