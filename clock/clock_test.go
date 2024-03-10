package clock

import (
	"fmt"
	"github.com/Jumpaku/gkoku/internal/tests/assert"
	assert2 "github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFixedClock(t *testing.T) {
	tests := []struct {
		in   Instant
		want Instant
	}{
		{
			in:   Unix(1, 0),
			want: Unix(1, 0),
		},
		{
			in:   Unix(0, 0),
			want: Unix(0, 0),
		},
		{
			in:   Unix(-1, 999_999_999),
			want: Unix(-1, 999_999_999),
		},
	}
	for number, tt := range tests {
		t.Run(fmt.Sprintf("%d", number), func(t *testing.T) {
			sut := Fixed(tt.in)
			gotS, gotN := sut.Now().Unix()
			wantS, wantN := tt.want.Unix()
			assert.Equal(t, wantS, gotS)
			assert.Equal(t, wantN, gotN)
		})
	}
}

func TestOffsetClock(t *testing.T) {
	tests := []struct {
		fix    Instant
		offset Duration
		want   Instant
	}{
		{
			fix:    Unix(1, 0),
			offset: Seconds(2, 500_000_000),
			want:   Unix(3, 500_000_000),
		},
		{
			fix:    Unix(1, 0),               // 1.0
			offset: Seconds(-3, 500_000_000), // -2.5
			want:   Unix(-2, 500_000_000),    //-1.5
		},
		{
			fix:    Unix(-1, 999_999_999),
			offset: Seconds(0, 0),
			want:   Unix(-1, 999_999_999),
		},
	}
	for number, tt := range tests {
		t.Run(fmt.Sprintf("%d", number), func(t *testing.T) {
			sut := Offset(Fixed(tt.fix), tt.offset)
			gotS, gotN := sut.Now().Unix()
			wantS, wantN := tt.want.Unix()
			assert.Equal(t, wantS, gotS)
			assert.Equal(t, wantN, gotN)
		})
	}
}

func TestWallClock(t *testing.T) {
	for number := 0; number < 5; number++ {
		t.Run(fmt.Sprintf("%d", number), func(t *testing.T) {
			sut := Wall()
			got, _ := sut.Now().Unix()
			want := time.Now().Unix()
			assert2.True(t, got <= want)
			assert2.True(t, got >= want-1)
		})
	}
}
