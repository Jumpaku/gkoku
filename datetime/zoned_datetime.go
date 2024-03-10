package datetime

import (
	"github.com/Jumpaku/gkoku/calendar"
	"github.com/Jumpaku/gkoku/clock"
)

type ZonedDateTime interface {
	Zone() Zone
	InstantCandidates() []clock.Instant
	Date() calendar.Date
	Time() Time
}

func NewZonedDateTime(date calendar.Date, time Time, zone Zone) ZonedDateTime {
	return nil
}
