package calender

type DateIterator interface {
	Add(days int)
	Sub(days int)
	Get() Date
	Year() YearIterator
	YearMonth() YearMonthIterator
	YearWeek() YearWeekIterator
}

type YearWeekIterator interface {
	Add(weeks int)
	Sub(weeks int)
	Get() YearWeek
	Year() YearIterator
	Date(dayOfWeek DayOfWeek) DateIterator
	FirstDate() DateIterator
	LastDate() DateIterator
}

type YearMonthIterator interface {
	Add(months int)
	Sub(months int)
	Get() YearMonth
	Year() YearIterator
	Date(dayOfMonth int) DateIterator
	FirstDate() DateIterator
	LastDate() DateIterator
}

type YearIterator interface {
	Add(years int)
	Sub(years int)
	Get() Year
	Date(dayOfYear int) DateIterator
	FirstDate() DateIterator
	BackDate() DateIterator
	Week(weekOfYear int) YearWeekIterator
	FirstWeek() YearWeekIterator
	BackWeek() YearWeekIterator
	Month(month Month) YearMonthIterator
	FirstMonth() YearMonthIterator
	LastMonth() YearMonthIterator
}
