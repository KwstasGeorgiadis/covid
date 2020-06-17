package mworld

type WorldTimeline struct {
	Cases          interface{} `json:"cases"`
	Deaths         interface{} `json:"deaths"`
	Recovered      interface{} `json:"recovered"`
	CasesDaily     interface{} `json:"casesDaily"`
	DeathsDaily    interface{} `json:"deathsDaily"`
	RecoveredDaily interface{} `json:"recoveredDaily"`
}
