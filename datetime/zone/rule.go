package zone

import (
	"github.com/Jumpaku/tokiope/calendar"
	"github.com/Jumpaku/tokiope/calendar/iter"
	"github.com/Jumpaku/tokiope/datetime"
)

type doNotImplement interface {
	doNotImplement(doNotImplement)
}

// Rule defines when the yearly timezone offset transition occurs in the future.
type Rule interface {
	doNotImplement
	// Transition returns the transition of the timezone offset in the given year.
	Transition(year int) Transition
}

type rule struct {
	OffsetMinutesBefore datetime.OffsetMinutes
	OffsetMinutesAfter  datetime.OffsetMinutes
	Month               calendar.Month
	BaseDay             int
	DayOfWeek           calendar.DayOfWeek
	TimeOfDay           datetime.Time
	TimeOffsetMinutes   datetime.OffsetMinutes
}

func (r rule) doNotImplement(doNotImplement) {}

func (r rule) Transition(year int) Transition {
	dateIter := iter.OfDate(calendar.DateOfYMD(year, r.Month, 1))
	dateIter.Move(r.BaseDay - 1)

	_, _, dow := dateIter.Get().YWD()
	addDays := int((r.DayOfWeek - dow + 7) % 7)
	dateIter.Move(addDays)

	transitionDateTime := datetime.NewOffsetDateTime(dateIter.Get(), r.TimeOfDay, r.TimeOffsetMinutes)

	return Transition{
		TransitionTimestamp: transitionDateTime.Instant(),
		OffsetMinutesBefore: r.OffsetMinutesBefore,
		OffsetMinutesAfter:  r.OffsetMinutesAfter,
	}
}
