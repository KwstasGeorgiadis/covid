package curve

//TODO add fucking caching you piece of shit and add expiration time

import (
	"fmt"
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
	serverConf = pconf.GetAppConfig("./config/covid.json")
)

// requestData does an HTTP GET request to the third party API that
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

	applogger.Log("INFO", "curve", "requestHistoryData",
		fmt.Sprintf("Get reqeust to %s and getting response %v", requestURL, keys))
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
		applogger.Log("INFO", "curve", "GetAllCountries", fmt.Sprintf("Setting cache data %v for key total", data))
		caching.SetCurveData(conn, data)
		applogger.Log("INFO", "curve", "GetAllCountries", fmt.Sprintf("Getting cache data %v instead of requesting it", cachedData))
		return data, nil
	}

	return cachedData, nil
}

// GetCountry seach through an array of structs.CountryCurve and
// gets COVID-19 per day stats for that specific country
// It returns structs.CountryCurve and any write error encountered.
func GetCountry(name string) (structs.CountryCurve, error) {
	allCountries, errGetAllCountries := GetAllCountries()
	if errGetAllCountries != nil {
		applogger.Log("ERROR", "curve", "GetCountry", errGetAllCountries.Error())
		return structs.CountryCurve{}, errGetAllCountries
	}

	for _, v := range allCountries {
		if name == "UK" && v.Country == "UK" {
			if len(v.Province) == 0 {
				applogger.Log("INFO", "curve", "GetCountry", fmt.Sprintf("Returning country %v", v))
				return v, nil
			}
			continue
		}

		if v.Country == name {
			applogger.Log("INFO", "curve", "GetCountry", fmt.Sprintf("Returning country %v", v))
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

	country, errGetCountryOne := GetCountry(nameOne)
	if errGetCountryOne != nil {
		applogger.Log("ERROR", "curve", "CompareDeathsCountries", errGetCountryOne.Error())
		return structs.Compare{}, errGetCountryOne
	}

	countryTwo, errGetCountryTwo := GetCountry(nameTwo)
	if errGetCountryTwo != nil {
		applogger.Log("ERROR", "curve", "CompareDeathsCountries", errGetCountryTwo.Error())
		return structs.Compare{}, errGetCountryTwo
	}

	var countrySortedDeath []float64
	var countryTwoSortedDeath []float64

	for _, v := range country.Timeline.Deaths.(map[string]interface{}) {
		countrySortedDeath = append(countrySortedDeath, v.(float64))
	}
	for _, v := range countryTwo.Timeline.Deaths.(map[string]interface{}) {
		countryTwoSortedDeath = append(countryTwoSortedDeath, v.(float64))
	}
	sort.Float64s(countrySortedDeath)
	sort.Float64s(countryTwoSortedDeath)

	var countryOneStruct structs.CompareData
	var countryTwoStruct structs.CompareData

	countryOneStruct.Country = nameOne
	countryOneStruct.Data = countrySortedDeath
	countryTwoStruct.Country = nameTwo
	countryTwoStruct.Data = countryTwoSortedDeath

	compareStructs := structs.Compare{CountryOne: countryOneStruct, CountryTwo: countryTwoStruct}

	applogger.Log("INFO", "curve", "CompareDeathsCountries", fmt.Sprintf("Returning country comperation %v", compareStructs))
	return compareStructs, nil
}

// CompareDeathsFromFirstDeathCountries returns two integer arrays (one per country passed
// in parameter) which contain total number of deaths from  the first confirm death.
// It returns structs.Compare and any write error encountered.
func CompareDeathsFromFirstDeathCountries(nameOne string, nameTwo string) (structs.Compare, error) {

	country, errGetCountryOne := GetCountry(nameOne)
	if errGetCountryOne != nil {
		applogger.Log("ERROR", "curve", "CompareDeathsFromFirstDeathCountries", errGetCountryOne.Error())
		return structs.Compare{}, errGetCountryOne
	}

	countryTwo, errGetCountryTwo := GetCountry(nameTwo)
	if errGetCountryTwo != nil {
		applogger.Log("ERROR", "curve", "CompareDeathsFromFirstDeathCountries", errGetCountryTwo.Error())
		return structs.Compare{}, errGetCountryTwo
	}

	var countrySortedDeath []float64
	var countryTwoSortedDeath []float64

	for _, v := range country.Timeline.Deaths.(map[string]interface{}) {
		if v.(float64) == 0 {
			continue
		}
		countrySortedDeath = append(countrySortedDeath, v.(float64))
	}
	for _, v := range countryTwo.Timeline.Deaths.(map[string]interface{}) {
		if v.(float64) == 0 {
			continue
		}
		countryTwoSortedDeath = append(countryTwoSortedDeath, v.(float64))
	}
	sort.Float64s(countrySortedDeath)
	sort.Float64s(countryTwoSortedDeath)

	var countryOneStruct structs.CompareData
	var countryTwoStruct structs.CompareData

	countryOneStruct.Country = nameOne
	countryOneStruct.Data = countrySortedDeath
	countryTwoStruct.Country = nameTwo
	countryTwoStruct.Data = countryTwoSortedDeath

	compareStructs := structs.Compare{CountryOne: countryOneStruct, CountryTwo: countryTwoStruct}

	applogger.Log("INFO", "curve", "CompareDeathsFromFirstDeathCountries", fmt.Sprintf("Returning country comperation %v", compareStructs))
	return compareStructs, nil
}

// ComparePerCentDeathsCountries returns two integer arrays (one per country passed
// in parameter) which contain incremental percentage of deaths from  the first confirm death.
// It returns structs.Compare and any write error encountered.
func ComparePerCentDeathsCountries(nameOne string, nameTwo string) (structs.Compare, error) {
	country, errGetCountryOne := GetCountry(nameOne)
	if errGetCountryOne != nil {
		applogger.Log("ERROR", "curve", "ComparePerCentDeathsCountries", errGetCountryOne.Error())
		return structs.Compare{}, errGetCountryOne
	}

	countryTwo, errGetCountryTwo := GetCountry(nameTwo)
	if errGetCountryTwo != nil {
		applogger.Log("ERROR", "curve", "ComparePerCentDeathsCountries", errGetCountryTwo.Error())
		return structs.Compare{}, errGetCountryTwo
	}

	var countrySortedDeath []float64
	var countryTwoSortedDeath []float64

	for _, v := range country.Timeline.Deaths.(map[string]interface{}) {

		if v.(float64) == 0 {
			continue
		}
		countrySortedDeath = append(countrySortedDeath, v.(float64))
	}
	for _, v := range countryTwo.Timeline.Deaths.(map[string]interface{}) {
		if v.(float64) == 0 {
			continue
		}
		countryTwoSortedDeath = append(countryTwoSortedDeath, v.(float64))
	}

	sort.Float64s(countrySortedDeath)
	sort.Float64s(countryTwoSortedDeath)

	var countryOneStruct structs.CompareData
	var countryTwoStruct structs.CompareData

	for _, v := range country.Timeline.Deaths.(map[string]interface{}) {

		countrySortedDeath = append(countrySortedDeath, v.(float64))
	}
	for _, v := range countryTwo.Timeline.Deaths.(map[string]interface{}) {
		countryTwoSortedDeath = append(countryTwoSortedDeath, v.(float64))
	}

	countryOneStruct.Country = nameOne
	countryOneStruct.Data = countrySortedDeath
	countryTwoStruct.Country = nameTwo
	countryTwoStruct.Data = countryTwoSortedDeath

	compareStructs := structs.Compare{CountryOne: countryOneStruct, CountryTwo: countryTwoStruct}
	applogger.Log("INFO", "curve", "CompareDeathsFromFirstDeathCountries", fmt.Sprintf("Returning country comperation %v", compareStructs))

	return compareStructs, nil
}

// ComparePerDayDeathsCountries returns two integer arrays (one per country passed
// in parameter) which contain unique per day number of deaths from first confrim death
// It returns structs.Compare and any write error encountered.
func ComparePerDayDeathsCountries(nameOne string, nameTwo string) (structs.Compare, error) {

	country, errGetCountryOne := GetCountry(nameOne)
	if errGetCountryOne != nil {
		applogger.Log("ERROR", "curve", "ComparePerDayDeathsCountries", errGetCountryOne.Error())
		fmt.Println(errGetCountryOne.Error())
		return structs.Compare{}, errGetCountryOne
	}

	countryTwo, errGetCountryTwo := GetCountry(nameTwo)
	if errGetCountryTwo != nil {
		applogger.Log("ERROR", "curve", "ComparePerDayDeathsCountries", errGetCountryTwo.Error())
		return structs.Compare{}, errGetCountryTwo
	}

	var countrySortedDeath []float64
	var countryTwoSortedDeath []float64

	for _, v := range country.Timeline.Deaths.(map[string]interface{}) {
		if v.(float64) == 0 {
			continue
		}
		countrySortedDeath = append(countrySortedDeath, v.(float64))
	}
	for _, v := range countryTwo.Timeline.Deaths.(map[string]interface{}) {
		if v.(float64) == 0 {
			continue
		}
		countryTwoSortedDeath = append(countryTwoSortedDeath, v.(float64))
	}
	sort.Float64s(countrySortedDeath)
	sort.Float64s(countryTwoSortedDeath)

	var tempCountryOneSortedDeath []float64
	for i := 0; i < len(countrySortedDeath); i++ {
		tempCountryOneSortedDeath = append(tempCountryOneSortedDeath, countrySortedDeath[i])
		if i == 0 {
			continue
		}

		countrySortedDeath[i] = countrySortedDeath[i] - tempCountryOneSortedDeath[i-1]
	}

	var tempCountryTwoSortedDeath []float64
	for i := 0; i < len(countryTwoSortedDeath); i++ {
		tempCountryTwoSortedDeath = append(tempCountryTwoSortedDeath, countryTwoSortedDeath[i])

		if i == 0 {
			continue
		}
		countryTwoSortedDeath[i] = countryTwoSortedDeath[i] - tempCountryTwoSortedDeath[i-1]
	}

	var countryOneStruct structs.CompareData
	var countryTwoStruct structs.CompareData

	countryOneStruct.Country = nameOne
	countryOneStruct.Data = countrySortedDeath
	countryTwoStruct.Country = nameTwo
	countryTwoStruct.Data = countryTwoSortedDeath

	compareStructs := structs.Compare{CountryOne: countryOneStruct, CountryTwo: countryTwoStruct}
	applogger.Log("INFO", "curve", "CompareDeathsFromFirstDeathCountries", fmt.Sprintf("Returning country comperation %v", compareStructs))
	return compareStructs, nil
}
