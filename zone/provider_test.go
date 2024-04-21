package zone_test

import (
	"fmt"
	"github.com/Jumpaku/gkoku/clock"
	"github.com/Jumpaku/gkoku/internal/tests/assert"
	"github.com/Jumpaku/gkoku/zone"
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadProvider(t *testing.T) {
	tests := []struct {
		inJSON    string
		inVersion string
		want      zone.Provider
		wantErr   bool
	}{
		{
			inJSON:  ``,
			wantErr: true,
		},
		{
			inJSON: `[
	{
        "zone": "Asia/Aden",
        "transitions": [
            {
                "transitionTimestamp": "invalid",
                "offsetSecondsBefore": 11212,
                "offsetSecondsAfter": 10800
            }
        ]
    }
]`,
			wantErr: true,
		},
		{
			inJSON: `[
	{
        "zone": "Asia/Aden",
         "rules": [
            {
                "offsetSecondsBefore": 3600,
                "offsetSecondsAfter": 7200,
                "month": 3,
                "baseDay": 25,
                "dayOfWeek": 7,
                "offsetTime": "invalid"
            }
        ]
    }
]`,
			wantErr: true,
		},
		{
			inJSON:    `[]`,
			inVersion: `empty`,
			want:      zone.CreateProvider(nil, `empty`),
		},
		{
			inJSON: `[
	{
        "id": "ExampleZone",
        "transitions": [
            {
                "transitionTimestamp": "1997-03-30T01:00:00Z",
                "offsetSecondsBefore": 3600,
                "offsetSecondsAfter": 7200
            },
            {
                "transitionTimestamp": "1997-10-26T01:00:00Z",
                "offsetSecondsBefore": 7200,
                "offsetSecondsAfter": 3600
            }
        ],
        "rules": [
            {
                "offsetSecondsBefore": 3600,
                "offsetSecondsAfter": 7200,
                "month": 3,
                "baseDay": 25,
                "dayOfWeek": 7,
                "offsetTime": "01:00:00Z"
            },
            {
                "offsetSecondsBefore": 7200,
                "offsetSecondsAfter": 3600,
                "month": 10,
                "baseDay": 25,
                "dayOfWeek": 7,
                "offsetTime": "01:00:00Z"
            }
        ]
    }
]`,
			inVersion: `example`,
			want: zone.CreateProvider([]zone.Zone{
				zone.Create("ExampleZone", []zone.Transition{
					{
						TransitionTimestamp: clock.Unix(859683600, 0),
						OffsetMinutesBefore: zone.OffsetMinutes(60),
						OffsetMinutesAfter:  zone.OffsetMinutes(120),
					}, {
						TransitionTimestamp: clock.Unix(877827600, 0),
						OffsetMinutesBefore: zone.OffsetMinutes(120),
						OffsetMinutesAfter:  zone.OffsetMinutes(60),
					},
				}, []zone.Rule{
					zone.NewRule(zone.RuleArg{
						OffsetMinutesBefore: 60,
						OffsetMinutesAfter:  120,
						Month:               3,
						BaseDay:             25,
						DayOfWeek:           7,
						SecondOfDay:         3600,
						TimeOffsetMinutes:   0,
					}),
					zone.NewRule(zone.RuleArg{
						OffsetMinutesBefore: 120,
						OffsetMinutesAfter:  60,
						Month:               10,
						BaseDay:             25,
						DayOfWeek:           7,
						SecondOfDay:         3600,
						TimeOffsetMinutes:   0,
					}),
				}),
			}, `example`),
		},
	}

	for number, tt := range tests {
		t.Run(fmt.Sprintf(`%d`, number), func(t *testing.T) {
			got, err := zone.LoadProvider([]byte(tt.inJSON), tt.inVersion)
			if tt.wantErr {
				assert2.NotNil(t, err)
			} else {
				assert.Equal(t, tt.want.Version(), got.Version())
				assert2.ElementsMatch(t, tt.want.AvailableZoneIDs(), got.AvailableZoneIDs())
				for _, zoneID := range tt.want.AvailableZoneIDs() {
					want, _ := tt.want.Get(zoneID)
					got, _ := got.Get(zoneID)
					zone.EqualZone(t, want, got)
				}
			}
		})
	}
}

func TestProvider_AvailableZoneIDs(t *testing.T) {
	tests := []struct {
		sut  zone.Provider
		want []string
	}{
		{
			sut:  zone.CreateProvider(nil, `nil`),
			want: []string{},
		},
		{
			sut:  zone.CreateProvider([]zone.Zone{}, `empty`),
			want: []string{},
		},
		{
			sut: zone.CreateProvider([]zone.Zone{
				zone.Create(`A`, nil, nil),
				zone.Create(`C`, nil, nil),
				zone.Create(`B`, nil, nil),
			}, `abc`),
			want: []string{"A", "B", "C"},
		},
	}

	for number, tt := range tests {
		t.Run(fmt.Sprintf(`%d`, number), func(t *testing.T) {
			got := tt.sut.AvailableZoneIDs()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestProvider_Get(t *testing.T) {
	exampleProvider := zone.CreateProvider([]zone.Zone{
		zone.Create(`C`, []zone.Transition{}, nil),
		zone.Create(`B`, nil, nil),
		zone.Create(`A`, []zone.Transition{
			{TransitionTimestamp: clock.Unix(10, 0), OffsetMinutesBefore: 1, OffsetMinutesAfter: 2},
			{TransitionTimestamp: clock.Unix(20, 0), OffsetMinutesBefore: 2, OffsetMinutesAfter: 3},
			{TransitionTimestamp: clock.Unix(30, 0), OffsetMinutesBefore: 3, OffsetMinutesAfter: 4},
		}, nil),
	}, `example`)

	tests := []struct {
		sut           zone.Provider
		zoneID        string
		wantZoneFound bool
	}{
		{
			sut:           zone.CreateProvider(nil, ``),
			zoneID:        "X",
			wantZoneFound: false,
		},
		{
			sut:           exampleProvider,
			zoneID:        "X",
			wantZoneFound: false,
		},
		{
			sut:           exampleProvider,
			zoneID:        "C",
			wantZoneFound: true,
		},
		{
			sut:           exampleProvider,
			zoneID:        "B",
			wantZoneFound: true,
		},
		{
			sut:           exampleProvider,
			zoneID:        "A",
			wantZoneFound: true,
		},
	}

	for number, tt := range tests {
		t.Run(fmt.Sprintf(`%d`, number), func(t *testing.T) {
			got, ok := tt.sut.Get(tt.zoneID)
			assert.Equal(t, tt.wantZoneFound, ok)
			if tt.wantZoneFound {
				assert.Equal(t, tt.zoneID, got.ID())
			}
		})
	}
}

func TestProvider_Version(t *testing.T) {
	tests := []struct {
		sut  zone.Provider
		want string
	}{
		{
			sut:  zone.CreateProvider(nil, ``),
			want: ``,
		},
		{
			sut: zone.CreateProvider([]zone.Zone{
				zone.Create(`A`, nil, nil),
				zone.Create(`C`, nil, nil),
				zone.Create(`B`, nil, nil),
			}, `abc`),
			want: `abc`,
		},
	}

	for number, tt := range tests {
		t.Run(fmt.Sprintf(`%d`, number), func(t *testing.T) {
			got := tt.sut.Version()
			assert.Equal(t, tt.want, got)
		})
	}
}
