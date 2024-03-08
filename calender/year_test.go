package calender_test

import (
	"fmt"
	"github.com/Jumpaku/gkoku/calender"
	"github.com/Jumpaku/gkoku/calender/testcases/cmd/date"
	"github.com/Jumpaku/gkoku/internal/tests/assert"
	"testing"
)

func TestYear_IsLeap(t *testing.T) {
	for sut := calender.Year(-9999); sut <= 9999; sut++ {
		t.Run(fmt.Sprintf("%d", sut), func(t *testing.T) {
			want := date.GoDate(int(sut), 12, 31).YearDay() == 366
			got := sut.IsLeap()
			assert.Equal(t, want, got)
		})
	}
}

func TestYear_Week(t *testing.T) {
	for sut := calender.Year(-9999); sut <= 9999; sut++ {
		t.Run(fmt.Sprintf("Weeks %d", sut), func(t *testing.T) {
			want := 52
			if _, w := date.GoDate(int(sut), 12, 31).ISOWeek(); w == 53 {
				want = 53
			}
			assert.Equal(t, want, sut.Weeks())
		})
		t.Run(fmt.Sprintf("whether %d contains -1", sut), func(t *testing.T) {
			assert.Equal(t, false, sut.ContainsWeek(-1))
		})
		t.Run(fmt.Sprintf("whether %d contains 0", sut), func(t *testing.T) {
			assert.Equal(t, false, sut.ContainsWeek(0))
		})
		t.Run(fmt.Sprintf("whether %d contains 1", sut), func(t *testing.T) {
			assert.Equal(t, true, sut.ContainsWeek(1))
		})
		t.Run(fmt.Sprintf("whether %d contains 51", sut), func(t *testing.T) {
			assert.Equal(t, true, sut.ContainsWeek(51))
		})
		t.Run(fmt.Sprintf("whether %d contains 52", sut), func(t *testing.T) {
			assert.Equal(t, true, sut.ContainsWeek(52))
		})
		t.Run(fmt.Sprintf("whether %d contains 53", sut), func(t *testing.T) {
			_, w := date.GoDate(int(sut), 12, 31).ISOWeek()
			assert.Equal(t, w == 53, sut.ContainsWeek(53))
		})
		t.Run(fmt.Sprintf("whether %d contains 54", sut), func(t *testing.T) {
			assert.Equal(t, false, sut.ContainsWeek(54))
		})
	}
}
func TestYear_Day(t *testing.T) {
	for sut := calender.Year(-9999); sut <= 9999; sut++ {
		t.Run(fmt.Sprintf("%d Days", sut), func(t *testing.T) {
			assert.Equal(t, date.GoDate(int(sut), 12, 31).YearDay(), sut.Days())
		})
		t.Run(fmt.Sprintf("whether %d contains -1", sut), func(t *testing.T) {
			assert.Equal(t, false, sut.ContainsDay(-1))
		})
		t.Run(fmt.Sprintf("whether %d contains 0", sut), func(t *testing.T) {
			assert.Equal(t, false, sut.ContainsDay(0))
		})
		t.Run(fmt.Sprintf("whether %d contains 1", sut), func(t *testing.T) {
			assert.Equal(t, true, sut.ContainsDay(1))
		})
		t.Run(fmt.Sprintf("whether %d contains 364", sut), func(t *testing.T) {
			assert.Equal(t, true, sut.ContainsDay(364))
		})
		t.Run(fmt.Sprintf("whether %d contains 365", sut), func(t *testing.T) {
			assert.Equal(t, true, sut.ContainsDay(365))
		})
		t.Run(fmt.Sprintf("whether %d contains 366", sut), func(t *testing.T) {
			yd := date.GoDate(int(sut), 12, 31).YearDay()
			assert.Equal(t, yd == 366, sut.ContainsDay(366))
		})
		t.Run(fmt.Sprintf("whether %d contains 367", sut), func(t *testing.T) {
			assert.Equal(t, false, sut.ContainsDay(367))
		})
	}
}
