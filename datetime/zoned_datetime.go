package datetime

import (
	"github.com/Jumpaku/gkoku/calendar"
	"github.com/Jumpaku/gkoku/clock"
	"github.com/Jumpaku/gkoku/datetime/zone"
	"slices"
)

type ZonedDateTime interface {
	Date() calendar.Date
	Time() Time
	Zone() zone.Zone
	InstantCandidates() []clock.Instant
}

func NewZonedDateTime(date calendar.Date, time Time, zone zone.Zone) ZonedDateTime {
	return zonedDateTime{
		date: date,
		time: time,
		zone: zone,
	}
}

type zonedDateTime struct {
	date calendar.Date
	time Time
	zone zone.Zone
}

func (d zonedDateTime) Zone() zone.Zone {
	return d.zone
}

func (d zonedDateTime) Date() calendar.Date {
	return d.date
}

func (d zonedDateTime) Time() Time {
	return d.time
}

func (d zonedDateTime) InstantCandidates() []clock.Instant {
	tu := NewOffsetDateTime(d.Date(), d.Time(), zone.OffsetMinutes(0)).Instant()
	tlo := zone.MinOffsetMinutes.AddTo(tu)
	thi := zone.MaxOffsetMinutes.AddTo(tu)
	z := d.Zone()
	ts := z.TransitionsBetween(tlo, thi)
	if len(ts) == 0 {
		o := z.FindOffset(tu)
		return []clock.Instant{o.AddTo(tu)}
	}

	var candidates []clock.Instant
	{
		t := ts[0]
		c := NewOffsetDateTime(d.Date(), d.Time(), t.OffsetMinutesBefore).Instant()
		if tlo.Cmp(c) <= 0 && c.Cmp(t.TransitionTimestamp) < 0 { // skip if c == t.TransitionTimestamp
			candidates = append(candidates, c)
		}
	}
	{
		t := ts[len(ts)-1]
		c := NewOffsetDateTime(d.Date(), d.Time(), t.OffsetMinutesAfter).Instant()
		if c.Between(t.TransitionTimestamp, thi) { // include even if c == thi
			candidates = append(candidates, c)
		}
	}
	for i := 1; i < len(ts); i++ {
		ta, tb := ts[i-1], ts[i]
		c := NewOffsetDateTime(d.Date(), d.Time(), ta.OffsetMinutesBefore).Instant()
		if ta.TransitionTimestamp.Cmp(c) <= 0 && c.Cmp(tb.TransitionTimestamp) < 0 { // skip if c == tb.TransitionTimestamp
			candidates = append(candidates, c)
		}
	}

	slices.SortFunc(candidates, clock.Instant.Cmp)
	uniqCandidates := []clock.Instant{}
	for _, c := range candidates {
		if len(uniqCandidates) == 0 || uniqCandidates[len(uniqCandidates)-1].Cmp(c) < 0 {
			uniqCandidates = append(uniqCandidates, c)
		}
	}

	return uniqCandidates
}
