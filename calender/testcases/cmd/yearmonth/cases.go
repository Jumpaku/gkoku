package yearmonth

type YearMonthField struct {
	Year  int
	Month int
}

func GetExampleYearMonths() []YearMonthField {
	yms := []YearMonthField{}
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
		for m := 1; m <= 12; m++ {
			yms = append(yms, YearMonthField{Year: y, Month: m})
		}
	}
	return yms
}
