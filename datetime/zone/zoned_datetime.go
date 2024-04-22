package zone

import (
	"fmt"
	"github.com/Jumpaku/gkoku"
	"github.com/Jumpaku/gkoku/date"
	"github.com/Jumpaku/gkoku/datetime"
	"slices"
)

type ZonedDateTime interface {
	Date() date.Date
	Time() datetime.Time
	Zone() Zone
	String() string
	InstantCandidates() []gkoku.Instant
}

func NewZonedDateTime(date date.Date, time datetime.Time, zone Zone) ZonedDateTime {
	return zonedDateTime{
		date: date,
		time: time,
		zone: zone,
	}
}

type zonedDateTime struct {
	date date.Date
	time datetime.Time
	zone Zone
}

func (d zonedDateTime) String() string {
	return fmt.Sprintf(`%s%s [%s]`, d.Date().String(), d.Time().String(), d.Zone().ID())
}

func (d zonedDateTime) Zone() Zone {
	return d.zone
}

func (d zonedDateTime) Date() date.Date {
	return d.date
}

func (d zonedDateTime) Time() datetime.Time {
	return d.time
}

func (d zonedDateTime) InstantCandidates() []gkoku.Instant {
	tu := datetime.NewOffsetDateTime(d.Date(), d.Time(), datetime.OffsetMinutes(0)).Instant()
	tlo := datetime.MinOffsetMinutes.AddTo(tu)
	thi := datetime.MaxOffsetMinutes.AddTo(tu)
	z := d.Zone()
	ts := z.transitionsBetween(tlo, thi)
	if len(ts) == 0 {
		offset := z.FindOffset(tu)
		return []gkoku.Instant{datetime.NewOffsetDateTime(d.Date(), d.Time(), offset).Instant()}
	}

	var candidates []gkoku.Instant
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

	slices.SortFunc(candidates, gkoku.Instant.Cmp)
	uniqCandidates := []gkoku.Instant{}
	for _, c := range candidates {
		if len(uniqCandidates) == 0 || uniqCandidates[len(uniqCandidates)-1].Cmp(c) < 0 {
			uniqCandidates = append(uniqCandidates, c)
		}
	}

	return uniqCandidates
}
