package clock

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

func System() Clock {
	return NewClock(func() Instant {
		now := time.Now()
		return Unix(now.Unix(), int64(now.Nanosecond()))
	})
}

func Fixed(fixAt Instant) Clock {
	return NewClock(func() Instant {
		return fixAt
	})
}

func Offset(original Clock, offset Duration) Clock {
	return NewClock(func() Instant {
		return original.Now().Add(offset)
	})
}
