package datetime

import (
	"fmt"
	"github.com/Jumpaku/gkoku/calendar"
	"github.com/Jumpaku/gkoku/clock"
	zone2 "github.com/Jumpaku/gkoku/zone"
	"slices"
)

type ZonedDateTime interface {
	Date() calendar.Date
	Time() Time
	Zone() zone2.Zone
	String() string
	InstantCandidates() []clock.Instant
}

func NewZonedDateTime(date calendar.Date, time Time, zone zone2.Zone) ZonedDateTime {
	return zonedDateTime{
		date: date,
		time: time,
		zone: zone,
	}
}

type zonedDateTime struct {
	date calendar.Date
	time Time
	zone zone2.Zone
}

func (d zonedDateTime) String() string {
	return fmt.Sprintf(`%s%s [%s]`, d.Date().String(), d.Time().String(), d.Zone().ID())
}

func (d zonedDateTime) Zone() zone2.Zone {
	return d.zone
}

func (d zonedDateTime) Date() calendar.Date {
	return d.date
}

func (d zonedDateTime) Time() Time {
	return d.time
}

func (d zonedDateTime) InstantCandidates() []clock.Instant {
	tu := NewOffsetDateTime(d.Date(), d.Time(), zone2.OffsetMinutes(0)).Instant()
	tlo := zone2.MinOffsetMinutes.AddTo(tu)
	thi := zone2.MaxOffsetMinutes.AddTo(tu)
	z := d.Zone()
	ts := z.TransitionsBetween(tlo, thi)
	if len(ts) == 0 {
		offset := z.FindOffset(tu)
		return []clock.Instant{NewOffsetDateTime(d.Date(), d.Time(), offset).Instant()}
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
