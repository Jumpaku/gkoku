package datetime

import (
	"github.com/Jumpaku/gkoku/calendar"
	"github.com/Jumpaku/gkoku/clock"
)

type OffsetDateTime interface {
	Offset() OffsetMinutes
	Instant() clock.Instant
	Date() calendar.Date
	Time() Time
}

func NewOffsetDateTime(date calendar.Date, time Time, offset OffsetMinutes) OffsetDateTime {
	return nil
}
