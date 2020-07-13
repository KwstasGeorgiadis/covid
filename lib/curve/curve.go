package curve

import (
	"sort"

	applogger "github.com/junkd0g/covid/lib/applogger"
	caching "github.com/junkd0g/covid/lib/caching"
	pconf "github.com/junkd0g/covid/lib/config"
	mcountry "github.com/junkd0g/covid/lib/model/country"

	"encoding/json"
	"io/ioutil"
	"net/http"
)

var (
	serverConf pconf.AppConf
	reqDataOB  requestAPI
	reqCacheOB requestCache
	redis      caching.RedisST
)

func init() {
	serverConf = pconf.GetAppConfig()
	reqDataOB = requestData{}
	reqCacheOB = requestCacheData{}
}

type requestData struct{}
type requestAPI interface {
	requestHistoryData() ([]mcountry.CountryCurve, error)
}

type requestCacheData struct{}
type requestCache interface {
	getCacheData() ([]mcountry.CountryCurve, error)
	setCacheData(ctn []mcountry.CountryCurve) error
}

func (r requestCacheData) setCacheData(ctn []mcountry.CountryCurve) error {
	err := redis.SetCurveData(ctn)
	return err
}

func (r requestCacheData) getCacheData() ([]mcountry.CountryCurve, error) {
	cachedData, cacheGetError := redis.GetCurveData()
	return cachedData, cacheGetError
}

// requestHistoryData does an HTTP GET request to the third party API that
// contains covid-9 stats ' history (per day from 22/01/2020)
// It returns []mcountry.Country and any write error encountered.
func (r requestData) requestHistoryData() ([]mcountry.CountryCurve, error) {
	client := &http.Client{}
	requestURL := serverConf.API.URLHistory

	req, reqErr := http.NewRequest("GET", requestURL, nil)
	if reqErr != nil {
		applogger.Log("ERROR", "curve", "requestHistoryData", reqErr.Error())
		return []mcountry.CountryCurve{}, reqErr
	}

	res, resError := client.Do(req)
	if resError != nil {
		applogger.Log("ERROR", "curve", "requestHistoryData", resError.Error())
		return []mcountry.CountryCurve{}, resError
	}
	defer res.Body.Close()

	b, errorReadAll := ioutil.ReadAll(res.Body)
	if errorReadAll != nil {
		applogger.Log("ERROR", "curve", "requestHistoryData", errorReadAll.Error())
		return []mcountry.CountryCurve{}, errorReadAll
	}

	keys := make([]mcountry.CountryCurve, 0)
	if errUnmarshal := json.Unmarshal(b, &keys); errUnmarshal != nil {
		applogger.Log("ERROR", "curve", "requestHistoryData", errUnmarshal.Error())
		return []mcountry.CountryCurve{}, errUnmarshal
	}

	return keys, nil
}

// GetAllCountries returns an array of all countries per day
// Covid-19 stats (data starts from date 22/01/2020)
// Check if there are cached data if not does a HTTP
// request to the 3rd party API (check requestHistoryData())
// It returns []structs.CountryCurve and any write error encountered.
func GetAllCountries() ([]mcountry.CountryCurve, error) {
	cachedData, cacheGetError := reqCacheOB.getCacheData()
	if cacheGetError != nil {
		applogger.Log("ERROR", "curve", "GetAllCountries", cacheGetError.Error())
		return []mcountry.CountryCurve{}, cacheGetError
	}

	if len(cachedData) == 0 {
		applogger.Log("INFO", "stats", "GetAllCountries", "Request data instead of getting cached data")
		data, err := reqDataOB.requestHistoryData()
		if err != nil {
			applogger.Log("ERROR", "curve", "GetAllCountries", err.Error())
			return []mcountry.CountryCurve{}, err
		}
		reqCacheOB.setCacheData(data)
		return data, nil
	}

	return cachedData, nil
}

// GetCountryBP seach through an array of structs.CountryCurve and
// gets COVID-19 per day stats for that specific country
// It returns mcountry.CountryCurve and any write error encountered.
func GetCountryBP(name string, allCountries []mcountry.CountryCurve) (mcountry.CountryCurve, error) {

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
	return mcountry.CountryCurve{}, nil
}

// CompareDeathsCountries returns two integer arrays (one per country passed
// in parameter) which contain total number of deaths from  22/01/2020
// It returns mcountry.Compare and any write error encountered.
func CompareDeathsCountries(nameOne string, nameTwo string) (mcountry.Compare, error) {
	countries, err := GetAllCountries()
	if err != nil {
		applogger.Log("ERROR", "curve", "ComparePerDayCasesCountries", err.Error())
		return mcountry.Compare{}, err
	}
	countryData, countryDataErr := GetCountryData(nameOne, countries)
	if countryDataErr != nil {
		applogger.Log("ERROR", "curve", "ComparePerDayCasesCountries", countryDataErr.Error())
		return mcountry.Compare{}, countryDataErr
	}

	countryTwoData, countryTwoDataErr := GetCountryData(nameTwo, countries)
	if countryTwoDataErr != nil {
		applogger.Log("ERROR", "curve", "ComparePerDayCasesCountries", countryTwoDataErr.Error())
		return mcountry.Compare{}, countryTwoDataErr
	}

	var countryOneStruct mcountry.CompareData
	var countryTwoStruct mcountry.CompareData

	countryOneStruct.Country = nameOne
	countryOneStruct.Data = countryData.Deaths
	countryTwoStruct.Country = nameTwo
	countryTwoStruct.Data = countryTwoData.Deaths

	compareStructs := mcountry.Compare{CountryOne: countryOneStruct, CountryTwo: countryTwoStruct}

	return compareStructs, nil
}

// CompareDeathsFromFirstDeathCountries returns two integer arrays (one per country passed
// in parameter) which contain total number of deaths from  the first confirm death.
// It returns mcountry.Compare and any write error encountered.
func CompareDeathsFromFirstDeathCountries(nameOne string, nameTwo string) (mcountry.Compare, error) {
	countries, err := GetAllCountries()
	if err != nil {
		applogger.Log("ERROR", "curve", "ComparePerDayCasesCountries", err.Error())
		return mcountry.Compare{}, err
	}

	countryData, countryDataErr := GetCountryData(nameOne, countries)
	if countryDataErr != nil {
		applogger.Log("ERROR", "curve", "ComparePerDayCasesCountries", countryDataErr.Error())
		return mcountry.Compare{}, countryDataErr
	}

	countryTwoData, countryTwoDataErr := GetCountryData(nameTwo, countries)
	if countryTwoDataErr != nil {
		applogger.Log("ERROR", "curve", "ComparePerDayCasesCountries", countryTwoDataErr.Error())
		return mcountry.Compare{}, countryTwoDataErr
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

	var countryOneStruct mcountry.CompareData
	var countryTwoStruct mcountry.CompareData

	countryOneStruct.Country = nameOne
	countryOneStruct.Data = countrySortedDeath
	countryTwoStruct.Country = nameTwo
	countryTwoStruct.Data = countryTwoSortedDeath

	compareStructs := mcountry.Compare{CountryOne: countryOneStruct, CountryTwo: countryTwoStruct}

	return compareStructs, nil
}

// ComparePerDayDeathsCountries returns two integer arrays (one per country passed
// in parameter) which contain unique per day number of deaths from first confrim death
// It returns mcountry.Compare and any write error encountered.
func ComparePerDayDeathsCountries(nameOne string, nameTwo string) (mcountry.Compare, error) {
	countries, err := GetAllCountries()
	if err != nil {
		applogger.Log("ERROR", "curve", "ComparePerDayCasesCountries", err.Error())
		return mcountry.Compare{}, err
	}
	countryData, countryDataErr := GetCountryData(nameOne, countries)
	if countryDataErr != nil {
		applogger.Log("ERROR", "curve", "ComparePerDayCasesCountries", countryDataErr.Error())
		return mcountry.Compare{}, countryDataErr
	}

	countryTwoData, countryTwoDataErr := GetCountryData(nameTwo, countries)
	if countryTwoDataErr != nil {
		applogger.Log("ERROR", "curve", "ComparePerDayCasesCountries", countryTwoDataErr.Error())
		return mcountry.Compare{}, countryTwoDataErr
	}

	var countryOneStruct mcountry.CompareData
	var countryTwoStruct mcountry.CompareData

	countryOneStruct.Country = nameOne
	countryOneStruct.Data = countryData.DeathsPerDay
	countryTwoStruct.Country = nameTwo
	countryTwoStruct.Data = countryTwoData.DeathsPerDay

	compareStructs := mcountry.Compare{CountryOne: countryOneStruct, CountryTwo: countryTwoStruct}
	return compareStructs, nil
}

// CompareRecoveryCountries returns two integer arrays (one per country passed
// in parameter) which contain total number of recovery patients from  22/01/2020
// It returns mcountry.Compare and any write error encountered.
func CompareRecoveryCountries(nameOne string, nameTwo string) (mcountry.Compare, error) {
	countries, err := GetAllCountries()
	if err != nil {
		applogger.Log("ERROR", "curve", "ComparePerDayCasesCountries", err.Error())
		return mcountry.Compare{}, err
	}
	countryData, countryDataErr := GetCountryData(nameOne, countries)
	if countryDataErr != nil {
		applogger.Log("ERROR", "curve", "ComparePerDayCasesCountries", countryDataErr.Error())
		return mcountry.Compare{}, countryDataErr
	}

	countryTwoData, countryTwoDataErr := GetCountryData(nameTwo, countries)
	if countryTwoDataErr != nil {
		applogger.Log("ERROR", "curve", "ComparePerDayCasesCountries", countryTwoDataErr.Error())
		return mcountry.Compare{}, countryTwoDataErr
	}

	var countryOneStruct mcountry.CompareData
	var countryTwoStruct mcountry.CompareData

	countryOneStruct.Country = nameOne
	countryOneStruct.Data = countryData.Recovered
	countryTwoStruct.Country = nameTwo
	countryTwoStruct.Data = countryTwoData.Recovered

	compareStructs := mcountry.Compare{CountryOne: countryOneStruct, CountryTwo: countryTwoStruct}

	return compareStructs, nil
}

// CompareCasesCountries returns two integer arrays (one per country passed
// in parameter) which contain total number of cases from  22/01/2020
// It returns mcountry.Compare and any write error encountered.
func CompareCasesCountries(nameOne string, nameTwo string) (mcountry.Compare, error) {
	countries, err := GetAllCountries()
	if err != nil {
		applogger.Log("ERROR", "curve", "ComparePerDayCasesCountries", err.Error())
		return mcountry.Compare{}, err
	}
	countryData, countryDataErr := GetCountryData(nameOne, countries)
	if countryDataErr != nil {
		applogger.Log("ERROR", "curve", "ComparePerDayCasesCountries", countryDataErr.Error())
		return mcountry.Compare{}, countryDataErr
	}

	countryTwoData, countryTwoDataErr := GetCountryData(nameTwo, countries)
	if countryTwoDataErr != nil {
		applogger.Log("ERROR", "curve", "ComparePerDayCasesCountries", countryTwoDataErr.Error())
		return mcountry.Compare{}, countryTwoDataErr
	}

	var countryOneStruct mcountry.CompareData
	var countryTwoStruct mcountry.CompareData

	countryOneStruct.Country = nameOne
	countryOneStruct.Data = countryData.Cases
	countryTwoStruct.Country = nameTwo
	countryTwoStruct.Data = countryTwoData.Cases

	compareStructs := mcountry.Compare{CountryOne: countryOneStruct, CountryTwo: countryTwoStruct}

	return compareStructs, nil
}

// ComparePerDayCasesCountries returns two integer arrays (one per country passed
// in parameter) which contain unique per day number of case from first confrim case
// It returns mcountry.Compare and any write error encountered.
func ComparePerDayCasesCountries(nameOne string, nameTwo string) (mcountry.Compare, error) {
	countries, err := GetAllCountries()
	if err != nil {
		applogger.Log("ERROR", "curve", "ComparePerDayCasesCountries", err.Error())
		return mcountry.Compare{}, err
	}
	countryData, countryDataErr := GetCountryData(nameOne, countries)
	if countryDataErr != nil {
		applogger.Log("ERROR", "curve", "ComparePerDayCasesCountries", countryDataErr.Error())
		return mcountry.Compare{}, countryDataErr
	}

	countryTwoData, countryTwoDataErr := GetCountryData(nameTwo, countries)
	if countryTwoDataErr != nil {
		applogger.Log("ERROR", "curve", "ComparePerDayCasesCountries", countryTwoDataErr.Error())
		return mcountry.Compare{}, countryTwoDataErr
	}

	var countryOneStruct mcountry.CompareData
	var countryTwoStruct mcountry.CompareData

	countryOneStruct.Country = nameOne
	countryOneStruct.Data = countryData.CasesPerDay
	countryTwoStruct.Country = nameTwo
	countryTwoStruct.Data = countryTwoData.CasesPerDay

	compareStructs := mcountry.Compare{CountryOne: countryOneStruct, CountryTwo: countryTwoStruct}
	return compareStructs, nil
}

func GetCountryData(countryName string, countries []mcountry.CountryCurve) (mcountry.MainCurveData, error) {
	country, err := GetCountryBP(countryName, countries)
	if err != nil {
		return mcountry.MainCurveData{}, err
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

	return mcountry.MainCurveData{deaths, deathsPerDay, deathsPerDayFromFirst, cases, casesPerDay, recovered, recoveredPerDay}, nil

}
