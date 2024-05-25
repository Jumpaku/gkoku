package datetime

import (
	"fmt"
	"github.com/Jumpaku/tokiope"
	"regexp"
	"strconv"
	"strings"
)

// OffsetMinutes represents the offset from UTC in minutes.
type OffsetMinutes int

const (
	// ZeroOffsetMinutes represents the offset of UTC.
	ZeroOffsetMinutes OffsetMinutes = 0
	// MaxOffsetMinutes represents the maximum offset in minutes.
	MaxOffsetMinutes OffsetMinutes = 14 * 60
	// MinOffsetMinutes represents the minimum offset in minutes.
	MinOffsetMinutes OffsetMinutes = -14 * 60
)

// ParseOffset parses the offset from the string.
// The format of the string is either "Z" or "[+-]HH:MM".
func ParseOffset(s string) (offset OffsetMinutes, err error) {
	if !regexp.MustCompile(`^(Z|([+-]\d\d(:?\d\d)?))$`).MatchString(s) {
		return 0, fmt.Errorf(`failed to parse offset: invalid format: %q`, s)
	}
	if s == "Z" {
		return 0, nil
	}

	sign := 1
	if s[0] == '-' {
		sign = -1
	}

	s = strings.ReplaceAll(s[1:], ":", "")
	n := len(s)
	var hour, minute int

	{
		if hour, err = strconv.Atoi(s[:2]); err != nil {
			return 0, fmt.Errorf(`failed to parse offset: invalid hour: %q`, s[:2])
		}
		if hour > 18 {
			return 0, fmt.Errorf(`failed to parse offset: invalid hour: %q`, s[:2])
		}
	}
	if n > 2 {
		if minute, err = strconv.Atoi(s[2:]); err != nil {
			return 0, fmt.Errorf(`failed to parse offset: invalid minute: %q`, s[2:])
		}
		if minute >= 60 {
			return 0, fmt.Errorf(`failed to parse offset: invalid minute: %q`, s[2:])
		}
		if hour == 14 && (minute != 0) {
			return 0, fmt.Errorf(`failed to parse offset: invalid minute: %q`, s)
		}
	}

	return OffsetMinutes(sign * (hour*60 + minute)), nil
}

// FormatOffset formats the offset to the string.
func FormatOffset(offset OffsetMinutes) string {
	s := "+"
	if offset < 0 {
		s = "-"
		offset = -offset
	}

	h, m := offset/60, offset%60

	return fmt.Sprintf(`%s%02d:%02d`, s, h, m)
}

var _ interface {
	String() string
	AddTo(i tokiope.Instant) tokiope.Instant
} = OffsetMinutes(0)

// String returns the string representation of the offset.
func (o OffsetMinutes) String() string {
	return FormatOffset(o)
}

// AddTo adds the offset to the instant.
func (o OffsetMinutes) AddTo(i tokiope.Instant) tokiope.Instant {
	return i.Add(tokiope.Minutes(int64(o)))
}
