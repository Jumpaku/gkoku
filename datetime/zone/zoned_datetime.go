package zone

import (
	"fmt"
	"github.com/Jumpaku/tokiope"
	"github.com/Jumpaku/tokiope/calendar"
	"github.com/Jumpaku/tokiope/datetime"
	"slices"
)

type ZonedDateTime interface {
	Date() calendar.Date
	Time() datetime.Time
	Zone() Zone
	String() string
	InstantCandidates() []tokiope.Instant
}

func NewZonedDateTime(date calendar.Date, time datetime.Time, zone Zone) ZonedDateTime {
	return zonedDateTime{
		date: date,
		time: time,
		zone: zone,
	}
}

type zonedDateTime struct {
	date calendar.Date
	time datetime.Time
	zone Zone
}

func (d zonedDateTime) String() string {
	return fmt.Sprintf(`%s%s [%s]`, d.Date().String(), d.Time().String(), d.Zone().ID())
}

func (d zonedDateTime) Zone() Zone {
	return d.zone
}

func (d zonedDateTime) Date() calendar.Date {
	return d.date
}

func (d zonedDateTime) Time() datetime.Time {
	return d.time
}

func (d zonedDateTime) InstantCandidates() []tokiope.Instant {
	tu := datetime.NewOffsetDateTime(d.Date(), d.Time(), datetime.OffsetMinutes(0)).Instant()
	tlo := datetime.MinOffsetMinutes.AddTo(tu)
	thi := datetime.MaxOffsetMinutes.AddTo(tu)
	z := d.Zone()
	ts := z.transitionsBetween(tlo, thi)
	if len(ts) == 0 {
		offset := z.FindOffset(tu)
		return []tokiope.Instant{datetime.NewOffsetDateTime(d.Date(), d.Time(), offset).Instant()}
	}

	var candidates []tokiope.Instant
	{
		t := ts[0]
		c := datetime.NewOffsetDateTime(d.Date(), d.Time(), t.OffsetMinutesBefore).Instant()
		if tlo.Cmp(c) <= 0 && c.Cmp(t.TransitionTimestamp) < 0 { // skip if c == t.TransitionTimestamp
			candidates = append(candidates, c)
		}
	}
	{
		t := ts[len(ts)-1]
		c := datetime.NewOffsetDateTime(d.Date(), d.Time(), t.OffsetMinutesAfter).Instant()
		if c.Between(t.TransitionTimestamp, thi) { // include even if c == thi
			candidates = append(candidates, c)
		}
	}
	for i := 1; i < len(ts); i++ {
		ta, tb := ts[i-1], ts[i]
		c := datetime.NewOffsetDateTime(d.Date(), d.Time(), ta.OffsetMinutesBefore).Instant()
		if ta.TransitionTimestamp.Cmp(c) <= 0 && c.Cmp(tb.TransitionTimestamp) < 0 { // skip if c == tb.TransitionTimestamp
			candidates = append(candidates, c)
		}
	}

	slices.SortFunc(candidates, tokiope.Instant.Cmp)
	uniqCandidates := []tokiope.Instant{}
	for _, c := range candidates {
		if len(uniqCandidates) == 0 || uniqCandidates[len(uniqCandidates)-1].Cmp(c) < 0 {
			uniqCandidates = append(uniqCandidates, c)
		}
	}

	return uniqCandidates
}
