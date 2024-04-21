package zone

import (
	"github.com/Jumpaku/gkoku/clock"
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
	OffsetMinutesBefore OffsetMinutes
	OffsetMinutesAfter  OffsetMinutes
	Month               int
	BaseDay             int
	DayOfWeek           int
	SecondOfDay         int
	TimeOffsetMinutes   OffsetMinutes
}

func (r rule) doNotImplement(doNotImplement) {}

func (r rule) Transition(year int) Transition {
	loc := time.FixedZone("", int(time.Duration(r.TimeOffsetMinutes)*time.Minute/time.Second))

	t := time.Date(year, time.Month(r.Month), r.BaseDay, 0, 0, 0, 0, loc)
	addDays := (r.DayOfWeek - int(t.Weekday()) + 7) % 7
	t = t.AddDate(0, 0, addDays)
	t = t.Add(time.Duration(r.SecondOfDay) * time.Second)

	return Transition{
		TransitionTimestamp: clock.Unix(t.Unix(), 0),
		OffsetMinutesBefore: r.OffsetMinutesBefore,
		OffsetMinutesAfter:  r.OffsetMinutesAfter,
	}
}
