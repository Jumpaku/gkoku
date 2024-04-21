package zone_test

import (
	"fmt"
	"github.com/Jumpaku/gkoku/clock"
	"github.com/Jumpaku/gkoku/internal/tests/assert"
	. "github.com/Jumpaku/gkoku/zone"
	"testing"
)

func TestParseOffset(t *testing.T) {
	tests := []struct {
		in         string
		wantOffset OffsetMinutes
		wantErr    bool
	}{
		{
			in:         `Z`,
			wantOffset: 0,
		},
		{
			in:         `+00:00`,
			wantOffset: 0,
		},
		{
			in:         `+12`,
			wantOffset: 12 * 60,
		},
		{
			in:         `+1234`,
			wantOffset: 12*60 + 34,
		},
		{
			in:         `+12:34`,
			wantOffset: 12*60 + 34,
		},
		{
			in:         `-12`,
			wantOffset: -12 * 60,
		},
		{
			in:         `-1234`,
			wantOffset: -(12*60 + 34),
		},
		{
			in:         `-12:34`,
			wantOffset: -(12*60 + 34),
		},
		{
			in:      `-1:34`,
			wantErr: true,
		},
		{
			in:      `-12:3`,
			wantErr: true,
		},
		{
			in:      `-123`,
			wantErr: true,
		},
		{
			in:      `1234`,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			gotOffset, err := ParseOffset(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseOffset() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOffset != tt.wantOffset {
				t.Errorf("ParseOffset() gotOffset = %v, want %v", gotOffset, tt.wantOffset)
			}
		})
	}
}

func TestFormatOffset(t *testing.T) {
	tests := []struct {
		in   OffsetMinutes
		want string
	}{
		{
			in:   0,
			want: `+00:00`,
		},
		{
			in:   12 * 60,
			want: `+12:00`,
		},
		{
			in:   12*60 + 34,
			want: `+12:34`,
		},
		{
			in:   -12 * 60,
			want: `-12:00`,
		},
		{
			in:   -(12*60 + 34),
			want: `-12:34`,
		},
		{
			in:   1 * 60,
			want: `+01:00`,
		},
		{
			in:   1*60 + 1,
			want: `+01:01`,
		},
		{
			in:   -1 * 60,
			want: `-01:00`,
		},
		{
			in:   -(1*60 + 1),
			want: `-01:01`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := FormatOffset(tt.in); got != tt.want {
				t.Errorf("FormatOffset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOffsetMinutes_AddTo(t *testing.T) {
	tests := []struct {
		sut  OffsetMinutes
		in   clock.Instant
		want clock.Instant
	}{
		{
			sut:  0,
			in:   clock.Unix(1, 1),
			want: clock.Unix(1, 1),
		},
		{
			sut:  1,
			in:   clock.Unix(1, 1),
			want: clock.Unix(61, 1),
		},
		{
			sut:  -1,
			in:   clock.Unix(1, 1),
			want: clock.Unix(-59, 1),
		},
		{
			sut:  60,
			in:   clock.Unix(1, 1),
			want: clock.Unix(3601, 1),
		},
		{
			sut:  -60,
			in:   clock.Unix(1, 1),
			want: clock.Unix(-3599, 1),
		},
		{
			sut:  0,
			in:   clock.Unix(-1, 1),
			want: clock.Unix(-1, 1),
		},
		{
			sut:  1,
			in:   clock.Unix(-1, 1),
			want: clock.Unix(59, 1),
		},
		{
			sut:  -1,
			in:   clock.Unix(-1, 1),
			want: clock.Unix(-61, 1),
		},
		{
			sut:  60,
			in:   clock.Unix(-1, 1),
			want: clock.Unix(3599, 1),
		},
		{
			sut:  -60,
			in:   clock.Unix(-1, 1),
			want: clock.Unix(-3601, 1),
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf(`%s%s`, tt.in.String(), tt.sut.String()), func(t *testing.T) {
			gotS, gotN := tt.sut.AddTo(tt.in).Unix()
			wantS, wantN := tt.want.Unix()
			assert.Equal(t, wantS, gotS)
			assert.Equal(t, wantN, gotN)
		})
	}
}
