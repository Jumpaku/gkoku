package datetime

import (
	"github.com/Jumpaku/gkoku/internal/tests/assert"
	"testing"
)

func TestFormatTime(t *testing.T) {
	tests := []struct {
		in   Time
		want string
	}{
		{
			in:   TimeOf(1, 2, 3, 4),
			want: `T01:02:03.000000004`,
		},
		{
			in:   TimeOf(0, 0, 0, 0),
			want: `T00:00:00.000000000`,
		},
		{
			in:   TimeOf(23, 59, 59, 999999999),
			want: `T23:59:59.999999999`,
		},
		{
			in:   TimeOf(24, 0, 0, 0),
			want: `T24:00:00.000000000`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			got := FormatTime(tt.in)
			if got != tt.want {
				t.Errorf("FormatTime() = %v, want %v", got, tt.want)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestParseTime(t *testing.T) {
	tests := []struct {
		in      string
		want    Time
		wantErr bool
	}{
		{
			in:   `T00`,
			want: TimeOf(0, 0, 0, 0),
		},
		{
			in:   `00`,
			want: TimeOf(0, 0, 0, 0),
		},
		{
			in:   `T00:00`,
			want: TimeOf(0, 0, 0, 0),
		},
		{
			in:   `00:00`,
			want: TimeOf(0, 0, 0, 0),
		},
		{
			in:   `T00:00:00`,
			want: TimeOf(0, 0, 0, 0),
		},
		{
			in:   `00:00:00`,
			want: TimeOf(0, 0, 0, 0),
		},
		{
			in:   `T00:00:00,0`,
			want: TimeOf(0, 0, 0, 0),
		},
		{
			in:   `00:00:00,0`,
			want: TimeOf(0, 0, 0, 0),
		},
		{
			in:   `T00:00:00.0`,
			want: TimeOf(0, 0, 0, 0),
		},
		{
			in:   `00:00:00.0`,
			want: TimeOf(0, 0, 0, 0),
		},
		{
			in:   `T00:00:00,000000000`,
			want: TimeOf(0, 0, 0, 0),
		},
		{
			in:   `00:00:00,000000000`,
			want: TimeOf(0, 0, 0, 0),
		},
		{
			in:   `T00:00:00.000000000`,
			want: TimeOf(0, 0, 0, 0),
		},
		{
			in:   `00:00:00.000000000`,
			want: TimeOf(0, 0, 0, 0),
		},
		{
			in:   `T24:00:00.000000000`,
			want: TimeOf(24, 0, 0, 0),
		},
		{
			in:   `T23:59:59.999999999`,
			want: TimeOf(23, 59, 59, 999_999_999),
		},
		{
			in:   `T01:01:01.100000000`,
			want: TimeOf(1, 1, 1, 100_000_000),
		},
		{
			in:   `T01:01:01.100000000`,
			want: TimeOf(1, 1, 1, 100_000_000),
		},
		{
			in:      `T01:01:01.1234567890`,
			wantErr: true,
		},
		{
			in:      `1:1:1`,
			wantErr: true,
		},
		{
			in:      `01:01:01.`,
			wantErr: true,
		},
		{
			in:      `01:01:60`,
			wantErr: true,
		},
		{
			in:      `01:60:01`,
			wantErr: true,
		},
		{
			in:      `24:00:01`,
			wantErr: true,
		},
		{
			in:      `25:00:00`,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got, err := ParseTime(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseTime() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, tt.want.Hour(), got.Hour())
			assert.Equal(t, tt.want.Minute(), got.Minute())
			assert.Equal(t, tt.want.Second(), got.Second())
			assert.Equal(t, tt.want.Nano(), got.Nano())
		})
	}
}
