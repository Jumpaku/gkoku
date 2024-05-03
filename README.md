# tokiope: An alternative Go library for basic time operations

- [Overview](#overview)
- [Getting Started](#getting-started)
    - [Installation](#installation)
    - [Timezone offset transitions](#timezone-offset-transitions)
- [Usage](#usage)
    - [Types Overview](#types-overview)
    - [API References](#api-references)
    - [Example Code](#example-code)
- [Design Policy](#design-policy)
    - [Focussing on basic time operations](#focussing-on-basic-time-operations)
    - [Addressing common pitfalls](#addressing-common-pitfalls)
        - [Implicit use of local timezones depending on environments](#implicit-use-of-local-timezones-depending-on-environments)
        - [Ambiguous calendar-based operations](#ambiguous-calendar-based-operations)
        - [Oversight of leap years](#oversight-of-leap-years)
        - [Storing datetimes without offset](#storing-datetimes-without-offset)
        - [Storing timestamps in the future](#storing-timestamps-in-the-future)
        - [Incorrect handling daylight saving time (DST)](#incorrect-handling-daylight-saving-time-dst)

## Overview

`tokiope` is a Go library to provide basic time operations, which includes types designed to empower developers to implement application-specific time manipulations with correctness and clarity.
The basic time operations in `tokiope` include:

- Representation of temporal values.
- Conversion of temporal values.
- Handling of timezones.
- Instant-based operations.
- Calendar-based operations.

The types in `tokiope` are designed with the following policy:

- Be well-defined.
- Be explicit.
- Be unambiguous.
- Follow standard.


## Getting started

### Installation

```shell
go get "github.com/Jumpaku/tokiope"
```


### Timezone offset transitions

`tokiope` requires the data for timezone offset transitions based on the IANA timezone database to handle timezones.
This data is managed in the separated repository https://github.com/Jumpaku/tz-offset-transitions because it will be updated independently of `tokiope`.
Therefore, `tokiope` provides a CLI tool to download this data as follows:

```shell
go run "github.com/Jumpaku/tokiope/cmd/tokiope-tzot" download -out-path=tzot.json
# The above command downloads the data and save it as a file 'tzot.json'.
```


## Usage

### Types overview

The `tokiope` package provides basic types to handle temporal values.

- `Instant`: Represents a point on the time series, which is compatible with the UNIX time seconds.
- `Duration`: Represents an amount of a difference between two instants.
- `Clock`: Obtains current instants.

The `tokiope/date` package provides types to represent values on the calendar.

- `Date`: Represents a day on the calendar in the format of `yyyy-mm-dd`, `yyyy-Www-dd`, or `yyyy-ddd`.
- `YearMonth`: Represents a month on the calendar in the format of `yyyy-mm`.
- `YearWeek`: Represents a week on the calendar in the format of `yyyy-Www`.
- `Year`: Represents a year on the calendar in the format of `yyyy`.

The `tokiope/date/iter` package provides iterators on the calendar.

- `DateIterator`: Iterates days on the calendar.
- `YearMonthIterator`: Iterates months on the calendar.
- `YearWeekIterator`: Iterates weeks on the calendar.
- `YearIterator`: Iterates years on the calendar.

The `tokiope/datetime` package provides types to handle datetimes.

- `OffsetDateTime`: Represents an instant as a combination of a date and a time with an offset.
- `OffsetMinutes`: Represents an offset from UTC in minutes.

The `tokiope/datetime/zone` package provides types to handle zoned datetimes.

- `ZonedDateTime`: Represents a combination of a date and a time with a timezone.
- `Zone`: Represents a timezone that has a timezone ID and provides a mapping from instants to offsets for a specific timezone.
- `Provider`: Provides timezones based on the information of the IANA timezone database.


### API references

Detailed API references are available at https://pkg.go.dev/github.com/Jumpaku/tokiope .


### Example code

Example code snippets demonstrating practical usage of `tokiope`'s functionalities are available at https://github.com/Jumpaku/tokiope/docs/examples-code.md .


## Design policy

`tokiope` focusing on basic time operations challenges to address to pitfalls by providing well-designed type framework with the following policy:

- Be well-defined: Valid calls of functions of time operations does not lead to invalid state or errors.
- Be explicit: Function behaviors depend on only object states and arguments but not external environments implicitly. 
- Be unambiguous: Only functions that can be named so that their behavior is clear and predictable are provided.
- Follow standard: Date and time representations follow the standard ISO 8601-1:2019 https://en.wikipedia.org/wiki/ISO_8601.


### Focussing on basic time operations

Since application-specific functionalities will depend on individual product requirements, `tokiope` focuses on providing a solid foundation for basic time operations which include:

- representation of temporal values including `Instant`, `Duration`, `Date`, `OffsetDateTime`, `ZonedDateTime`.
- conversion of temporal values. For example, conversion between `Instant` and `OffsetDateTime`.
- handling timezones by `tokiope/datetime/zone` package.
- instant-based temporal operations by `Instant` and `Duration`.
- calendar-based temporal operations by iterators in `tokiope/date/iter` package.


### Addressing common pitfalls

`tokiope`'s types are designed according to the above policy, which can address the following common pitfalls about time operations:

- Implicit use of local timezones that depends on environments
- Ambiguous calendar-based operations.
- Ignoring leap years on calendar-based operations.
- Saving datetimes of events that occurred in the past without offsets at those instants.
- Saving timestamps of events that scheduled in the future.
- Incorrect handling daylight saving time (DST).


#### Implicit use of local timezones depending on environments

Implicit reliance on local timezones can lead to inconsistencies and bugs when dealing with time operations across different environments.
This often occurs when applications use the system's default timezone without explicitly specifying it.
The timezone should be always explicitly specified to ensure consistent behavior across environments.
`tokiope` explicitly obtains `Zone` via `zone.Provider.Get` method or `zone.Create` function while `tokiope` has no API to obtain `Zone` from the local timezone implicitly.


#### Ambiguous calendar-based operations

Ambiguities arise in calendar-based operations, for example:

- What date is one month after 2024-01-31 although 2024-02-31 does not exist?
- What date is one year after 2024-02-29 although 2025-02-29 does not exist?
- How many months are there between 2024-03-30 and 2024-04-30 or 2024-03-31 and 2024-04-30?.
- How many years are there between 2024-02-28 and 2025-02-28 or 2024-02-29 and 2025-02-28?.

To avoid the confusion caused by the ambiguities in date calculations, API to provide calendar-based operations should be well-defined with consistency.
`tokiope` is based on the ISO calendar system, which is the de facto world calendar following the proleptic Gregorian rules.
In addition, the iterators on the calendar provided by `tokiope/date/iter` package are designed to realize calendar-based operations by only unambiguous operations.


#### Oversight of leap years

The existence of leap years is often overlooked when calculating dates.
Also, the rules of the leap year understood by developers may be not correct.
You should utilize libraries to implement calendar-based operations to handle leap years appropriately and test the implementation.
`tokiope` supports leap years and is fully tested by unit tests.


#### Storing datetimes without offset

Datetime stored without offset may cause difficulties in restoring the corresponding timestamp because the offset from UTC is not clear.
You should store a timestamp of the past or physical event as a representation convertible to the original timestamp such as datetime with offset or UNIX seconds.
To represent timestamps, `tokiope` provides `Instant` for UNIX seconds and `OffsetDateTime` for datetime with offset both of which are always convertible to each other.


#### Storing timestamps in the future

Storing a timestamp for the event scheduled on the calendar in the future as UNIX seconds or datetime with offset may cause problems because the timestamp is not deterministic.
For example, the timestamp corresponding to the original datetime may differ from the stored timestamp when the timezone offset transitions due to daylight saving time or political changes in rules.
You can store the datetime with a timezone to respond to the offset transitions.
`tokiope` provides `ZonedDateTime` to represent datetime with timezone.


#### Incorrect handling daylight saving time (DST)

DST may not be aware and handled inappropriately if a developer is a beginner or works in a timezone without DST.
`tokiope` supports to appropriately handle offset transitions including DST.
For example:

- `func (Zone) FindOffset(at Instant) OffsetMinutes` returns an offset in a specific timezone at a specific instant.
- `func (ZonedDateTime) InstantCandidates() []Instant` returns all possible timestamps that a specific datetime may represent. The returned timestamps may be empty by gaps or multiple timestamps by overlaps.

