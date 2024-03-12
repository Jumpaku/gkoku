package zone

import (
	"encoding/json"
	"fmt"
	"github.com/Jumpaku/gkoku/clock"
	"github.com/Jumpaku/go-assert"
	"github.com/Jumpaku/go-tzot"
	"github.com/samber/lo"
	"slices"
	"strings"
	"time"
)

type Provider interface {
	Version() string
	AvailableZoneIDs() []string
	Get(zoneID string) (zone Zone, found bool)
}

func DefaultProvider() Provider {
	return defaultProvider{}
}

type defaultProvider struct{}

var _ Provider = defaultProvider{}

func (p defaultProvider) Version() string {
	return tzot.GetTZVersion()
}

func (p defaultProvider) AvailableZoneIDs() []string {
	return tzot.AvailableZoneIDs()
}

func (p defaultProvider) Get(zoneID string) (zone Zone, found bool) {
	z, found := tzot.GetZone(zoneID)
	ts := lo.Map(z.Transitions, func(t tzot.Transition, _ int) Transition {
		return Transition{
			TransitionTimestamp: clock.Unix(t.When.Unix(), 0),
			OffsetMinutesBefore: OffsetMinutes(t.OffsetBefore / time.Minute),
			OffsetMinutesAfter:  OffsetMinutes(t.OffsetAfter / time.Minute),
		}
	})
	zone = Create(z.ID, ts)
	return
}

// LoadProvider parses a JSON and returns a timezone Provider.
// tzotJSONBytes must be a JSON that is an array of zone objects.
//
// A zone object has the following properties:
//   - zone: zone is a string which represents a timezone ID.
//   - transitions: transitions is an array of transition object.
//
// A transition object has the following properties:
//   - transition_timestamp: transition_timestamp is a string which represents a timestamp at which timezone offset transition occurred, which is in form of yyyy-mm-ddThh:mm:ssZ.
//   - offset_seconds_before: offset_seconds_before is a number which represents a UTC offset in seconds that is effective before the transition.
//   - offset_seconds_after: offset_seconds_after is a number which represents a UTC offset in seconds that is effective after the transition.
//
// An example is as shown below:
//
//		[
//	   		{
//	       		"zone": "Asia/Tokyo",
//	       		"transitions": [
//	           		{
//	               		"transition_timestamp": "1887-12-31T15:00:00Z",
//	               		"offset_seconds_before": 33539,
//	           		    "offset_seconds_after": 32400
//	           		},
//	           		{
//	           		    "transition_timestamp": "1948-05-01T15:00:00Z",
//	           		    "offset_seconds_before": 32400,
//	           		    "offset_seconds_after": 36000
//	           		},
//	           		{
//	           		    "transition_timestamp": "1948-09-11T15:00:00Z",
//	           		    "offset_seconds_before": 36000,
//	           		    "offset_seconds_after": 32400
//	           		},
//	           		{
//	           		    "transition_timestamp": "1949-04-02T15:00:00Z",
//	           		    "offset_seconds_before": 32400,
//	           		    "offset_seconds_after": 36000
//	           		},
//	           		{
//	           		    "transition_timestamp": "1949-09-10T15:00:00Z",
//	           		    "offset_seconds_before": 36000,
//	           		    "offset_seconds_after": 32400
//	           		},
//	           		{
//	           		    "transition_timestamp": "1950-05-06T15:00:00Z",
//	           		    "offset_seconds_before": 32400,
//	           		    "offset_seconds_after": 36000
//	           		},
//	           		{
//	           		    "transition_timestamp": "1950-09-09T15:00:00Z",
//	           		    "offset_seconds_before": 36000,
//	           		    "offset_seconds_after": 32400
//	           		},
//	           		{
//	           		    "transition_timestamp": "1951-05-05T15:00:00Z",
//	           		    "offset_seconds_before": 32400,
//	           		    "offset_seconds_after": 36000
//	           		},
//	           		{
//	           		    "transition_timestamp": "1951-09-08T15:00:00Z",
//	           		    "offset_seconds_before": 36000,
//	           		    "offset_seconds_after": 32400
//	           		}
//	       		]
//	   		}
//		]
func LoadProvider(tzotJSONBytes []byte, version string) (Provider, error) {
	var zoneJSONSlice []tzot.ZoneJSON
	if err := json.Unmarshal(tzotJSONBytes, &zoneJSONSlice); err != nil {
		return nil, fmt.Errorf("failed to unmarshal tzot JSON: %w", err)
	}

	zoneIDs := []string{}
	zoneMap := map[string]Zone{}
	for _, zj := range zoneJSONSlice {
		zoneID := zj.Zone
		transitions := []Transition{}
		for _, tj := range zj.Transitions {
			t, err := time.Parse(time.RFC3339Nano, tj.TransitionTimestamp)
			if err != nil {
				return nil, fmt.Errorf("failed to parse transition timestamp of zone %q: %w", zoneID, err)
			}
			transitions = append(transitions, Transition{
				TransitionTimestamp: clock.Unix(t.Unix(), 0),
				OffsetMinutesBefore: OffsetMinutes(tj.OffsetSecondsBefore / 60),
				OffsetMinutesAfter:  OffsetMinutes(tj.OffsetSecondsAfter / 60),
			})
		}
		slices.SortFunc(transitions, func(a, b Transition) int { return a.TransitionTimestamp.Cmp(b.TransitionTimestamp) })

		if err := validateTransitions(transitions); err != nil {
			return nil, fmt.Errorf("failed to validate transitions: %w", err)
		}

		zoneIDs = append(zoneIDs, zoneID)
		zoneMap[zoneID] = Zone{id: zoneID, transitions: transitions}
	}

	slices.Sort(zoneIDs)

	return provider{version: version, zoneIDs: zoneIDs, zoneMap: zoneMap}, nil
}

// CreateProvider creates a timezone Provider from given zones and version.
// Each zone ID in the given zones must be unique.
// For each zone of the zones, the transitions of the zone must be sorted in ascending order with respect to TransitionTimestamp field.
// OffsetMinutesAfter field of each transition of the transitions must match OffsetMinutesBefore field of the successor transition.
func CreateProvider(zones []Zone, version string) Provider {
	p := provider{version: version, zoneIDs: []string{}, zoneMap: map[string]Zone{}}
	slices.SortFunc(zones, func(a, b Zone) int { return strings.Compare(a.id, b.id) })
	for i, z := range zones {
		err := validateTransitions(z.transitions)
		assert.Params(err == nil, "invalid transitions: %+v", err)
		if i > 0 {
			assert.Params(p.zoneIDs[len(p.zoneIDs)-1] != z.id, "zone ID must be unique in zones")
		}
		p.zoneIDs = append(p.zoneIDs, z.id)
		p.zoneMap[z.id] = Zone{id: z.id, transitions: append([]Transition{}, z.transitions...)}
	}

	return p
}

type provider struct {
	version string
	zoneIDs []string
	zoneMap map[string]Zone
}

var _ Provider = provider{}

func (p provider) Version() string {
	return p.version
}

func (p provider) AvailableZoneIDs() []string {
	return p.zoneIDs
}

func (p provider) Get(zoneID string) (zone Zone, found bool) {
	zone, found = p.zoneMap[zoneID]
	return
}
