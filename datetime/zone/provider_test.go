package zone_test

import (
	"fmt"
	"github.com/Jumpaku/gkoku/clock"
	"github.com/Jumpaku/gkoku/datetime/zone"
	"github.com/Jumpaku/gkoku/internal/tests/assert"
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultProvider_Get(t *testing.T) {
	sut := zone.DefaultProvider()
	tests := []struct {
		id       string
		wantFail bool
	}{
		{
			id: "Zulu",
		},
		{
			id: "Asia/Tokyo",
		},
		{
			id: "Pacific/Apia",
		},
		{
			id: "Europe/Zurich",
		},
		{
			id:       "not-found",
			wantFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			got, found := sut.Get(tt.id)
			if tt.wantFail {
				if found {
					t.Errorf("Get() found = %v, want %v", found, tt.wantFail)
				}
			} else {
				assert.Equal(t, tt.id, got.ID())
			}
		})
	}
}

func TestDefaultProvider_AvailableZoneIDs(t *testing.T) {
	sut := zone.DefaultProvider()
	tests := []struct {
		want     string
		wantFail bool
	}{
		{
			want: "Zulu",
		},
		{
			want: "Asia/Tokyo",
		},
		{
			want: "Pacific/Apia",
		},
		{
			want: "Europe/Zurich",
		},
		{
			want:     "not-found",
			wantFail: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			got := sut.AvailableZoneIDs()
			if tt.wantFail {
				assert2.NotContains(t, got, tt.want)
			} else {
				assert2.Contains(t, got, tt.want)
			}
		})
	}
}

func TestDefaultProvider_Version(t *testing.T) {
	sut := zone.DefaultProvider()
	assert2.NotEmpty(t, sut.Version())
}

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
                "transition_timestamp": "invalid",
                "offset_seconds_before": 11212,
                "offset_seconds_after": 10800
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
        "zone": "ExampleZone",
        "transitions": [
            {
                "transition_timestamp": "1970-01-01T00:00:00Z",
                "offset_seconds_before": -60,
                "offset_seconds_after": 60
            }
        ]
    }
]`,
			inVersion: `example`,
			want: zone.CreateProvider([]zone.Zone{
				zone.Create("ExampleZone", []zone.Transition{{
					TransitionTimestamp: clock.Unix(0, 0),
					OffsetMinutesBefore: zone.OffsetMinutes(1),
					OffsetMinutesAfter:  zone.OffsetMinutes(-1),
				}}),
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
				assert.Equal(t, tt.want.AvailableZoneIDs(), got.AvailableZoneIDs())

				got, ok := got.Get("ExampleZone")
				want, wantOK := tt.want.Get("ExampleZone")
				if !wantOK {
					assert.Equal(t, false, ok)
				} else {
					assert.Equal(t, true, ok)
					assert.Equal(t, want.ID(), got.ID())
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
				zone.Create(`A`, nil),
				zone.Create(`C`, nil),
				zone.Create(`B`, nil),
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
		zone.Create(`C`, []zone.Transition{}),
		zone.Create(`B`, nil),
		zone.Create(`A`, []zone.Transition{
			{TransitionTimestamp: clock.Unix(10, 0), OffsetMinutesBefore: 1, OffsetMinutesAfter: 2},
			{TransitionTimestamp: clock.Unix(20, 0), OffsetMinutesBefore: 2, OffsetMinutesAfter: 3},
			{TransitionTimestamp: clock.Unix(30, 0), OffsetMinutesBefore: 3, OffsetMinutesAfter: 4},
		}),
	}, `example`)

	tests := []struct {
		sut              zone.Provider
		now              clock.Instant
		zoneID           string
		wantOffset       zone.OffsetMinutes
		wantZoneNotFound bool
	}{
		{
			sut:              zone.CreateProvider(nil, ``),
			zoneID:           "X",
			wantZoneNotFound: true,
		},
		{
			sut:              exampleProvider,
			now:              clock.Unix(5, 0),
			zoneID:           "X",
			wantZoneNotFound: true,
		},
		{
			sut:        exampleProvider,
			now:        clock.Unix(0, 0),
			zoneID:     "C",
			wantOffset: 0,
		},
		{
			sut:        exampleProvider,
			now:        clock.Unix(5, 0),
			zoneID:     "A",
			wantOffset: 1,
		},
		{
			sut:        exampleProvider,
			now:        clock.Unix(10, 0),
			zoneID:     "A",
			wantOffset: 2,
		},
		{
			sut:        exampleProvider,
			now:        clock.Unix(15, 0),
			zoneID:     "A",
			wantOffset: 2,
		},
		{
			sut:        exampleProvider,
			now:        clock.Unix(20, 0),
			zoneID:     "A",
			wantOffset: 3,
		},
		{
			sut:        exampleProvider,
			now:        clock.Unix(25, 0),
			zoneID:     "A",
			wantOffset: 3,
		},
		{
			sut:        exampleProvider,
			now:        clock.Unix(30, 0),
			zoneID:     "A",
			wantOffset: 4,
		},
		{
			sut:        exampleProvider,
			now:        clock.Unix(35, 0),
			zoneID:     "A",
			wantOffset: 4,
		},
	}

	for number, tt := range tests {
		t.Run(fmt.Sprintf(`%d`, number), func(t *testing.T) {
			got, ok := tt.sut.Get(tt.zoneID)
			if tt.wantZoneNotFound {
				assert.Equal(t, false, ok)
			} else {
				got := got.FindOffset(tt.now)
				assert.Equal(t, tt.wantOffset, got)
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
				zone.Create(`A`, nil),
				zone.Create(`C`, nil),
				zone.Create(`B`, nil),
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
