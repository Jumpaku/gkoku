package calendar

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
func (y Year) Days() int {
	if y.IsLeap() {
		return 366
	}
	return 365
}
func (y Year) Weeks() int {
	m1 := YearWeekOf(int(y), 1).Date(DayOfWeekMonday)
	m2 := YearWeekOf(int(y+1), 1).Date(DayOfWeekMonday)
	return int(m1.DaysUntil(m2) / 7)
}

func (y Year) Date(dayOfYear int) Date {
	return DateOfYD(int(y), dayOfYear)
}

func (y Year) FirstDate() Date {
	return DateOfYD(int(y), 1)
}

func (y Year) LastDate() Date {
	return DateOfYD(int(y), y.Days())
}

func (y Year) Week(weekOfYear int) YearWeek {
	return YearWeekOf(int(y), weekOfYear)

}

func (y Year) FirstWeek() YearWeek {
	return YearWeekOf(int(y), 1)
}

func (y Year) LastWeek() YearWeek {
	return YearWeekOf(int(y), y.Weeks())
}

func (y Year) Month(month Month) YearMonth {
	return YearMonthOf(int(y), month)
}

func (y Year) FirstMonth() YearMonth {
	return YearMonthOf(int(y), MonthJanuary)
}

func (y Year) LastMonth() YearMonth {
	return YearMonthOf(int(y), MonthDecember)
}

func (y Year) ContainsDay(dayOfYear int) bool {
	return 1 <= dayOfYear && dayOfYear <= y.Days()
}

func (y Year) ContainsWeek(week int) bool {
	return 1 <= week && week <= y.Weeks()
}

func (y Year) String() string {
	return FormatYear(y)
}
