package date

import (
	"time"
)

type DateField struct {
	Year  int
	Month int
	Day   int
}

type PeriodField struct {
	Years  int
	Months int
	Days   int
}

func IsLeap(year int) bool {
	return time.Date(year, 12, 31, 0, 0, 0, 0, time.UTC).YearDay() == 366
}

func GoDate(y, m, d int) time.Time {
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)

}

func GetExampleDates() []DateField {
	dates := []DateField{}
	year := []int{
		-9999,
		-401, -400, -399,
		-101, -100, -99,
		-1, 0, 1,
		99, 100, 101,
		1999, 2000, 2001,
		2020, 2021, 2022, 2023, 2024, 2025, 2026, 2027, 2028, 2029, 2030, 2031, 2032, 2033, 2034, 2035, 2036, 2037, 2038, 2039, 2040,
		2099, 2100, 2101,
		9999,
	}
	monthDay := [][]int{
		{},
		{1, 2, 30, 31},
		{1, 2, 27, 28, 29},
		{1, 2, 30, 31},
		{1, 2, 29, 30},
		{1, 2, 30, 31},
		{1, 2, 29, 30},
		{1, 2, 30, 31},
		{1, 2, 30, 31},
		{1, 2, 29, 30},
		{1, 2, 30, 31},
		{1, 2, 29, 30},
		{1, 2, 30, 31},
	}

	for _, y := range year {
		for m, ds := range monthDay {
			for _, d := range ds {
				if m == 2 && d == 29 && !IsLeap(y) {
					continue
				}
				dates = append(dates, DateField{Year: y, Month: m, Day: d})
			}
		}
	}
	return dates
}
