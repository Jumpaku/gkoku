package zone_test

import (
	"github.com/Jumpaku/gkoku"
	"github.com/Jumpaku/gkoku/datetime"
	zone2 "github.com/Jumpaku/gkoku/datetime/zone"
	"github.com/Jumpaku/gkoku/internal/tests/assert"
	"testing"
)

func TestCreateFixed(t *testing.T) {
	tests := []struct {
		id     string
		offset datetime.OffsetMinutes
	}{
		{
			id:     "",
			offset: 0,
		},
		{
			id:     "example1",
			offset: 1,
		},
		{
			id:     "example2",
			offset: -1,
		},
		{
			id:     "example3",
			offset: 14 * 60,
		},
		{
			id:     "example4",
			offset: -14 * 60,
		},
		{
			id:     "example5",
			offset: 60,
		},
		{
			id:     "example6",
			offset: -60,
		},
	}
	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			got := zone2.CreateFixed(tt.id, tt.offset)
			assert.Equal(t, tt.id, got.ID())
			assert.Equal(t, tt.offset, got.FindOffset(gkoku.MinInstant))
			assert.Equal(t, tt.offset, got.FindOffset(gkoku.MaxInstant))
		})
	}
}

func TestZone_FindOffset(t *testing.T) {
	provider := getTestProvider()
	load := func(zoneID string) zone2.Zone {
		z, ok := provider.Get(zoneID)
		if !ok {
			t.Fatal()
		}
		return z
	}

	tests := []struct {
		sut  zone2.Zone
		in   gkoku.Instant
		want datetime.OffsetMinutes
	}{
		{
			sut:  load("Asia/Tokyo"),
			in:   gkoku.Unix(946684800, 0),
			want: 540,
		},
		{
			sut:  load("Pacific/Apia"),
			in:   gkoku.Unix(946684800, 0),
			want: -660,
		},
		{
			sut:  load("Europe/Zurich"),
			in:   gkoku.Unix(946684800, 0),
			want: 60,
		},
		{
			sut:  load("Zulu"),
			in:   gkoku.Unix(946684800, 0),
			want: 0,
		},
		{
			sut:  zone2.Create("example1", nil, nil),
			in:   gkoku.Unix(946684800, 0),
			want: 0,
		},
		{
			sut: zone2.Create("example2", []zone2.Transition{
				{TransitionTimestamp: gkoku.MinInstant, OffsetMinutesBefore: -1, OffsetMinutesAfter: 1},
			}, nil),
			in:   gkoku.Unix(946684800, 0),
			want: 1,
		},
		{
			sut: zone2.Create("example3", []zone2.Transition{
				{TransitionTimestamp: gkoku.MaxInstant, OffsetMinutesBefore: -1, OffsetMinutesAfter: 1},
			}, nil),
			in:   gkoku.Unix(946684800, 0),
			want: -1,
		},
		{
			sut: zone2.Create("example4", []zone2.Transition{
				{TransitionTimestamp: gkoku.Unix(946684800, 0), OffsetMinutesBefore: -1, OffsetMinutesAfter: 1},
			}, nil),
			in:   gkoku.Unix(946684800, 0),
			want: 1,
		},
		{
			sut: zone2.Create("example5", []zone2.Transition{
				{TransitionTimestamp: gkoku.Unix(946684799, 0), OffsetMinutesBefore: -1, OffsetMinutesAfter: 1},
			}, nil),
			in:   gkoku.Unix(946684800, 0),
			want: 1,
		},
		{
			sut: zone2.Create("example6", []zone2.Transition{
				{TransitionTimestamp: gkoku.Unix(946684801, 0), OffsetMinutesBefore: -1, OffsetMinutesAfter: 1},
			}, nil),
			in:   gkoku.Unix(946684800, 0),
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.sut.ID(), func(t *testing.T) {
			got := tt.sut.FindOffset(tt.in)
			assert.Equal(t, tt.want, got)
		})
	}
}
