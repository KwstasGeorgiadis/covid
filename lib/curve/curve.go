package curve

import (
	"sort"

	applogger "github.com/junkd0g/covid/lib/applogger"
	caching "github.com/junkd0g/covid/lib/caching"
	pconf "github.com/junkd0g/covid/lib/config"
	structs "github.com/junkd0g/covid/lib/structs"

	"encoding/json"
	"io/ioutil"
	"net/http"
)

var (
	serverConf pconf.AppConf
)

func init() {
	serverConf = pconf.GetAppConfig()
}

// requestHistoryData does an HTTP GET request to the third party API that
// contains covid-9 stats ' history (per day from 22/01/2020)
// It returns []structs.Country and any write error encountered.
func requestHistoryData() ([]structs.CountryCurve, error) {
	client := &http.Client{}
	requestURL := serverConf.API.URLHistory

	req, reqErr := http.NewRequest("GET", requestURL, nil)
	if reqErr != nil {
		applogger.Log("ERROR", "curve", "requestHistoryData", reqErr.Error())
		return []structs.CountryCurve{}, reqErr
	}

	res, resError := client.Do(req)
	if resError != nil {
		applogger.Log("ERROR", "curve", "requestHistoryData", resError.Error())
		return []structs.CountryCurve{}, resError
	}
	defer res.Body.Close()

	b, errorReadAll := ioutil.ReadAll(res.Body)
	if errorReadAll != nil {
		applogger.Log("ERROR", "curve", "requestHistoryData", errorReadAll.Error())
		return []structs.CountryCurve{}, errorReadAll
	}

	keys := make([]structs.CountryCurve, 0)
	if errUnmarshal := json.Unmarshal(b, &keys); errUnmarshal != nil {
		applogger.Log("ERROR", "curve", "requestHistoryData", errUnmarshal.Error())
		return []structs.CountryCurve{}, errUnmarshal
	}

	return keys, nil
}

// GetAllCountries returns an array of all countries per day
// Covid-19 stats (data starts from date 22/01/2020)
// Check if there are cached data if not does a HTTP
// request to the 3rd party API (check requestHistoryData())
// It returns []structs.CountryCurve and any write error encountered.
func GetAllCountries() ([]structs.CountryCurve, error) {

	pool := caching.NewPool()
	conn := pool.Get()
	defer conn.Close()

	cachedData, cacheGetError := caching.GetCurveData(conn)
	if cacheGetError != nil {
		applogger.Log("ERROR", "curve", "GetAllCountries", cacheGetError.Error())
		return []structs.CountryCurve{}, cacheGetError
	}

	if len(cachedData) == 0 {
		applogger.Log("INFO", "stats", "GetAllCountries", "Request data instead of getting cached data")
		data, err := requestHistoryData()
		if err != nil {
			applogger.Log("ERROR", "curve", "GetAllCountries", err.Error())
			return []structs.CountryCurve{}, err
		}
		caching.SetCurveData(conn, data)
		return data, nil
	}

	return cachedData, nil
}

// GetCountry seach through an array of structs.CountryCurve and
// gets COVID-19 per day stats for that specific country
// It returns structs.CountryCurve and any write error encountered.

// GetCountryBP
func GetCountryBP(name string, allCountries []structs.CountryCurve) (structs.CountryCurve, error) {

	for _, v := range allCountries {
		if name == "UK" && v.Country == "UK" {
			if len(v.Province) == 0 {
				return v, nil
			}
			continue
		}
		if name == "France" && v.Country == "France" {
			if v.Province == "" {
				return v, nil
			}
			continue
		}

		if v.Country == name {
			return v, nil
		}
	}

	applogger.Log("WARN", "curve", "GetCountry", "Returning empty country")
	return structs.CountryCurve{}, nil
}

// CompareDeathsCountries returns two integer arrays (one per country passed
// in parameter) which contain total number of deaths from  22/01/2020
// It returns structs.Compare and any write error encountered.
func CompareDeathsCountries(nameOne string, nameTwo string) (structs.Compare, error) {
	countries, err := GetAllCountries()
	if err != nil {
		applogger.Log("ERROR", "curve", "ComparePerDayCasesCountries", err.Error())
		return structs.Compare{}, err
	}
	countryData, countryDataErr := GetCountryData(nameOne, countries)
	if countryDataErr != nil {
		applogger.Log("ERROR", "curve", "ComparePerDayCasesCountries", countryDataErr.Error())
		return structs.Compare{}, countryDataErr
	}

	countryTwoData, countryTwoDataErr := GetCountryData(nameTwo, countries)
	if countryTwoDataErr != nil {
		applogger.Log("ERROR", "curve", "ComparePerDayCasesCountries", countryTwoDataErr.Error())
		return structs.Compare{}, countryTwoDataErr
	}

	var countryOneStruct structs.CompareData
	var countryTwoStruct structs.CompareData

	countryOneStruct.Country = nameOne
	countryOneStruct.Data = countryData.Deaths
	countryTwoStruct.Country = nameTwo
	countryTwoStruct.Data = countryTwoData.Deaths

	compareStructs := structs.Compare{CountryOne: countryOneStruct, CountryTwo: countryTwoStruct}

	return compareStructs, nil
}

// CompareDeathsFromFirstDeathCountries returns two integer arrays (one per country passed
// in parameter) which contain total number of deaths from  the first confirm death.
// It returns structs.Compare and any write error encountered.
func CompareDeathsFromFirstDeathCountries(nameOne string, nameTwo string) (structs.Compare, error) {
	countries, err := GetAllCountries()
	if err != nil {
		applogger.Log("ERROR", "curve", "ComparePerDayCasesCountries", err.Error())
		return structs.Compare{}, err
	}

	countryData, countryDataErr := GetCountryData(nameOne, countries)
	if countryDataErr != nil {
		applogger.Log("ERROR", "curve", "ComparePerDayCasesCountries", countryDataErr.Error())
		return structs.Compare{}, countryDataErr
	}

	countryTwoData, countryTwoDataErr := GetCountryData(nameTwo, countries)
	if countryTwoDataErr != nil {
		applogger.Log("ERROR", "curve", "ComparePerDayCasesCountries", countryTwoDataErr.Error())
		return structs.Compare{}, countryTwoDataErr
	}

	var countrySortedDeath []float64
	var countryTwoSortedDeath []float64

	for _, v := range countryData.Deaths {
		if v == 0 {
			continue
		}
		countrySortedDeath = append(countrySortedDeath, v)
	}
	for _, v := range countryTwoData.Deaths {
		if v == 0 {
			continue
		}
		countryTwoSortedDeath = append(countryTwoSortedDeath, v)
	}

	var countryOneStruct structs.CompareData
	var countryTwoStruct structs.CompareData

	countryOneStruct.Country = nameOne
	countryOneStruct.Data = countrySortedDeath
	countryTwoStruct.Country = nameTwo
	countryTwoStruct.Data = countryTwoSortedDeath

	compareStructs := structs.Compare{CountryOne: countryOneStruct, CountryTwo: countryTwoStruct}

	return compareStructs, nil
}

// ComparePerDayDeathsCountries returns two integer arrays (one per country passed
// in parameter) which contain unique per day number of deaths from first confrim death
// It returns structs.Compare and any write error encountered.
func ComparePerDayDeathsCountries(nameOne string, nameTwo string) (structs.Compare, error) {
	countries, err := GetAllCountries()
	if err != nil {
		applogger.Log("ERROR", "curve", "ComparePerDayCasesCountries", err.Error())
		return structs.Compare{}, err
	}
	countryData, countryDataErr := GetCountryData(nameOne, countries)
	if countryDataErr != nil {
		applogger.Log("ERROR", "curve", "ComparePerDayCasesCountries", countryDataErr.Error())
		return structs.Compare{}, countryDataErr
	}

	countryTwoData, countryTwoDataErr := GetCountryData(nameTwo, countries)
	if countryTwoDataErr != nil {
		applogger.Log("ERROR", "curve", "ComparePerDayCasesCountries", countryTwoDataErr.Error())
		return structs.Compare{}, countryTwoDataErr
	}

	var countryOneStruct structs.CompareData
	var countryTwoStruct structs.CompareData

	countryOneStruct.Country = nameOne
	countryOneStruct.Data = countryData.DeathsPerDay
	countryTwoStruct.Country = nameTwo
	countryTwoStruct.Data = countryTwoData.DeathsPerDay

	compareStructs := structs.Compare{CountryOne: countryOneStruct, CountryTwo: countryTwoStruct}
	return compareStructs, nil
}

// CompareRecoveryCountries returns two integer arrays (one per country passed
// in parameter) which contain total number of recovery patients from  22/01/2020
// It returns structs.Compare and any write error encountered.
func CompareRecoveryCountries(nameOne string, nameTwo string) (structs.Compare, error) {
	countries, err := GetAllCountries()
	if err != nil {
		applogger.Log("ERROR", "curve", "ComparePerDayCasesCountries", err.Error())
		return structs.Compare{}, err
	}
	countryData, countryDataErr := GetCountryData(nameOne, countries)
	if countryDataErr != nil {
		applogger.Log("ERROR", "curve", "ComparePerDayCasesCountries", countryDataErr.Error())
		return structs.Compare{}, countryDataErr
	}

	countryTwoData, countryTwoDataErr := GetCountryData(nameTwo, countries)
	if countryTwoDataErr != nil {
		applogger.Log("ERROR", "curve", "ComparePerDayCasesCountries", countryTwoDataErr.Error())
		return structs.Compare{}, countryTwoDataErr
	}

	var countryOneStruct structs.CompareData
	var countryTwoStruct structs.CompareData

	countryOneStruct.Country = nameOne
	countryOneStruct.Data = countryData.Recovered
	countryTwoStruct.Country = nameTwo
	countryTwoStruct.Data = countryTwoData.Recovered

	compareStructs := structs.Compare{CountryOne: countryOneStruct, CountryTwo: countryTwoStruct}

	return compareStructs, nil
}

// CompareCasesCountries returns two integer arrays (one per country passed
// in parameter) which contain total number of cases from  22/01/2020
// It returns structs.Compare and any write error encountered.
func CompareCasesCountries(nameOne string, nameTwo string) (structs.Compare, error) {
	countries, err := GetAllCountries()
	if err != nil {
		applogger.Log("ERROR", "curve", "ComparePerDayCasesCountries", err.Error())
		return structs.Compare{}, err
	}
	countryData, countryDataErr := GetCountryData(nameOne, countries)
	if countryDataErr != nil {
		applogger.Log("ERROR", "curve", "ComparePerDayCasesCountries", countryDataErr.Error())
		return structs.Compare{}, countryDataErr
	}

	countryTwoData, countryTwoDataErr := GetCountryData(nameTwo, countries)
	if countryTwoDataErr != nil {
		applogger.Log("ERROR", "curve", "ComparePerDayCasesCountries", countryTwoDataErr.Error())
		return structs.Compare{}, countryTwoDataErr
	}

	var countryOneStruct structs.CompareData
	var countryTwoStruct structs.CompareData

	countryOneStruct.Country = nameOne
	countryOneStruct.Data = countryData.Cases
	countryTwoStruct.Country = nameTwo
	countryTwoStruct.Data = countryTwoData.Cases

	compareStructs := structs.Compare{CountryOne: countryOneStruct, CountryTwo: countryTwoStruct}

	return compareStructs, nil
}

// ComparePerDayCasesCountries returns two integer arrays (one per country passed
// in parameter) which contain unique per day number of case from first confrim case
// It returns structs.Compare and any write error encountered.
func ComparePerDayCasesCountries(nameOne string, nameTwo string) (structs.Compare, error) {
	countries, err := GetAllCountries()
	if err != nil {
		applogger.Log("ERROR", "curve", "ComparePerDayCasesCountries", err.Error())
		return structs.Compare{}, err
	}
	countryData, countryDataErr := GetCountryData(nameOne, countries)
	if countryDataErr != nil {
		applogger.Log("ERROR", "curve", "ComparePerDayCasesCountries", countryDataErr.Error())
		return structs.Compare{}, countryDataErr
	}

	countryTwoData, countryTwoDataErr := GetCountryData(nameTwo, countries)
	if countryTwoDataErr != nil {
		applogger.Log("ERROR", "curve", "ComparePerDayCasesCountries", countryTwoDataErr.Error())
		return structs.Compare{}, countryTwoDataErr
	}

	var countryOneStruct structs.CompareData
	var countryTwoStruct structs.CompareData

	countryOneStruct.Country = nameOne
	countryOneStruct.Data = countryData.CasesPerDay
	countryTwoStruct.Country = nameTwo
	countryTwoStruct.Data = countryTwoData.CasesPerDay

	compareStructs := structs.Compare{CountryOne: countryOneStruct, CountryTwo: countryTwoStruct}
	return compareStructs, nil
}

func GetCountryData(countryName string, countries []structs.CountryCurve) (structs.MainCurveData, error) {
	country, err := GetCountryBP(countryName, countries)
	if err != nil {
		return structs.MainCurveData{}, err
	}

	deaths := make([]float64, 0)
	cases := make([]float64, 0)
	recovered := make([]float64, 0)

	for _, vv := range country.Timeline.Deaths.(map[string]interface{}) {
		deaths = append(deaths, vv.(float64))
	}

	for _, vv := range country.Timeline.Cases.(map[string]interface{}) {
		cases = append(cases, vv.(float64))
	}

	for _, v := range country.Timeline.Recovered.(map[string]interface{}) {
		recovered = append(recovered, v.(float64))
	}

	sort.Float64s(deaths)
	sort.Float64s(cases)
	sort.Float64s(recovered)

	var tempCountryOneSortedDeath []float64
	deathsPerDay := make([]float64, 0)

	for i := 0; i < len(deaths); i++ {
		tempCountryOneSortedDeath = append(tempCountryOneSortedDeath, deaths[i])
		if i == 0 {
			continue
		}
		deathsPerDay = append(deathsPerDay, (deaths[i] - tempCountryOneSortedDeath[i-1]))
	}

	var tempCountryOneSortedCases []float64
	casesPerDay := make([]float64, 0)
	for i := 0; i < len(cases); i++ {
		tempCountryOneSortedCases = append(tempCountryOneSortedCases, cases[i])
		if i == 0 {
			continue
		}
		casesPerDay = append(casesPerDay, (cases[i] - tempCountryOneSortedCases[i-1]))
	}

	var tempCountryOneSortedRecovered []float64
	recoveredPerDay := make([]float64, 0)
	for i := 0; i < len(deaths); i++ {
		tempCountryOneSortedRecovered = append(tempCountryOneSortedRecovered, recovered[i])
		if i == 0 {
			continue
		}

		recoveredPerDay = append(recoveredPerDay, (recovered[i] - tempCountryOneSortedRecovered[i-1]))
	}

	deathsPerDayFromFirst := make([]float64, 0)

	for i := 0; i < len(deathsPerDay); i++ {
		findFirstDeath := false
		if deathsPerDay[i] == 0 && findFirstDeath == false {
			continue
		}

		if deathsPerDay[i] != 0 {
			findFirstDeath = true
		}

		deathsPerDayFromFirst = append(deathsPerDayFromFirst, deathsPerDay[i])
	}

	return structs.MainCurveData{deaths, deathsPerDay, deathsPerDayFromFirst, cases, casesPerDay, recovered, recoveredPerDay}, nil

}
