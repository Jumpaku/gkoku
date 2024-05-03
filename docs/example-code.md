
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
		date.YyyyMmDd(2000, 1, 1),
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
		date.YyyyMmDd(2000, 1, 1),
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
	d1 := date.YyyyMmDd(2000, 1, 1)
	d2 := date.YyyyMmDd(2000, 1, 2)
	d3 := date.YyyyMmDd(2000, 1, 3)
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
	di := iter.OfDate(date.YyyyMmDd(2000, 1, 1))
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