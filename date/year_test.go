package date_test

import (
	"fmt"
	"github.com/Jumpaku/gkoku/date"
	test_date "github.com/Jumpaku/gkoku/date/internal/cmd/testcases/date"
	"github.com/Jumpaku/gkoku/internal/tests/assert"
	calendar_test "github.com/Jumpaku/gkoku/internal/tests/date"
	"testing"
)

func TestYear_IsLeap(t *testing.T) {
	for sut := date.Year(-9999); sut <= 9999; sut++ {
		t.Run(fmt.Sprintf("%d", sut), func(t *testing.T) {
			want := test_date.GoDate(int(sut), 12, 31).YearDay() == 366
			got := sut.IsLeap()
			assert.Equal(t, want, got)
		})
	}
}

func TestYear_Weeks(t *testing.T) {
	for sut := date.Year(-9999); sut <= 9999; sut++ {
		t.Run(fmt.Sprintf("Weeks %d", sut), func(t *testing.T) {
			want := 52
			if _, w := test_date.GoDate(int(sut), 12, 31).ISOWeek(); w == 53 {
				want = 53
			}
			assert.Equal(t, want, sut.Weeks())
		})
	}
}

func TestYear_ContainsWeek(t *testing.T) {
	for sut := date.Year(-9999); sut <= 9999; sut++ {
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
			_, w := test_date.GoDate(int(sut), 12, 31).ISOWeek()
			assert.Equal(t, w == 53, sut.ContainsWeek(53))
		})
		t.Run(fmt.Sprintf("whether %d contains 54", sut), func(t *testing.T) {
			assert.Equal(t, false, sut.ContainsWeek(54))
		})
	}
}

func TestYear_Week(t *testing.T) {
	for sut := date.Year(-9999); sut <= 9999; sut++ {
		for w := 1; w <= sut.Weeks(); w++ {
			t.Run(fmt.Sprintf("Week %d", sut), func(t *testing.T) {
				gotY, gotW := sut.Week(w).YyyyWw()
				assert.Equal(t, int(sut), gotY)
				assert.Equal(t, w, gotW)
			})
		}
	}
}

func TestYear_FirstWeek(t *testing.T) {
	for sut := date.Year(-9999); sut <= 9999; sut++ {
		t.Run(fmt.Sprintf("FirstWee %d", sut), func(t *testing.T) {
			gotY, gotW := sut.FirstWeek().YyyyWw()
			assert.Equal(t, int(sut), gotY)
			assert.Equal(t, 1, gotW)
		})
	}
}

func TestYear_LastWeek(t *testing.T) {
	for sut := date.Year(-9999); sut <= 9999; sut++ {
		t.Run(fmt.Sprintf("LastWeek %d", sut), func(t *testing.T) {
			want := 52
			if _, w := test_date.GoDate(int(sut), 12, 31).ISOWeek(); w == 53 {
				want = 53
			}
			assert.Equal(t, want, sut.Weeks())
			gotY, gotW := sut.LastWeek().YyyyWw()
			assert.Equal(t, int(sut), gotY)
			assert.Equal(t, want, gotW)
		})
	}
}

func TestYear_Days(t *testing.T) {
	for sut := date.Year(-9999); sut <= 9999; sut++ {
		t.Run(fmt.Sprintf("%d Days", sut), func(t *testing.T) {
			assert.Equal(t, test_date.GoDate(int(sut), 12, 31).YearDay(), sut.Days())
		})
	}
}

func TestYear_ContainsDay(t *testing.T) {
	for sut := date.Year(-9999); sut <= 9999; sut++ {
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
			yd := test_date.GoDate(int(sut), 12, 31).YearDay()
			assert.Equal(t, yd == 366, sut.ContainsDay(366))
		})
		t.Run(fmt.Sprintf("whether %d contains 367", sut), func(t *testing.T) {
			assert.Equal(t, false, sut.ContainsDay(367))
		})
	}
}

func TestYear_FirstDate(t *testing.T) {
	for sut := date.Year(-9999); sut <= 9999; sut++ {
		t.Run(fmt.Sprintf(" %d", sut), func(t *testing.T) {
			got := sut.FirstDate()
			want := date.YyyyDdd(int(sut), 1)
			calendar_test.AssertEqualDate(t, want, got)
		})
	}
}

func TestYear_LastDate(t *testing.T) {
	for sut := date.Year(-9999); sut <= 9999; sut++ {
		t.Run(fmt.Sprintf(" %d", sut), func(t *testing.T) {
			got := sut.LastDate()
			want := date.YyyyDdd(int(sut), sut.Days())
			calendar_test.AssertEqualDate(t, want, got)
		})
	}
}

func TestYear_Date(t *testing.T) {
	for sut := date.Year(-9999); sut <= 9999; sut++ {
		for _, doy := range []int{1, 2, sut.Days() - 1, sut.Days()} {
			t.Run(fmt.Sprintf("%d-%d", sut, doy), func(t *testing.T) {
				got := sut.Date(doy)
				want := date.YyyyDdd(int(sut), doy)
				calendar_test.AssertEqualDate(t, want, got)
			})
		}
	}
}

func TestYear_Month(t *testing.T) {
	for sut := date.Year(-9999); sut <= 9999; sut++ {
		for m := date.MonthJanuary; m <= date.MonthDecember; m++ {
			t.Run(fmt.Sprintf("Month %d", sut), func(t *testing.T) {
				gotY, gotM := sut.Month(m).YyyyMm()
				assert.Equal(t, int(sut), gotY)
				assert.Equal(t, m, gotM)
			})
		}
	}
}

func TestYear_FirstMonth(t *testing.T) {
	for sut := date.Year(-9999); sut <= 9999; sut++ {
		t.Run(fmt.Sprintf("FirstMonth %d", sut), func(t *testing.T) {
			gotY, gotM := sut.FirstMonth().YyyyMm()
			assert.Equal(t, int(sut), gotY)
			assert.Equal(t, date.MonthJanuary, gotM)
		})
	}
}

func TestYear_LastMonth(t *testing.T) {
	for sut := date.Year(-9999); sut <= 9999; sut++ {
		t.Run(fmt.Sprintf("LastMonth %d", sut), func(t *testing.T) {
			gotY, gotM := sut.LastMonth().YyyyMm()
			assert.Equal(t, int(sut), gotY)
			assert.Equal(t, date.MonthDecember, gotM)
		})
	}
}
