package datetime

import (
	"fmt"
	"github.com/Jumpaku/go-assert"
	"github.com/Jumpaku/tokiope"
	"regexp"
	"strconv"
	"strings"
)

// Time represents a time of day.
type Time struct {
	hour   int
	minute int
	second int
	nano   int
}

// TimeOf creates a Time from the hour, minute, second, and nano.
func TimeOf(hour, minute, second, nano int) Time {
	assert.Params(0 <= hour && hour <= 24, "hour must be in [0,24]: %d", hour)
	assert.Params(0 <= minute && minute < 60, "minute must be in [0,60): %d", minute)
	assert.Params(0 <= second && second < 60, "second must be in [0,60): %d", second)
	assert.Params(0 <= nano && nano < 1_000_000_000, "nano must be in [0,1_000_000_000): %d", nano)
	if hour == 24 {
		assert.Params(minute == 0, "minute must be 0 if hour is 24: %d", minute)
		assert.Params(second == 0, "second must be 0 if hour is 24: %d", second)
		assert.Params(nano == 0, "nano must be 0 if hour is 24: %d", nano)
	}
	return Time{hour: hour, minute: minute, second: second, nano: nano}
}

// TimeFromSeconds creates a Time from the seconds of day and nano.
// It panics if the secondsOfDay and nano represents an invalid time.
func TimeFromSeconds(secondsOfDay, nano int) Time {
	assert.Params(0 <= secondsOfDay && secondsOfDay <= tokiope.SecondsPerDay, "secondsOfDay must be in [0,%d]: %d", tokiope.SecondsPerDay, secondsOfDay)
	assert.Params(0 <= nano && nano < 1_000_000_000, "nano must be in [0,1_000_000_000): %d", nano)
	if secondsOfDay == tokiope.SecondsPerDay {
		assert.Params(nano == 0, "nano must be 0 if hour is 24: %d", nano)
	}
	wholeMinutes, second := secondsOfDay/tokiope.SecondsPerMinute, secondsOfDay%tokiope.SecondsPerMinute
	hour, minute := wholeMinutes/tokiope.MinutesPerHour, wholeMinutes%tokiope.MinutesPerHour
	return Time{hour: hour, minute: minute, second: second, nano: nano}
}

// ParseTime parses the Time from the string.
// The format of the string is "T?hh:mm:ss[.SSSSSSSSS]".
func ParseTime(s string) (Time, error) {
	if !regexp.MustCompile(`^T?\d\d(:\d\d(:\d\d([,.]\d{1,9})?)?)?$`).MatchString(s) {
		return Time{}, fmt.Errorf("fail to parse Time: invalid format: %q", s)
	}
	s = strings.ReplaceAll(s, "T", "")
	s = strings.ReplaceAll(s, ".", ":")
	s = strings.ReplaceAll(s, ",", ":")
	arr := strings.Split(s, ":")
	n := len(arr)
	var hour, minute, second, nano int

	hour, err := strconv.Atoi(arr[0])
	if err != nil {
		return Time{}, fmt.Errorf("fail to parse Time: invalid hour format: %q", arr[0])
	}
	if hour < 0 || hour > 24 {
		return Time{}, fmt.Errorf("fail to parse Time: hour must be in [0,24]: %d", hour)
	}

	if n > 1 {
		minute, err = strconv.Atoi(arr[1])
		if err != nil {
			return Time{}, fmt.Errorf("fail to parse Time: invalid minute format: %q", arr[1])
		}
		if minute < 0 || minute >= 60 {
			return Time{}, fmt.Errorf("fail to parse Time: minute must be in [0,60): %d", minute)
		}
	}

	if n > 2 {
		second, err = strconv.Atoi(arr[2])
		if err != nil {
			return Time{}, fmt.Errorf("fail to parse Time: invalid second format: %q", arr[2])
		}
		if second < 0 || second >= 60 {
			return Time{}, fmt.Errorf("fail to parse Time: second must be in [0,60): %d", second)
		}
	}

	if n > 3 {
		nano, err = strconv.Atoi((arr[3] + "000000000")[:9])
		if err != nil {
			return Time{}, fmt.Errorf("fail to parse Time: invalid nano format: %q", arr[3])
		}
		if nano < 0 || nano >= 1_000_000_000 {
			return Time{}, fmt.Errorf("fail to parse Time: nano must be in [0,60): %d", nano)
		}
	}

	if (hour == 24) && (minute != 0 || second != 0 || nano != 0) {
		return Time{}, fmt.Errorf(
			"fail to parse Time: minute, second, and nano must be 0 if hour is 24: %02d:%02d:%02d.%09d",
			hour, minute, second, nano)
	}

	return TimeOf(hour, minute, second, nano), nil
}

// FormatTime formats the Time to a string.
func FormatTime(t Time) string {
	return fmt.Sprintf(`T%02d:%02d:%02d.%09d`, t.Hour(), t.Minute(), t.Second(), t.Nano())
}

var _ interface {
	Hour() int
	Minute() int
	Second() int
	Nano() int
	String() string
} = Time{}

// Hour returns the hour of the Time.
func (t Time) Hour() int {
	return t.hour
}

// Minute returns the minute of the Time.
func (t Time) Minute() int {
	return t.minute
}

// Second returns the second of the Time.
func (t Time) Second() int {
	return t.second
}

// Nano returns the nano of the Time.
func (t Time) Nano() int {
	return t.nano
}

// String returns the string representation of the Time.
func (t Time) String() string {
	return FormatTime(t)
}
