package structs

import (
	"encoding/xml"
)

/*
	structs that are used across the service
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
	Test               int     `json:"tests"`
	TestPerOneMillion  int     `json:"testsPerOneMillion"`
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

// CompareAll is being used in lib/curve/curve.go
type CompareAll struct {
	CountryOne CompareAllData `json:"countryOne"`
	CountryTwo CompareAllData `json:"countryTwo"`
}

// CompareAllData is being used in lib/curve/curve.go
type CompareAllData struct {
	Country             string    `json:"country"`
	DataDeaths          []float64 `json:"dataDeaths"`
	DataDeathsFromFirst []float64 `json:"dataDeathsFromFirst"`
	DataDeathsPerDay    []float64 `json:"dataDeathsPerDay"`
	DataRecovered       []float64 `json:"dataRecoverd"`
	DataCases           []float64 `json:"dataCases"`
	DataCasesFromFist   []float64 `json:"dataCasesFromFirst"`
}

// CountryCurve is being used in lib/curve/curve.go
type CountryCurve struct {
	Country  string         `json:"country"`
	Timeline TimelineStruct `json:"timeline"`
	Province string         `json:"province"`
}

// TimelineStruct is being used in lib/curve/curve.go
type TimelineStruct struct {
	Cases     interface{} `json:"cases"`
	Deaths    interface{} `json:"deaths"`
	Recovered interface{} `json:"recovered"`
}

type WorldTimeline struct {
	Cases          interface{} `json:"cases"`
	Deaths         interface{} `json:"deaths"`
	Recovered      interface{} `json:"recovered"`
	CasesDaily     interface{} `json:"casesDaily"`
	DeathsDaily    interface{} `json:"deathsDaily"`
	RecoveredDaily interface{} `json:"recoveredDaily"`
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

// ErrorMessage is being used across all controllers
type ErrorMessage struct {
	ErrorMessage string `json:"message"`
	Code         int    `json:"code"`
}

// ReponseNews response we are getting for the news api, being used in lib/news/news.go
type ReponseNews struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Media   string   `xml:"media,attr"`
	Channel struct {
		Text          string `xml:",chardata"`
		Generator     string `xml:"generator"`
		Title         string `xml:"title"`
		Link          string `xml:"link"`
		Language      string `xml:"language"`
		WebMaster     string `xml:"webMaster"`
		Copyright     string `xml:"copyright"`
		LastBuildDate string `xml:"lastBuildDate"`
		Description   string `xml:"description"`
		Item          []struct {
			Text  string `xml:",chardata"`
			Title string `xml:"title"`
			Link  string `xml:"link"`
			GUUID struct {
				Text        string `xml:",chardata"`
				IsPermaLink string `xml:"isPermaLink,attr"`
			} `xml:"guid"`
			PubDate     string `xml:"pubDate"`
			Description string `xml:"description"`
			Source      struct {
				Text string `xml:",chardata"`
				URL  string `xml:"url,attr"`
			} `xml:"source"`
		} `xml:"item"`
	} `xml:"channel"`
}

// AllArticlesData is being used in lib/news/news.go
type AllArticlesData struct {
	VaccineArticles   ArticlesData `json:"vaccine"`
	TreatmentArticles ArticlesData `json:"treament"`
	NewsArticles      ArticlesData `json:"news"`
}

// ArticlesData is being used in lib/news/news.go
type ArticlesData struct {
	Articles []Article `json:"data"`
}

// Article is being used in lib/news/news.go
type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	URLToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
	Source      string `json:"source"`
	SourceURL   string `json:"sourceURL"`
}

type Hotspot struct {
	MostCases    CompareData `json:"mostCases"`
	SecondCases  CompareData `json:"secondCases"`
	ThirdCases   CompareData `json:"thirdCases"`
	MostDeaths   CompareData `json:"mostDeaths"`
	SecondDeaths CompareData `json:"secondDeaths"`
	ThirdDeaths  CompareData `json:"thirdDeaths"`
}

type MainCurveData struct {
	Deaths          []float64
	DeathsPerDay    []float64
	Cases           []float64
	CasesPerDay     []float64
	Recovered       []float64
	RecoveredPerDay []float64
}
