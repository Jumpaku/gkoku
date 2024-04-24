# tokiope

An alternative Go library for basic time operations.


## Types overview

- package `tokiope` provides basic types to handle temporal values.
    - `Instant` represents a point on the time series, which is compatible with the UNIX time seconds.
    - `Duration` represents an amount of a difference between two instants.
    - `Clock` obtains current instants.

- package `tokiope/date` provides types to represent calendrical values.
    - `Date` represents a day on the calendar in the format of `yyyy-mm-dd`, `yyyy-Www-dd`, or `yyyy-ddd`.
    - `YearMonth` represents a month on the calendar in the format of `yyyy-mm`.
    - `YearWeek` represents a week on the calendar in the format of `yyyy-Www`.
    - `Year` represents a year on the calendar in the format of `yyyy`.

- package `tokiope/date/iter` provides iterators to iterate on calendar.
    - `DateIterator` iterates days on the calendar.
    - `YearMonthIterator` iterates months on the calendar.
    - `YearWeekIterator` iterates weeks on the calendar.
    - `YearIterator` iterates years on the calendar.

- package `tokiope/datetime` provides types to represent a date and a time.
    - `OffsetDateTime` represents an instant as a combination of date and time with offset.

- package `tokiope/datetime/zone`
    - `ZonedDateTime` represents a combination a date and a time with a timezone.
    - `Zone` represents a timezone that is a mapping from instants to offsets and has a timezone ID.
    - `Provider` provides timezones based on the information of the IANA timezone database.


## Motivation

`tokiope` provides a typed framework that is designed to avoid mistakes on time operations:

### Examples of mistakes

- Implicit use of local timezones that depends on environments.
- Ambiguous calendrical operations.
- Ignoring leap years on calendrical operations.
- Saving datetimes of events that occurred in the past without offsets at those instants.
- Saving timestamps of events that scheduled in the future.
- Ignoring daylight saving times on conversions from zoned datetimes to timestamps.


## Mainly used types

`Instant` represents a timestamp based on Unix time.
It can be used for the following usecases:

1. Events that occurred in the past.
2. Physical events that occurred in the past or will occur in the future.

`OffsetDateTime` represents a timestamp using a date and an offset.
It is useful for the datetime operations according to the calendar.

`ZonedDateTime` represents a datetime (the timestamp may not be specific) with a timezone.
It can be used to specify datetime (according to the calendar) of events scheduled in the future at specific regions.
Note that you should specify a datetime in the future with not an offset but a timezone because offset in the future may
vary due to changes of rules.

`Offset` represents a offset.

`Zone` represents a timezone. Specified a timestamp, `Zone` obtains `Offset` at the timestamp.

`Duration` represents the duration between two `Instance`s.

`DateTimePeriod` represents the amount of the difference between two datetime.


## Operations

### Instant-based operations



### Calendrical operations



### Conversions

`OffsetDateTime` can always be converted to `Instant`.

With an offset, `Instant` can always be converted to `OffsetDateTime`.
