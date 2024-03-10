package datetime_test

import (
	"github.com/Jumpaku/gkoku/clock"
	"github.com/Jumpaku/gkoku/datetime"
	"github.com/Jumpaku/gkoku/internal/tests/assert"
	"testing"
)

func TestFixedOffsetZone(t *testing.T) {
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
			got := datetime.FixedZone(tt.id, tt.offset)
			assert.Equal(t, tt.id, got.ID())
			assert.Equal(t, tt.offset, got.FindOffset(clock.MinInstant))
			assert.Equal(t, tt.offset, got.FindOffset(clock.MaxInstant))
		})
	}
}

func TestLoad(t *testing.T) {
	tests := []struct {
		id      string
		wantErr bool
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
			id:      "not-found",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			got, err := datetime.LoadZone(tt.id)
			if tt.wantErr {
				if err == nil {
					t.Errorf("LoadZone() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			} else {
				assert.Equal(t, tt.id, got.ID())
			}
		})
	}
}

func TestZone_FindOffset(t *testing.T) {
	load := func(zoneID string) datetime.Zone {
		z, _ := datetime.LoadZone(zoneID)
		return z
	}
	tests := []struct {
		sut  datetime.Zone
		in   clock.Instant
		want datetime.OffsetMinutes
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
			sut:  datetime.CreateZone("example1", nil),
			in:   clock.Unix(946684800, 0),
			want: 0,
		},
		{
			sut: datetime.CreateZone("example2", []datetime.Transition{
				{When: clock.MinInstant, OffsetBefore: -1, OffsetAfter: 1},
			}),
			in:   clock.Unix(946684800, 0),
			want: 1,
		},
		{
			sut: datetime.CreateZone("example3", []datetime.Transition{
				{When: clock.MaxInstant, OffsetBefore: -1, OffsetAfter: 1},
			}),
			in:   clock.Unix(946684800, 0),
			want: -1,
		},
		{
			sut: datetime.CreateZone("example4", []datetime.Transition{
				{When: clock.Unix(946684800, 0), OffsetBefore: -1, OffsetAfter: 1},
			}),
			in:   clock.Unix(946684800, 0),
			want: 1,
		},
		{
			sut: datetime.CreateZone("example5", []datetime.Transition{
				{When: clock.Unix(946684799, 0), OffsetBefore: -1, OffsetAfter: 1},
			}),
			in:   clock.Unix(946684800, 0),
			want: 1,
		},
		{
			sut: datetime.CreateZone("example6", []datetime.Transition{
				{When: clock.Unix(946684801, 0), OffsetBefore: -1, OffsetAfter: 1},
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
