package zone

import (
	"github.com/Jumpaku/gkoku/date"
	"github.com/Jumpaku/gkoku/date/iter"
	"github.com/Jumpaku/gkoku/datetime"
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
	Month               date.Month
	BaseDay             int
	DayOfWeek           date.DayOfWeek
	TimeOfDay           datetime.Time
	TimeOffsetMinutes   datetime.OffsetMinutes
}

func (r rule) doNotImplement(doNotImplement) {}

func (r rule) Transition(year int) Transition {
	dateIter := iter.OfDate(date.YyyyMmDd(year, r.Month, 1))
	dateIter.Move(r.BaseDay - 1)

	_, _, dow := dateIter.Get().YyyyWwD()
	addDays := int((r.DayOfWeek - dow + 7) % 7)
	dateIter.Move(addDays)

	transitionDateTime := datetime.NewOffsetDateTime(dateIter.Get(), r.TimeOfDay, r.TimeOffsetMinutes)

	return Transition{
		TransitionTimestamp: transitionDateTime.Instant(),
		OffsetMinutesBefore: r.OffsetMinutesBefore,
		OffsetMinutesAfter:  r.OffsetMinutesAfter,
	}
}
