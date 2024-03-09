package calender

type YearWeek struct {
	year int
	week int
}

func YearWeekOf(year int, week int) YearWeek {
	return YearWeek{year: year, week: week}
}

var _ interface {
	YyyyWw() (year int, week int)
	Year() Year
	Date(dayOfWeek DayOfWeek) Date
	FirstDate() Date
	LastDate() Date
	Cmp(other YearWeek) int
	Equal(other YearWeek) bool
	Before(other YearWeek) bool
	After(other YearWeek) bool
} = YearWeek{}

func (yw YearWeek) YyyyWw() (year int, week int) {
	return yw.year, yw.week
}

func (yw YearWeek) Year() Year {
	return Year(yw.year)
}

func (yw YearWeek) Date(dayOfWeek DayOfWeek) Date {
	return YyyyWwD(yw.year, yw.week, dayOfWeek)
}

func (yw YearWeek) FirstDate() Date {
	return YyyyWwD(yw.year, yw.week, DayOfWeekMonday)

}

func (yw YearWeek) LastDate() Date {
	return YyyyWwD(yw.year, yw.week, DayOfWeekSunday)
}

func (yw YearWeek) Cmp(other YearWeek) int {
	if yw.Before(other) {
		return -1
	}
	if yw.After(other) {
		return 1
	}
	return 0
}

func (yw YearWeek) Equal(other YearWeek) bool {
	return yw.year == other.year && yw.week == other.week
}

func (yw YearWeek) Before(other YearWeek) bool {
	if yw.year == other.year {
		return yw.week < other.week
	}
	return yw.year < other.year
}

func (yw YearWeek) After(other YearWeek) bool {
	if yw.year == other.year {
		return yw.week > other.week
	}
	return yw.year > other.year
}
