package stats

import (
	pconf "../config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
)

type Countries struct {
	Data []Country `json:"data"`
}

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

var (
	serverConf = pconf.GetAppConfig("./config/covid.json")
)

func requestData() []Country {

	client := &http.Client{}
	req, err := http.NewRequest("GET", serverConf.API.URL, nil)

	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	defer res.Body.Close()
	b, _ := ioutil.ReadAll(res.Body)

	//var obj Countries
	keys := make([]Country, 0)
	if err := json.Unmarshal(b, &keys); err != nil {
		panic(err)
	}

	return keys
}

func GetAllCountries() Countries {
	s := Countries{Data: requestData()}
	return s
}

func GetCountry(name string) Country {
	allCountries := requestData()

	for _, v := range allCountries {
		if v.Country == name {
			return v
		}
	}

	return Country{}
}

func SortByCases() Countries {
	allCountries := requestData()

	sort.Slice(allCountries, func(i, j int) bool {
		if allCountries[i].Cases != allCountries[j].Cases {
			return allCountries[i].Cases > allCountries[j].Cases
		} else {
			return allCountries[i].Deaths > allCountries[j].Deaths
		}
	})

	s := Countries{Data: allCountries}
	return s
}

func SortByDeaths() Countries {
	allCountries := requestData()

	sort.Slice(allCountries, func(i, j int) bool {
		if allCountries[i].Deaths != allCountries[j].Deaths {
			return allCountries[i].Deaths > allCountries[j].Deaths
		} else {
			return allCountries[i].Cases > allCountries[j].Cases
		}
	})

	s := Countries{Data: allCountries}
	return s
}

func SortByTodayCases() Countries {
	allCountries := requestData()

	sort.Slice(allCountries, func(i, j int) bool {
		return allCountries[i].TodayCases > allCountries[j].TodayCases
	})

	s := Countries{Data: allCountries}
	return s
}

func SortByTodayDeaths() Countries {
	allCountries := requestData()

	sort.Slice(allCountries, func(i, j int) bool {
		return allCountries[i].TodayDeaths > allCountries[j].TodayDeaths
	})

	s := Countries{Data: allCountries}
	return s
}

func SortByRecovered() Countries {
	allCountries := requestData()

	sort.Slice(allCountries, func(i, j int) bool {
		return allCountries[i].Recovered > allCountries[j].Recovered
	})

	s := Countries{Data: allCountries}
	return s
}

func SortByActive() Countries {
	allCountries := requestData()

	sort.Slice(allCountries, func(i, j int) bool {
		return allCountries[i].Active > allCountries[j].Active
	})

	s := Countries{Data: allCountries}
	return s
}

func SortByCritical() Countries {
	allCountries := requestData()

	sort.Slice(allCountries, func(i, j int) bool {
		return allCountries[i].Critical > allCountries[j].Critical
	})

	s := Countries{Data: allCountries}
	return s
}

func SortByCasesPerOneMillion() Countries {
	allCountries := requestData()

	sort.Slice(allCountries, func(i, j int) bool {
		return allCountries[i].CasesPerOneMillion > allCountries[j].CasesPerOneMillion
	})

	s := Countries{Data: allCountries}
	return s
}

type CountryStats struct {
	Country                   string `json:"country"`
	TodayPerCentOfTotalCases  int    `json:"todayPerCentOfTotalCases"`
	TodayPerCentOfTotalDeaths int    `json:"todayPerCentOfTotalDeaths"`
}

func StatsPerCountry(name string) CountryStats {
	country := GetCountry(name)

	var todayPerCentOfTotalCases = country.TodayCases * 100 / country.Cases
	var todayPerCentOfTotalDeaths = country.TodayDeaths * 100 / country.Deaths

	return CountryStats{Country: country.Country,
		TodayPerCentOfTotalCases:  todayPerCentOfTotalCases,
		TodayPerCentOfTotalDeaths: todayPerCentOfTotalDeaths}
}

type TotalStats struct {
	TodayPerCentOfTotalCases  int `json:"todayPerCentOfTotalCases"`
	TodayPerCentOfTotalDeaths int `json:"todayPerCentOfTotalDeaths"`
	TotalCases                int `json:"totalCases"`
	TotalDeaths               int `json:"totalDeaths"`
	TodayTotalCases           int `json:"todayTotalCases"`
	TodayTotalDeaths          int `json:"todayTotalDeaths"`
}

func GetTotalStats() TotalStats {
	var totalDeaths = 0
	var totalCases = 0
	var todayTotalDeaths = 0
	var todayTotalCases = 0

	allCountries := requestData()

	for _, v := range allCountries {
		totalDeaths = totalDeaths + v.Deaths
		totalCases = totalCases + v.Cases
		todayTotalDeaths = todayTotalDeaths + v.TodayDeaths
		todayTotalCases = todayTotalCases + v.TodayCases
	}

	var todayPerCentOfTotalCases = todayTotalDeaths * 100 / totalDeaths
	var todayPerCentOfTotalDeaths = todayTotalCases * 100 / totalCases

	return TotalStats{
		TodayPerCentOfTotalCases:  todayPerCentOfTotalCases,
		TodayPerCentOfTotalDeaths: todayPerCentOfTotalDeaths,
		TotalCases:                totalCases,
		TotalDeaths:               totalDeaths,
		TodayTotalCases:           todayTotalCases,
		TodayTotalDeaths:          todayTotalDeaths,
	}
}
