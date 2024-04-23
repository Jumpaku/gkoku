package datetime

import (
	"fmt"
	"github.com/Jumpaku/tokiope"
	"github.com/Jumpaku/tokiope/date"
	"github.com/Jumpaku/tokiope/internal/exact"
	"regexp"
	"strings"
)

type OffsetDateTime interface {
	String() string
	Offset() OffsetMinutes
	Date() date.Date
	Time() Time
	Instant() tokiope.Instant
}

func NewOffsetDateTime(date date.Date, time Time, offset OffsetMinutes) OffsetDateTime {
	return offsetDateTime{
		date:   date,
		time:   time,
		offset: offset,
	}
}

func FromInstant(at tokiope.Instant, offset OffsetMinutes) OffsetDateTime {
	sec, nano := offset.AddTo(at).Unix()
	unixDays, secondsOfDay, _ := exact.DivFloor(sec, tokiope.SecondsPerDay)
	date := date.UnixDay(unixDays)
	time := TimeFromSeconds(int(secondsOfDay), int(nano))
	return NewOffsetDateTime(date, time, offset)
}

func ParseOffsetDateTime(s string) (d OffsetDateTime, err error) {
	if !regexp.MustCompile(`^.*T.*(Z|([-+].*))$`).MatchString(s) {
		return nil, fmt.Errorf(`failed to parse offset datetime: invalid format: %q`, s)
	}

	arr := strings.Split(s, "T")

	date, err := date.ParseDate(arr[0], date.DateFormatAny)
	if err != nil {
		return nil, fmt.Errorf(`failed to parse offset datetime: invalid date: %w`, err)
	}

	s = arr[1]

	so := regexp.MustCompile(`[-+Z].*$`).FindString(s)
	offset, err := ParseOffset(so)
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
	date   date.Date
	time   Time
	offset OffsetMinutes
}

func (d offsetDateTime) String() string {
	return FormatOffsetDateTime(d)
}

func (d offsetDateTime) Offset() OffsetMinutes {
	return d.offset
}

func (d offsetDateTime) Date() date.Date {
	return d.date
}

func (d offsetDateTime) Time() Time {
	return d.time
}

func (d offsetDateTime) Instant() tokiope.Instant {
	t := d.Time()
	o, h, m, s, n := int64(d.Offset()), int64(t.Hour()), int64(t.Minute()), int64(t.Second()), int64(t.Nano())
	secondsOfDay := tokiope.Hours(h).Add(tokiope.Minutes(m)).Add(tokiope.Seconds(s, n))

	offset := tokiope.Minutes(o)

	unixDays := d.Date().UnixDay()

	return tokiope.Unix(tokiope.Days(unixDays).Add(secondsOfDay).Sub(offset).Seconds())
}
