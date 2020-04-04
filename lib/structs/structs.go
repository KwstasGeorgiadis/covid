package structs

/*
	structs that are used through the service
*/

// Countries is being used in controller/sort/sort.go,
// lib/caching/caching.go and lib/stats/stats.go
type Countries struct {
	Data []Country `json:"data"`
}

// Country is being use in lib/curve/curve.go and lib/stats/stats.go
type Country struct {
	Country            string  `json:"country"`
	Cases              int     `json:"cases"`
	TodayCases         int     `json:"todayCases"`
	Deaths             int     `json:"deaths"`
	TodayDeaths        int     `json:"todayDeaths"`
	Recovered          int     `json:"recovered"`
	Active             int     `json:"active"`
	Critical           int     `json:"critical"`
	CasesPerOneMillion float64 `json:"casesPerOneMillion"`
}

// Compare is being used in lib/curve/curve.go
type Compare struct {
	CountryOne CompareData `json:"countryOne"`
	CountryTwo CompareData `json:"countryTwo"`
}

// CompareData is being used in lib/curve/curve.go
type CompareData struct {
	Country string    `json:"country"`
	Data    []float64 `json:"data"`
}

// CountryCurve is being used in lib/curve/curve.go
type CountryCurve struct {
	Country  string         `json:"country"`
	Timeline TimelineStruct `json:"timeline"`
}

// TimelineStruct is being used in lib/curve/curve.go
type TimelineStruct struct {
	Cases     interface{} `json:"cases"`
	Deaths    interface{} `json:"deaths"`
	Recovered interface{} `json:"recovered"`
}

// CountryStats is being used in lib/curve/stats.go
type CountryStats struct {
	Country                   string `json:"country"`
	TodayPerCentOfTotalCases  int    `json:"todayPerCentOfTotalCases"`
	TodayPerCentOfTotalDeaths int    `json:"todayPerCentOfTotalDeaths"`
}

// TotalStats is being used in lib/curve/stats.go
type TotalStats struct {
	TodayPerCentOfTotalCases  int `json:"todayPerCentOfTotalCases"`
	TodayPerCentOfTotalDeaths int `json:"todayPerCentOfTotalDeaths"`
	TotalCases                int `json:"totalCases"`
	TotalDeaths               int `json:"totalDeaths"`
	TodayTotalCases           int `json:"todayTotalCases"`
	TodayTotalDeaths          int `json:"todayTotalDeaths"`
}

// AllCountriesName is being used in lib/curve/stats.go
type AllCountriesName struct {
	Countries []string `json:"countries"`
}
