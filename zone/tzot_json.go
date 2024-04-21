package zone

type tzotJSON []zoneJSON
type zoneJSON struct {
	ID          string           `json:"id"`
	Transitions []transitionJSON `json:"transitions"`
	Rules       []ruleJSON       `json:"rules"`
}
type transitionJSON struct {
	OffsetSecondsBefore int    `json:"offsetSecondsBefore"`
	OffsetSecondsAfter  int    `json:"offsetSecondsAfter"`
	TransitionTimestamp string `json:"transitionTimestamp"`
}
type ruleJSON struct {
	OffsetSecondsBefore int    `json:"offsetSecondsBefore"`
	OffsetSecondsAfter  int    `json:"offsetSecondsAfter"`
	Month               int    `json:"month"`
	BaseDay             int    `json:"baseDay"`
	DayOfWeek           int    `json:"dayOfWeek"`
	OffsetTime          string `json:"offsetTime"`
}
