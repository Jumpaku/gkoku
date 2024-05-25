package tokiope

// Instant represents a point in time.
type Instant struct {
	unixSeconds Duration
}

// Unix returns an Instant since the Unix epoch.
func Unix(unixSeconds int64, nano int64) Instant {
	return Instant{unixSeconds: Seconds(unixSeconds, nano)}
}

var MinInstant = Instant{unixSeconds: MinDuration}
var MaxInstant = Instant{unixSeconds: MaxDuration}

var _ interface {
	Unix() (seconds int64, nano int64)
	Add(o Duration) Instant
	AddNano(nanoseconds int64) Instant
	Sub(o Duration) Instant
	SubNano(nanoseconds int64) Instant
	Diff(o Instant) Duration
	Cmp(o Instant) int
	Before(o Instant) bool
	After(o Instant) bool
	Equal(o Instant) bool
	String() string
	State() State
	OK() bool
} = Instant{}

// Unix returns the number of seconds and nanoseconds since the Unix epoch.
func (i Instant) Unix() (seconds int64, nano int64) {
	return i.unixSeconds.Seconds()
}

// Add returns the Instant going forward by the amount of duration.
func (i Instant) Add(duration Duration) Instant {
	return Instant{unixSeconds: i.unixSeconds.Add(duration)}
}

// AddNano returns the Instant going forward by the amount of nanoseconds.
func (i Instant) AddNano(nanoseconds int64) Instant {
	return Instant{unixSeconds: i.unixSeconds.AddNano(nanoseconds)}
}

// Sub returns the Instant going backward by the amount of duration.
func (i Instant) Sub(duration Duration) Instant {
	return Instant{unixSeconds: i.unixSeconds.Sub(duration)}
}

// SubNano returns the Instant going backward by the amount of nanoseconds.
func (i Instant) SubNano(nanoseconds int64) Instant {
	return Instant{unixSeconds: i.unixSeconds.SubNano(nanoseconds)}
}

// Diff returns the duration of difference from the instant.
func (i Instant) Diff(instant Instant) Duration {
	return i.unixSeconds.Sub(instant.unixSeconds)
}

// Between returns whether the instant is between the lo and hi.
func (i Instant) Between(lo, hi Instant) bool {
	return lo.Cmp(i) <= 0 && i.Cmp(hi) <= 0
}

// Cmp compares the instant with the other instant.
func (i Instant) Cmp(o Instant) int {
	return i.unixSeconds.Cmp(o.unixSeconds)
}

// Before returns whether the instant is before the other instant.
func (i Instant) Before(o Instant) bool {
	return i.unixSeconds.Cmp(o.unixSeconds) < 0
}

// After returns whether the instant is after the other instant.
func (i Instant) After(o Instant) bool {
	return i.unixSeconds.Cmp(o.unixSeconds) > 0
}

// Equal returns whether the instant is equal to the other instant.
func (i Instant) Equal(o Instant) bool {
	return i.unixSeconds.Cmp(o.unixSeconds) == 0
}

// String returns the string representation of the instant.
func (i Instant) String() string {
	return i.unixSeconds.String()
}

func (i Instant) State() State {
	return i.unixSeconds.State()
}

func (i Instant) OK() bool {
	return i.unixSeconds.OK()
}
