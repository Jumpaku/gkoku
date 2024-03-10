package datetime

import (
	"fmt"
	"github.com/Jumpaku/gkoku/clock"
	"github.com/Jumpaku/go-tzot"
	"github.com/samber/lo"
	"sort"
	"time"
)

type Zone struct {
	id          string
	transitions []Transition
}

type Transition struct {
	When         clock.Instant
	OffsetBefore OffsetMinutes
	OffsetAfter  OffsetMinutes
}

func CreateZone(zoneID string, transitions []Transition) Zone {
	return Zone{
		id:          zoneID,
		transitions: transitions,
	}
}

func LoadZone(zoneID string) (Zone, error) {
	zone, found := tzot.GetZone(zoneID)
	if !found {
		return Zone{}, fmt.Errorf("timezone %q not found", zoneID)
	}

	ts := lo.Map(zone.Transitions, func(t tzot.Transition, _ int) Transition {
		return Transition{
			When:         clock.Unix(t.When.Unix(), 0),
			OffsetBefore: OffsetMinutes(t.OffsetBefore / time.Minute),
			OffsetAfter:  OffsetMinutes(t.OffsetAfter / time.Minute),
		}
	})

	return CreateZone(zone.ID, ts), nil
}

func FixedZone(zoneID string, offset OffsetMinutes) Zone {
	return CreateZone(zoneID, []Transition{
		{
			When:         clock.MinInstant,
			OffsetBefore: offset,
			OffsetAfter:  offset,
		},
	})
}

func (z Zone) FindOffset(at clock.Instant) OffsetMinutes {
	n := len(z.transitions)
	if n == 0 {
		return 0
	}

	idx := sort.Search(n, func(i int) bool {
		t := z.transitions[i]
		return t.When.Cmp(at) >= 0
	})
	if idx == 0 {
		t := z.transitions[0]
		if t.When.After(at) {
			return t.OffsetBefore
		}
		return t.OffsetAfter
	}

	return z.transitions[idx-1].OffsetAfter
}

func (z Zone) ID() string {
	return z.id
}
