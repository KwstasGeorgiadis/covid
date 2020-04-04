package stats

//TODO add expiration time

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"

	caching "../caching"
	pconf "../config"
	structs "../structs"
)

var (
	serverConf = pconf.GetAppConfig("./config/covid.json")
)

func requestData() []structs.Country {

	client := &http.Client{}
	req, err := http.NewRequest("GET", serverConf.API.URL, nil)

	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	defer res.Body.Close()
	b, _ := ioutil.ReadAll(res.Body)

	//var obj Countries
	keys := make([]structs.Country, 0)
	if err := json.Unmarshal(b, &keys); err != nil {
		panic(err)
	}

	return keys
}

func GetAllCountries() structs.Countries {
	pool := caching.NewPool()
	conn := pool.Get()
	defer conn.Close()

	cachedData, _ := caching.Get(conn)

	var s structs.Countries

	if len(cachedData.Data) == 0 {
		n := requestData()
		s = structs.Countries{Data: n}
		caching.Set(conn, s)
	} else {
		return cachedData
	}

	return s
}

func GetCountry(name string) structs.Country {
	allCountries := GetAllCountries().Data

	for _, v := range allCountries {
		if v.Country == name {
			return v
		}
	}

	return structs.Country{}
}

func SortByCases() structs.Countries {
	allCountries := GetAllCountries().Data
	fmt.Println(allCountries)
	sort.Slice(allCountries, func(i, j int) bool {
		if allCountries[i].Cases != allCountries[j].Cases {
			return allCountries[i].Cases > allCountries[j].Cases
		} else {
			return allCountries[i].Deaths > allCountries[j].Deaths
		}
	})

	s := structs.Countries{Data: allCountries}
	return s
}

func SortByDeaths() structs.Countries {
	allCountries := GetAllCountries().Data

	sort.Slice(allCountries, func(i, j int) bool {
		if allCountries[i].Deaths != allCountries[j].Deaths {
			return allCountries[i].Deaths > allCountries[j].Deaths
		} else {
			return allCountries[i].Cases > allCountries[j].Cases
		}
	})

	s := structs.Countries{Data: allCountries}
	return s
}

func SortByTodayCases() structs.Countries {
	allCountries := GetAllCountries().Data

	sort.Slice(allCountries, func(i, j int) bool {
		return allCountries[i].TodayCases > allCountries[j].TodayCases
	})

	s := structs.Countries{Data: allCountries}
	return s
}

func SortByTodayDeaths() structs.Countries {
	allCountries := GetAllCountries().Data

	sort.Slice(allCountries, func(i, j int) bool {
		return allCountries[i].TodayDeaths > allCountries[j].TodayDeaths
	})

	s := structs.Countries{Data: allCountries}
	return s
}

func SortByRecovered() structs.Countries {
	allCountries := GetAllCountries().Data

	sort.Slice(allCountries, func(i, j int) bool {
		return allCountries[i].Recovered > allCountries[j].Recovered
	})

	s := structs.Countries{Data: allCountries}
	return s
}

func SortByActive() structs.Countries {
	allCountries := GetAllCountries().Data

	sort.Slice(allCountries, func(i, j int) bool {
		return allCountries[i].Active > allCountries[j].Active
	})

	s := structs.Countries{Data: allCountries}
	return s
}

func SortByCritical() structs.Countries {
	allCountries := GetAllCountries().Data

	sort.Slice(allCountries, func(i, j int) bool {
		return allCountries[i].Critical > allCountries[j].Critical
	})

	s := structs.Countries{Data: allCountries}
	return s
}

func SortByCasesPerOneMillion() structs.Countries {
	allCountries := GetAllCountries().Data

	sort.Slice(allCountries, func(i, j int) bool {
		return allCountries[i].CasesPerOneMillion > allCountries[j].CasesPerOneMillion
	})

	s := structs.Countries{Data: allCountries}
	return s
}

func StatsPerCountry(name string) structs.CountryStats {
	country := GetCountry(name)

	var todayPerCentOfTotalCases = country.TodayCases * 100 / country.Cases
	var todayPerCentOfTotalDeaths = country.TodayDeaths * 100 / country.Deaths

	return structs.CountryStats{Country: country.Country,
		TodayPerCentOfTotalCases:  todayPerCentOfTotalCases,
		TodayPerCentOfTotalDeaths: todayPerCentOfTotalDeaths}
}

func GetTotalStats() structs.TotalStats {
	var totalDeaths = 0
	var totalCases = 0
	var todayTotalDeaths = 0
	var todayTotalCases = 0

	allCountries := GetAllCountries().Data

	for _, v := range allCountries {
		if v.Country == "World" {
			continue
		}
		totalDeaths = totalDeaths + v.Deaths
		totalCases = totalCases + v.Cases
		todayTotalDeaths = todayTotalDeaths + v.TodayDeaths
		todayTotalCases = todayTotalCases + v.TodayCases
	}

	var todayPerCentOfTotalCases = todayTotalDeaths * 100 / totalDeaths
	var todayPerCentOfTotalDeaths = todayTotalCases * 100 / totalCases

	return structs.TotalStats{
		TodayPerCentOfTotalCases:  todayPerCentOfTotalCases,
		TodayPerCentOfTotalDeaths: todayPerCentOfTotalDeaths,
		TotalCases:                totalCases,
		TotalDeaths:               totalDeaths,
		TodayTotalCases:           todayTotalCases,
		TodayTotalDeaths:          todayTotalDeaths,
	}
}

func GetAllCountriesName() structs.AllCountriesName {
	allCountries := GetAllCountries().Data
	var counties []string

	for _, v := range allCountries {
		counties = append(counties, v.Country)
	}

	return structs.AllCountriesName{Countries: counties}
}
