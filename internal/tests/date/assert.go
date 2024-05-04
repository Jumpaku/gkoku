package date

import (
	"github.com/Jumpaku/tokiope/calendar"
	"github.com/Jumpaku/tokiope/internal/tests/assert"
	"testing"
)

func AssertEqualDate(t *testing.T, want calendar.Date, got calendar.Date) {
	t.Helper()
	wy, wm, wd := want.YMD()
	gy, gm, gd := got.YMD()
	assert.Equal(t, wy, gy)
	assert.Equal(t, wm, gm)
	assert.Equal(t, wd, gd)
}

func AssertEqualYearMonth(t *testing.T, want calendar.YearMonth, got calendar.YearMonth) {
	t.Helper()
	wy, wm := want.YyyyMm()
	gy, gm := got.YyyyMm()
	assert.Equal(t, wy, gy)
	assert.Equal(t, wm, gm)

}

func AssertEqualYearWeek(t *testing.T, want calendar.YearWeek, got calendar.YearWeek) {
	t.Helper()
	wy, ww := want.YyyyWw()
	gy, gw := got.YyyyWw()
	assert.Equal(t, wy, gy)
	assert.Equal(t, ww, gw)

}
