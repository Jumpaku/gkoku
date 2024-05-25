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
	wy, wm := want.YM()
	gy, gm := got.YM()
	assert.Equal(t, wy, gy)
	assert.Equal(t, wm, gm)

}

func AssertEqualYearWeek(t *testing.T, want calendar.YearWeek, got calendar.YearWeek) {
	t.Helper()
	wy, ww := want.YW()
	gy, gw := got.YW()
	assert.Equal(t, wy, gy)
	assert.Equal(t, ww, gw)

}
