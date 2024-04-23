package zone_test

import (
	"github.com/Jumpaku/tokiope"
	"github.com/Jumpaku/tokiope/date"
	. "github.com/Jumpaku/tokiope/datetime"
	. "github.com/Jumpaku/tokiope/datetime/zone"
	"github.com/Jumpaku/tokiope/internal/tests/assert"
	"testing"
)

func TestZonedDateTime_InstantCandidates(t *testing.T) {
	testProvider := getTestProvider()
	instantOf := func(s string) tokiope.Instant {
		d, _ := ParseOffsetDateTime(s)
		return d.Instant()
	}
	dateOf := func(s string) date.Date {
		d, _ := date.ParseDate(s, date.DateFormatYyyyMmDd)
		return d
	}
	timeOf := func(s string) Time {
		t, _ := ParseTime(s)
		return t
	}
	getZone := func(zoneID string) Zone {
		z, _ := testProvider.Get(zoneID)
		return z
	}

	tests := []struct {
		sut  ZonedDateTime
		want []tokiope.Instant
	}{
		{
			sut:  NewZonedDateTime(dateOf(`2024-03-13`), timeOf(`00:00:00`), getZone(`Zulu`)),
			want: []tokiope.Instant{instantOf(`2024-03-13T00:00:00Z`)},
		},
		{
			sut:  NewZonedDateTime(dateOf(`2011-09-01`), timeOf(`12:00:00`), getZone(`Pacific/Apia`)),
			want: []tokiope.Instant{instantOf(`2011-09-01T12:00:00-11:00`)},
		},
		{
			sut:  NewZonedDateTime(dateOf(`2011-12-29`), timeOf(`12:00:00`), getZone(`Pacific/Apia`)),
			want: []tokiope.Instant{instantOf(`2011-12-29T12:00:00-10:00`)},
		},
		{
			sut:  NewZonedDateTime(dateOf(`2011-12-30`), timeOf(`12:00:00`), getZone(`Pacific/Apia`)),
			want: []tokiope.Instant{},
		},
		{
			sut:  NewZonedDateTime(dateOf(`2011-12-31`), timeOf(`12:00:00`), getZone(`Pacific/Apia`)),
			want: []tokiope.Instant{instantOf(`2011-12-31T12:00:00+14:00`)},
		},
		{
			sut:  NewZonedDateTime(dateOf(`2012-09-01`), timeOf(`12:00:00`), getZone(`Pacific/Apia`)),
			want: []tokiope.Instant{instantOf(`2012-09-01T12:00:00+13:00`)},
		},
		{
			sut:  NewZonedDateTime(dateOf(`2023-01-01`), timeOf(`12:00:00`), getZone(`Europe/Zurich`)),
			want: []tokiope.Instant{instantOf(`2023-01-01T12:00:00+01:00`)},
		},
		{
			sut:  NewZonedDateTime(dateOf(`2023-07-01`), timeOf(`12:00:00`), getZone(`Europe/Zurich`)),
			want: []tokiope.Instant{instantOf(`2023-07-01T12:00:00+02:00`)},
		},

		// Gap
		{
			sut:  NewZonedDateTime(dateOf(`2023-03-26`), timeOf(`01:30:00`), getZone(`Europe/Zurich`)),
			want: []tokiope.Instant{instantOf(`2023-03-26T01:30:00+01:00`)},
		},
		{
			sut:  NewZonedDateTime(dateOf(`2023-03-26`), timeOf(`02:00:00`), getZone(`Europe/Zurich`)),
			want: []tokiope.Instant{},
		},
		{
			sut:  NewZonedDateTime(dateOf(`2023-03-26`), timeOf(`02:30:00`), getZone(`Europe/Zurich`)),
			want: []tokiope.Instant{},
		},
		{
			sut:  NewZonedDateTime(dateOf(`2023-03-26`), timeOf(`03:00:00`), getZone(`Europe/Zurich`)),
			want: []tokiope.Instant{instantOf(`2023-03-26T03:00:00+02:00`)},
		},
		{
			sut:  NewZonedDateTime(dateOf(`2023-03-26`), timeOf(`03:30:00`), getZone(`Europe/Zurich`)),
			want: []tokiope.Instant{instantOf(`2023-03-26T03:30:00+02:00`)},
		},

		// Overlap
		{
			sut:  NewZonedDateTime(dateOf(`2023-10-29`), timeOf(`01:30:00`), getZone(`Europe/Zurich`)),
			want: []tokiope.Instant{instantOf(`2023-10-29T01:30:00+02:00`)},
		},
		{
			sut:  NewZonedDateTime(dateOf(`2023-10-29`), timeOf(`02:00:00`), getZone(`Europe/Zurich`)),
			want: []tokiope.Instant{instantOf(`2023-10-29T02:00:00+02:00`), instantOf(`2023-10-29T02:00:00+01:00`)},
		},
		{
			sut:  NewZonedDateTime(dateOf(`2023-10-29`), timeOf(`02:30:00`), getZone(`Europe/Zurich`)),
			want: []tokiope.Instant{instantOf(`2023-10-29T02:30:00+02:00`), instantOf(`2023-10-29T02:30:00+01:00`)},
		},
		{
			sut:  NewZonedDateTime(dateOf(`2023-10-29`), timeOf(`03:00:00`), getZone(`Europe/Zurich`)),
			want: []tokiope.Instant{instantOf(`2023-10-29T03:00:00+01:00`)},
		},
		{
			sut:  NewZonedDateTime(dateOf(`2023-10-29`), timeOf(`03:30:00`), getZone(`Europe/Zurich`)),
			want: []tokiope.Instant{instantOf(`2023-10-29T03:30:00+01:00`)},
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
