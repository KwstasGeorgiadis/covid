package stats

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"

	applogger "github.com/junkd0g/covid/lib/applogger"
	caching "github.com/junkd0g/covid/lib/caching"
	pconf "github.com/junkd0g/covid/lib/config"
	mcountry "github.com/junkd0g/covid/lib/model/country"
)

var (
	serverConf = pconf.GetAppConfig()
	redis      caching.RedisST
)

// requestData does an HTTP GET request to the third party API that
// contains covid-9 stats
// It returns []mcountry.Country and any write error encountered.
func requestData() ([]mcountry.Country, error) {

	client := &http.Client{}
	requestURL := serverConf.API.URL
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		applogger.Log("ERROR", "stats", "requestData", err.Error())
		return []mcountry.Country{}, err
	}

	res, resError := client.Do(req)
	if resError != nil {
		applogger.Log("ERROR", "stats", "requestData", resError.Error())
		return []mcountry.Country{}, resError
	}
	defer res.Body.Close()

	b, readBoyError := ioutil.ReadAll(res.Body)
	if readBoyError != nil {
		applogger.Log("ERROR", "stats", "requestData", readBoyError.Error())
		return []mcountry.Country{}, readBoyError
	}

	keys := make([]mcountry.Country, 0)
	if errUnmarshal := json.Unmarshal(b, &keys); err != nil {
		applogger.Log("ERROR", "stats", "requestData", errUnmarshal.Error())
		return []mcountry.Country{}, errUnmarshal
	}

	return keys, nil
}

// GetAllCountries get an array of all countries that have
// Covid-19 stats (data starts from date 22/01/2020)
// Check if there are cached data if not does a HTTP
// request to the 3rd party API (check requestData())
// It returns mcountry.Countries ([] Country) and any write error encountered.
func GetAllCountries() (mcountry.Countries, error) {

	cachedData, cacheGetError := redis.GetCountriesData()
	if cacheGetError != nil {
		applogger.Log("ERROR", "stats", "GetAllCountries", cacheGetError.Error())
		return mcountry.Countries{}, cacheGetError
	}
	var s mcountry.Countries

	if len(cachedData.Data) == 0 {
		applogger.Log("INFO", "stats", "GetAllCountries", "Request data instead of getting cached data")
		response, responseError := requestData()
		if responseError != nil {
			applogger.Log("ERROR", "stats", "GetAllCountries", responseError.Error())
			return mcountry.Countries{}, responseError
		}

		s = mcountry.Countries{Data: response}

		redis.SetCountriesData(s)

	} else {
		applogger.Log("INFO", "stats", "GetAllCountries", "Getting cache data %v instead of requesting it")
		return cachedData, nil
	}

	return s, nil
}

// GetCountry seach through an array of mcountry.Country and
// gets COVID-19 stats for that specific country
// It returns mcountry.Country and any write error encountered.
func GetCountry(name string) (mcountry.Country, error) {
	allCountries, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		applogger.Log("ERROR", "stats", "GetCountry", allCountriesError.Error())
		return mcountry.Country{}, allCountriesError
	}

	for _, v := range allCountries.Data {
		if v.Country == name {
			return v, nil
		}
	}

	applogger.Log("WARN", "stats", "GetCountry", "Returning empty country")
	return mcountry.Country{}, nil
}

// SortByCases sorts an array of Country structs by Country.Cases
// It returns structs.Countries ([] Country) and any write error encountered.
func SortByCases() (mcountry.Countries, error) {

	allCountriesArr, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		applogger.Log("ERROR", "stats", "SortByCases", allCountriesError.Error())
		return mcountry.Countries{}, allCountriesError
	}

	allCountries := allCountriesArr.Data

	sort.Slice(allCountries, func(i, j int) bool {
		if allCountries[i].Cases != allCountries[j].Cases {
			return allCountries[i].Cases > allCountries[j].Cases
		}
		return allCountries[i].Deaths > allCountries[j].Deaths
	})

	s := mcountry.Countries{Data: allCountries}

	return s, nil
}

// SortByDeaths sorts an array of Country structs by Country.Deaths
// It returns structs.Countries ([] Country) and any write error encountered.
func SortByDeaths() (mcountry.Countries, error) {
	allCountriesArr, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		applogger.Log("ERROR", "stats", "SortByDeaths", allCountriesError.Error())
		return mcountry.Countries{}, allCountriesError
	}

	allCountries := allCountriesArr.Data

	sort.Slice(allCountries, func(i, j int) bool {
		if allCountries[i].Deaths != allCountries[j].Deaths {
			return allCountries[i].Deaths > allCountries[j].Deaths
		}
		return allCountries[i].Cases > allCountries[j].Cases
	})

	s := mcountry.Countries{Data: allCountries}

	return s, nil
}

// SortByTodayCases sorts an array of Country mcountry by Country.TodayCases
// It returns mcountry.Countries ([] Country) and any write error encountered.
func SortByTodayCases() (mcountry.Countries, error) {
	allCountriesArr, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		applogger.Log("ERROR", "stats", "SortByTodayCases", allCountriesError.Error())
		return mcountry.Countries{}, allCountriesError
	}

	allCountries := allCountriesArr.Data

	sort.Slice(allCountries, func(i, j int) bool {
		return allCountries[i].TodayCases > allCountries[j].TodayCases
	})

	s := mcountry.Countries{Data: allCountries}
	return s, nil
}

// SortByTodayDeaths sorts an array of Country structs by Country.TodayDeaths
// It returns structs.Countries ([] Country) and any write error encountered.
func SortByTodayDeaths() (mcountry.Countries, error) {
	allCountriesArr, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		applogger.Log("ERROR", "stats", "SortByTodayDeaths", allCountriesError.Error())
		return mcountry.Countries{}, allCountriesError
	}

	allCountries := allCountriesArr.Data

	sort.Slice(allCountries, func(i, j int) bool {
		return allCountries[i].TodayDeaths > allCountries[j].TodayDeaths
	})

	s := mcountry.Countries{Data: allCountries}
	return s, nil
}

// SortByRecovered sorts an array of Country structs by Country.Recovered
// It returns mcountry.Countries ([] Country) and any write error encountered.
func SortByRecovered() (mcountry.Countries, error) {
	allCountriesArr, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		applogger.Log("ERROR", "stats", "SortByRecovered", allCountriesError.Error())
		return mcountry.Countries{}, allCountriesError
	}

	allCountries := allCountriesArr.Data

	sort.Slice(allCountries, func(i, j int) bool {
		return allCountries[i].Recovered > allCountries[j].Recovered
	})

	s := mcountry.Countries{Data: allCountries}
	return s, nil
}

// SortByActive sorts an array of Country structs by Country.Active
// It returns mcountry.Countries ([] Country) and any write error encountered.
func SortByActive() (mcountry.Countries, error) {
	allCountriesArr, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		applogger.Log("ERROR", "stats", "SortByActive", allCountriesError.Error())
		return mcountry.Countries{}, allCountriesError
	}

	allCountries := allCountriesArr.Data

	sort.Slice(allCountries, func(i, j int) bool {
		return allCountries[i].Active > allCountries[j].Active
	})

	s := mcountry.Countries{Data: allCountries}
	return s, nil
}

// SortByCritical sorts an array of Country structs by Country.Critical
// It returns mcountry.Countries ([] Country) and any write error encountered.
func SortByCritical() (mcountry.Countries, error) {
	allCountriesArr, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		applogger.Log("ERROR", "stats", "SortByCritical", allCountriesError.Error())
		return mcountry.Countries{}, allCountriesError
	}

	allCountries := allCountriesArr.Data

	sort.Slice(allCountries, func(i, j int) bool {
		return allCountries[i].Critical > allCountries[j].Critical
	})

	s := mcountry.Countries{Data: allCountries}
	return s, nil
}

// SortByCasesPerOneMillion sorts an array of Country structs by Country.CasesPerOneMillion
// It returns mcountry.Countries ([] Country) and any write error encountered.
func SortByCasesPerOneMillion() (mcountry.Countries, error) {
	allCountriesArr, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		applogger.Log("ERROR", "stats", "SortByCasesPerOneMillion", allCountriesError.Error())
		return mcountry.Countries{}, allCountriesError
	}

	allCountries := allCountriesArr.Data

	sort.Slice(allCountries, func(i, j int) bool {
		return allCountries[i].CasesPerOneMillion > allCountries[j].CasesPerOneMillion
	})

	s := mcountry.Countries{Data: allCountries}
	return s, nil
}

// PercentancePerCountry gets a country's COVID-19 stats (getting the from GetCountry)
// and calculate today's total cases percentance and today's death percentance
// It returns mcountry.CountryStats and any write error encountered.
func PercentancePerCountry(name string) (mcountry.CountryStats, error) {
	country, countryError := GetCountry(name)
	if countryError != nil {
		applogger.Log("ERROR", "stats", "PercentancePerCountry", countryError.Error())
		return mcountry.CountryStats{}, nil
	}

	var todayPerCentOfTotalCases = country.TodayCases * 100 / country.Cases
	var todayPerCentOfTotalDeaths = country.TodayDeaths * 100 / country.Deaths

	countryStats := mcountry.CountryStats{Country: country.Country,
		TodayPerCentOfTotalCases:  todayPerCentOfTotalCases,
		TodayPerCentOfTotalDeaths: todayPerCentOfTotalDeaths}

	return countryStats, nil
}

// GetTotalStats gets worlds COVID-19 total statistics.
// The statistics are total cases, total deaths today's total deaths
// totltoal cases, percentace totay increase in deaths and cases
// It returns mcountry.TotalStats and any write error encountered.
func GetTotalStats() (mcountry.TotalStats, error) {
	var totalDeaths = 0
	var totalCases = 0
	var todayTotalDeaths = 0
	var todayTotalCases = 0

	allCountriesArr, errorAllCountries := GetAllCountries()
	if errorAllCountries != nil {
		applogger.Log("ERROR", "stats", "GetTotalStats", errorAllCountries.Error())
		return mcountry.TotalStats{}, nil
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

	totalStatsStuct := mcountry.TotalStats{
		TodayPerCentOfTotalCases:  todayPerCentOfTotalCases,
		TodayPerCentOfTotalDeaths: todayPerCentOfTotalDeaths,
		TotalCases:                totalCases,
		TotalDeaths:               totalDeaths,
		TodayTotalCases:           todayTotalCases,
		TodayTotalDeaths:          todayTotalDeaths,
	}

	return totalStatsStuct, nil
}

// GetAllCountriesName get names of the countries that we have Covid-19 stats
// It returns mcountry.AllCountriesName and any write error encountered.
func GetAllCountriesName() (mcountry.AllCountriesName, error) {
	allCountriesArr, allCountriesError := GetAllCountries()
	if allCountriesError != nil {
		applogger.Log("ERROR", "stats", "GetAllCountriesName", allCountriesError.Error())
		return mcountry.AllCountriesName{}, allCountriesError
	}

	allCountries := allCountriesArr.Data

	var counties []string

	for _, v := range allCountries {
		counties = append(counties, v.Country)
	}

	allCountriesStruct := mcountry.AllCountriesName{Countries: counties}

	return allCountriesStruct, nil
}
