package calendar

import (
	"fmt"
	"regexp"
	"strconv"
)

type DateFormat int

const (
	DateFormatAny DateFormat = iota
	DateFormatYMD
	DateFormatYWD
	DateFormatYD
)

// ParseDate parses a formatted string and returns the Date value it represents.
// The following formats are supported:
//
// - year-month-dayOfMonth:
//   - yyyy-mm-dd:
//
// - year-dayOfYear:
//   - yyyy-ddd
//
// - year-week-dayOfWeek:
//   - yyyy-Www-d
//
// where y, m, d, and w are decimal digits and W is a rune 'W'.
// Each of the above formats may have a prefix of '-' or '+' for the sign of the year.
func ParseDate(s string, format DateFormat) (d Date, err error) {
	reYMD := regexp.MustCompile(`^[-+]?\d{4,}-\d{2}-\d{2}$`)
	reYWD := regexp.MustCompile(`^[-+]?\d{4,}-W\d{2}-\d$`)
	reYD := regexp.MustCompile(`^[-+]?\d{4,}-\d{3}$`)
	switch format {
	default:
		return Date{}, fmt.Errorf("invalid date format: %q", s)
	case DateFormatYMD:
		if !reYMD.MatchString(s) {
			return Date{}, fmt.Errorf("invalid date format: %q", s)
		}
		return parseYMD(s)
	case DateFormatYWD:
		if !reYWD.MatchString(s) {
			return Date{}, fmt.Errorf("invalid date format: %q", s)
		}
		return parseYWD(s)
	case DateFormatYD:
		if !reYD.MatchString(s) {
			return Date{}, fmt.Errorf("invalid date format: %q", s)
		}
		return parseYD(s)
	case DateFormatAny:
		switch {
		default:
			return Date{}, fmt.Errorf("invalid date format: %q", s)
		case reYMD.MatchString(s):
			return parseYMD(s)
		case reYWD.MatchString(s):
			return parseYWD(s)
		case reYD.MatchString(s):
			return parseYD(s)
		}
	}
}

// parseYMD parses a formatted string that matches `^[-+]?\d{4}-\d{2}-\d{2}$` and returns the Date value it represents.
func parseYMD(s string) (d Date, err error) {
	n := len(s)

	year, err := strconv.Atoi(s[:n-6])
	if err != nil {
		return Date{}, fmt.Errorf("fail to parse year: %w", err)
	}

	month, err := strconv.Atoi(s[n-5 : n-3])
	if err != nil {
		return Date{}, fmt.Errorf("fail to parse month: %w", err)
	}

	dayOfMonth, err := strconv.Atoi(s[n-2:])
	if err != nil {
		return Date{}, fmt.Errorf("fail to parse day of month: %w", err)
	}

	if err := validateYMD(year, month, dayOfMonth); err != nil {
		return Date{}, fmt.Errorf("fail to parse date: %w", err)
	}

	return Date{days: daysFromYMD(year, Month(month), dayOfMonth)}, nil
}

// parseYWD parses a formatted string that matches `^[-+]?\d{4}-W\d{2}-\d{2}$` and returns the Date value it represents.
func parseYWD(s string) (d Date, err error) {
	n := len(s)

	year, err := strconv.Atoi(s[:n-6])
	if err != nil {
		return Date{}, fmt.Errorf("fail to parse year: %w", err)
	}

	week, err := strconv.Atoi(s[n-4 : n-2])
	if err != nil {
		return Date{}, fmt.Errorf("fail to parse week: %w", err)
	}

	dayOfWeek, err := strconv.Atoi(s[n-1:])
	if err != nil {
		return Date{}, fmt.Errorf("fail to parse day of week: %w", err)
	}

	if err := validateYWD(year, week, dayOfWeek); err != nil {
		return Date{}, fmt.Errorf("fail to parse date: %w", err)
	}

	return Date{days: daysFromYWD(year, week, DayOfWeek(dayOfWeek))}, nil
}

// parseYD parses a formatted string that matches `^[-+]?\d{4}-\d{3}$` and returns the Date value it represents.
func parseYD(s string) (d Date, err error) {
	n := len(s)

	year, err := strconv.Atoi(s[:n-4])
	if err != nil {
		return Date{}, fmt.Errorf("fail to parse year: %w", err)
	}

	dayOfYear, err := strconv.Atoi(s[n-3:])
	if err != nil {
		return Date{}, fmt.Errorf("fail to parse day of month: %w", err)
	}

	if err := validateYD(year, dayOfYear); err != nil {
		return Date{}, fmt.Errorf("fail to parse date: %w", err)
	}

	return Date{days: daysFromYD(year, dayOfYear)}, nil
}

// FormatDate returns a textual representation of the Date value formatted according to the format defined by the argument.
func FormatDate(d Date, format DateFormat) string {
	switch format {
	default:
		panic(fmt.Sprintf("invalid format: %v", format))
	case DateFormatYMD:
		y, m, dom := d.YMD()
		sign := ""
		if y < 0 {
			sign = "-"
			y = -y
		}
		return fmt.Sprintf(`%s%04d-%02d-%02d`, sign, y, m, dom)
	case DateFormatYD:
		y, doy := d.YD()
		sign := ""
		if y < 0 {
			sign = "-"
			y = -y
		}
		return fmt.Sprintf(`%s%04d-%03d`, sign, y, doy)
	case DateFormatYWD:
		y, w, dow := d.YWD()
		sign := ""
		if y < 0 {
			sign = "-"
			y = -y
		}
		return fmt.Sprintf(`%s%04d-W%02d-%1d`, sign, y, w, dow)
	}
}

// ParseYearMonth parses a formatted string and returns the YearMonth value it represents.
// The following format is supported:
//
// - year-month:
//   - yyyy-mm:
//
// where y and m are decimal digits.
// The above format may have a prefix of '-' or '+' for the sign of the year.
func ParseYearMonth(s string) (YearMonth, error) {
	if !regexp.MustCompile(`^[-+]?\d{4,}-\d{2}$`).MatchString(s) {
		return YearMonth{}, fmt.Errorf("invalid year-month format: %q", s)
	}
	n := len(s)

	year, err := strconv.Atoi(s[:n-3])
	if err != nil {
		return YearMonth{}, fmt.Errorf("fail to parse year: %w", err)
	}

	month, err := strconv.Atoi(s[n-2:])
	if err != nil {
		return YearMonth{}, fmt.Errorf("fail to parse month: %w", err)
	}

	if !(1 <= month && month <= 12) {
		return YearMonth{}, fmt.Errorf("fail to parse year-month: %w", err)
	}

	return YearMonthOf(year, Month(month)), nil
}

// FormatYearMonth returns a textual representation of the YearMonth value formatted.
func FormatYearMonth(ym YearMonth) string {
	y, m := ym.YyyyMm()
	sign := ""
	if y < 0 {
		sign = "-"
		y = -y
	}
	return fmt.Sprintf(`%s%04d-%02d`, sign, y, m)
}

// ParseYearWeek parses a formatted string and returns the YearWeek value it represents.
// The following format is supported:
//
// - year-week:
//   - yyyy-Www:
//
// where y and w are a decimal digits.
// The above format may have a prefix of '-' or '+' for the sign of the year.
func ParseYearWeek(s string) (YearWeek, error) {
	if !regexp.MustCompile(`^[-+]?\d{4,}-W\d{2}$`).MatchString(s) {
		return YearWeek{}, fmt.Errorf("invalid year-week format: %q", s)
	}

	n := len(s)

	year, err := strconv.Atoi(s[:n-4])
	if err != nil {
		return YearWeek{}, fmt.Errorf("fail to parse year: %w", err)
	}

	week, err := strconv.Atoi(s[n-2:])
	if err != nil {
		return YearWeek{}, fmt.Errorf("fail to parse week: %w", err)
	}

	if err := validateYWD(year, week, 1); err != nil {
		return YearWeek{}, fmt.Errorf("fail to parse date: %w", err)
	}

	return YearWeekOf(year, week), nil
}

// FormatYearWeek returns a textual representation of the YearWeek value formatted.
func FormatYearWeek(yw YearWeek) string {
	y, w := yw.YyyyWw()
	sign := ""
	if y < 0 {
		sign = "-"
		y = -y
	}
	return fmt.Sprintf(`%s%04d-W%02d`, sign, y, w)
}

// ParseYear parses a formatted string and returns the Year value it represents.
// The following format is supported:
//
// - year:
//   - yyyy:
//
// where y is a decimal digit.
// The above format may have a prefix of '-' or '+' for the sign of the year.
func ParseYear(s string) (Year, error) {
	if !regexp.MustCompile(`^[-+]?\d{4,}$`).MatchString(s) {
		return 0, fmt.Errorf("invalid year format: %q", s)
	}

	year, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("fail to parse year: %w", err)
	}

	return Year(year), nil
}

// FormatYear returns a textual representation of the Year value formatted.
func FormatYear(y Year) string {
	sign := ""
	if y < 0 {
		sign = "-"
		y = -y
	}
	return fmt.Sprintf(`%s%04d`, sign, y)
}
