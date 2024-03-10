package calendar

import (
	"cmp"
	"fmt"
	"github.com/Jumpaku/gkoku/internal/exact"
	"github.com/Jumpaku/go-assert"
)

// Date . The zero value represents 0001-01-01
type Date struct {
	days int64
}

var _ interface {
	Unix() int64
	YyyyMmDd() (year int, month Month, dayOfMonth int)
	YyyyWwD() (year int, week int, dayOfWeek DayOfWeek)
	YyyyDdd() (year int, dayOfYear int)
	Add(days int) Date
	Sub(days int) Date
	Year() Year
	YearMonth() YearMonth
	YearWeek() YearWeek
	DaysUntil(endExclusive Date) int64
	WholeMonthsUntil(endExclusive Date) int64
	WholeWeeksUntil(endExclusive Date) int64
	WholeYearsUntil(endExclusive Date) int64
	Cmp(other Date) int
	Equal(other Date) bool
	Before(other Date) bool
	After(other Date) bool
	String() string
} = Date{}

func YyyyMmDd(year int, month Month, dayOfMonth int) Date {
	assert.Params(MonthJanuary <= month && month <= MonthDecember, "month must be in [%d, %d]: %d", MonthJanuary, MonthDecember, month)
	lastDayOfMonth := YearMonthOf(year, month).Days()
	assert.Params(1 <= dayOfMonth && dayOfMonth <= lastDayOfMonth, "dayOfMonth must be in [%d, %d]: %d", 1, lastDayOfMonth, dayOfMonth)
	return Date{days: daysFromYyyyMmDd(year, month, dayOfMonth)}
}

func YyyyWwD(year int, week int, dayOfWeek DayOfWeek) Date {
	assert.Params(DayOfWeekMonday <= dayOfWeek && dayOfWeek <= DayOfWeekSunday, "dayOfWeek must be in [%d, %d]: %d", DayOfWeekMonday, DayOfWeekSunday, dayOfWeek)
	assert.Params(1 <= week && week <= 53, "week must be in [%d, %d]: %d", 1, 53, week)
	days := daysFromYyyyWwD(year, week, dayOfWeek)
	if week == 53 {
		assert.Params(days < daysFromYyyyWwD(year+1, 1, dayOfWeek),
			"week must be in [%d, %d]: %d", 1, 52, week)
	}
	return Date{days: days}
}
func YyyyDdd(year int, dayOfYear int) Date {
	lastDayOfYear := Year(year).Days()
	assert.Params(1 <= dayOfYear && dayOfYear <= lastDayOfYear, "day must be in [%d, %d]: %d", 1, lastDayOfYear, dayOfYear)
	return Date{days: daysFromYyyyDdd(year, dayOfYear)}
}

func (d Date) Unix() int64 {
	return toEpochDays(d.YyyyMmDd())
}

func (d Date) YyyyMmDd() (year int, month Month, day int) {
	return daysToYyyyMmDd(d.days)
}

func (d Date) YyyyWwD() (year int, week int, dayOfWeek DayOfWeek) {
	return daysToYyyyWwD(d.days)
}

func (d Date) YyyyDdd() (year int, dayOfYear int) {
	return daysToYyyyDdd(d.days)
}

func (d Date) Year() Year {
	y, _ := d.YyyyDdd()
	return Year(y)
}

func (d Date) YearMonth() YearMonth {
	y, m, _ := d.YyyyMmDd()
	return YearMonthOf(y, m)
}

func (d Date) YearWeek() YearWeek {
	y, w, _ := d.YyyyWwD()
	return YearWeekOf(y, w)
}

func (d Date) Cmp(other Date) int {
	return cmp.Compare(d.days, other.days)
}

func (d Date) Equal(other Date) bool {
	return d.days == other.days
}

func (d Date) Before(other Date) bool {
	return d.days < other.days
}

func (d Date) After(other Date) bool {
	return d.days > other.days
}

func (d Date) Add(days int) Date {
	return Date{d.days + int64(days)}
}

func (d Date) Sub(days int) Date {
	return Date{d.days - int64(days)}

}

func (d Date) DaysUntil(endExclusive Date) int64 {
	return endExclusive.days - d.days
}

func (d Date) WholeMonthsUntil(endExclusive Date) int64 {
	by, bm, bd := d.YyyyMmDd()
	ey, em, ed := endExclusive.YyyyMmDd()
	wm := YearMonthOf(by, bm).MonthsUntil(YearMonthOf(ey, em))

	if d.After(endExclusive) && bd < ed {
		return wm + 1
	}
	if d.Before(endExclusive) && bd > ed {
		return wm - 1
	}
	return wm
}

func (d Date) WholeWeeksUntil(endExclusive Date) int64 {
	days := d.DaysUntil(endExclusive)
	ww, _, _ := exact.DivTrunc(days, 7)
	return ww
}

func (d Date) WholeYearsUntil(endExclusive Date) int64 {
	wy, _, _ := exact.DivTrunc(d.WholeMonthsUntil(endExclusive), 12)
	return wy
}

func (d Date) String() string {
	return FormatDate(d, DateFormatYyyyMmDd)
}

func toEpochDays(year int, month Month, day int) int64 {
	y := int64(year)
	m := int64(month)
	d := int64(day)
	var total int64 = 0
	total += 365 * y
	if y >= 0 {
		total += (y+3)/4 - (y+99)/100 + (y+399)/400
	} else {
		total -= y/-4 - y/-100 + y/-400
	}
	total += (367*m - 362) / 12
	total += d - 1
	if m > 2 {
		total--
		if !Year(year).IsLeap() {
			total--
		}
	}
	return total - days0000To1970
}

func fromEpochDays(epochDays int64) (int, Month, int) {
	var zeroDay = epochDays + days0000To1970
	// find the march-based year
	zeroDay -= 60 // adjust to 0000-03-01 so leap day is at end of four years cycle

	var adjust int64 = 0
	if zeroDay < 0 { // adjust negative years to positive for calculation
		adjustCycles := (zeroDay+1)/daysPerCycle - 1
		adjust = adjustCycles * 400
		zeroDay += -adjustCycles * daysPerCycle
	}
	var yearEst = (400*zeroDay + 591) / daysPerCycle
	var doyEst = zeroDay - (365*yearEst + yearEst/4 - yearEst/100 + yearEst/400)
	if doyEst < 0 { // fix estimate
		yearEst--
		doyEst = zeroDay - (365*yearEst + yearEst/4 - yearEst/100 + yearEst/400)
	}
	yearEst += adjust // reset any negative year

	marchDoy0 := doyEst

	// convert march-based values back to january-based
	marchMonth0 := (marchDoy0*5 + 2) / 153
	month := (marchMonth0+2)%12 + 1
	dom := marchDoy0 - (marchMonth0*306+5)/10 + 1
	yearEst += marchMonth0 / 10

	return int(yearEst), Month(month), int(dom)
}

func daysToYyyyMmDd(days int64) (year int, month Month, day int) {
	return fromEpochDays(days - days0000To1970 + 366)
}

func firstMondayIn(year int) int64 {
	jan4 := daysFromYyyyMmDd(year, MonthJanuary, 4)
	_, dow, _ := exact.DivFloor(jan4-firstMondayIn2000, 7)
	return jan4 - dow
}

func daysToYyyyWwD(days int64) (year int, week int, dayOfWeek DayOfWeek) {
	year, _, _ = daysToYyyyMmDd(days)
	firstMonday := firstMondayIn(year)
	if days < firstMonday {
		year--
		firstMonday = firstMondayIn(year)
	} else if nextMonday := firstMondayIn(year + 1); days >= nextMonday {
		year++
		firstMonday = nextMonday
	}

	week = int((days-firstMonday)/7) + 1
	dayOfWeek = DayOfWeek((days-firstMonday)%7) + 1
	return year, week, dayOfWeek
}

var yearDays = []int{
	0,
	0 + 31,
	0 + 31 + 28,
	0 + 31 + 28 + 31,
	0 + 31 + 28 + 31 + 30,
	0 + 31 + 28 + 31 + 30 + 31,
	0 + 31 + 28 + 31 + 30 + 31 + 30,
	0 + 31 + 28 + 31 + 30 + 31 + 30 + 31,
	0 + 31 + 28 + 31 + 30 + 31 + 30 + 31 + 31,
	0 + 31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30,
	0 + 31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31,
	0 + 31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31 + 30,
	0 + 31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31 + 30 + 31,
}

func daysToYyyyDdd(days int64) (year int, day int) {
	y, m, d := daysToYyyyMmDd(days)
	leap := 0
	if Year(y).IsLeap() && m > 2 {
		leap = 1
	}
	return y, yearDays[m-1] + d + leap
}
func daysFromYyyyMmDd(year int, month Month, day int) (days int64) {
	return toEpochDays(year, month, day) + days0000To1970 - 366
}
func daysFromYyyyWwD(year int, week int, dayOfWeek DayOfWeek) (days int64) {
	firstMonday := firstMondayIn(year)
	return firstMonday + 7*int64(week-1) + int64(dayOfWeek-1)
}
func daysFromYyyyDdd(year int, dayOfYear int) (days int64) {
	return daysFromYyyyMmDd(year, MonthJanuary, 1) - 1 + int64(dayOfYear)
}

func validateYyyyMmDd(year, month, dayOfMonth int) error {
	//if YearMin <= Year(year) && Year(year) <= YearMax {
	//	return fmt.Errorf("year must be in [%d, %d]: %d", YearMin, YearMax, year)
	//}
	if !(MonthJanuary <= Month(month) && Month(month) <= MonthDecember) {
		return fmt.Errorf("month must be in [%d, %d]: %d", MonthJanuary, MonthDecember, month)
	}
	lastDayOfMonth := YearMonthOf(year, Month(month)).Days()
	if !(1 <= dayOfMonth && dayOfMonth <= lastDayOfMonth) {
		return fmt.Errorf("day of month must be in [%d, %d]: %d", 1, lastDayOfMonth, dayOfMonth)
	}
	return nil
}

func validateYyyyWwD(year, week, dayOfWeek int) error {
	//if YearMin <= Year(year) && Year(year) <= YearMax {
	//	return fmt.Errorf("year must be in [%d, %d]: %d", YearMin, YearMax, year)
	//}
	if !Year(year).ContainsWeek(week) {
		return fmt.Errorf("month must be in [%d, %d]: %d", 1, Year(year).Weeks(), week)
	}

	if !(DayOfWeekMonday <= DayOfWeek(dayOfWeek) && DayOfWeek(dayOfWeek) <= DayOfWeekSunday) {
		return fmt.Errorf("day of month must be in [%d, %d]: %d", DayOfWeekMonday, DayOfWeekSunday, dayOfWeek)
	}
	return nil
}

func validateYyyyDdd(year, dayOfYear int) error {
	//if YearMin <= Year(year) && Year(year) <= YearMax {
	//	return fmt.Errorf("year must be in [%d, %d]: %d", YearMin, YearMax, year)
	//}
	lastDayOfYear := Year(year).Days()
	if !(1 <= dayOfYear && dayOfYear <= lastDayOfYear) {
		return fmt.Errorf("month must be in [%d, %d]: %d", 1, lastDayOfYear, dayOfYear)
	}
	return nil
}

const (
	// The number of days in a 400 years cycle.
	daysPerCycle = 146097

	// The number of days from year zero to year 1970.
	// There are five 400 year cycles from year zero to 2000.
	// There are 7 leap years from 1970 to 2000.
	days0000To1970 = daysPerCycle*5 - (30*365 + 7)

	// The number of days from year 0001-01-01 to year 2000-01-03.
	firstMondayIn2000 = 730121
)
