package tokiope

import "time"

type Clock interface {
	Now() Instant
}

type NowFunc func() Instant

func NewClock(now NowFunc) Clock {
	return now
}

func (f NowFunc) Now() Instant {
	return f()
}

func WallClock() Clock {
	return NewClock(func() Instant {
		now := time.Now()
		return Unix(now.Unix(), int64(now.Nanosecond()))
	})
}

func FixedClock(fixAt Instant) Clock {
	return NewClock(func() Instant {
		return fixAt
	})
}

func OffsetClock(original Clock, offset Duration) Clock {
	return NewClock(func() Instant {
		return original.Now().Add(offset)
	})
}
