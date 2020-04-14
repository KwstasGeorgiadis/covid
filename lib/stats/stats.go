package stats

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"

	applogger "../applogger"
	caching "../caching"
	pconf "../config"
	structs "../structs"
)

var (
	serverConf = pconf.GetAppConfig("./config/covid.json")
)

// requestData does an HTTP GET request to the third party API that
// contains covid-9 stats
// It returns []structs.Country and any write error encountered.
func requestData() ([]structs.Country, error) {

	client := &http.Client{}
	requestURL := serverConf.API.URL
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		applogger.Log("ERROR", "stats", "requestData", err.Error())
		return []structs.Country{}, err
	}

	res, resError := client.Do(req)
	if resError != nil {
		applogger.Log("ERROR", "stats", "requestData", resError.Error())
		return []structs.Country{}, resError
	}
	defer res.Body.Close()

	b, readBoyError := ioutil.ReadAll(res.Body)
	if readBoyError != nil {
		applogger.Log("ERROR", "stats", "requestData", readBoyError.Error())
		return []structs.Country{}, readBoyError
	}

	keys := make([]structs.Country, 0)
	if errUnmarshal := json.Unmarshal(b, &keys); err != nil {
		applogger.Log("ERROR", "stats", "requestData", errUnmarshal.Error())
		return []structs.Country{}, errUnmarshal
	}

	applogger.Log("INFO", "stats", "requestData",
		fmt.Sprintf("Get reqeust to %s and getting response %v", requestURL, keys))
	return keys, nil
}

// GetAllCountries get an array of all countries that have
// Covid-19 stats (data starts from date 22/01/2020)
// Check if there are cached data if not does a HTTP
// request to the 3rd party API (check requestData())
// It returns structs.Countries ([] Country) and any write error encountered.
func GetAllCountries() (structs.Countries, error) {
	pool := caching.NewPool()
	conn := pool.Get()
	defer conn.Close()

	cachedData, cacheGetError := caching.Get(conn, "total")
	if cacheGetError != nil {
		applogger.Log("ERROR", "stats", "GetAllCountries", cacheGetError.Error())
		return structs.Countries{}, cacheGetError
	}

	var s structs.Countries

	if len(cachedData.Data) == 0 {
		applogger.Log("INFO", "stats", "GetAllCountries", "Request data instead of getting cached data")
		response, responseError := requestData()

		if responseError != nil {
			applogger.Log("ERROR", "stats", "GetAllCountries", responseError.Error())
			return structs.Countries{}, responseError
		}

		s = structs.Countries{Data: response}
		caching.Set(conn, s, "total")
		applogger.Log("INFO", "stats", "GetAllCountries", fmt.Sprintf("Setting cache data %v for key total", s))

	} else {
		applogger.Log("INFO", "stats", "GetAllCountries", fmt.Sprintf("Getting cache data %v instead of requesting it", cachedData))
		return cachedData, nil
	}

	return s, nil
}

// GetCountry seach through an array of structs.Country and
// gets COVID-19 stats for that specific country
// It returns structs.Country and any write error encountered.
func GetCountry(name string) (structs.Country, error) {
	allCountries, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		applogger.Log("ERROR", "stats", "GetCountry", allCountriesError.Error())
		return structs.Country{}, allCountriesError
	}

	for _, v := range allCountries.Data {
		if v.Country == name {
			applogger.Log("INFO", "stats", "GetCountry", fmt.Sprintf("Returning country %v", v))
			return v, nil
		}
	}

	applogger.Log("WARN", "stats", "GetCountry", "Returning empty country")
	return structs.Country{}, nil
}

// SortByCases sorts an array of Country structs by Country.Cases
// It returns structs.Countries ([] Country) and any write error encountered.
func SortByCases() (structs.Countries, error) {

	allCountriesArr, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		applogger.Log("ERROR", "stats", "SortByCases", allCountriesError.Error())
		return structs.Countries{}, allCountriesError
	}

	allCountries := allCountriesArr.Data

	sort.Slice(allCountries, func(i, j int) bool {
		if allCountries[i].Cases != allCountries[j].Cases {
			return allCountries[i].Cases > allCountries[j].Cases
		}
		return allCountries[i].Deaths > allCountries[j].Deaths
	})

	s := structs.Countries{Data: allCountries}
	applogger.Log("INFO", "stats", "SortByCases", fmt.Sprintf("Returning sorted countries %v", s))
	return s, nil
}

// SortByDeaths sorts an array of Country structs by Country.Deaths
// It returns structs.Countries ([] Country) and any write error encountered.
func SortByDeaths() (structs.Countries, error) {
	allCountriesArr, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		applogger.Log("ERROR", "stats", "SortByDeaths", allCountriesError.Error())
		return structs.Countries{}, allCountriesError
	}

	allCountries := allCountriesArr.Data

	sort.Slice(allCountries, func(i, j int) bool {
		if allCountries[i].Deaths != allCountries[j].Deaths {
			return allCountries[i].Deaths > allCountries[j].Deaths
		}
		return allCountries[i].Cases > allCountries[j].Cases
	})

	s := structs.Countries{Data: allCountries}
	applogger.Log("INFO", "stats", "SortByDeaths", fmt.Sprintf("Returning sorted countries %v", s))
	return s, nil
}

// SortByTodayCases sorts an array of Country structs by Country.TodayCases
// It returns structs.Countries ([] Country) and any write error encountered.
func SortByTodayCases() (structs.Countries, error) {
	allCountriesArr, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		applogger.Log("ERROR", "stats", "SortByTodayCases", allCountriesError.Error())
		return structs.Countries{}, allCountriesError
	}

	allCountries := allCountriesArr.Data

	sort.Slice(allCountries, func(i, j int) bool {
		return allCountries[i].TodayCases > allCountries[j].TodayCases
	})

	s := structs.Countries{Data: allCountries}
	applogger.Log("INFO", "stats", "SortByTodayCases", fmt.Sprintf("Returning sorted countries %v", s))
	return s, nil
}

// SortByTodayDeaths sorts an array of Country structs by Country.TodayDeaths
// It returns structs.Countries ([] Country) and any write error encountered.
func SortByTodayDeaths() (structs.Countries, error) {
	allCountriesArr, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		applogger.Log("ERROR", "stats", "SortByTodayDeaths", allCountriesError.Error())
		return structs.Countries{}, allCountriesError
	}

	allCountries := allCountriesArr.Data

	sort.Slice(allCountries, func(i, j int) bool {
		return allCountries[i].TodayDeaths > allCountries[j].TodayDeaths
	})

	s := structs.Countries{Data: allCountries}
	applogger.Log("INFO", "stats", "SortByTodayDeaths", fmt.Sprintf("Returning sorted countries %v", s))
	return s, nil
}

// SortByRecovered sorts an array of Country structs by Country.Recovered
// It returns structs.Countries ([] Country) and any write error encountered.
func SortByRecovered() (structs.Countries, error) {
	allCountriesArr, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		applogger.Log("ERROR", "stats", "SortByRecovered", allCountriesError.Error())
		return structs.Countries{}, allCountriesError
	}

	allCountries := allCountriesArr.Data

	sort.Slice(allCountries, func(i, j int) bool {
		return allCountries[i].Recovered > allCountries[j].Recovered
	})

	s := structs.Countries{Data: allCountries}
	applogger.Log("INFO", "stats", "SortByRecovered", fmt.Sprintf("Returning sorted countries %v", s))
	return s, nil
}

// SortByActive sorts an array of Country structs by Country.Active
// It returns structs.Countries ([] Country) and any write error encountered.
func SortByActive() (structs.Countries, error) {
	allCountriesArr, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		applogger.Log("ERROR", "stats", "SortByActive", allCountriesError.Error())
		return structs.Countries{}, allCountriesError
	}

	allCountries := allCountriesArr.Data

	sort.Slice(allCountries, func(i, j int) bool {
		return allCountries[i].Active > allCountries[j].Active
	})

	s := structs.Countries{Data: allCountries}
	applogger.Log("INFO", "stats", "SortByActive", fmt.Sprintf("Returning sorted countries %v", s))
	return s, nil
}

// SortByCritical sorts an array of Country structs by Country.Critical
// It returns structs.Countries ([] Country) and any write error encountered.
func SortByCritical() (structs.Countries, error) {
	allCountriesArr, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		applogger.Log("ERROR", "stats", "SortByCritical", allCountriesError.Error())
		return structs.Countries{}, allCountriesError
	}

	allCountries := allCountriesArr.Data

	sort.Slice(allCountries, func(i, j int) bool {
		return allCountries[i].Critical > allCountries[j].Critical
	})

	s := structs.Countries{Data: allCountries}
	applogger.Log("INFO", "stats", "SortByCritical", fmt.Sprintf("Returning sorted countries %v", s))
	return s, nil
}

// SortByCasesPerOneMillion sorts an array of Country structs by Country.CasesPerOneMillion
// It returns structs.Countries ([] Country) and any write error encountered.
func SortByCasesPerOneMillion() (structs.Countries, error) {
	allCountriesArr, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		applogger.Log("ERROR", "stats", "SortByCasesPerOneMillion", allCountriesError.Error())
		return structs.Countries{}, allCountriesError
	}

	allCountries := allCountriesArr.Data

	sort.Slice(allCountries, func(i, j int) bool {
		return allCountries[i].CasesPerOneMillion > allCountries[j].CasesPerOneMillion
	})

	s := structs.Countries{Data: allCountries}
	applogger.Log("INFO", "stats", "SortByCasesPerOneMillion", fmt.Sprintf("Returning sorted countries %v", s))
	return s, nil
}

// PercentancePerCountry gets a country's COVID-19 stats (getting the from GetCountry)
// and calculate today's total cases percentance and today's death percentance
// It returns structs.CountryStats and any write error encountered.
func PercentancePerCountry(name string) (structs.CountryStats, error) {
	country, countryError := GetCountry(name)
	if countryError != nil {
		applogger.Log("ERROR", "stats", "PercentancePerCountry", countryError.Error())
		return structs.CountryStats{}, nil
	}

	var todayPerCentOfTotalCases = country.TodayCases * 100 / country.Cases
	var todayPerCentOfTotalDeaths = country.TodayDeaths * 100 / country.Deaths

	countryStats := structs.CountryStats{Country: country.Country,
		TodayPerCentOfTotalCases:  todayPerCentOfTotalCases,
		TodayPerCentOfTotalDeaths: todayPerCentOfTotalDeaths}

	applogger.Log("INFO", "stats", "PercentancePerCountry", fmt.Sprintf("Percentanceper country  %v", countryStats))
	return countryStats, nil
}

// GetTotalStats gets worlds COVID-19 total statistics.
// The statistics are total cases, total deaths today's total deaths
// totltoal cases, percentace totay increase in deaths and cases
// It returns structs.TotalStats and any write error encountered.
func GetTotalStats() (structs.TotalStats, error) {
	var totalDeaths = 0
	var totalCases = 0
	var todayTotalDeaths = 0
	var todayTotalCases = 0

	allCountriesArr, errorAllCountries := GetAllCountries()
	if errorAllCountries != nil {
		applogger.Log("ERROR", "stats", "GetTotalStats", errorAllCountries.Error())
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

	applogger.Log("INFO", "stats", "GetTotalStats", fmt.Sprintf("Total stats  %v", totalStatsStuct))
	return totalStatsStuct, nil
}

// GetAllCountriesName get names of the countries that we have Covid-19 stats
// It returns structs.AllCountriesName and any write error encountered.
func GetAllCountriesName() (structs.AllCountriesName, error) {
	allCountriesArr, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		applogger.Log("ERROR", "stats", "GetAllCountriesName", allCountriesError.Error())
		return structs.AllCountriesName{}, allCountriesError
	}

	allCountries := allCountriesArr.Data

	var counties []string

	for _, v := range allCountries {
		counties = append(counties, v.Country)
	}

	allCountriesStruct := structs.AllCountriesName{Countries: counties}

	applogger.Log("INFO", "stats", "GetAllCountriesName", fmt.Sprintf("Total stats  %v", allCountriesStruct))
	return allCountriesStruct, nil
}
