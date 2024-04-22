package zone

import (
	"github.com/Jumpaku/gkoku/date"
	"github.com/Jumpaku/gkoku/datetime"
	"github.com/stretchr/testify/assert"
	"testing"
)

type RuleArg struct {
	OffsetMinutesBefore datetime.OffsetMinutes
	OffsetMinutesAfter  datetime.OffsetMinutes
	Month               date.Month
	BaseDay             int
	DayOfWeek           date.DayOfWeek
	TimeOfDay           datetime.Time
	TimeOffsetMinutes   datetime.OffsetMinutes
}

func NewRule(args RuleArg) Rule {
	return rule{
		OffsetMinutesBefore: args.OffsetMinutesBefore,
		OffsetMinutesAfter:  args.OffsetMinutesAfter,
		Month:               args.Month,
		BaseDay:             args.BaseDay,
		DayOfWeek:           args.DayOfWeek,
		TimeOfDay:           args.TimeOfDay,
		TimeOffsetMinutes:   args.TimeOffsetMinutes,
	}
}

func EqualRule(t *testing.T, want, got Rule) {
	t.Helper()
	{
		w, okW := want.(rule)
		g, okG := got.(rule)
		if okW && okG {
			assert.Equal(t, w, g)
			return
		}
	}
	t.Fail()
}

func EqualZone(t *testing.T, want, got Zone) {
	t.Helper()
	assert.Equal(t, want.id, got.id)
	assert.Equal(t, len(want.transitions), len(got.transitions), "the number of transitions are not equal")
	for i, w := range want.transitions {
		g := got.transitions[i]
		assert.Equal(t, w, g)
	}
	assert.Equalf(t, len(want.rules), len(got.rules), "the number of rules are not equal")
	for i, w := range want.rules {
		g := got.rules[i]
		EqualRule(t, w, g)
	}
}
