package date

//go:generate stringer -type=Month -linecomment
type Month int

const (
	MonthUnspecified Month = iota // Unspecified
	MonthJanuary                  // January
	MonthFebruary                 // February
	MonthMarch                    // March
	MonthApril                    // April
	MonthMay                      // May
	MonthJune                     // June
	MonthJuly                     // July
	MonthAugust                   // August
	MonthSeptember                // September
	MonthOctober                  // October
	MonthNovember                 // November
	MonthDecember                 // December
)

var monthDays = [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
