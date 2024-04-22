package zone_test

import (
	"fmt"
	"github.com/Jumpaku/gkoku"
	"github.com/Jumpaku/gkoku/datetime"
	zone2 "github.com/Jumpaku/gkoku/datetime/zone"
	"github.com/Jumpaku/gkoku/internal/tests/assert"
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadProvider(t *testing.T) {
	tests := []struct {
		inJSON    string
		inVersion string
		want      zone2.Provider
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
			want:      zone2.CreateProvider(nil, `empty`),
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
			want: zone2.CreateProvider([]zone2.Zone{
				zone2.Create("ExampleZone", []zone2.Transition{
					{
						TransitionTimestamp: gkoku.Unix(859683600, 0),
						OffsetMinutesBefore: datetime.OffsetMinutes(60),
						OffsetMinutesAfter:  datetime.OffsetMinutes(120),
					}, {
						TransitionTimestamp: gkoku.Unix(877827600, 0),
						OffsetMinutesBefore: datetime.OffsetMinutes(120),
						OffsetMinutesAfter:  datetime.OffsetMinutes(60),
					},
				}, []zone2.Rule{
					zone2.NewRule(zone2.RuleArg{
						OffsetMinutesBefore: 60,
						OffsetMinutesAfter:  120,
						Month:               3,
						BaseDay:             25,
						DayOfWeek:           7,
						TimeOfDay:           datetime.TimeOf(1, 0, 0, 0),
						TimeOffsetMinutes:   0,
					}),
					zone2.NewRule(zone2.RuleArg{
						OffsetMinutesBefore: 120,
						OffsetMinutesAfter:  60,
						Month:               10,
						BaseDay:             25,
						DayOfWeek:           7,
						TimeOfDay:           datetime.TimeOf(1, 0, 0, 0),
						TimeOffsetMinutes:   0,
					}),
				}),
			}, `example`),
		},
	}

	for number, tt := range tests {
		t.Run(fmt.Sprintf(`%d`, number), func(t *testing.T) {
			got, err := zone2.LoadProvider([]byte(tt.inJSON), tt.inVersion)
			if tt.wantErr {
				assert2.NotNil(t, err)
			} else {
				assert.Equal(t, tt.want.Version(), got.Version())
				assert2.ElementsMatch(t, tt.want.AvailableZoneIDs(), got.AvailableZoneIDs())
				for _, zoneID := range tt.want.AvailableZoneIDs() {
					want, _ := tt.want.Get(zoneID)
					got, _ := got.Get(zoneID)
					zone2.EqualZone(t, want, got)
				}
			}
		})
	}
}

func TestProvider_AvailableZoneIDs(t *testing.T) {
	tests := []struct {
		sut  zone2.Provider
		want []string
	}{
		{
			sut:  zone2.CreateProvider(nil, `nil`),
			want: []string{},
		},
		{
			sut:  zone2.CreateProvider([]zone2.Zone{}, `empty`),
			want: []string{},
		},
		{
			sut: zone2.CreateProvider([]zone2.Zone{
				zone2.Create(`A`, nil, nil),
				zone2.Create(`C`, nil, nil),
				zone2.Create(`B`, nil, nil),
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
	exampleProvider := zone2.CreateProvider([]zone2.Zone{
		zone2.Create(`C`, []zone2.Transition{}, nil),
		zone2.Create(`B`, nil, nil),
		zone2.Create(`A`, []zone2.Transition{
			{TransitionTimestamp: gkoku.Unix(10, 0), OffsetMinutesBefore: 1, OffsetMinutesAfter: 2},
			{TransitionTimestamp: gkoku.Unix(20, 0), OffsetMinutesBefore: 2, OffsetMinutesAfter: 3},
			{TransitionTimestamp: gkoku.Unix(30, 0), OffsetMinutesBefore: 3, OffsetMinutesAfter: 4},
		}, nil),
	}, `example`)

	tests := []struct {
		sut           zone2.Provider
		zoneID        string
		wantZoneFound bool
	}{
		{
			sut:           zone2.CreateProvider(nil, ``),
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
		sut  zone2.Provider
		want string
	}{
		{
			sut:  zone2.CreateProvider(nil, ``),
			want: ``,
		},
		{
			sut: zone2.CreateProvider([]zone2.Zone{
				zone2.Create(`A`, nil, nil),
				zone2.Create(`C`, nil, nil),
				zone2.Create(`B`, nil, nil),
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
