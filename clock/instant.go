package clock

type Instant struct {
	unixSeconds Duration
}

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

func (i Instant) Unix() (seconds int64, nano int64) {
	return i.unixSeconds.Seconds()
}

func (i Instant) Add(o Duration) Instant {
	return Instant{unixSeconds: i.unixSeconds.Add(o)}
}

func (i Instant) AddNano(nanoseconds int64) Instant {
	return Instant{unixSeconds: i.unixSeconds.AddNano(nanoseconds)}
}

func (i Instant) Sub(o Duration) Instant {
	return Instant{unixSeconds: i.unixSeconds.Sub(o)}
}

func (i Instant) SubNano(nanoseconds int64) Instant {
	return Instant{unixSeconds: i.unixSeconds.SubNano(nanoseconds)}
}

func (i Instant) Diff(o Instant) Duration {
	return i.unixSeconds.Sub(o.unixSeconds)
}

func (i Instant) Cmp(o Instant) int {
	return i.unixSeconds.Cmp(o.unixSeconds)
}

func (i Instant) Before(o Instant) bool {
	return i.unixSeconds.Cmp(o.unixSeconds) < 0
}

func (i Instant) After(o Instant) bool {
	return i.unixSeconds.Cmp(o.unixSeconds) > 0
}

func (i Instant) Equal(o Instant) bool {
	return i.unixSeconds.Cmp(o.unixSeconds) == 0
}

func (i Instant) String() string {
	return i.unixSeconds.String()
}

func (i Instant) State() State {
	return i.unixSeconds.State()
}

func (i Instant) OK() bool {
	return i.unixSeconds.OK()
}
