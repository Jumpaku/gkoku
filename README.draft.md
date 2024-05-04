# tokiope

An alternative Go library for basic time operations.


- [Usage](#usage)
  - [Types overview](#types-overview)
  - [Conversions](#conversions)
  - [Instant-based operations](#instant-based-operations)
  - [Calendar-based operations](#calendar-based-operations)
  - [Detailed documentation](#detailed-documentation)
  - [Installation](#installation)
  - [Timezone offset transitions](#timezone-offset-transitions)
- [Motivation](#motivation)
  - [Example mistakes to be avoided](#example-mistakes-to-be-avoided)

## Usage

### Types overview

- package `tokiope` provides basic types to handle temporal values.
  - `Instant` represents a point on the time series, which is compatible with the UNIX time seconds.
  - `Duration` represents an amount of a difference between two instants.
  - `Clock` obtains current instants.

- package `tokiope/date` provides types to represent values on the calendar.
  - `Date` represents a day on the calendar in the format of `yyyy-mm-dd`, `yyyy-Www-dd`, or `yyyy-ddd`.
  - `YearMonth` represents a month on the calendar in the format of `yyyy-mm`.
  - `YearWeek` represents a week on the calendar in the format of `yyyy-Www`.
  - `Year` represents a year on the calendar in the format of `yyyy`.

- package `tokiope/date/iter` provides iterators on the calendar.
  - `DateIterator` iterates days on the calendar.
  - `YearMonthIterator` iterates months on the calendar.
  - `YearWeekIterator` iterates weeks on the calendar.
  - `YearIterator` iterates years on the calendar.

- package `tokiope/datetime` provides types to handle datetimes.
  - `OffsetDateTime` represents an instant as a combination of a date and a time with an offset.
  - `OffsetMinutes` represents an offset from UTC in minutes.

- package `tokiope/datetime/zone` provides types to handle zoned datetimes.
  - `ZonedDateTime` represents a combination a date and a time with a timezone.
  - `Zone` represents a timezone that has a timezone ID and is a mapping from instants to offsets.
  - `Provider` provides timezones based on the information of the IANA timezone database.


### Conversions

#### Convert `Instant` to `OffsetDateTime`

```go
	t := tokiope.Unix(946684800, 0) // 2000-01-01T00:00:00+09:00
	od := datetime.FromInstant(t, datetime.OffsetMinutes(9*60))
	fmt.Println(od)
	fmt.Println(od.Date())
	fmt.Println(od.Time())
	fmt.Println(od.Offset())
	// Output:
	// 2000-01-01T09:00:00.000000000+09:00
	// 2000-01-01
	// T09:00:00.000000000
	// +09:00
```

#### Convert `OffsetDateTime` to `Instant`

```go
	od := datetime.NewOffsetDateTime(
		date.DateOfYMD(2000, 1, 1),
		datetime.TimeOf(9, 0, 0, 0),
		datetime.OffsetMinutes(9*60),
	) // 2000-01-01T09:00:00+09:00
	t := od.Instant()
	fmt.Println(t)
	// Output: 946684800.000000000
```

#### Convert `ZonedDateTime` to `Instant`

```go
	zd := zone.NewZonedDateTime(
		date.DateOfYMD(2000, 1, 1),
		datetime.TimeOf(9, 0, 0, 0),
		zone.CreateFixed("Asia/Tokyo", datetime.OffsetMinutes(9*60)),
	) // 2000-01-01T09:00:00[Asia/Tokyo]
	ts := zd.InstantCandidates()
	fmt.Println(ts)
	// Output: [946684800.000000000]
```

### Instant-based operations

```go
	t1 := tokiope.Unix(1, 0)
	t2 := tokiope.Unix(2, 0)
	t3 := tokiope.Unix(3, 0)
	d := tokiope.Seconds(10, 0)
	fmt.Println(t2.Diff(t1))
	fmt.Println(t3.Add(d))
	fmt.Println(t3.Sub(d))
	// Output:
	// 1.000000000
	// 13.000000000
	// -7.000000000
```

### Calendar-based operations

#### Calculate dates

```go
	d1 := date.DateOfYMD(2000, 1, 1)
	d2 := date.DateOfYMD(2000, 1, 2)
	d3 := date.DateOfYMD(2000, 1, 3)
	days := 10
	fmt.Println(d1.DaysUntil(d2))
	fmt.Println(d3.Add(days))
	fmt.Println(d3.Sub(days))
	// Output:
	// 1
	// 2000-01-13
	// 1999-12-24
```

#### Iterate on the calendar

```go
	di := iter.OfDate(date.DateOfYMD(2000, 1, 1))
	di.Move(1)
	fmt.Println(di.Get())

	ymi := iter.OfYearMonth(date.YearMonthOf(2000, 1))
	ymi.Move(2)
	fmt.Println(ymi.Get())

	ywi := iter.OfYearWeek(date.YearWeekOf(2000, 1))
	ywi.Move(3)
	fmt.Println(ywi.Get())

	yi := iter.OfYear(date.Year(2000))
	yi.Move(4)
	fmt.Println(yi.Get())

	// Output:
	// 2000-01-01
	// 2000-02
	// 2000-W01
	// 2004
```

### Detailed documentation


### Installation

```shell
go get github.com/Jumpaku/tokiope
```

### Timezone offset transitions

`tokiope` needs information with respect to timezone offset transitions based on the IANA timezone database to handle timezones.
The data of the information is available on https://github.com/Jumpaku/tz-offset-transitions .


## Motivation

`tokiope` provides a typed framework that is designed to avoid mistakes on time operations as follows:

- Implicit use of local timezones that depends on environments
- Ambiguous calendar-based operations.
- Ignoring leap years on calendar-based operations.
- Saving datetimes of events that occurred in the past without offsets at those instants.
- Saving timestamps of events that scheduled in the future.
- Handling daylight saving time (DST) incorrectly.

### Example mistakes to be avoided

#### Implicit use of local timezones depending on environments

Implicit reliance on local timezones can lead to inconsistencies and errors when dealing with time-related operations across different environments.
This often occurs when applications use the system's default timezone without explicitly specifying it.
The timezone should be always explicitly specified to ensure consistent behavior across environments.
`tokiope` has no API to obtain `Zone` from local timezone and obtains `Zone` via `zone.Provider.Get` method or `zone.Create` function explicitly.  

#### Ambiguous calendar-based operations

Ambiguities arise in calendar-based operations, for example:

- What date is one month after 2024-01-31 although 2024-02-31 does not exist?
- What date is one year after 2024-02-29 although 2025-02-29 does not exist?
- How many months are there between 2024-03-30 and 2024-04-30 or 2024-03-31 and 2024-04-30?.
- How many years are there between 2024-02-28 and 2025-02-28 or 2024-02-29 and 2025-02-28?.

To avoid the confusions caused by the ambiguities in date calculations and interpretations, API to provide calendar-based operations should be well-desigbed with consistency.
`tokiope` is based on the ISO calendar system, which is the de facto world calendar following the proleptic Gregorian rules.
In addition, the iterators on the calendar provided by `tokiope/date/iter` package are designed to realize calendar-based operations by only unambiguous operations.

#### Oversight of leap years

The existence of leap years is often overlooked when calculating dates.
Also, the rules of the leap year understood by developers may be not correct.
You should utilize libraries to implement calender-based operations to handle leap years appropriately and test the implementation.
`tokiope` supports leap years and is fully-tested by unit tests.

#### Storing datetimes without offset

Datetime stored without offset may cause difficulties to restore the corresponding timestamp because the offset from UTC is not clear.
You should store timestamp of the past or physical event as a representation convertible to original timestamp such as datetime with offset or UNIX seconds.
To represent timestamps, `tokiope` provides `Instant` for UNIX seconds and `OffsetDateTime` for datetime with offset, which are always uniquely convertible each other.

#### Storing timestamps in the future

Storing timestamp for the event scheduled on the calendar in the future as UNIX seconds or datetime with offset may cause problems because the timestamp is not deterministic.
For example, the timestamp corresponding to the original datetime may differ from the stored timestamp when the timezone offset transitions due to daylight saving time or political changes of rules.
You can store the datetime with a timezone to respond to the offset transitions.
`tokiope` provides `ZonedDateTime` to represent datetime with timezone.

#### Handling daylight saving time (DST) incorrectly

DST may not be aware and handled inappropriately if a developer is a beginner or works in a  timezone without DST.
`tokiope` supports to appropriately handle offset transitions including DST.
For example:

- `func (Zone) FindOffset(at Instant) OffsetMinutes` returns an offset in a specific timezone at a specific instant.
- `func (ZonedDateTime) InstantCandidates() []Instant` returns all possible timestamps that a specific datetime may represent. The returned timestamps may be empty by gaps or multiple timestamps by overlaps.

