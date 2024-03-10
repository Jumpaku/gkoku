package datetime

import (
	"github.com/Jumpaku/gkoku/calendar"
	"github.com/Jumpaku/gkoku/clock"
)

type ZonedDateTime interface {
	Date() calendar.Date
	Time() Time
	Zone() Zone
	InstantCandidates() []clock.Instant
}

func NewZonedDateTime(date calendar.Date, time Time, zone Zone) ZonedDateTime {
	return zonedDateTime{
		date: date,
		time: time,
		zone: zone,
	}
}

type zonedDateTime struct {
	date calendar.Date
	time Time
	zone Zone
}

func (d zonedDateTime) Zone() Zone {
	return d.zone
}

func (d zonedDateTime) Date() calendar.Date {
	return d.date
}

func (d zonedDateTime) Time() Time {
	return d.time
}

func (d zonedDateTime) InstantCandidates() []clock.Instant {
	//TODO implement me
	panic("implement me")
}
