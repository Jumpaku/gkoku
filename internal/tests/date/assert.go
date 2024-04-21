package date

import (
	"github.com/Jumpaku/gkoku/date"
	"github.com/Jumpaku/gkoku/internal/tests/assert"
	"testing"
)

func AssertEqualDate(t *testing.T, want date.Date, got date.Date) {
	t.Helper()
	wy, wm, wd := want.YyyyMmDd()
	gy, gm, gd := got.YyyyMmDd()
	assert.Equal(t, wy, gy)
	assert.Equal(t, wm, gm)
	assert.Equal(t, wd, gd)
}

func AssertEqualYearMonth(t *testing.T, want date.YearMonth, got date.YearMonth) {
	t.Helper()
	wy, wm := want.YyyyMm()
	gy, gm := got.YyyyMm()
	assert.Equal(t, wy, gy)
	assert.Equal(t, wm, gm)

}

func AssertEqualYearWeek(t *testing.T, want date.YearWeek, got date.YearWeek) {
	t.Helper()
	wy, ww := want.YyyyWw()
	gy, gw := got.YyyyWw()
	assert.Equal(t, wy, gy)
	assert.Equal(t, ww, gw)

}
