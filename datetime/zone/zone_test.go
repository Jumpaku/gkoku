package zone_test

import (
	"fmt"
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

func TestZone_TransitionsBetween(t *testing.T) {
	unix := func(s int64) gkoku.Instant { return gkoku.Unix(s, 0) }
	tests := []struct {
		sut           zone2.Zone
		inBeginAt     gkoku.Instant
		inEndAt       gkoku.Instant
		wantTimestamp []gkoku.Instant
	}{
		{
			sut: zone2.Create("", []zone2.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
			}, nil),
			inBeginAt:     unix(30),
			inEndAt:       unix(0),
			wantTimestamp: []gkoku.Instant{},
		},
		{
			sut: zone2.Create("", []zone2.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
			}, nil),
			inBeginAt:     unix(5),
			inEndAt:       unix(5),
			wantTimestamp: []gkoku.Instant{},
		},
		{
			sut: zone2.Create("", []zone2.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
			}, nil),
			inBeginAt:     unix(10),
			inEndAt:       unix(10),
			wantTimestamp: []gkoku.Instant{unix(10)},
		},
		{
			sut: zone2.Create("", []zone2.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
			}, nil),
			inBeginAt:     unix(15),
			inEndAt:       unix(15),
			wantTimestamp: []gkoku.Instant{},
		},
		{
			sut: zone2.Create("", []zone2.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
			}, nil),
			inBeginAt:     unix(20),
			inEndAt:       unix(20),
			wantTimestamp: []gkoku.Instant{unix(20)},
		},
		{
			sut: zone2.Create("", []zone2.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
			}, nil),
			inBeginAt:     unix(25),
			inEndAt:       unix(25),
			wantTimestamp: []gkoku.Instant{},
		},
		{
			sut: zone2.Create("", []zone2.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
			}, nil),
			inBeginAt:     unix(0),
			inEndAt:       unix(5),
			wantTimestamp: []gkoku.Instant{},
		},
		{
			sut: zone2.Create("", []zone2.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
			}, nil),
			inBeginAt:     unix(5),
			inEndAt:       unix(10),
			wantTimestamp: []gkoku.Instant{unix(10)},
		},
		{
			sut: zone2.Create("", []zone2.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
			}, nil),
			inBeginAt:     unix(7),
			inEndAt:       unix(12),
			wantTimestamp: []gkoku.Instant{unix(10)},
		},
		{
			sut: zone2.Create("", []zone2.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
			}, nil),
			inBeginAt:     unix(10),
			inEndAt:       unix(15),
			wantTimestamp: []gkoku.Instant{unix(10)},
		},
		{
			sut: zone2.Create("", []zone2.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
			}, nil),
			inBeginAt:     unix(12),
			inEndAt:       unix(17),
			wantTimestamp: []gkoku.Instant{},
		},
		{
			sut: zone2.Create("", []zone2.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
			}, nil),
			inBeginAt:     unix(15),
			inEndAt:       unix(20),
			wantTimestamp: []gkoku.Instant{unix(20)},
		},
		{
			sut: zone2.Create("", []zone2.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
			}, nil),
			inBeginAt:     unix(17),
			inEndAt:       unix(22),
			wantTimestamp: []gkoku.Instant{unix(20)},
		},
		{
			sut: zone2.Create("", []zone2.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
			}, nil),
			inBeginAt:     unix(20),
			inEndAt:       unix(25),
			wantTimestamp: []gkoku.Instant{unix(20)},
		},
		{
			sut: zone2.Create("", []zone2.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
			}, nil),
			inBeginAt:     unix(25),
			inEndAt:       unix(30),
			wantTimestamp: []gkoku.Instant{},
		},
		{
			sut: zone2.Create("", []zone2.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
			}, nil),
			inBeginAt:     unix(10),
			inEndAt:       unix(20),
			wantTimestamp: []gkoku.Instant{unix(10), unix(20)},
		},
		{
			sut: zone2.Create("", []zone2.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
			}, nil),
			inBeginAt:     unix(5),
			inEndAt:       unix(20),
			wantTimestamp: []gkoku.Instant{unix(10), unix(20)},
		},
		{
			sut: zone2.Create("", []zone2.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
			}, nil),
			inBeginAt:     unix(10),
			inEndAt:       unix(25),
			wantTimestamp: []gkoku.Instant{unix(10), unix(20)},
		},
		{
			sut: zone2.Create("", []zone2.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
			}, nil),
			inBeginAt:     unix(5),
			inEndAt:       unix(25),
			wantTimestamp: []gkoku.Instant{unix(10), unix(20)},
		},
		{
			sut:           zone2.Create("", []zone2.Transition{}, nil),
			inBeginAt:     unix(10),
			inEndAt:       unix(20),
			wantTimestamp: []gkoku.Instant{},
		},
		{
			sut: zone2.Create("", []zone2.Transition{
				{TransitionTimestamp: unix(15)},
			}, nil),
			inBeginAt:     unix(5),
			inEndAt:       unix(10),
			wantTimestamp: []gkoku.Instant{},
		},
		{
			sut: zone2.Create("", []zone2.Transition{
				{TransitionTimestamp: unix(15)},
			}, nil),
			inBeginAt:     unix(10),
			inEndAt:       unix(15),
			wantTimestamp: []gkoku.Instant{unix(15)},
		},
		{
			sut: zone2.Create("", []zone2.Transition{
				{TransitionTimestamp: unix(15)},
			}, nil),
			inBeginAt:     unix(12),
			inEndAt:       unix(17),
			wantTimestamp: []gkoku.Instant{unix(15)},
		},
		{
			sut: zone2.Create("", []zone2.Transition{
				{TransitionTimestamp: unix(15)},
			}, nil),
			inBeginAt:     unix(15),
			inEndAt:       unix(20),
			wantTimestamp: []gkoku.Instant{unix(15)},
		},
		{
			sut: zone2.Create("", []zone2.Transition{
				{TransitionTimestamp: unix(15)},
			}, nil),
			inBeginAt:     unix(20),
			inEndAt:       unix(25),
			wantTimestamp: []gkoku.Instant{},
		},
		{
			sut: zone2.Create("", []zone2.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
				{TransitionTimestamp: unix(30)},
			}, nil),
			inBeginAt:     unix(5),
			inEndAt:       unix(15),
			wantTimestamp: []gkoku.Instant{unix(10)},
		},
		{
			sut: zone2.Create("", []zone2.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
				{TransitionTimestamp: unix(30)},
			}, nil),
			inBeginAt:     unix(5),
			inEndAt:       unix(20),
			wantTimestamp: []gkoku.Instant{unix(10), unix(20)},
		},
		{
			sut: zone2.Create("", []zone2.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
				{TransitionTimestamp: unix(30)},
			}, nil),
			inBeginAt:     unix(5),
			inEndAt:       unix(25),
			wantTimestamp: []gkoku.Instant{unix(10), unix(20)},
		},
		{
			sut: zone2.Create("", []zone2.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
				{TransitionTimestamp: unix(30)},
			}, nil),
			inBeginAt:     unix(10),
			inEndAt:       unix(25),
			wantTimestamp: []gkoku.Instant{unix(10), unix(20)},
		},
		{
			sut: zone2.Create("", []zone2.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
				{TransitionTimestamp: unix(30)},
			}, nil),
			inBeginAt:     unix(15),
			inEndAt:       unix(25),
			wantTimestamp: []gkoku.Instant{unix(20)},
		},
		{
			sut: zone2.Create("", []zone2.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
				{TransitionTimestamp: unix(30)},
			}, nil),
			inBeginAt:     unix(15),
			inEndAt:       unix(30),
			wantTimestamp: []gkoku.Instant{unix(20), unix(30)},
		},
		{
			sut: zone2.Create("", []zone2.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
				{TransitionTimestamp: unix(30)},
			}, nil),
			inBeginAt:     unix(15),
			inEndAt:       unix(35),
			wantTimestamp: []gkoku.Instant{unix(20), unix(30)},
		},
		{
			sut: zone2.Create("", []zone2.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
				{TransitionTimestamp: unix(30)},
			}, nil),
			inBeginAt:     unix(20),
			inEndAt:       unix(35),
			wantTimestamp: []gkoku.Instant{unix(20), unix(30)},
		},
		{
			sut: zone2.Create("", []zone2.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
				{TransitionTimestamp: unix(30)},
			}, nil),
			inBeginAt:     unix(25),
			inEndAt:       unix(35),
			wantTimestamp: []gkoku.Instant{unix(30)},
		},
	}
	for number, tt := range tests {
		t.Run(fmt.Sprintf(`%d`, number), func(t *testing.T) {
			got := tt.sut.TransitionsBetween(tt.inBeginAt, tt.inEndAt)
			if len(tt.wantTimestamp) != len(got) {
				t.Error("the number of want transitions != the number of got transitions")
			} else {
				for i := 0; i < len(tt.wantTimestamp); i++ {
					want, _ := tt.wantTimestamp[i].Unix()
					got, _ := got[i].TransitionTimestamp.Unix()
					assert.Equal(t, want, got)
				}
			}
		})
	}
}
