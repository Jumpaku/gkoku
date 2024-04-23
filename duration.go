package tokiope

import (
	"fmt"
	"math"
)

type State int8

const (
	// StateOK represents no error.
	StateOK State = (1 << iota) >> 1
	// StateOverflow bit is set if overflow occurred
	StateOverflow
)

const NanosPerSecond = 1_000_000_000

const (
	SecondsPerMinute = 60
	MinutesPerHour   = 60
	HoursPerDay      = 24
	SecondsPerHour   = SecondsPerMinute * MinutesPerHour
	SecondsPerDay    = SecondsPerHour * HoursPerDay
)

type Duration struct {
	state   State
	seconds int64
	nano    int
}

var MinDuration = Duration{seconds: math.MinInt64}
var MaxDuration = Duration{seconds: math.MaxInt64, nano: NanosPerSecond - 1}

func Seconds(seconds int64, nano int64) (d Duration) {
	secs, nanos, state := divFloor(nano, NanosPerSecond)
	d.state |= state

	secs, state = add(secs, seconds)
	d.state |= state

	d.seconds, d.nano = secs, int(nanos)
	return
}

func Minutes(minutes int64) (d Duration) {
	d.seconds, d.state = mul(minutes, SecondsPerMinute)
	return
}

func Hours(hours int64) (d Duration) {
	d.seconds, d.state = mul(hours, SecondsPerHour)
	return
}

func Days(days int64) (d Duration) {
	d.seconds, d.state = mul(days, SecondsPerDay)
	return
}

func Nanoseconds(nanoseconds int64) (d Duration) {
	secs, nano, state := divFloor(nanoseconds, NanosPerSecond)
	d.seconds, d.nano, d.state = secs, int(nano), state
	return
}

var _ interface {
	Seconds() (seconds int64, nano int64)
	Add(o Duration) Duration
	AddNano(nanoseconds int64) Duration
	Sub(o Duration) Duration
	SubNano(nanoseconds int64) Duration
	Abs() Duration
	Sign() int
	Neg() Duration
	Cmp(o Duration) int
	State() State
	OK() bool
	String() string
} = Duration{}

func (d Duration) Seconds() (seconds int64, nano int64) {
	return d.seconds, int64(d.nano)
}

func (d Duration) Add(o Duration) (out Duration) {
	out.state = d.State() | o.State()

	nanos := int64(d.nano + o.nano)
	secs, nanos := nanos/NanosPerSecond, nanos%NanosPerSecond

	secs, state := add(secs, d.seconds)
	out.state |= state

	secs, state = add(secs, o.seconds)
	out.state |= state

	out.seconds, out.nano = secs, int(nanos)
	return
}

func (d Duration) AddNano(nanoseconds int64) (out Duration) {
	secs, nanos, _ := divFloor(nanoseconds, NanosPerSecond)
	return d.Add(Duration{seconds: secs, nano: int(nanos)})
}

func (d Duration) Sub(o Duration) (out Duration) {
	out.state = d.State() | o.State()

	nanos := int64(d.nano - o.nano)
	secs, nanos, state := divFloor(nanos, NanosPerSecond)
	out.state |= state

	secs, state = add(secs, d.seconds)
	out.state |= state

	secs, state = sub(secs, o.seconds)
	out.state |= state

	out.seconds, out.nano = secs, int(nanos)
	return
}

func (d Duration) SubNano(nanoseconds int64) (out Duration) {
	secs, nanos, _ := divFloor(nanoseconds, NanosPerSecond)
	return d.Sub(Duration{seconds: secs, nano: int(nanos)})
}

func (d Duration) Abs() Duration {
	if d.Sign() < 0 {
		return d.Neg()
	}
	return d
}

func (d Duration) Sign() int {
	if d.seconds == 0 && d.nano == 0 {
		return 0
	}
	if d.seconds >= 0 {
		return 1
	}
	return -1
}

func (d Duration) Neg() (out Duration) {
	out = d

	just, nanos, _ := divFloor(int64(NanosPerSecond-d.nano), NanosPerSecond)
	secs, state := mul(out.seconds, -1)
	out.state |= state

	if just != 1 {
		secs, state = mul(d.seconds, -1)
		out.state |= state

		secs, state = sub(secs, 1)
		out.state |= state
	}

	out.seconds, out.nano = secs, int(nanos)
	return
}

func (d Duration) Cmp(o Duration) int {
	if d.less(o) {
		return -1
	}
	if d.greater(o) {
		return 1
	}
	return 0
}

func (d Duration) less(o Duration) bool {
	if d.seconds == o.seconds {
		return d.nano < o.nano
	}
	return d.seconds < o.seconds
}

func (d Duration) greater(o Duration) bool {
	if d.seconds == o.seconds {
		return d.nano > o.nano
	}
	return d.seconds > o.seconds
}

func (d Duration) String() string {
	return fmt.Sprintf(`%d.%09d`, d.seconds, d.nano)
}

func (d Duration) State() State {
	return d.state
}

func (d Duration) OK() bool {
	return d.state == StateOK
}
