package calendar

// YearWeek represents a year and a week.
type YearWeek struct {
	year int
	week int
}

// YearWeekOf returns a YearWeek of the given year and week.
func YearWeekOf(year int, week int) YearWeek {
	return YearWeek{year: year, week: week}
}

var _ interface {
	YW() (year int, week int)
	Year() Year
	Date(dayOfWeek DayOfWeek) Date
	FirstDate() Date
	LastDate() Date
	Cmp(other YearWeek) int
	Equal(other YearWeek) bool
	Before(other YearWeek) bool
	After(other YearWeek) bool
	String() string
} = YearWeek{}

// YW returns the year and the week.
func (yw YearWeek) YW() (year int, week int) {
	return yw.year, yw.week
}

// Year returns the year.
func (yw YearWeek) Year() Year {
	return Year(yw.year)
}

// Date returns a Date of the dayOfWeek in this year and week.
func (yw YearWeek) Date(dayOfWeek DayOfWeek) Date {
	return DateOfYWD(yw.year, yw.week, dayOfWeek)
}

// FirstDate returns the first date in this year and week.
func (yw YearWeek) FirstDate() Date {
	return DateOfYWD(yw.year, yw.week, DayOfWeekMonday)

}

// LastDate returns the last date in this year and week.
func (yw YearWeek) LastDate() Date {
	return DateOfYWD(yw.year, yw.week, DayOfWeekSunday)
}

// Cmp compares two YearWeeks.
func (yw YearWeek) Cmp(other YearWeek) int {
	if yw.Before(other) {
		return -1
	}
	if yw.After(other) {
		return 1
	}
	return 0
}

// Equal returns true if the YearWeek is equal to the other.
func (yw YearWeek) Equal(other YearWeek) bool {
	return yw.year == other.year && yw.week == other.week
}

// Before returns true if the YearWeek is before the other.
func (yw YearWeek) Before(other YearWeek) bool {
	if yw.year == other.year {
		return yw.week < other.week
	}
	return yw.year < other.year
}

// After returns true if the YearWeek is after the other.
func (yw YearWeek) After(other YearWeek) bool {
	if yw.year == other.year {
		return yw.week > other.week
	}
	return yw.year > other.year
}

// String returns the string representation of the YearWeek.
func (yw YearWeek) String() string {
	return FormatYearWeek(yw)
}
