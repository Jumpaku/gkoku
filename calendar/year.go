package calendar

// Year represents a year.
type Year int

var _ interface {
	IsLeap() bool
	Days() int
	ContainsDay(dayOfYear int) bool
	ContainsWeek(week int) bool
	Date(dayOfYear int) Date
	FirstDate() Date
	LastDate() Date
	Week(weekOfYear int) YearWeek
	FirstWeek() YearWeek
	LastWeek() YearWeek
	Month(month Month) YearMonth
	FirstMonth() YearMonth
	LastMonth() YearMonth
	String() string
} = Year(0)

// IsLeap returns true if the year is a leap year.
func (y Year) IsLeap() bool {
	switch {
	case y%400 == 0:
		return true
	case y%100 == 0:
		return false
	case y%4 == 0:
		return true
	default:
		return false
	}
}

// Days returns the number of days in the year.
func (y Year) Days() int {
	if y.IsLeap() {
		return 366
	}
	return 365
}

// Weeks returns the number of weeks in the year.
func (y Year) Weeks() int {
	m1 := YearWeekOf(int(y), 1).Date(DayOfWeekMonday)
	m2 := YearWeekOf(int(y+1), 1).Date(DayOfWeekMonday)
	return int(m1.DaysUntil(m2) / 7)
}

// Date returns a Date of the dayOfYear in this year.
func (y Year) Date(dayOfYear int) Date {
	return DateOfYD(int(y), dayOfYear)
}

// FirstDate returns the first date in this year.
func (y Year) FirstDate() Date {
	return DateOfYD(int(y), 1)
}

// LastDate returns the last date in this year.
func (y Year) LastDate() Date {
	return DateOfYD(int(y), y.Days())
}

// Week returns a YearWeek of the weekOfYear in this year.
func (y Year) Week(weekOfYear int) YearWeek {
	return YearWeekOf(int(y), weekOfYear)

}

// FirstWeek returns the first week in this year.
func (y Year) FirstWeek() YearWeek {
	return YearWeekOf(int(y), 1)
}

// LastWeek returns the last week in this year.
func (y Year) LastWeek() YearWeek {
	return YearWeekOf(int(y), y.Weeks())
}

// Month returns a YearMonth of the month in this year.
func (y Year) Month(month Month) YearMonth {
	return YearMonthOf(int(y), month)
}

// FirstMonth returns the first month in this year.
func (y Year) FirstMonth() YearMonth {
	return YearMonthOf(int(y), MonthJanuary)
}

// LastMonth returns the last month in this year.
func (y Year) LastMonth() YearMonth {
	return YearMonthOf(int(y), MonthDecember)
}

// ContainsDay returns true if the dayOfYear is in this year.
func (y Year) ContainsDay(dayOfYear int) bool {
	return 1 <= dayOfYear && dayOfYear <= y.Days()
}

// ContainsWeek returns true if the week is in this year.
func (y Year) ContainsWeek(week int) bool {
	return 1 <= week && week <= y.Weeks()
}

// String returns the string representation of the year.
func (y Year) String() string {
	return FormatYear(y)
}
