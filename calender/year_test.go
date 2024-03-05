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
		t.Run(fmt.Sprintf("%d", sut), func(t *testing.T) {
			{
				want := 52
				if _, w := date.GoDate(int(sut), 12, 31).ISOWeek(); w == 53 {
					want = 53
				}
				assert.Equal(t, want, sut.Weeks())
			}
			{
				assert.Equal(t, false, sut.ContainsWeek(-1))
			}
			{
				assert.Equal(t, false, sut.ContainsWeek(0))
			}
			{
				assert.Equal(t, true, sut.ContainsWeek(1))
			}
			{
				assert.Equal(t, true, sut.ContainsWeek(52))
			}
			{
				_, w := date.GoDate(int(sut), 12, 31).ISOWeek()
				assert.Equal(t, w == 53, sut.ContainsWeek(53))
			}
			{
				assert.Equal(t, false, sut.ContainsWeek(54))
			}
		})
	}
}
func TestYear_Day(t *testing.T) {
	for sut := calender.Year(-9999); sut <= 9999; sut++ {
		t.Run(fmt.Sprintf("%d", sut), func(t *testing.T) {
			{
				assert.Equal(t, date.GoDate(int(sut), 12, 31).YearDay(), sut.Days())
			}
			{
				assert.Equal(t, false, sut.ContainsDay(-1))
			}
			{
				assert.Equal(t, false, sut.ContainsDay(0))
			}
			{
				assert.Equal(t, true, sut.ContainsDay(1))
			}
			{
				assert.Equal(t, true, sut.ContainsDay(365))
			}
			{
				yd := date.GoDate(int(sut), 12, 31).YearDay()
				assert.Equal(t, yd == 366, sut.ContainsDay(366))
			}
			{
				assert.Equal(t, false, sut.ContainsDay(367))
			}
		})
	}
}
