package tokiope

func NGDuration() Duration {
	return Duration{state: StateOverflow}
}

func NGInstant() Instant {
	return Instant{unixSeconds: NGDuration()}
}
