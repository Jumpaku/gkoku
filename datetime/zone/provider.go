package zone

import (
	"encoding/json"
	"fmt"
	"github.com/Jumpaku/gkoku"
	"github.com/Jumpaku/gkoku/datetime"
	"github.com/Jumpaku/go-assert"
	"slices"
	"strconv"
	"strings"
	"time"
)

// LoadProvider parses a JSON and returns a timezone Provider.
// tzotJSONBytes must be a JSON that is an array of zone objects.
func LoadProvider(tzotJSONBytes []byte, version string) (Provider, error) {
	var tzotJSON tzotJSON
	if err := json.Unmarshal(tzotJSONBytes, &tzotJSON); err != nil {
		return Provider{}, fmt.Errorf("failed to unmarshal tzot JSON: %w", err)
	}

	zoneIDs := []string{}
	zoneMap := map[string]Zone{}
	for _, zj := range tzotJSON {
		zoneID := zj.ID

		transitions := []Transition{}
		for _, tj := range zj.Transitions {
			t, err := time.Parse(time.RFC3339Nano, tj.TransitionTimestamp)
			if err != nil {
				return Provider{}, fmt.Errorf("failed to parse transition timestamp of zone %q: %w", zoneID, err)
			}

			transitions = append(transitions, Transition{
				TransitionTimestamp: gkoku.Unix(t.Unix(), 0),
				OffsetMinutesBefore: datetime.OffsetMinutes(tj.OffsetSecondsBefore / 60),
				OffsetMinutesAfter:  datetime.OffsetMinutes(tj.OffsetSecondsAfter / 60),
			})
		}
		slices.SortFunc(transitions, func(a, b Transition) int { return a.TransitionTimestamp.Cmp(b.TransitionTimestamp) })
		if err := validateTransitions(transitions); err != nil {
			return Provider{}, fmt.Errorf("failed to validate transitions: %w", err)
		}

		rules := []Rule{}
		for _, rj := range zj.Rules {
			sod, om, err := parseOffsetTime(rj.OffsetTime)
			if err != nil {
				return Provider{}, fmt.Errorf("failed to parse transition timestamp of zone %q: %w", zoneID, err)
			}

			rules = append(rules, rule{
				OffsetMinutesBefore: datetime.OffsetMinutes(rj.OffsetSecondsBefore / 60),
				OffsetMinutesAfter:  datetime.OffsetMinutes(rj.OffsetSecondsAfter / 60),
				Month:               rj.Month,
				BaseDay:             rj.BaseDay,
				DayOfWeek:           rj.DayOfWeek,
				SecondOfDay:         sod,
				TimeOffsetMinutes:   datetime.OffsetMinutes(om),
			})
		}

		zoneIDs = append(zoneIDs, zoneID)
		zoneMap[zoneID] = Create(zoneID, transitions, rules)
	}

	slices.Sort(zoneIDs)

	return Provider{version: version, zoneIDs: zoneIDs, zoneMap: zoneMap}, nil
}

func parseOffsetTime(s string) (secondsOfDay int, offsetMinutes int, err error) {
	s = strings.ReplaceAll(s, "Z", "+00:00")
	if len(s) != 14 {
		return 0, 0, fmt.Errorf("failed to parse offset time: invalid format")
	}
	{
		hh, err := strconv.Atoi(s[0:2])
		if err != nil {
			return 0, 0, fmt.Errorf("failed to parse offset time: invalid hour format")
		}

		mm, err := strconv.Atoi(s[3:5])
		if err != nil {
			return 0, 0, fmt.Errorf("failed to parse offset time: invalid minute format")
		}

		ss, err := strconv.Atoi(s[6:8])
		if err != nil {
			return 0, 0, fmt.Errorf("failed to parse offset time: invalid second format")
		}

		secondsOfDay = hh*3600 + mm*60 + ss
	}
	{
		sign := 1
		if s[8] == '-' {
			sign = -1
		}

		hh, err := strconv.Atoi(s[9:11])
		if err != nil {
			return 0, 0, fmt.Errorf("failed to parse offset time: invalid offset hour format")
		}

		mm, err := strconv.Atoi(s[12:14])
		if err != nil {
			return 0, 0, fmt.Errorf("failed to parse offset time: invalid offset minute format")
		}

		offsetMinutes = sign * (hh*60 + mm)
	}

	return secondsOfDay, offsetMinutes, nil
}

// CreateProvider creates a timezone Provider from given zones and version.
// Each zone ID in the given zones must be unique.
// For each zone of the zones, the transitions of the zone must be sorted in ascending order with respect to TransitionTimestamp field.
// OffsetMinutesAfter field of each transition of the transitions must match OffsetMinutesBefore field of the successor transition.
func CreateProvider(zones []Zone, version string) Provider {
	p := Provider{version: version, zoneIDs: []string{}, zoneMap: map[string]Zone{}}
	slices.SortFunc(zones, func(a, b Zone) int { return strings.Compare(a.id, b.id) })
	for i, z := range zones {
		err := validateTransitions(z.transitions)
		assert.Params(err == nil, "invalid transitions: %+v", err)
		if i > 0 {
			assert.Params(p.zoneIDs[len(p.zoneIDs)-1] != z.id, "zone ID must be unique in zones")
		}
		p.zoneIDs = append(p.zoneIDs, z.id)
		p.zoneMap[z.id] = Zone{id: z.id, transitions: append([]Transition{}, z.transitions...), rules: append([]Rule{}, z.rules...)}
	}

	return p
}

type Provider struct {
	version string
	zoneIDs []string
	zoneMap map[string]Zone
}

func (p Provider) Version() string {
	return p.version
}

func (p Provider) AvailableZoneIDs() []string {
	return p.zoneIDs
}

func (p Provider) Get(zoneID string) (zone Zone, found bool) {
	zone, found = p.zoneMap[zoneID]
	return
}
