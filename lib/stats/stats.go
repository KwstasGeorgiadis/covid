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

func GetAllCountries() (structs.Countries, error){
	pool := caching.NewPool()
	conn := pool.Get()
	defer conn.Close()

	cachedData, cacheGetError := caching.Get(conn, "total")
	if cacheGetError != nil {
		return structs.Countries{}, cacheGetError
	}

	var s structs.Countries

	if len(cachedData.Data) == 0 {
		n := requestData()
		s = structs.Countries{Data: n}
		caching.Set(conn, s, "total")
	} else {
		return cachedData,nil
	}

	return s,nil
}

func GetCountry(name string) (structs.Country,error) {
	allCountries, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		return structs.Country{}, allCountriesError
	}

	for _, v := range allCountries.Data {
		if v.Country == name {
			return v, nil
		}
	}

	return structs.Country{}, nil
}

func SortByCases() (structs.Countries,error) {

	allCountriesArr, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		return structs.Countries{}, allCountriesError
	}

	allCountries := allCountriesArr.Data

	sort.Slice(allCountries, func(i, j int) bool {
		if allCountries[i].Cases != allCountries[j].Cases {
			return allCountries[i].Cases > allCountries[j].Cases
		} else {
			return allCountries[i].Deaths > allCountries[j].Deaths
		}
	})

	s := structs.Countries{Data: allCountries}
	return s, nil
}

func SortByDeaths() (structs.Countries,error) {
	allCountriesArr, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		return structs.Countries{}, allCountriesError
	}

	allCountries := allCountriesArr.Data

	sort.Slice(allCountries, func(i, j int) bool {
		if allCountries[i].Deaths != allCountries[j].Deaths {
			return allCountries[i].Deaths > allCountries[j].Deaths
		} else {
			return allCountries[i].Cases > allCountries[j].Cases
		}
	})

	s := structs.Countries{Data: allCountries}
	return s,nil 
}

func SortByTodayCases() (structs.Countries,error) {
	allCountriesArr, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		return structs.Countries{}, allCountriesError
	}

	allCountries := allCountriesArr.Data

	sort.Slice(allCountries, func(i, j int) bool {
		return allCountries[i].TodayCases > allCountries[j].TodayCases
	})

	s := structs.Countries{Data: allCountries}
	return s,nil
}

func SortByTodayDeaths() (structs.Countries,error) {
	allCountriesArr, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		return structs.Countries{}, allCountriesError
	}

	allCountries := allCountriesArr.Data

	sort.Slice(allCountries, func(i, j int) bool {
		return allCountries[i].TodayDeaths > allCountries[j].TodayDeaths
	})

	s := structs.Countries{Data: allCountries}
	return s,nil
}

func SortByRecovered() (structs.Countries,error) {
	allCountriesArr, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		return structs.Countries{}, allCountriesError
	}

	allCountries := allCountriesArr.Data

	sort.Slice(allCountries, func(i, j int) bool {
		return allCountries[i].Recovered > allCountries[j].Recovered
	})

	s := structs.Countries{Data: allCountries}
	return s,nil
}

func SortByActive() (structs.Countries,error) {
	allCountriesArr, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		return structs.Countries{}, allCountriesError
	}

	allCountries := allCountriesArr.Data

	sort.Slice(allCountries, func(i, j int) bool {
		return allCountries[i].Active > allCountries[j].Active
	})

	s := structs.Countries{Data: allCountries}
	return s, nil
}

func SortByCritical() (structs.Countries,error) {
	allCountriesArr, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		return structs.Countries{}, allCountriesError
	}

	allCountries := allCountriesArr.Data


	sort.Slice(allCountries, func(i, j int) bool {
		return allCountries[i].Critical > allCountries[j].Critical
	})

	s := structs.Countries{Data: allCountries}
	return s, nil
}

func SortByCasesPerOneMillion() (structs.Countries,error) {
	allCountriesArr, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		return structs.Countries{}, allCountriesError
	}

	allCountries := allCountriesArr.Data

	sort.Slice(allCountries, func(i, j int) bool {
		return allCountries[i].CasesPerOneMillion > allCountries[j].CasesPerOneMillion
	})

	s := structs.Countries{Data: allCountries}
	return s, nil
}

func StatsPerCountry(name string) (structs.CountryStats,error) {
	country, countryError := GetCountry(name)
	if countryError!= nil {
		return structs.CountryStats{}, nil 
	}

	var todayPerCentOfTotalCases = country.TodayCases * 100 / country.Cases
	var todayPerCentOfTotalDeaths = country.TodayDeaths * 100 / country.Deaths

	countryStats := structs.CountryStats{Country: country.Country,
		TodayPerCentOfTotalCases:  todayPerCentOfTotalCases,
		TodayPerCentOfTotalDeaths: todayPerCentOfTotalDeaths}

	return countryStats, nil
}

func GetTotalStats() (structs.TotalStats,error) {
	var totalDeaths = 0
	var totalCases = 0
	var todayTotalDeaths = 0
	var todayTotalCases = 0

	allCountriesArr, errorAllCountries := GetAllCountries()
	if errorAllCountries!= nil {
		return structs.TotalStats{}, nil 
	}

	allCountries := allCountriesArr.Data

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

	totalStatsStuct := structs.TotalStats{
		TodayPerCentOfTotalCases:  todayPerCentOfTotalCases,
		TodayPerCentOfTotalDeaths: todayPerCentOfTotalDeaths,
		TotalCases:                totalCases,
		TotalDeaths:               totalDeaths,
		TodayTotalCases:           todayTotalCases,
		TodayTotalDeaths:          todayTotalDeaths,
	}

	return totalStatsStuct, nil
}

func GetAllCountriesName() (structs.AllCountriesName,error) {
	allCountriesArr, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		return structs.AllCountriesName{}, allCountriesError
	}

	allCountries := allCountriesArr.Data

	var counties []string

	for _, v := range allCountries {
		counties = append(counties, v.Country)
	}

	return structs.AllCountriesName{Countries: counties}, nil
}
