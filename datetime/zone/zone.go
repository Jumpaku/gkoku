package zone

import (
	"fmt"
	"github.com/Jumpaku/gkoku/clock"
	"github.com/Jumpaku/go-assert"
	"sort"
)

type Zone struct {
	id          string
	transitions []Transition
}

type Transition struct {
	TransitionTimestamp clock.Instant
	OffsetMinutesBefore OffsetMinutes
	OffsetMinutesAfter  OffsetMinutes
}

// Create creates a Zone with the given zone ID and transitions.
// The transitions must be sorted in ascending order with respect to TransitionTimestamp field.
// OffsetMinutesAfter field of each transition of the transitions must match OffsetMinutesBefore field of the successor transition.
func Create(zoneID string, transitions []Transition) Zone {
	err := validateTransitions(transitions)
	assert.Params(err == nil, `invalid transitions: %+v`, err)
	return Zone{
		id:          zoneID,
		transitions: transitions,
	}
}

func CreateFixed(zoneID string, offset OffsetMinutes) Zone {
	return Create(zoneID, []Transition{
		{
			TransitionTimestamp: clock.MinInstant,
			OffsetMinutesBefore: offset,
			OffsetMinutesAfter:  offset,
		},
	})
}

func (z Zone) FindOffset(at clock.Instant) OffsetMinutes {
	ts := z.transitions
	n := len(ts)
	if n == 0 {
		return 0
	}

	if t := ts[n-1]; t.TransitionTimestamp.Cmp(at) <= 0 {
		return t.OffsetMinutesAfter
	}

	idx := sort.Search(n, func(i int) bool {
		t := z.transitions[i]
		return at.Cmp(t.TransitionTimestamp) < 0
	})

	return z.transitions[idx].OffsetMinutesBefore
}

func (z Zone) TransitionsBetween(beginAt, endAt clock.Instant) []Transition {
	ts := z.transitions
	n := len(ts)
	if n == 0 || beginAt.After(endAt) {
		return []Transition{}
	}

	var narrowed []Transition
	idx := sort.Search(n, func(i int) bool {
		t := ts[i]
		return beginAt.Cmp(t.TransitionTimestamp) <= 0
	})
	for ; idx < n; idx++ {
		t := ts[idx]
		if t.TransitionTimestamp.After(endAt) {
			break
		}

		narrowed = append(narrowed, t)
	}

	return narrowed
}

func (z Zone) ID() string {
	return z.id
}

func validateTransitions(transitions []Transition) error {
	for i := 0; i < len(transitions)-1; i++ {
		ti, tj := transitions[i], transitions[i+1]
		if !ti.TransitionTimestamp.Before(tj.TransitionTimestamp) {
			return fmt.Errorf(`transitions must be sorted in ascending order with respect to TransitionTimestamp field: transitions[%d].TransitionTimestamp = %v, transitions[%d+1].TransitionTimestamp = %v`, i, ti.TransitionTimestamp, i+1, tj.TransitionTimestamp)
		}
		if ti.OffsetMinutesAfter != tj.OffsetMinutesBefore {
			return fmt.Errorf(`OffsetMinutesAfter field of a transition must match OffsetMinutesBefore field of a successor transition: transitions[%d].OffsetMinutesAfter = %v, transitions[%d+1].OffsetMinutesBefore = %v`, i, ti.OffsetMinutesAfter, i+1, tj.OffsetMinutesBefore)
		}
	}
	return nil
}
