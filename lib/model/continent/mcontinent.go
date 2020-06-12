package mcontinent

//Response geeting it from get request https://corona.lmao.ninja​/v2/continents
type Response []struct {
	Updated                int64    `json:"updated"`
	Cases                  int      `json:"cases"`
	TodayCases             int      `json:"todayCases"`
	Deaths                 int      `json:"deaths"`
	TodayDeaths            int      `json:"todayDeaths"`
	Recovered              int      `json:"recovered"`
	TodayRecovered         int      `json:"todayRecovered"`
	Active                 int      `json:"active"`
	Critical               int      `json:"critical"`
	CasesPerOneMillion     float64  `json:"casesPerOneMillion"`
	DeathsPerOneMillion    float64  `json:"deathsPerOneMillion"`
	Tests                  int      `json:"tests"`
	TestsPerOneMillion     float64  `json:"testsPerOneMillion"`
	Population             int      `json:"population"`
	Continent              string   `json:"continent"`
	ActivePerOneMillion    float64  `json:"activePerOneMillion"`
	RecoveredPerOneMillion float64  `json:"recoveredPerOneMillion"`
	CriticalPerOneMillion  float64  `json:"criticalPerOneMillion"`
	Countries              []string `json:"countries"`
}
