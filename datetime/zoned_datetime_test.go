package datetime_test

import (
	"github.com/Jumpaku/gkoku/calendar"
	"github.com/Jumpaku/gkoku/clock"
	. "github.com/Jumpaku/gkoku/datetime"
	"github.com/Jumpaku/gkoku/internal/tests/assert"
	zone2 "github.com/Jumpaku/gkoku/zone"
	"testing"
)

func TestZonedDateTime_InstantCandidates(t *testing.T) {
	instant := func(s string) clock.Instant {
		d, _ := ParseOffsetDateTime(s)
		return d.Instant()
	}
	date := func(s string) calendar.Date {
		d, _ := calendar.ParseDate(s, calendar.DateFormatYyyyMmDd)
		return d
	}
	time := func(s string) Time {
		t, _ := ParseTime(s)
		return t
	}
	getZone := func(zoneID string) zone2.Zone {
		//z, _ := zone2.DefaultProvider().Get(zoneID)
		return zone2.Zone{}
	}

	tests := []struct {
		sut  ZonedDateTime
		want []clock.Instant
	}{
		{
			sut:  NewZonedDateTime(date(`2024-03-13`), time(`00:00:00`), getZone(`Zulu`)),
			want: []clock.Instant{instant(`2024-03-13T00:00:00Z`)},
		},
		{
			sut:  NewZonedDateTime(date(`2011-09-01`), time(`12:00:00`), getZone(`Pacific/Apia`)),
			want: []clock.Instant{instant(`2011-09-01T12:00:00-11:00`)},
		},
		{
			sut:  NewZonedDateTime(date(`2011-12-29`), time(`12:00:00`), getZone(`Pacific/Apia`)),
			want: []clock.Instant{instant(`2011-12-29T12:00:00-10:00`)},
		},
		{
			sut:  NewZonedDateTime(date(`2011-12-30`), time(`12:00:00`), getZone(`Pacific/Apia`)),
			want: []clock.Instant{},
		},
		{
			sut:  NewZonedDateTime(date(`2011-12-31`), time(`12:00:00`), getZone(`Pacific/Apia`)),
			want: []clock.Instant{instant(`2011-12-31T12:00:00+14:00`)},
		},
		{
			sut:  NewZonedDateTime(date(`2012-09-01`), time(`12:00:00`), getZone(`Pacific/Apia`)),
			want: []clock.Instant{instant(`2012-09-01T12:00:00+13:00`)},
		},
		{
			sut:  NewZonedDateTime(date(`2023-01-01`), time(`12:00:00`), getZone(`Europe/Zurich`)),
			want: []clock.Instant{instant(`2023-01-01T12:00:00+01:00`)},
		},
		{
			sut:  NewZonedDateTime(date(`2023-07-01`), time(`12:00:00`), getZone(`Europe/Zurich`)),
			want: []clock.Instant{instant(`2023-07-01T12:00:00+02:00`)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.sut.String(), func(t *testing.T) {
			got := tt.sut.InstantCandidates()
			if len(tt.want) != len(got) {
				t.Error("the number of want instants != the number of got instants")
			} else {
				for i := 0; i < len(tt.want); i++ {
					want, _ := tt.want[i].Unix()
					got, _ := got[i].Unix()
					assert.Equal(t, want, got)
				}
			}
		})
	}
}
