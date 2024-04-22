package zone

import (
	"fmt"
	"github.com/Jumpaku/gkoku"
	"github.com/Jumpaku/gkoku/datetime"
	"github.com/Jumpaku/go-assert"
	"github.com/samber/lo"
	"slices"
	"sort"
)

type Zone struct {
	id          string
	transitions []Transition
	rules       []Rule
}

type Transition struct {
	TransitionTimestamp gkoku.Instant
	OffsetMinutesBefore datetime.OffsetMinutes
	OffsetMinutesAfter  datetime.OffsetMinutes
}

// Create creates a Zone with the given zone ID and transitions.
// The transitions must be sorted in ascending order with respect to TransitionTimestamp field.
// OffsetMinutesAfter field of each transition of the transitions must match OffsetMinutesBefore field of the successor transition.
func Create(zoneID string, transitions []Transition, rules []Rule) Zone {
	err := validateTransitions(transitions)
	assert.Params(err == nil, `invalid transitions: %+v`, err)
	return Zone{
		id:          zoneID,
		transitions: append([]Transition{}, transitions...),
		rules:       append([]Rule{}, rules...),
	}
}

func CreateFixed(zoneID string, offset datetime.OffsetMinutes) Zone {
	return Create(zoneID, []Transition{
		{
			TransitionTimestamp: gkoku.MinInstant,
			OffsetMinutesBefore: offset,
			OffsetMinutesAfter:  offset,
		},
	}, nil)
}

func (z Zone) FindOffset(at gkoku.Instant) datetime.OffsetMinutes {
	ts, rs := z.transitions, z.rules
	if len(ts) == 0 && len(rs) == 0 {
		return 0
	}

	if n := len(ts); n > 0 {
		if at.Cmp(ts[n-1].TransitionTimestamp) < 0 || len(rs) == 0 {
			return findOffset(ts, at)
		}
	}

	minYear, _, _ := datetime.FromInstant(at, datetime.MinOffsetMinutes).Date().YyyyMmDd()
	maxYear, _, _ := datetime.FromInstant(at, datetime.MaxOffsetMinutes).Date().YyyyMmDd()
	rts := collectRuledTransitions(z.rules, minYear, maxYear)
	rts = lo.Filter(rts, func(t Transition, _ int) bool {
		n := len(ts)
		return n == 0 || t.TransitionTimestamp.Cmp(ts[n-1].TransitionTimestamp) > 0
	})

	return findOffset(rts, at)
}

func findOffset(transitions []Transition, at gkoku.Instant) datetime.OffsetMinutes {
	n := len(transitions)
	if n == 0 {
		return 0
	}

	if t := transitions[n-1]; t.TransitionTimestamp.Cmp(at) <= 0 {
		return t.OffsetMinutesAfter
	}

	idx := sort.Search(n, func(i int) bool {
		return at.Cmp(transitions[i].TransitionTimestamp) < 0
	})

	return transitions[idx].OffsetMinutesBefore
}

func collectRuledTransitions(rules []Rule, minYear, maxYear int) []Transition {
	ts := []Transition{}
	for _, r := range rules {
		for year := minYear; year <= maxYear; year++ {
			ts = append(ts, r.Transition(year))
		}
	}

	slices.SortFunc(ts, func(a, b Transition) int { return a.TransitionTimestamp.Cmp(b.TransitionTimestamp) })

	return ts
}

func (z Zone) transitionsBetween(beginAt, endAt gkoku.Instant) []Transition {
	ts := append([]Transition{}, z.transitions...)
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
