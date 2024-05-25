package tokiope

import "time"

// Clock represents a source of instants.
type Clock interface {
	Now() Instant
}

// NowFunc is a function that returns an instant.
type NowFunc func() Instant

// NewClock creates a Clock from a NowFunc.
func NewClock(now NowFunc) Clock {
	return now
}

// Now returns a current instant.
func (f NowFunc) Now() Instant {
	return f()
}

// WallClock returns a Clock that returns the current time based on time.Now().
func WallClock() Clock {
	return NewClock(func() Instant {
		now := time.Now()
		return Unix(now.Unix(), int64(now.Nanosecond()))
	})
}

// FixedClock returns a Clock that always returns the instant fixAt.
func FixedClock(fixAt Instant) Clock {
	return NewClock(func() Instant {
		return fixAt
	})
}

// OffsetClock returns a Clock that returns the instant offset from the original Clock.
func OffsetClock(original Clock, offset Duration) Clock {
	return NewClock(func() Instant {
		return original.Now().Add(offset)
	})
}
