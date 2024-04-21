package date_test

import (
	"github.com/Jumpaku/gkoku/date"
	"time"
)

func DateFromTime(t time.Time) date.Date {
	return date.YyyyMmDd(t.Year(), date.Month(t.Month()), t.Day())
}
func ToTime(y int, m date.Month, d int) time.Time {
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
}
