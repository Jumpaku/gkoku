# tokiope: An alternative Go library for basic time operations

- [Overview](#overview)
- [Getting Started](#getting-started)
    - [Installation](#installation)
    - [Time Zone Offset Transitions](#time-zone-offset-transitions)
- [Usage](#usage)
    - [Types Overview](#types-overview)
    - [API References](#api-references)
    - [Example Code](#example-code)
- [Design Policy](#design-policy)
    - [Focusing on Basic Time Operations](#focusing-on-basic-time-operations)
    - [Addressing Common Pitfalls](#addressing-common-pitfalls)

## Overview

`tokiope` is a Go library to provide basic time operations, consisting of types designed to empower developers to implement application-specific time manipulations with correctness and clarity.
The basic time operations in `tokiope` include:

- Representation of temporal values.
- Conversion of temporal values.
- Handling of time zones.
- Instant-based operations.
- Calendar-based operations.

The types in `tokiope` are designed with the following policies:

- Be well-defined.
- Be explicit.
- Be unambiguous.
- Follow standard.


## Getting Started

### Installation

```shell
go get "github.com/Jumpaku/tokiope"
```


### Time Zone Offset Transitions

`tokiope` requires data for time zone offset transitions based on the IANA timezone database to handle time zones.
This data is managed in the separate repository https://github.com/Jumpaku/tz-offset-transitions because it will be updated independently of `tokiope`.
Therefore, `tokiope` provides a CLI tool to download this data as follows:

```shell
go run "github.com/Jumpaku/tokiope/cmd/tokiope-tzot" download -out-path=tzot.json
# The above command downloads the data and saves it as a file 'tzot.json'.
```


## Usage

### Types Overview

The `tokiope` package provides basic types to handle temporal values:

- `Instant`: Represents an instantaneous point on the time-line, which is compatible with the UNIX time seconds.
- `Duration`: Represents an amount of a difference between two instants.
- `Clock`: Obtains current instants.

The tokiope/date package provides types to represent values on the calendar:

- `Date`: Represents a day on the calendar in the format of `yyyy-mm-dd`, `yyyy-Www-dd`, or `yyyy-ddd`.
- `YearMonth`: Represents a month on the calendar in the format of `yyyy-mm`.
- `YearWeek`: Represents a week on the calendar in the format of `yyyy-Www`.
- `Year`: Represents a year on the calendar in the format of `yyyy`.

The `tokiope/date/iter` package provides iterators on the calendar:

- `DateIterator`: Iterates days on the calendar.
- `YearMonthIterator`: Iterates months on the calendar.
- `YearWeekIterator`: Iterates weeks on the calendar.
- `YearIterator`: Iterates years on the calendar.

The `tokiope/datetime` package provides types to handle date-times:

- `OffsetDateTime`: Represents an instant as a combination of a date and a time with an offset.
- `OffsetMinutes`: Represents an offset from UTC in minutes.

The `tokiope/datetime/zone` package provides types to handle zoned date-times:

- `ZonedDateTime`: Represents a combination of a date and a time with a time zone.
- `Zone`: Represents a time zone that has a time zone ID and provides a mapping from instants to offsets for a specific time zone.
- `Provider`: Provides time zones according to the data for time zone offset transitions based on the IANA timezone database.


### API References

Detailed API references are available at https://pkg.go.dev/github.com/Jumpaku/tokiope .


### Example Code

Example code snippets demonstrating practical usage of `tokiope`'s functionalities are available at https://github.com/Jumpaku/tokiope/docs/examples-code.md .


## Design Policy

`tokiope`, which focuses on basic time operations, challenges to address pitfalls by providing a well-designed type framework with the following policies:

- Be well-defined: Valid calls of functions for time operations do not lead to invalid states or errors.
- Be explicit: Function behaviors depend only on object states and arguments, but not on external environments implicitly.
- Be unambiguous: Only functions that can be named so that their behavior is clear and predictable are provided.
- Follow standard: Date and time representations follow the standard ISO 8601-1:2019 https://en.wikipedia.org/wiki/ISO_8601.


### Focusing on Basic Time Operations

Since application-specific functionalities depend on individual product requirements, `tokiope` focuses on providing a solid foundation for basic time operations, which include:

- Representation of temporal values, including `Instant`, `Duration`, `Date`, `OffsetDateTime`, `ZonedDateTime`.
- Conversion of temporal values, such as conversion between `Instant` and `OffsetDateTime`.
- Handling time zones by the `tokiope/datetime/zone` package.
- Instant-based temporal operations using `Instant` and `Duration`.
- Calendar-based temporal operations using iterators in the `tokiope/date/iter` package.


### Addressing Common Pitfalls

The types in `tokiope` are designed according to the above policies, enabling developers to address the following common pitfalls in time operations:

- Implicit use of local time zones depending on environments.
- Ambiguities in calendar-based operations.
- Oversight of leap years.
- Storing date-times of past events without offsets.
- Storing timestamps of events scheduled in the future.
- Incorrect handling of daylight saving time (DST).

#### Implicit Use of Local Time Zones Depending on Environments

Implicit reliance on local time zones can lead to inconsistencies and bugs when dealing with time operations across different environments.
This often occurs when applications use the system's default time zone without explicitly specifying it.
Therefore, the time zone should always be explicitly specified to ensure consistent behavior across environments.

`tokiope` explicitly obtains `Zone` via the `zone.Provider.Get` method or `zone.Create` function.
Notably, `tokiope` does not provide an API to obtain `Zone` from the local time zone implicitly.


#### Ambiguities in Calendar-Based Operations

Ambiguities often arise in calendar-based operations, presenting challenges such as:

- Determining the date one month after 2024-01-31, even though 2024-02-31 does not exist.
- Determining the date one year after 2024-02-29, even though 2025-02-29 does not exist.
- Counting the number of months between 2024-03-30 and 2024-04-30 or 2024-03-31 and 2024-04-30.
- Counting the number of years between 2024-02-28 and 2025-02-28 or 2024-02-29 and 2025-02-28.

To avoid these confusions, APIs for calendar-based operations should be designed clearly.

`tokiope`'s iterators provided by the `tokiope/date/iter` package are specifically designed to execute calendar-based operations unambiguously.


#### Oversight of Leap Years

Leap years are frequently overlooked when calculating dates, and developers may not always understand the rules governing leap years correctly.
It is essential to use libraries implementing calendar-based operations with leap years appropriately and to test the implementation.

`tokiope` fully supports leap years and undergoes unit tests to ensure reliability.


#### Storing Date-Times of Past Events without Offsets

Date-times stored without offsets may cause difficulties when attempting to restore the corresponding timestamp because the offsets from UTC are not clear and may vary due to time zones or DST.
It is advisable to store the timestamps of past or physical events using a representation convertible to the original timestamps, such as date-times with offsets or UNIX seconds.

To represent timestamps consistently, `tokiope` offers `Instant` for UNIX seconds and `OffsetDateTime` for date-times with offsets, both of which are always convertible to timestamps.


#### Storing Timestamps of Events Scheduled in the Future

Storing a timestamp for the event scheduled on the calendar in the future as UNIX seconds or date-time with offset can lead to issues because the timestamp is not deterministic.
For instance, the timestamp corresponding to the original date-time may change from the stored timestamp due to time zone offset transitions resulting from DST or changes in political rules.
To address these concerns, it's recommended to store the date-time with a time zone to accommodate the offset transitions.

`tokiope` offers `ZonedDateTime` to represent date-times with time zones.


#### Incorrect Handling of Daylight Saving Time (DST)

Daylight Saving Time (DST) may not be handled properly, especially by inexperienced developers or those working in time zones without DST.

`tokiope` is designed to appropriately handle offset transitions, including DST.
For example:

- `func (Zone) FindOffset(at Instant) OffsetMinutes` returns the corresponding offset at a particular instant in a specific time zone.
- `func (ZonedDateTime) InstantCandidates() []Instant` returns all possible timestamps that a specific date-time may represent, where the returned timestamps may be empty due to gaps or include multiple timestamps due to overlaps.
