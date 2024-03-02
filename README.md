# gkoku

gkoku (gee-koh-koo) is a Go library that provides basic time operations.

## Motivation

This library wraps the standard `time` package and provides a type safe API to handle date and time.

## Definition

- A timestamp is a point on the time axis.
- A duration is the amount of the difference between two timestamps.

- A date is a representation of a day according to a calendar, which consists of year, month, and day.
- A time of day is a combination of hour, minute, and second.
- A datetime is a combination of a date and a time of day.

- A timezone is a collection of offsets assigned based on region.
- An offset is the amount of difference from a datetime in UTC to a datetime of a region.

## Mainly used types

`Instant` represents a timestamp based on Unix time.
It can be used for the following usecases:

1. Events that occurred in the past.
2. Physical events that occurred in the past or will occur in the future.

`OffsetDateTime` represents a timestamp using a date and an offset.
It is useful for the datetime operations according to the calendar.

`ZonedDateTime` represents a datetime (the timestamp may not be specific) with a timezone.
It can be used to specify datetime (according to the calendar) of events scheduled in the future at specific regions.
Note that you should specify a datetime in the future with not an offset but a timezone because offset in the future may vary due to changes of rules.

`Offset` represents a offset.

`Zone` represents a timezone. Specified a timestamp, `Zone` obtains `Offset` at the timestamp.

`Duration` represents the duration between two `Instance`s.

`DateTimePeriod` represents the amount of the difference between two datetime.

## Conversion

`OffsetDateTime` can always be converted to `Instant`.
With an offset, `Instant` can be converted to `OffsetDateTime`.

With a timestamp, `ZonedDateTime` can be converted to `Instant` using obtained offset at the timestamp.
With a timezone, `Instant` can be converted to `ZonedDateTime`.
