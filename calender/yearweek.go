package calender

import "github.com/Jumpaku/gkoku/exact"

type YearWeek struct {
	year int
	week int
}

func (yw YearWeek) WeeksUntil(endExclusive YearWeek) int64 {
	bd := daysFromYyyyWwD(yw.year, yw.week, DayOfWeekMonday)
	ed := daysFromYyyyWwD(endExclusive.year, endExclusive.week, DayOfWeekMonday)
	w, _, _ := exact.DivTrunc(ed-bd, 7)
	return w
}

func (yw YearWeek) WholeYearsUntil(endExclusive YearWeek) int64 {
	if yw.year == endExclusive.year {
		return 0
	}
	hy := int64(yw.year - endExclusive.year)
	if yw.After(endExclusive) && yw.week < endExclusive.week {
		return hy + 1
	}
	if yw.Before(endExclusive) && yw.week > endExclusive.week {
		return hy - 1
	}
	return hy
}

func (yw YearWeek) Iter() YearWeekIterator {
	return nil
}

func YearWeekOf(year int, week int) YearWeek {
	return YearWeek{year: year, week: week}
}

var _ interface {
	YyyyWw() (year int, week int)
	Iter() YearWeekIterator
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
