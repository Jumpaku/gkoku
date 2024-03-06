package calender

import "github.com/Jumpaku/gkoku/exact"

type YearWeek struct {
	year int
	week int
}

func YearWeekOf(year int, week int) YearWeek {
	return YearWeek{year: year, week: week}
}

var _ interface {
	YyyyWw() (year int, week int)
	Add(weeks int) YearWeek
	Sub(weeks int) YearWeek
	Year() Year
	Date(dayOfWeek DayOfWeek) Date
	FirstDate() Date
	LastDate() Date
	WeeksUntil(endExclusive YearWeek) int64
	WholeYearsUntil(endExclusive YearWeek) int64
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

func (yw YearWeek) Add(weeks int) YearWeek {
	return yw.Date(DayOfWeekMonday).Add(7 * weeks).YearWeek()
}

func (yw YearWeek) Sub(weeks int) YearWeek {
	return yw.Date(DayOfWeekMonday).Sub(7 * weeks).YearWeek()
}

func (yw YearWeek) WeeksUntil(endExclusive YearWeek) int64 {
	bd := daysFromYyyyWwD(yw.year, yw.week, DayOfWeekMonday)
	ed := daysFromYyyyWwD(endExclusive.year, endExclusive.week, DayOfWeekMonday)
	w, _, _ := exact.DivTrunc(ed-bd, 7)
	return w
}

func (yw YearWeek) WholeYearsUntil(endExclusive YearWeek) int64 {
	by, bw := yw.YyyyWw()
	ey, ew := endExclusive.YyyyWw()
	wy := int64(Year(ey) - Year(by))

	if yw.After(endExclusive) && bw < ew {
		return wy + 1
	}
	if yw.Before(endExclusive) && bw > ew {
		return wy - 1
	}
	return wy
}
