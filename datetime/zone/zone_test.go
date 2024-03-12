package zone_test

import (
	"fmt"
	"github.com/Jumpaku/gkoku/clock"
	"github.com/Jumpaku/gkoku/datetime/zone"
	"github.com/Jumpaku/gkoku/internal/tests/assert"
	"testing"
)

func TestCreateFixed(t *testing.T) {
	tests := []struct {
		id     string
		offset zone.OffsetMinutes
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
			got := zone.CreateFixed(tt.id, tt.offset)
			assert.Equal(t, tt.id, got.ID())
			assert.Equal(t, tt.offset, got.FindOffset(clock.MinInstant))
			assert.Equal(t, tt.offset, got.FindOffset(clock.MaxInstant))
		})
	}
}

func TestZone_FindOffset(t *testing.T) {
	provider := zone.DefaultProvider()
	load := func(zoneID string) zone.Zone {
		z, ok := provider.Get(zoneID)
		if !ok {
			t.Fatal()
		}
		return z
	}

	tests := []struct {
		sut  zone.Zone
		in   clock.Instant
		want zone.OffsetMinutes
	}{
		{
			sut:  load("Asia/Tokyo"),
			in:   clock.Unix(946684800, 0),
			want: 540,
		},
		{
			sut:  load("Pacific/Apia"),
			in:   clock.Unix(946684800, 0),
			want: -660,
		},
		{
			sut:  load("Europe/Zurich"),
			in:   clock.Unix(946684800, 0),
			want: 60,
		},
		{
			sut:  load("Zulu"),
			in:   clock.Unix(946684800, 0),
			want: 0,
		},
		{
			sut:  zone.Create("example1", nil),
			in:   clock.Unix(946684800, 0),
			want: 0,
		},
		{
			sut: zone.Create("example2", []zone.Transition{
				{TransitionTimestamp: clock.MinInstant, OffsetMinutesBefore: -1, OffsetMinutesAfter: 1},
			}),
			in:   clock.Unix(946684800, 0),
			want: 1,
		},
		{
			sut: zone.Create("example3", []zone.Transition{
				{TransitionTimestamp: clock.MaxInstant, OffsetMinutesBefore: -1, OffsetMinutesAfter: 1},
			}),
			in:   clock.Unix(946684800, 0),
			want: -1,
		},
		{
			sut: zone.Create("example4", []zone.Transition{
				{TransitionTimestamp: clock.Unix(946684800, 0), OffsetMinutesBefore: -1, OffsetMinutesAfter: 1},
			}),
			in:   clock.Unix(946684800, 0),
			want: 1,
		},
		{
			sut: zone.Create("example5", []zone.Transition{
				{TransitionTimestamp: clock.Unix(946684799, 0), OffsetMinutesBefore: -1, OffsetMinutesAfter: 1},
			}),
			in:   clock.Unix(946684800, 0),
			want: 1,
		},
		{
			sut: zone.Create("example6", []zone.Transition{
				{TransitionTimestamp: clock.Unix(946684801, 0), OffsetMinutesBefore: -1, OffsetMinutesAfter: 1},
			}),
			in:   clock.Unix(946684800, 0),
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
	unix := func(s int64) clock.Instant { return clock.Unix(s, 0) }
	tests := []struct {
		sut           zone.Zone
		inBeginAt     clock.Instant
		inEndAt       clock.Instant
		wantTimestamp []clock.Instant
	}{
		{
			sut: zone.Create("", []zone.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
			}),
			inBeginAt:     unix(30),
			inEndAt:       unix(0),
			wantTimestamp: []clock.Instant{},
		},
		{
			sut: zone.Create("", []zone.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
			}),
			inBeginAt:     unix(5),
			inEndAt:       unix(5),
			wantTimestamp: []clock.Instant{},
		},
		{
			sut: zone.Create("", []zone.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
			}),
			inBeginAt:     unix(10),
			inEndAt:       unix(10),
			wantTimestamp: []clock.Instant{unix(10)},
		},
		{
			sut: zone.Create("", []zone.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
			}),
			inBeginAt:     unix(15),
			inEndAt:       unix(15),
			wantTimestamp: []clock.Instant{},
		},
		{
			sut: zone.Create("", []zone.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
			}),
			inBeginAt:     unix(20),
			inEndAt:       unix(20),
			wantTimestamp: []clock.Instant{unix(20)},
		},
		{
			sut: zone.Create("", []zone.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
			}),
			inBeginAt:     unix(25),
			inEndAt:       unix(25),
			wantTimestamp: []clock.Instant{},
		},
		{
			sut: zone.Create("", []zone.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
			}),
			inBeginAt:     unix(0),
			inEndAt:       unix(5),
			wantTimestamp: []clock.Instant{},
		},
		{
			sut: zone.Create("", []zone.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
			}),
			inBeginAt:     unix(5),
			inEndAt:       unix(10),
			wantTimestamp: []clock.Instant{unix(10)},
		},
		{
			sut: zone.Create("", []zone.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
			}),
			inBeginAt:     unix(7),
			inEndAt:       unix(12),
			wantTimestamp: []clock.Instant{unix(10)},
		},
		{
			sut: zone.Create("", []zone.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
			}),
			inBeginAt:     unix(10),
			inEndAt:       unix(15),
			wantTimestamp: []clock.Instant{unix(10)},
		},
		{
			sut: zone.Create("", []zone.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
			}),
			inBeginAt:     unix(12),
			inEndAt:       unix(17),
			wantTimestamp: []clock.Instant{},
		},
		{
			sut: zone.Create("", []zone.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
			}),
			inBeginAt:     unix(15),
			inEndAt:       unix(20),
			wantTimestamp: []clock.Instant{unix(20)},
		},
		{
			sut: zone.Create("", []zone.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
			}),
			inBeginAt:     unix(17),
			inEndAt:       unix(22),
			wantTimestamp: []clock.Instant{unix(20)},
		},
		{
			sut: zone.Create("", []zone.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
			}),
			inBeginAt:     unix(20),
			inEndAt:       unix(25),
			wantTimestamp: []clock.Instant{unix(20)},
		},
		{
			sut: zone.Create("", []zone.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
			}),
			inBeginAt:     unix(25),
			inEndAt:       unix(30),
			wantTimestamp: []clock.Instant{},
		},
		{
			sut: zone.Create("", []zone.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
			}),
			inBeginAt:     unix(10),
			inEndAt:       unix(20),
			wantTimestamp: []clock.Instant{unix(10), unix(20)},
		},
		{
			sut: zone.Create("", []zone.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
			}),
			inBeginAt:     unix(5),
			inEndAt:       unix(20),
			wantTimestamp: []clock.Instant{unix(10), unix(20)},
		},
		{
			sut: zone.Create("", []zone.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
			}),
			inBeginAt:     unix(10),
			inEndAt:       unix(25),
			wantTimestamp: []clock.Instant{unix(10), unix(20)},
		},
		{
			sut: zone.Create("", []zone.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
			}),
			inBeginAt:     unix(5),
			inEndAt:       unix(25),
			wantTimestamp: []clock.Instant{unix(10), unix(20)},
		},
		{
			sut:           zone.Create("", []zone.Transition{}),
			inBeginAt:     unix(10),
			inEndAt:       unix(20),
			wantTimestamp: []clock.Instant{},
		},
		{
			sut: zone.Create("", []zone.Transition{
				{TransitionTimestamp: unix(15)},
			}),
			inBeginAt:     unix(5),
			inEndAt:       unix(10),
			wantTimestamp: []clock.Instant{},
		},
		{
			sut: zone.Create("", []zone.Transition{
				{TransitionTimestamp: unix(15)},
			}),
			inBeginAt:     unix(10),
			inEndAt:       unix(15),
			wantTimestamp: []clock.Instant{unix(15)},
		},
		{
			sut: zone.Create("", []zone.Transition{
				{TransitionTimestamp: unix(15)},
			}),
			inBeginAt:     unix(12),
			inEndAt:       unix(17),
			wantTimestamp: []clock.Instant{unix(15)},
		},
		{
			sut: zone.Create("", []zone.Transition{
				{TransitionTimestamp: unix(15)},
			}),
			inBeginAt:     unix(15),
			inEndAt:       unix(20),
			wantTimestamp: []clock.Instant{unix(15)},
		},
		{
			sut: zone.Create("", []zone.Transition{
				{TransitionTimestamp: unix(15)},
			}),
			inBeginAt:     unix(20),
			inEndAt:       unix(25),
			wantTimestamp: []clock.Instant{},
		},
		{
			sut: zone.Create("", []zone.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
				{TransitionTimestamp: unix(30)},
			}),
			inBeginAt:     unix(5),
			inEndAt:       unix(15),
			wantTimestamp: []clock.Instant{unix(10)},
		},
		{
			sut: zone.Create("", []zone.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
				{TransitionTimestamp: unix(30)},
			}),
			inBeginAt:     unix(5),
			inEndAt:       unix(20),
			wantTimestamp: []clock.Instant{unix(10), unix(20)},
		},
		{
			sut: zone.Create("", []zone.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
				{TransitionTimestamp: unix(30)},
			}),
			inBeginAt:     unix(5),
			inEndAt:       unix(25),
			wantTimestamp: []clock.Instant{unix(10), unix(20)},
		},
		{
			sut: zone.Create("", []zone.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
				{TransitionTimestamp: unix(30)},
			}),
			inBeginAt:     unix(10),
			inEndAt:       unix(25),
			wantTimestamp: []clock.Instant{unix(10), unix(20)},
		},
		{
			sut: zone.Create("", []zone.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
				{TransitionTimestamp: unix(30)},
			}),
			inBeginAt:     unix(15),
			inEndAt:       unix(25),
			wantTimestamp: []clock.Instant{unix(20)},
		},
		{
			sut: zone.Create("", []zone.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
				{TransitionTimestamp: unix(30)},
			}),
			inBeginAt:     unix(15),
			inEndAt:       unix(30),
			wantTimestamp: []clock.Instant{unix(20), unix(30)},
		},
		{
			sut: zone.Create("", []zone.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
				{TransitionTimestamp: unix(30)},
			}),
			inBeginAt:     unix(15),
			inEndAt:       unix(35),
			wantTimestamp: []clock.Instant{unix(20), unix(30)},
		},
		{
			sut: zone.Create("", []zone.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
				{TransitionTimestamp: unix(30)},
			}),
			inBeginAt:     unix(20),
			inEndAt:       unix(35),
			wantTimestamp: []clock.Instant{unix(20), unix(30)},
		},
		{
			sut: zone.Create("", []zone.Transition{
				{TransitionTimestamp: unix(10)},
				{TransitionTimestamp: unix(20)},
				{TransitionTimestamp: unix(30)},
			}),
			inBeginAt:     unix(25),
			inEndAt:       unix(35),
			wantTimestamp: []clock.Instant{unix(30)},
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
