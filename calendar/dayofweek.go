package calendar

//go:generate stringer -type=DayOfWeek -linecomment
type DayOfWeek int

const (
	DayOfWeekUnspecified DayOfWeek = iota // Unspecified
	DayOfWeekMonday                       // Monday
	DayOfWeekTuesday                      // Tuesday
	DayOfWeekWednesday                    // Wednesday
	DayOfWeekThursday                     // Thursday
	DayOfWeekFriday                       // Friday
	DayOfWeekSaturday                     // Saturday
	DayOfWeekSunday                       // Sunday
)
