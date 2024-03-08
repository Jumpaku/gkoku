package calender_test

import (
	"github.com/Jumpaku/gkoku/calender"
	"github.com/Jumpaku/gkoku/internal/tests/assert"
	"testing"
	"time"
)

func DateFromTime(t time.Time) calender.Date {
	return calender.YyyyMmDd(t.Year(), calender.Month(t.Month()), t.Day())
}
func ToTime(y int, m calender.Month, d int) time.Time {
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
}

func AssertEqualDate(t *testing.T, want calender.Date, got calender.Date) {
	t.Helper()
	wy, wm, wd := want.YyyyMmDd()
	gy, gm, gd := got.YyyyMmDd()
	assert.Equal(t, wy, gy)
	assert.Equal(t, wm, gm)
	assert.Equal(t, wd, gd)

}

func AssertEqualYearMonth(t *testing.T, want calender.YearMonth, got calender.YearMonth) {
	t.Helper()
	wy, wm := want.YyyyMm()
	gy, gm := got.YyyyMm()
	assert.Equal(t, wy, gy)
	assert.Equal(t, wm, gm)

}
