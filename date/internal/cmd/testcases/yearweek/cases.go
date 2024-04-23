package yearweek

import (
	"github.com/Jumpaku/tokiope/date/internal/cmd/testcases/date"
)

type YearWeekField struct {
	Year int
	Week int
}

func GetExampleYearWeeks() []YearWeekField {
	yws := []YearWeekField{}
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

	for _, y := range year {
		yws = append(yws, YearWeekField{Year: y, Week: 1})
		yws = append(yws, YearWeekField{Year: y, Week: 2})
		yws = append(yws, YearWeekField{Year: y, Week: 51})
		yws = append(yws, YearWeekField{Year: y, Week: 52})
		if yy, w := date.GoDate(y, 12, 31).ISOWeek(); y == yy && w == 53 {
			yws = append(yws, YearWeekField{Year: y, Week: w})
		}
	}
	return yws
}
