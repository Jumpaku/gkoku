package datetime

import (
	"fmt"
	"github.com/Jumpaku/gkoku/calendar"
	"github.com/Jumpaku/gkoku/clock"
	"github.com/Jumpaku/gkoku/internal/exact"
	"github.com/Jumpaku/gkoku/zone"
	"regexp"
	"strings"
)

type OffsetDateTime interface {
	String() string
	Offset() zone.OffsetMinutes
	Date() calendar.Date
	Time() Time
	Instant() clock.Instant
}

func NewOffsetDateTime(date calendar.Date, time Time, offset zone.OffsetMinutes) OffsetDateTime {
	return offsetDateTime{
		date:   date,
		time:   time,
		offset: offset,
	}
}

func FromInstant(at clock.Instant, offset zone.OffsetMinutes) OffsetDateTime {
	sec, nano := offset.AddTo(at).Unix()
	unixDays, secondsOfDay, _ := exact.DivFloor(sec, clock.SecondsPerDay)
	date := calendar.UnixDay(unixDays)
	time := TimeFromSeconds(int(secondsOfDay), int(nano))
	return NewOffsetDateTime(date, time, offset)
}

func ParseOffsetDateTime(s string) (d OffsetDateTime, err error) {
	if !regexp.MustCompile(`^.*T.*(Z|([-+].*))$`).MatchString(s) {
		return nil, fmt.Errorf(`failed to parse offset datetime: invalid format: %q`, s)
	}

	arr := strings.Split(s, "T")

	date, err := calendar.ParseDate(arr[0], calendar.DateFormatAny)
	if err != nil {
		return nil, fmt.Errorf(`failed to parse offset datetime: invalid date: %w`, err)
	}

	s = arr[1]

	so := regexp.MustCompile(`[-+Z].*$`).FindString(s)
	offset, err := zone.ParseOffset(so)
	if err != nil {
		return nil, fmt.Errorf(`failed to parse offset datetime: invalid offset: %w`, err)
	}

	time, err := ParseTime(s[:len(s)-len(so)])
	if err != nil {
		return nil, fmt.Errorf(`failed to parse offset datetime: invalid time: %w`, err)
	}

	return NewOffsetDateTime(date, time, offset), nil
}

func FormatOffsetDateTime(d OffsetDateTime) string {
	return fmt.Sprintf(`%s%s%s`, d.Date().String(), d.Time().String(), d.Offset().String())
}

type offsetDateTime struct {
	date   calendar.Date
	time   Time
	offset zone.OffsetMinutes
}

func (d offsetDateTime) String() string {
	return FormatOffsetDateTime(d)
}

func (d offsetDateTime) Offset() zone.OffsetMinutes {
	return d.offset
}

func (d offsetDateTime) Date() calendar.Date {
	return d.date
}

func (d offsetDateTime) Time() Time {
	return d.time
}

func (d offsetDateTime) Instant() clock.Instant {
	t := d.Time()
	o, h, m, s, n := int64(d.Offset()), int64(t.Hour()), int64(t.Minute()), int64(t.Second()), int64(t.Nano())
	secondsOfDay := clock.Hours(h).Add(clock.Minutes(m)).Add(clock.Seconds(s, n))

	offset := clock.Minutes(o)

	unixDays := d.Date().UnixDay()

	return clock.Unix(clock.Days(unixDays).Add(secondsOfDay).Sub(offset).Seconds())
}
