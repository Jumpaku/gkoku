package zone

import (
	"github.com/Jumpaku/gkoku"
	"github.com/Jumpaku/gkoku/datetime"
	"time"
)

type doNotImplement interface {
	doNotImplement(doNotImplement)
}

type Rule interface {
	doNotImplement
	Transition(year int) Transition
}

type rule struct {
	OffsetMinutesBefore datetime.OffsetMinutes
	OffsetMinutesAfter  datetime.OffsetMinutes
	Month               int
	BaseDay             int
	DayOfWeek           int
	SecondOfDay         int
	TimeOffsetMinutes   datetime.OffsetMinutes
}

func (r rule) doNotImplement(doNotImplement) {}

func (r rule) Transition(year int) Transition {
	loc := time.FixedZone("", int(time.Duration(r.TimeOffsetMinutes)*time.Minute/time.Second))

	t := time.Date(year, time.Month(r.Month), r.BaseDay, 0, 0, 0, 0, loc)
	addDays := (r.DayOfWeek - int(t.Weekday()) + 7) % 7
	t = t.AddDate(0, 0, addDays)
	t = t.Add(time.Duration(r.SecondOfDay) * time.Second)

	return Transition{
		TransitionTimestamp: gkoku.Unix(t.Unix(), 0),
		OffsetMinutesBefore: r.OffsetMinutesBefore,
		OffsetMinutesAfter:  r.OffsetMinutesAfter,
	}
}
